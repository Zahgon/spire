package main

import (
	"net/http"
	"net/url"

	"github.com/go-jose/go-jose/v4"
	"github.com/sirupsen/logrus"
)

const (
	keyUse = "sig"
)

type Handler struct {
	source              JWKSSource
	domainPolicy        DomainPolicy
	allowInsecureScheme bool
	setKeyUse           bool
	log                 logrus.FieldLogger
	jwtIssuer           *url.URL
	jwksURI             *url.URL
	serverPathPrefix    string

	http.Handler
}

func NewHandler(log logrus.FieldLogger, domainPolicy DomainPolicy, source JWKSSource, allowInsecureScheme, setKeyUse bool, jwtIssuer, jwksURI *url.URL, serverPathPrefix string) (*Handler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *Handler) serveWellKnown(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// If jwksIsser is set but not jwksURI, fall back to 1.11.1 behavior until we can remove jwksIssuer leaking into jwksURI in 1.13.0

// The following are required fields that we'll just hardcode response
// to based on SPIRE capabilities, etc.

func (h *Handler) serveKeys(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Disable caching

func (h *Handler) verifyHost(host string) error {
	_ = "STUB: not implemented"
	// Obtain the domain name from the host value, which comes from the
	// request, or is pulled from the X-Forwarded-Host header (via the
	// ProxyHeaders middleware). The value may be in host or host:port form.
	return nil
}

// `Host` was not in the host:port form.

func (h *Handler) enrichJwksKeys(jwkKeys []jose.JSONWebKey) []jose.JSONWebKey {
	_ = "STUB: not implemented"
	return nil
}
