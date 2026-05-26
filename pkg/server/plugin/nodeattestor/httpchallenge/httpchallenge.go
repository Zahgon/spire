package httpchallenge

import (
	"context"
	"net/http"
	"regexp"
	"sync"

	"github.com/hashicorp/go-hclog"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	nodeattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/nodeattestor/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
	nodeattestorbase "github.com/spiffe/spire/pkg/server/plugin/nodeattestor/base"
)

const (
	pluginName = "http_challenge"
)

var (
	agentNamePattern = regexp.MustCompile("^[a-zA-z]+[a-zA-Z0-9-]$")
)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func disableHTTPRedirects(req *http.Request, via []*http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func BuiltInTesting(client *http.Client, forceNonce string) catalog.BuiltIn {
	_ = "STUB: not implemented"
	return *new(catalog.BuiltIn)
}

func builtin(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

type configuration struct {
	trustDomain       spiffeid.TrustDomain
	requiredPort      *int
	allowNonRootPorts bool
	dnsPatterns       []*regexp.Regexp
	tofu              bool
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *configuration {
	_ = "STUB: not implemented"
	return nil
}

// User has explicitly asked for a required port that is untrusted

// User has just chosen the defaults, any port is allowed

// User explicitly set AllowNonRootPorts to true and no required port specified

type Config struct {
	AllowedDNSPatterns []string `hcl:"allowed_dns_patterns"`
	RequiredPort       *int     `hcl:"required_port"`
	AllowNonRootPorts  *bool    `hcl:"allow_non_root_ports"`
	TOFU               *bool    `hcl:"tofu"`
}

type Plugin struct {
	nodeattestorbase.Base
	nodeattestorv1.UnsafeNodeAttestorServer
	configv1.UnsafeConfigServer

	m      sync.Mutex
	config *configuration

	log hclog.Logger

	client     *http.Client
	forceNonce string
}

func New() *Plugin { _ = "STUB: not implemented"; return nil }

func (p *Plugin) Attest(stream nodeattestorv1.NodeAttestor_AttestServer) error {
	_ = "STUB: not implemented"
	return nil
}

// receive the response. We don't really care what it is but the plugin system requires it.

func (p *Plugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetLogger sets this plugin's logger
func (p *Plugin) SetLogger(log hclog.Logger) { _ = "STUB: not implemented"; return }

func (p *Plugin) getConfig() (*configuration, error) { _ = "STUB: not implemented"; return nil, nil }

func buildSelectorValues(hostName string) []string { _ = "STUB: not implemented"; return nil }

func validateAgentName(agentName string) error { _ = "STUB: not implemented"; return nil }

func validateHostName(hostName string, dnsPatterns []*regexp.Regexp) error {
	_ = "STUB: not implemented"
	return nil
}
