package spireplugin

import (
	"context"
	"sync"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/hashicorp/go-hclog"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	upstreamauthorityv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/upstreamauthority/v1"
	plugintypes "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/types"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

const (
	pluginName       = "spire"
	upstreamPollFreq = 5 * time.Second
	internalPollFreq = time.Second
)

type Configuration struct {
	ServerAddr        string             `hcl:"server_address" json:"server_address"`
	ServerPort        string             `hcl:"server_port" json:"server_port"`
	WorkloadAPISocket string             `hcl:"workload_api_socket" json:"workload_api_socket"`
	Experimental      experimentalConfig `hcl:"experimental"`
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *Configuration {
	_ = "STUB: not implemented"
	return nil
}

// TODO: add field validation

type experimentalConfig struct {
	WorkloadAPINamedPipeName string `hcl:"workload_api_named_pipe_name" json:"workload_api_named_pipe_name"`
	RequirePQKEM             bool   `hcl:"require_pq_kem" json:"require_pq_kem"`
}

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

type Plugin struct {
	upstreamauthorityv1.UnsafeUpstreamAuthorityServer
	configv1.UnsafeConfigServer

	clk clock.Clock
	log hclog.Logger

	mtx         sync.RWMutex
	trustDomain spiffeid.TrustDomain
	config      *Configuration

	// Server's client. It uses an X509 source to fetch SVIDs from Workload API
	serverClient *serverClient

	pollMtx                sync.Mutex
	stopPolling            context.CancelFunc
	pollDone               <-chan struct{}
	currentPollSubscribers uint64

	bundleMtx     sync.RWMutex
	bundleVersion uint64
	currentBundle *plugintypes.Bundle
	// bundleUpdated is closed and replaced each time the bundle changes,
	// allowing subscribers to be notified immediately rather than waiting for
	// the internalPollFreq ticker.
	bundleUpdated chan struct{}
}

func New() *Plugin { _ = "STUB: not implemented"; return nil }

func (p *Plugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Swap Running Config

// Create spire-server client

func (p *Plugin) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) SetLogger(log hclog.Logger) { _ = "STUB: not implemented"; return }

func (p *Plugin) MintX509CAAndSubscribe(request *upstreamauthorityv1.MintX509CARequest, stream upstreamauthorityv1.UpstreamAuthority_MintX509CAAndSubscribeServer) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: downstream RPC is not returning authority metadata, like tainted bit
// avoid using it for now in favor of a call to get bundle RPC

// Set X509 Authorities

func (p *Plugin) SubscribeToLocalBundle(req *upstreamauthorityv1.SubscribeToLocalBundleRequest, stream upstreamauthorityv1.UpstreamAuthority_SubscribeToLocalBundleServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Capture the notification channel before reading the bundle.
// If the bundle is updated after getBundleUpdateCh() returns but
// before the select below, the old channel will already be closed
// and the select will fire immediately rather than sleeping until
// the next internalPollFreq tick.

func (p *Plugin) PublishJWTKeyAndSubscribe(req *upstreamauthorityv1.PublishJWTKeyRequest, stream upstreamauthorityv1.UpstreamAuthority_PublishJWTKeyAndSubscribeServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Publish JWT authority

// Set JWT authority

func (p *Plugin) pollBundleUpdates(ctx context.Context) { _ = "STUB: not implemented"; return }

// setBundleIfVersionMatches updates currentBundle only when bundleVersion
// still equals expectedVersion. This prevents a fetch that started before a
// local mutation (setBundleJWTAuthorities / setBundleX509Authorities) from
// overwriting the newer local state. bundleVersion is intentionally not
// incremented here; it is only incremented by the local-mutation helpers so
// that it remains a reliable guard against concurrent upstream fetches.
func (p *Plugin) setBundleIfVersionMatches(b *types.Bundle, expectedVersion uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) getBundle() *plugintypes.Bundle { _ = "STUB: not implemented"; return nil }

// getBundleUpdateCh returns the current bundle update notification channel.
// When the bundle is updated, this channel is closed and a new one is created.
func (p *Plugin) getBundleUpdateCh() <-chan struct{} { _ = "STUB: not implemented"; return nil }

// signalBundleUpdate closes the current bundleUpdated channel and replaces it
// with a new one, waking any goroutines waiting on the previous channel.
// Must be called with bundleMtx write lock held.
func (p *Plugin) signalBundleUpdate() { _ = "STUB: not implemented"; return }

func (p *Plugin) setBundleJWTAuthorities(keys []*plugintypes.JWTKey) {
	_ = "STUB: not implemented"
	return
}

func (p *Plugin) setBundleX509Authorities(rootCAs []*plugintypes.X509Certificate) {
	_ = "STUB: not implemented"
	return
}

func (p *Plugin) getBundleVersion() uint64 { _ = "STUB: not implemented"; return 0 }

func (p *Plugin) subscribeToPolling(streamCtx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) unsubscribeToPolling() { _ = "STUB: not implemented"; return }

// Wait for the polling goroutine to fully exit so its deferred
// release() call — which nils out bundleClient — cannot race with
// a future startPolling call.

func (p *Plugin) startPolling(streamCtx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func areRootsEqual(a, b []*plugintypes.X509Certificate) bool {
	_ = "STUB: not implemented"
	return false
}

func arePublicKeysEqual(a, b []*plugintypes.JWTKey) bool { _ = "STUB: not implemented"; return false }
