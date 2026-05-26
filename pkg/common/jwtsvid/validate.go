package jwtsvid

import (
	"context"
	"crypto"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
)

type KeyStore interface {
	FindPublicKey(ctx context.Context, td spiffeid.TrustDomain, kid string) (crypto.PublicKey, error)
}

type keyStore struct {
	trustDomainKeys map[spiffeid.TrustDomain]map[string]crypto.PublicKey
}

func NewKeyStore(trustDomainKeys map[spiffeid.TrustDomain]map[string]crypto.PublicKey) KeyStore {
	_ = "STUB: not implemented"
	return *new(KeyStore)
}

func (t *keyStore) FindPublicKey(_ context.Context, td spiffeid.TrustDomain, keyID string) (crypto.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey), nil
}

func ValidateToken(ctx context.Context, token string, keyStore KeyStore, audience []string) (spiffeid.ID, map[string]any, error) {
	_ = "STUB: not implemented"
	return *new(spiffeid.ID), nil, nil
}

// Obtain the key ID from the header

// Parse out the unverified claims. We need to look up the key by the trust
// domain of the SPIFFE ID. We'll verify the signature on the claims below
// when creating the generic map of claims that we return to the caller.

// Construct the trust domain id from the SPIFFE ID and look up key by ID

// Now obtain the generic claims map verified using the obtained key

// Now that the signature over the claims has been verified, validate the
// standard claims.

// Convert expected validation errors for pretty errors
