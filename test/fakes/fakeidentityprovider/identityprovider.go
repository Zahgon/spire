package fakeidentityprovider

import (
	"context"
	"sync"

	identityproviderv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/hostservice/server/identityprovider/v1"
	plugintypes "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/types"
)

type IdentityProvider struct {
	identityproviderv1.UnsafeIdentityProviderServer

	mu      sync.Mutex
	bundles []*plugintypes.Bundle
}

func New() *IdentityProvider { _ = "STUB: not implemented"; return nil }

func (c *IdentityProvider) FetchX509Identity(context.Context, *identityproviderv1.FetchX509IdentityRequest) (*identityproviderv1.FetchX509IdentityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: support sending back the identity

func (c *IdentityProvider) AppendBundle(bundle *plugintypes.Bundle) {
	_ = "STUB: not implemented"
	return
}
