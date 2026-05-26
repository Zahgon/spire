package httpchallenge

import (
	"context"
	"net"
	"sync"

	"github.com/hashicorp/go-hclog"
	nodeattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/agent/nodeattestor/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

const (
	pluginName = "http_challenge"
)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

type configData struct {
	port           int
	advertisedPort int
	hostName       string
	agentName      string
}

type Config struct {
	HostName       string `hcl:"hostname"`
	AgentName      string `hcl:"agentname"`
	Port           int    `hcl:"port"`
	AdvertisedPort int    `hcl:"advertised_port"`
}

func (p *Plugin) buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *configData {
	_ = "STUB: not implemented"
	return nil
}

// hostname unset, autodetect hostname

// if unset, advertised port is same as hcl:"port"

type Plugin struct {
	nodeattestorv1.UnsafeNodeAttestorServer
	configv1.UnsafeConfigServer

	m sync.RWMutex
	c *configData

	log hclog.Logger

	hooks struct {
		// Controls which interface to bind to ("" in production, "localhost"
		// in tests) and acts as the default HostName value when not provided
		// via configuration.
		bindHost string
	}
}

func New() *Plugin { _ = "STUB: not implemented"; return nil }

func (p *Plugin) AidAttestation(stream nodeattestorv1.NodeAttestor_AidAttestationServer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// send the attestation data back to the agent

// receive challenge

// due to https://github.com/spiffe/spire/blob/8f9fa036e182a2fab968e03cd25a7fdb2d8c88bb/pkg/agent/plugin/nodeattestor/v1.go#L63, we must respond with a non-blank challenge response

func (p *Plugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) serveNonce(ctx context.Context, l net.Listener, agentName string, nonce string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // G118: complaint about context.Background()

// SetLogger sets this plugin's logger
func (p *Plugin) SetLogger(log hclog.Logger) { _ = "STUB: not implemented"; return }

func (p *Plugin) getConfig() (*configData, error) { _ = "STUB: not implemented"; return nil, nil }
