package sshpop

import (
	"context"
	"sync"

	nodeattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/agent/nodeattestor/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/plugin/sshpop"
)

type Plugin struct {
	nodeattestorv1.UnsafeNodeAttestorServer
	configv1.UnsafeConfigServer

	mu        sync.RWMutex
	sshclient *sshpop.Client
}

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func New() *Plugin { _ = "STUB: not implemented"; return nil }

func (p *Plugin) AidAttestation(stream nodeattestorv1.NodeAttestor_AidAttestationServer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Configure configures the Plugin.
func (p *Plugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
