package uniqueid

import (
	"context"

	credentialcomposerv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/credentialcomposer/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtIn(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

type Plugin struct {
	credentialcomposerv1.UnsafeCredentialComposerServer
}

func New() *Plugin { _ = "STUB: not implemented"; return nil }

func (p *Plugin) ComposeServerX509CA(context.Context, *credentialcomposerv1.ComposeServerX509CARequest) (*credentialcomposerv1.ComposeServerX509CAResponse, error) {
	_ = "STUB: not implemented"
	// Intentionally not implemented.
	return nil, nil
}

func (p *Plugin) ComposeServerX509SVID(context.Context, *credentialcomposerv1.ComposeServerX509SVIDRequest) (*credentialcomposerv1.ComposeServerX509SVIDResponse, error) {
	_ = "STUB: not implemented"
	// Intentionally not implemented.
	return nil, nil
}

func (p *Plugin) ComposeAgentX509SVID(context.Context, *credentialcomposerv1.ComposeAgentX509SVIDRequest) (*credentialcomposerv1.ComposeAgentX509SVIDResponse, error) {
	_ = "STUB: not implemented"
	// Intentionally not implemented.
	return nil, nil
}

func (p *Plugin) ComposeWorkloadX509SVID(_ context.Context, req *credentialcomposerv1.ComposeWorkloadX509SVIDRequest) (*credentialcomposerv1.ComposeWorkloadX509SVIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No need to clone

// Add the attribute if it does not already exist. Otherwise, replace the old value.

func (p *Plugin) ComposeWorkloadJWTSVID(context.Context, *credentialcomposerv1.ComposeWorkloadJWTSVIDRequest) (*credentialcomposerv1.ComposeWorkloadJWTSVIDResponse, error) {
	_ = "STUB: not implemented"
	// Intentionally not implemented.
	return nil, nil
}

func uniqueIDAttributeTypeAndValue(id string) (*credentialcomposerv1.AttributeTypeAndValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// purely defensive.
