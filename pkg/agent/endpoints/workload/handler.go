package workload

import (
	"context"
	"crypto/x509"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/bundle/spiffebundle"
	"github.com/spiffe/go-spiffe/v2/proto/spiffe/workload"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/agent/client"
	"github.com/spiffe/spire/pkg/agent/manager/cache"
	"github.com/spiffe/spire/pkg/common/jwtsvid"
	"github.com/spiffe/spire/proto/spire/common"
	"google.golang.org/protobuf/types/known/structpb"
)

type Manager interface {
	SubscribeToCacheChanges(ctx context.Context, key cache.Selectors) (cache.Subscriber, error)
	MatchingRegistrationEntries(selectors []*common.Selector) []*common.RegistrationEntry
	FetchJWTSVID(ctx context.Context, entry *common.RegistrationEntry, audience []string) (*client.JWTSVID, error)
	FetchWorkloadUpdate([]*common.Selector) *cache.WorkloadUpdate
}

type Attestor interface {
	Attest(ctx context.Context) ([]*common.Selector, error)
}

type Config struct {
	Manager                       Manager
	Attestor                      Attestor
	AllowUnauthenticatedVerifiers bool
	AllowedForeignJWTClaims       map[string]struct{}
	TrustDomain                   spiffeid.TrustDomain
}

// Handler implements the Workload API interface
type Handler struct {
	workload.UnsafeSpiffeWorkloadAPIServer
	c Config
}

func New(c Config) *Handler { _ = "STUB: not implemented"; return nil }

// FetchJWTSVID processes request for a JWT-SVID. In case of multiple fetched SVIDs with same hint, the SVID that has the oldest
// associated entry will be returned.
func (h *Handler) FetchJWTSVID(ctx context.Context, req *workload.JWTSVIDRequest) (resp *workload.JWTSVIDResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FetchJWTBundles processes request for JWT bundles
func (h *Handler) FetchJWTBundles(_ *workload.JWTBundlesRequest, stream workload.SpiffeWorkloadAPI_FetchJWTBundlesServer) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateJWTSVID processes request for JWT-SVID validation
func (h *Handler) ValidateJWTSVID(ctx context.Context, req *workload.ValidateJWTSVIDRequest) (*workload.ValidateJWTSVIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RFC 7519 structures `aud` as an array of StringOrURIs but has a special
// case where it MAY be specified as a single StringOrURI if there is only
// one audience. We have traditionally always returned it as an array but
// the JWT library we use now returns a single string when there is only
// one. To maintain backcompat, convert a single string value for the
// audience to a list.

// FetchX509SVID processes request for a x509 SVID. In case of multiple fetched SVIDs with same hint, the SVID that has the oldest
// associated entry will be returned.
func (h *Handler) FetchX509SVID(_ *workload.X509SVIDRequest, stream workload.SpiffeWorkloadAPI_FetchX509SVIDServer) error {
	_ = "STUB: not implemented"
	return nil
}

// The agent health check currently exercises the Workload API.
// Only log if it is not the agent itself.

// FetchX509Bundles processes request for x509 bundles
func (h *Handler) FetchX509Bundles(_ *workload.X509BundlesRequest, stream workload.SpiffeWorkloadAPI_FetchX509BundlesServer) error {
	_ = "STUB: not implemented"
	return nil
}

// The agent health check currently exercises the Workload API.
// Only log if it is not the agent itself.

func (h *Handler) fetchJWTSVID(ctx context.Context, log logrus.FieldLogger, entry *common.RegistrationEntry, audience []string, start time.Time) (*workload.JWTSVID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sendX509BundlesResponse(update *cache.WorkloadUpdate, stream workload.SpiffeWorkloadAPI_FetchX509BundlesServer, log logrus.FieldLogger, allowUnauthenticatedVerifiers bool, previousResponse *workload.X509BundlesResponse, quietLogging bool, start time.Time) (*workload.X509BundlesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func composeX509BundlesResponse(update *cache.WorkloadUpdate) (*workload.X509BundlesResponse, error) {
	_ = "STUB: not implemented"
	return nil,

		// This should be purely defensive since the cache should always supply
		// a bundle.
		nil
}

func sendX509SVIDResponse(update *cache.WorkloadUpdate, stream workload.SpiffeWorkloadAPI_FetchX509SVIDServer, log logrus.FieldLogger, quietLogging bool, start time.Time) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// log and emit telemetry on each SVID
// a response has already been sent so nothing is
// blocked on this logic

func composeX509SVIDResponse(update *cache.WorkloadUpdate) (*workload.X509SVIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sendJWTBundlesResponse(update *cache.WorkloadUpdate, stream workload.SpiffeWorkloadAPI_FetchJWTBundlesServer, log logrus.FieldLogger, allowUnauthenticatedVerifiers bool, previousResponse *workload.JWTBundlesResponse, start time.Time) (*workload.JWTBundlesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func composeJWTBundlesResponse(update *cache.WorkloadUpdate) (*workload.JWTBundlesResponse, error) {
	_ = "STUB: not implemented"
	return nil,

		// This should be purely defensive since the cache should always supply
		// a bundle.
		nil
}

// isAgent returns true if the caller PID from the provided context is the
// agent's process ID.
func isAgent(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func loggerWithContextInfo(ctx context.Context, log logrus.FieldLogger, start time.Time, err error) logrus.FieldLogger {
	_ = "STUB: not implemented"
	return *

	// If the context is done, include a cancellation cause if present
	// and the request duration (since it likely means the client has gone
	// away so it's useful to know how long it waited before giving up).
	new(logrus.FieldLogger)
}

//nolint:errorlint

// If no specific error was given, use the context error;
// it will be context.Canceled or context.DeadlineExceeded.

func logNoIdentityIssued(ctx context.Context, log logrus.FieldLogger, start time.Time) {
	_ = "STUB: not implemented"
	return
}

func (h *Handler) getWorkloadBundles(selectors []*common.Selector) (bundles []*spiffebundle.Bundle) {
	_ = "STUB: not implemented"
	return nil
}

func marshalBundle(certs []*x509.Certificate) []byte { _ = "STUB: not implemented"; return nil }

func keyStoreFromBundles(bundles []*spiffebundle.Bundle) (jwtsvid.KeyStore, error) {
	_ = "STUB: not implemented"
	return *new(jwtsvid.KeyStore), nil
}

func structFromValues(values map[string]any) (*structpb.Struct, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isClaimAllowed(claim string, allowedClaims map[string]struct{}) bool {
	_ = "STUB: not implemented"
	return false
}

func filterIdentities(identities []cache.Identity, log logrus.FieldLogger) []cache.Identity {
	_ = "STUB: not implemented"
	return nil
}

func filterRegistrations(entries []*common.RegistrationEntry, log logrus.FieldLogger) []*common.RegistrationEntry {
	_ = "STUB: not implemented"
	return nil
}

func getEntriesToRemove(entries []*common.RegistrationEntry, log logrus.FieldLogger) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func hintTieBreaking(entryA *common.RegistrationEntry, entryB *common.RegistrationEntry) (maintain *common.RegistrationEntry, remove *common.RegistrationEntry) {
	_ = "STUB: not implemented"
	return nil, nil
}
