package main

import (
	"net/http"
	"time"
)

const (
	ThresholdMultiplicator = 5
	ThresholdMinTime       = time.Minute * 3
)

type HealthChecksHandler struct {
	source       JWKSSource
	healthChecks HealthChecksConfig
	jwkThreshold time.Duration
	initTime     time.Time

	http.Handler
}

func NewHealthChecksHandler(source JWKSSource, config *Config) *HealthChecksHandler {
	_ = "STUB: not implemented"
	return nil
}

// jwkThreshold determines the duration from the last successful poll before the server is considered unhealthy
func jwkThreshold(config *Config) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// readyCheck is a health check that returns 200 if the server can successfully fetch a jwt keyset
func (h *HealthChecksHandler) readyCheck(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// liveCheck is a health check that returns 200 if the server is able to reply to http requests
func (h *HealthChecksHandler) liveCheck(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
