package client

import (
	"context"
	"crypto/x509"
	"io"
	"net/http"

	"github.com/spiffe/go-spiffe/v2/bundle/spiffebundle"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/common/tlspolicy"
)

type SPIFFEAuthConfig struct {
	// EndpointSpiffeID is the expected SPIFFE ID of the bundle endpoint server.
	EndpointSpiffeID spiffeid.ID

	// RootCAs is the set of root CA certificates used to authenticate the
	// endpoint server.
	RootCAs []*x509.Certificate
}

type ClientConfig struct { //revive:disable-line:exported name stutter is intentional
	// TrustDomain is the federated trust domain (i.e. domain.test)
	TrustDomain spiffeid.TrustDomain

	// EndpointURL is the URL used to fetch the bundle of the federated
	// trust domain. Is served by a SPIFFE bundle endpoint server.
	EndpointURL string

	// SPIFFEAuth contains required configuration to authenticate the endpoint
	// using SPIFFE authentication. If unset, it is assumed that the endpoint
	// is authenticated via Web PKI.
	SPIFFEAuth *SPIFFEAuthConfig

	// TLSPolicy specifies the post-quantum-security policy used for TLS
	// connections.
	TLSPolicy tlspolicy.Policy

	// mutateTransportHook is a hook to influence the transport used during
	// tests.
	mutateTransportHook func(*http.Transport)
}

// Client is used to fetch a bundle and metadata from a bundle endpoint
type Client interface {
	FetchBundle(context.Context) (*spiffebundle.Bundle, error)
}

type client struct {
	c      ClientConfig
	client *http.Client
}

func NewClient(config ClientConfig) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

func (c *client) FetchBundle(context.Context) (*spiffebundle.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func tryRead(r io.Reader) string { _ = "STUB: not implemented"; return "" }

func newTransport() *http.Transport { _ = "STUB: not implemented"; return nil }
