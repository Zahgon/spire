package bundle

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/bundle/spiffebundle"
)

type Getter interface {
	GetBundle(ctx context.Context) (*spiffebundle.Bundle, error)
}

type GetterFunc func(ctx context.Context) (*spiffebundle.Bundle, error)

func (fn GetterFunc) GetBundle(ctx context.Context) (*spiffebundle.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ServerAuth interface {
	GetTLSConfig() *tls.Config
}

type ServerConfig struct {
	Log         logrus.FieldLogger
	Address     string
	Getter      Getter
	ServerAuth  ServerAuth
	RefreshHint time.Duration

	// test hooks
	listen func(network, address string) (net.Listener, error)
}

type Server struct {
	c ServerConfig
}

func NewServer(config ServerConfig) *Server { _ = "STUB: not implemented"; return nil }

func (s *Server) ListenAndServe(ctx context.Context) error {
	_ = "STUB: not implemented"
	// create the listener explicitly instead of using ListenAndServeTLS since
	// it gives us the ability to use/inspect an ephemeral port during testing.
	return nil
}

// Set up the TLS config, setting TLS 1.2 as the minimum.

func (s *Server) WaitForListening() {
	_ = "STUB: not implemented"
	// This method is a no-op for the bundle server since it does not have a
	// separate listening hook like the agent endpoints.
	// If needed, this can be implemented to signal when the server starts
	// listening.
	return
}

func (s *Server) serveHTTP(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// TODO: bundle sequence number?

func chainDER(chain []*x509.Certificate) [][]byte { _ = "STUB: not implemented"; return nil }
