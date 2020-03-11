package main

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

// Throttle is the interface to get throttlig handler
type Throttle interface {
	NewThrottle(handler http.Handler) http.Handler
}

// AutoThrustFactory news the AutoThrust instance
type AutoThrustFactory struct {
	logger                        *log.Logger
	defaultRetryTryAfterInSeconds int
}

//NewThrottle constructs a http handler for throttling
func (f *AutoThrustFactory) NewThrottle(handlerToWrap http.Handler) http.Handler {
	return &AutoThrust{f.logger, handlerToWrap, sync.Map{}, f.defaultRetryTryAfterInSeconds}
}

//AutoThrust is a middleware handler that does request throttling
type AutoThrust struct {
	logger                        *log.Logger
	handler                       http.Handler
	cache                         sync.Map
	defaultRetryTryAfterInSeconds int
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

type cacheItem struct {
	until time.Time
}

//ServeHTTP handles the request by passing it to the real
//handler and logging the request details
func (t *AutoThrust) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if c, ok := t.cache.Load(r.RequestURI); ok {
		shouldAfter := c.(*cacheItem).until
		if time.Now().After(shouldAfter) {
			t.cache.Delete(r.RequestURI)
		} else {
			t.logger.Infof("Request %s will be throttled until %v", r.RequestURI, shouldAfter)

			// Add 5 seconds delay to avoid client side spin
			time.Sleep(5 * time.Second)

			w.WriteHeader(429)
			return
		}
	}

	rec := statusRecorder{w, 0}
	t.handler.ServeHTTP(&rec, r)

	if rec.status == 429 {
		// Default retry after is 3 minutes
		retryAfter := time.Now().Add(time.Duration(t.defaultRetryTryAfterInSeconds) * time.Second)
		retryAfterStr := w.Header().Get("Retry-After")
		if retryAfterStr != "" {
			if retrySec, _ := strconv.Atoi(retryAfterStr); retrySec > 0 {
				retryAfter = time.Now().Add(time.Duration(retrySec) * time.Second)
			} else if retryTime, err := time.Parse(time.RFC1123, retryAfterStr); err == nil {
				retryAfter = retryTime
			}
		}

		t.logger.Infof("Setting retry after to %s for request %v", retryAfter, r.RequestURI)

		ci := &cacheItem{until: retryAfter}
		t.cache.Store(r.RequestURI, ci)
	}
}
