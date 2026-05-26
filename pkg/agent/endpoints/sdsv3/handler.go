package sdsv3

import (
	"context"

	discovery_v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	secret_v3 "github.com/envoyproxy/go-control-plane/envoy/service/secret/v3"
	"github.com/spiffe/go-spiffe/v2/bundle/spiffebundle"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/agent/manager/cache"
	"github.com/spiffe/spire/proto/spire/common"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	disableSPIFFECertValidationKey = "disable_spiffe_cert_validation"
)

type Attestor interface {
	Attest(ctx context.Context) ([]*common.Selector, error)
}

type Manager interface {
	SubscribeToCacheChanges(ctx context.Context, key cache.Selectors) (cache.Subscriber, error)
	FetchWorkloadUpdate(selectors []*common.Selector) *cache.WorkloadUpdate
}

type Config struct {
	Attestor                    Attestor
	Manager                     Manager
	DefaultAllBundlesName       string
	DefaultBundleName           string
	DefaultSVIDName             string
	DisableSPIFFECertValidation bool
}

type Handler struct {
	c Config

	hooks struct {
		// test hook used to synchronize receipt of a stream request
		received chan struct{}
	}
}

func New(config Config) *Handler { _ = "STUB: not implemented"; return nil }

func (h *Handler) StreamSecrets(stream secret_v3.SecretDiscoveryService_StreamSecretsServer) error {
	_ = "STUB: not implemented"
	return nil
}

// If there's error detail, always log it

// If we've previously sent a nonce, this must be a reply

// The nonce should match the last sent nonce, otherwise
// it's stale and the request should be ignored.

// The caller has failed to apply the last update.
// A NACK might also contain an update to the resource hint, so we need to continue processing.

// If the current request does not contain node information, use the information from a previous request (if any)

// We need to send updates if the requested resource list has changed
// either explicitly, or implicitly because this is the first request.

// save request so that all future workload updates lead to SDS updates for the last request

// Workload update has not been received yet, defer sending updates until then

// Nothing has been requested yet.

// remember the last nonce

// Remember Node info if it exists

func subListChanged(oldSubs []string, newSubs []string) (b bool) {
	_ = "STUB: not implemented"
	return false
}

func (h *Handler) DeltaSecrets(secret_v3.SecretDiscoveryService_DeltaSecretsServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *Handler) FetchSecrets(ctx context.Context, req *discovery_v3.DiscoveryRequest) (*discovery_v3.DiscoveryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *Handler) buildResponse(versionInfo string, req *discovery_v3.DiscoveryRequest, upd *cache.WorkloadUpdate) (resp *discovery_v3.DiscoveryResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// provide a nonce for streaming requests

// build a convenient set of names for lookups

// TODO: verify the type url

func (h *Handler) triggerReceivedHook() { _ = "STUB: not implemented"; return }

type validationContextBuilder interface {
	buildOne(resourceName, trustDomainID string) (*anypb.Any, error)
	buildAll(resourceName string) (*anypb.Any, error)
}

func (h *Handler) getValidationContextBuilder(req *discovery_v3.DiscoveryRequest, upd *cache.WorkloadUpdate) (validationContextBuilder, error) {
	_ = "STUB: not implemented"
	return *new(validationContextBuilder), nil
}

type rootCABuilder struct {
	bundles map[string]*spiffebundle.Bundle
}

func newRootCABuilder(bundle *spiffebundle.Bundle, federatedBundles map[spiffeid.TrustDomain]*spiffebundle.Bundle) validationContextBuilder {
	_ = "STUB: not implemented"
	return *new(validationContextBuilder)
}

// Only include tdBundle if it is not nil, which shouldn't ever be the case. This is purely defensive.

func (b *rootCABuilder) buildOne(resourceName, trustDomain string) (*anypb.Any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *rootCABuilder) buildAll(string) (*anypb.Any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type spiffeBuilder struct {
	bundles map[spiffeid.TrustDomain]*spiffebundle.Bundle
}

func newSpiffeBuilder(tdBundle *spiffebundle.Bundle, federatedBundles map[spiffeid.TrustDomain]*spiffebundle.Bundle) (validationContextBuilder, error) {
	_ = "STUB: not implemented"
	return *new(validationContextBuilder), nil
}

// Only include tdBundle if it is not nil, which shouldn't ever be the case. This is purely defensive.

// Add all federated bundles

func (b *spiffeBuilder) buildOne(resourceName, trustDomainID string) (*anypb.Any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *spiffeBuilder) buildAll(resourceName string) (*anypb.Any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create SPIFFE validator config

// bundle := bundles[td]

// // Order by trustdomain name to return in consistent order

func supportsSPIFFEAuthExtension(req *discovery_v3.DiscoveryRequest) bool {
	_ = "STUB: not implemented"
	return false
}

// Support as default except

func (h *Handler) isSPIFFECertValidationDisabled(req *discovery_v3.DiscoveryRequest) bool {
	_ = "STUB: not implemented"
	return false
}

// error means that field have some unexpected value
// so it would be safer to assume that key doesn't exist in envoy node metadata

func parseBool(v *structpb.Value) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func buildTLSCertificate(identity cache.Identity, defaultSVIDName string) (*anypb.Any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func nextNonce() (string, error) { _ = "STUB: not implemented"; return "", nil }

func sortedNames(names map[string]bool) []string { _ = "STUB: not implemented"; return nil }
