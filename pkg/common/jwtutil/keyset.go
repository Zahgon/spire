package jwtutil

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/go-jose/go-jose/v4"
)

const (
	wellKnownOpenIDConfiguration = "/.well-known/openid-configuration"
)

type KeySetProvider interface {
	GetKeySet(context.Context) (*jose.JSONWebKeySet, error)
}

type KeySetProviderFunc func(context.Context) (*jose.JSONWebKeySet, error)

func (fn KeySetProviderFunc) GetKeySet(ctx context.Context) (*jose.JSONWebKeySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type OIDCIssuer string

func (c OIDCIssuer) GetKeySet(ctx context.Context) (*jose.JSONWebKeySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type CachingKeySetProvider struct {
	provider        KeySetProvider
	refreshInterval time.Duration

	mu      sync.Mutex
	updated time.Time
	jwks    *jose.JSONWebKeySet

	hooks struct {
		now func() time.Time
	}
}

func NewCachingKeySetProvider(provider KeySetProvider, refreshInterval time.Duration) *CachingKeySetProvider {
	_ = "STUB: not implemented"
	return nil
}

func (c *CachingKeySetProvider) GetKeySet(ctx context.Context) (*jose.JSONWebKeySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// refresh key set. if there is a failure, log and return the old set if
// available.

func DiscoverKeySetURI(ctx context.Context, configURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func FetchKeySet(ctx context.Context, jwksURI string) (*jose.JSONWebKeySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func tryRead(r io.Reader) string { _ = "STUB: not implemented"; return "" }
