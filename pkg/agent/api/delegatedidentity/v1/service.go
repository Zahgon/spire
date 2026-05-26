package delegatedidentity

import (
	"context"
	"crypto/x509"

	"github.com/sirupsen/logrus"
	delegatedidentityv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/agent/delegatedidentity/v1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	workloadattestor "github.com/spiffe/spire/pkg/agent/attestor/workload"
	"github.com/spiffe/spire/pkg/agent/manager"
	"github.com/spiffe/spire/pkg/agent/manager/cache"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/proto/spire/common"
	"google.golang.org/grpc"
)

// RegisterService registers the delegated identity service on the provided server
func RegisterService(s *grpc.Server, service *Service) { _ = "STUB: not implemented"; return }

type attestor interface {
	Attest(ctx context.Context) ([]*common.Selector, error)
}

type Config struct {
	Log                 logrus.FieldLogger
	Metrics             telemetry.Metrics
	Manager             manager.Manager
	Attestor            workloadattestor.Attestor
	AuthorizedDelegates []string
}

func New(config Config) *Service { _ = "STUB: not implemented"; return nil }

// Service implements the delegated identity server
type Service struct {
	delegatedidentityv1.UnsafeDelegatedIdentityServer

	manager                  manager.Manager
	peerAttestor             attestor
	delegateWorkloadAttestor workloadattestor.Attestor
	metrics                  telemetry.Metrics

	// SPIFFE IDs of delegates that are authorized to use this API
	authorizedDelegates map[string]bool
}

// isCallerAuthorized attests the caller based on the authorized delegates map.
func (s *Service) isCallerAuthorized(ctx context.Context, log logrus.FieldLogger, cachedSelectors []*common.Selector) ([]*common.Selector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// caller has identity associated with but none is authorized

func (s *Service) constructValidSelectorsFromReq(ctx context.Context, log logrus.FieldLogger, reqPid int32, reqSelectors []*types.Selector) ([]*common.Selector, error) {
	_ = "STUB: not implemented"
	// If you set
	// - both pid and selector args
	// - neither of them
	// it's an error
	// NOTE: the default value of int32 is naturally 0 in protobuf, which is also a valid PID.
	// However, we will still treat that as an error, as we do not expect to ever be asked to attest
	// pid 0.
	return nil, nil
}

// Delegate authorized, if the delegate gives us selectors, we treat them as attested.

// Delegate authorized, use PID the delegate gave us to try and attest on-behalf-of

// Attempt to attest and authorize the delegate, and then
//
// - Take a pre-atttested set of selectors from the delegate
// - the PID the delegate gave us and attempt to attest that into a set of selectors
//
// and provide a SVID subscription for those selectors.
//
// NOTE:
// - If supplying a PID, the trusted delegate is responsible for ensuring the PID is valid and not recycled,
// from initiation of this call until the termination of the response stream, and if it is,
// must discard any stream contents provided by this call as invalid.
// - If supplying selectors, the trusted delegate is responsible for ensuring they are correct.
func (s *Service) SubscribeToX509SVIDs(req *delegatedidentityv1.SubscribeToX509SVIDsRequest, stream delegatedidentityv1.DelegatedIdentity_SubscribeToX509SVIDsServer) error {
	_ = "STUB: not implemented"
	return nil
}

// emit latency metric for first update containing an SVID.

func sendX509SVIDResponse(update *cache.WorkloadUpdate, stream delegatedidentityv1.DelegatedIdentity_SubscribeToX509SVIDsServer, log logrus.FieldLogger) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// log details on each SVID
// a response has already been sent so nothing is
// blocked on this logic

// Ideally ID Proto parsing should succeed, but if it fails,
// ignore the error and still log with empty spiffe_id.

func composeX509SVIDBySelectors(update *cache.WorkloadUpdate) (*delegatedidentityv1.SubscribeToX509SVIDsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sort list to give a stable response instead of one dependent on the map
// iteration order above.

// Do not send admin nor downstream SVIDs to the caller

// check if SVIDs exist for the identity

func (s *Service) SubscribeToX509Bundles(_ *delegatedidentityv1.SubscribeToX509BundlesRequest, stream delegatedidentityv1.DelegatedIdentity_SubscribeToX509BundlesServer) error {
	_ = "STUB: not implemented"
	return nil
}

// send initial update....

// Attempt to attest and authorize the delegate, and then
//
// - Take a pre-atttested set of selectors from the delegate
// - the PID the delegate gave us and attempt to attest that into a set of selectors
//
// and provide a JWT SVID for those selectors.
//
// NOTE:
// - If supplying a PID, the trusted delegate is responsible for ensuring the PID is valid and not recycled,
// from initiation of this call until the response is returned, and if it is,
// must discard any response provided by this call as invalid.
// - If supplying selectors, the trusted delegate is responsible for ensuring they are correct.
func (s *Service) FetchJWTSVIDs(ctx context.Context, req *delegatedidentityv1.FetchJWTSVIDsRequest) (resp *delegatedidentityv1.FetchJWTSVIDsResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Do not send admin nor downstream SVIDs to the caller

func (s *Service) SubscribeToJWTBundles(_ *delegatedidentityv1.SubscribeToJWTBundlesRequest, stream delegatedidentityv1.DelegatedIdentity_SubscribeToJWTBundlesServer) error {
	_ = "STUB: not implemented"
	return nil
}

// send initial update....

func marshalBundle(certs []*x509.Certificate) []byte { _ = "STUB: not implemented"; return nil }

func logNoIdentityIssued(ctx context.Context, log logrus.FieldLogger) {
	_ = "STUB: not implemented"
	return
}
