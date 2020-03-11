package main

import (
	"flag"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var (
	port                       = flag.Int("port", 7788, "The proxy server port")
	cloud                      = flag.String("cloud", "public", "The cloud environment")
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

	proxy := &ProxyServer{
		port:     *port,
		cloud:    *cloud,
		logger:   logger,
		client:   &http.Client{},
		throttle: &AutoThrustFactory{logger: logger, defaultRetryTryAfterInSeconds: *defaultRetryAfterInSeconds},
	}

	proxy.ListenAndServe()
}
