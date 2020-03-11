package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	reqURI := fmt.Sprintf("https://management.azure.com%s", req.RequestURI)
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
