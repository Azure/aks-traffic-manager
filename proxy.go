package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	"github.com/Azure/go-autorest/autorest/azure"
	log "github.com/sirupsen/logrus"
)

// ProxyServer serves the client request and proxy to ARM endpoint
type ProxyServer struct {
	port     int
	cloud    string
	logger   *log.Logger
	client   *http.Client
	throttle Throttle
}

// ListenAndServe starts the proxy server
func (s *ProxyServer) ListenAndServe() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handleAzureRequests)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: s.throttle.NewThrottle(mux),
	}

	go func() {
		s.logger.Info("Start listening")
		if err := server.ListenAndServe(); err != nil {
			s.logger.Fatal("Failed to listen on proxy address: ", err)
		}
	}()

	// listening to OS shutdown singal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	s.logger.Infof("Got shutdown signal, shutting down webhook server gracefully...")

	server.Shutdown(context.Background()) // nolint: errcheck
}

func (s *ProxyServer) handleAzureRequests(w http.ResponseWriter, req *http.Request) {
	logger := s.logger
	logger.Infof("Proxy %s %s", req.Method, req.RequestURI)

	defer req.Body.Close()
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Errorf("Failed to read request: %v", err.Error())
		s.handleProxyFailure(w)
		return
	}

	apiVersion := req.URL.Query()["api-version"]
	buf, err = enableTCPReset(req.Method, req.RequestURI, apiVersion[0], buf)
	if err != nil {
		logger.Errorf("Failed to enable TCP Reset: %v", err.Error())
		s.handleProxyFailure(w)
		return
	}

	env, err := azure.EnvironmentFromName(s.cloud)
	if err != nil {
		logger.Errorf("error in getting %s environment: %v", s.cloud, err)
		return
	}

	reqURI := fmt.Sprintf("%s%s", env.ResourceManagerEndpoint, req.RequestURI)
	proxyReq, err := http.NewRequest(req.Method, reqURI, bytes.NewReader(buf))
	if err != nil {
		logger.Errorf("Unable to construct request: %s", err.Error())
		s.handleProxyFailure(w)
		return
	}

	copyHeader(proxyReq.Header, req.Header)

	resp, err := s.client.Do(proxyReq)
	if err != nil {
		logger.Errorf("Failed to proxy request: %v", err.Error())
		s.handleProxyFailure(w)
		return
	}

	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)

	defer resp.Body.Close()

	if _, err := io.Copy(w, resp.Body); err != nil {
		logger.Errorf("Failed to copy response body: %v", err.Error())
		s.handleProxyFailure(w)
	}
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func (s *ProxyServer) handleProxyFailure(w http.ResponseWriter) {
	w.WriteHeader(502)
	_, err := w.Write([]byte("Internal Server Error"))
	if err != nil {
		s.logger.Errorf("Failed to handle proxy failure: %v", err.Error())
	}
}

func enableTCPReset(httpMethod, requestURI, apiVersion string, input []byte) (output []byte, err error) {
	if !strings.EqualFold(httpMethod, "PUT") ||
		strings.Compare(apiVersion, "2018-07-01") < 0 {
		return input, nil
	}

	var resourceType, _ = getResourceType(requestURI)
	if !strings.EqualFold(resourceType, "Microsoft.Network/loadBalancers") {
		return input, nil
	}

	var jsonBody map[string]interface{}
	if err := json.Unmarshal(input, &jsonBody); err != nil {
		return input, nil
	}

	sku := jsonBody["sku"]
	if sku == nil || !strings.EqualFold(sku.(map[string]interface{})["name"].(string), "Standard") {
		return input, nil
	}

	if jsonBody["properties"] == nil {
		return input, nil
	}

	properties := jsonBody["properties"].(map[string]interface{})

	if properties != nil && properties["loadBalancingRules"] != nil {
		loadBalancingRules := properties["loadBalancingRules"].([]interface{})
		for _, lbrule := range loadBalancingRules {
			rule := lbrule.(map[string]interface{})
			setEnableTCPReset(rule["properties"].(map[string]interface{}))
		}
	}

	if properties != nil && properties["outboundRules"] != nil {
		outboundRules := properties["outboundRules"].([]interface{})
		for _, obrule := range outboundRules {
			rule := obrule.(map[string]interface{})
			setEnableTCPReset(rule["properties"].(map[string]interface{}))
		}
	}

	if properties != nil && properties["inboundNatRules"] != nil {
		inboundNatRules := properties["inboundNatRules"].([]interface{})
		for _, natrule := range inboundNatRules {
			rule := natrule.(map[string]interface{})
			setEnableTCPReset(rule["properties"].(map[string]interface{}))
		}
	}

	return json.Marshal(jsonBody)
}

func setEnableTCPReset(properties map[string]interface{}) {
	if properties != nil {
		properties["enableTcpReset"] = true
	}
}

// based on the ParseResourceID method defined here
// https://github.com/Azure/go-autorest/blob/master/autorest/azure/azure.go#L176
func getResourceType(resourceID string) (string, error) {
	const resourceIDPatternText = `(?i)subscriptions/(.+)/resourceGroups/(.+)/providers/(.+?)/(.+?)/(.+)`
	resourceIDPattern := regexp.MustCompile(resourceIDPatternText)
	match := resourceIDPattern.FindStringSubmatch(resourceID)

	if len(match) == 0 || len(match) < 5 {
		return "", fmt.Errorf("parsing failed for %s. Invalid resource Id format", resourceID)
	}

	return fmt.Sprintf("%s/%s", match[3], match[4]), nil
}
