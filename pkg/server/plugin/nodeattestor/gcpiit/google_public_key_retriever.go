package gcpiit

import (
	"context"
	"sync"
	"time"

	"github.com/go-jose/go-jose/v4"
)

type googlePublicKeyRetriever struct {
	url    string
	expiry time.Time

	mtx  sync.Mutex
	jwks *jose.JSONWebKeySet
}

func newGooglePublicKeyRetriever(url string) *googlePublicKeyRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (r *googlePublicKeyRetriever) retrieveJWKS(ctx context.Context) (*jose.JSONWebKeySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *googlePublicKeyRetriever) downloadJWKS(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
