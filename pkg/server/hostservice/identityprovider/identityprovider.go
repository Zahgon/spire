package identityprovider

import (
	"context"
	"crypto"
	"crypto/x509"
	"sync"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	identityproviderv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/hostservice/server/identityprovider/v1"
	"github.com/spiffe/spire/pkg/server/datastore"
)

type X509Identity struct {
	CertChain  []*x509.Certificate
	PrivateKey crypto.PrivateKey
}

type X509IdentityFetcher interface {
	FetchX509Identity(context.Context) (*X509Identity, error)
}

type X509IdentityFetcherFunc func(context.Context) (*X509Identity, error)

func (fn X509IdentityFetcherFunc) FetchX509Identity(ctx context.Context) (*X509Identity, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Config struct {
	// TrustDomain is the server trust domain.
	TrustDomain spiffeid.TrustDomain
}

type Deps struct {
	// DataStore is used to retrieve the latest bundle. It MUST be set.
	DataStore datastore.DataStore

	// X509IdentityFetcher is used to fetch the X509 identity. It MUST be set.
	X509IdentityFetcher X509IdentityFetcher
}

type IdentityProvider struct {
	config Config

	mu   sync.RWMutex
	deps *Deps
}

func New(config Config) *IdentityProvider { _ = "STUB: not implemented"; return nil }

func (s *IdentityProvider) SetDeps(deps Deps) error { _ = "STUB: not implemented"; return nil }

func (s *IdentityProvider) getDeps() (*Deps, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *IdentityProvider) V1() identityproviderv1.IdentityProviderServer {
	_ = "STUB: not implemented"
	return *new(identityproviderv1.IdentityProviderServer)
}

type identityProviderV1 struct {
	identityproviderv1.UnsafeIdentityProviderServer

	s *IdentityProvider
}

func (v1 *identityProviderV1) FetchX509Identity(ctx context.Context, _ *identityproviderv1.FetchX509IdentityRequest) (*identityproviderv1.FetchX509IdentityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
