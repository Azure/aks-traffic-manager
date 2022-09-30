package main

import (
	"crypto/tls"
	"flag"
	"net"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	port                       = flag.Int("port", 7788, "The proxy server port")
	cloud                      = flag.String("cloud", "AzurePublicCloud", "The cloud environment")
	defaultRetryAfterInSeconds = flag.Int("default-retry-after-in-seconds", 25, "The default value for retry after if no retry-after header is present ")
)

func main() {
	flag.Parse()

	logger := log.New()
	if port == nil {
		logger.Fatal("port is not specified")
	}
	if cloud == nil || *cloud == "" {
		logger.Fatal("The cloud specified is invalid or not specified")
	}

	logger.Infof("port: %d", *port)
	logger.Infof("cloud: %s", *cloud)
	logger.Infof("default-retry-after-in-seconds: %d", *defaultRetryAfterInSeconds)

	httpTransport := http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // the same as default transport
			KeepAlive: 30 * time.Second, // the same as default transport
		}).DialContext,
		ForceAttemptHTTP2:     true,             // always attempt HTTP/2 even though custom dialer is provided
		MaxIdleConns:          100,              // Zero means no limit, the same as default transport
		MaxIdleConnsPerHost:   100,              // Default is 2, ref:https://cs.opensource.google/go/go/+/go1.18.4:src/net/http/transport.go;l=58
		IdleConnTimeout:       90 * time.Second, // the same as default transport
		TLSHandshakeTimeout:   10 * time.Second, // the same as default transport
		ExpectContinueTimeout: 1 * time.Second,  // the same as default transport
		TLSClientConfig: &tls.Config{
			MinVersion:    tls.VersionTLS12,     //force to use TLS 1.2
			Renegotiation: tls.RenegotiateNever, // the same as default transport https://pkg.go.dev/crypto/tls#RenegotiationSupport
		},
	}
	proxy := &ProxyServer{
		port:   *port,
		cloud:  *cloud,
		logger: logger,
		client: &http.Client{
			Transport: &httpTransport,
		},
		throttle: &AutoThrustFactory{logger: logger, defaultRetryTryAfterInSeconds: *defaultRetryAfterInSeconds},
	}

	proxy.ListenAndServe()
}
