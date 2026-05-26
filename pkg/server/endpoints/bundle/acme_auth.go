package bundle

import (
	"context"
	"crypto"
	"crypto/tls"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire/pkg/server/endpoints/bundle/internal/autocert"
	"github.com/spiffe/spire/pkg/server/plugin/keymanager"
)

const (
	acmeKeyPrefix = "bundle-acme-"
)

// ACMECache implements a cache for the autocert manager. It makes some
// simplifying assumptions based on our usage for the bundle endpoint. Namely,
// it assumes there is going to be a single cache entry, since we only support
// a single domain. It assumes PEM encoded blocks of data and strips out the
// private key to be stored in the key manager instead of on disk with the rest
// of the data.
type ACMEConfig struct {
	// DirectoryURL is the ACME directory URL
	DirectoryURL string

	// DomainName is the domain name of the certificate to obtain.
	DomainName string

	// CacheDir is the directory on disk where we cache certificates.
	CacheDir string

	// Email is the email address of the account to register with ACME
	Email string

	// ToSAccepted is whether the terms of service have been accepted. If
	// not true, and the provider requires acceptance, then certificate
	// retrieval will fail.
	ToSAccepted bool
}

func ACMEAuth(log logrus.FieldLogger, km keymanager.KeyManager, config ACMEConfig) ServerAuth {
	_ = "STUB: not implemented"
	// The acme client already defaulting to Let's Encrypt if the URL is unset,
	// but we want it populated for logging purposes.
	return *new(ServerAuth)
}

type acmeAuth struct {
	m *autocert.Manager
}

func (a *acmeAuth) GetTLSConfig() *tls.Config { _ = "STUB: not implemented"; return nil }

type acmeKeyStore struct {
	log logrus.FieldLogger
	km  keymanager.KeyManager
}

func (ks *acmeKeyStore) GetPrivateKey(ctx context.Context, id string) (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}

func (ks *acmeKeyStore) NewPrivateKey(ctx context.Context, id string, keyType autocert.KeyType) (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}
