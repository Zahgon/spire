package manager

import (
	"context"
	"crypto"
	"crypto/x509"
	"sync"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/common/coretypes/x509certificate"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/pkg/server/ca"
	"github.com/spiffe/spire/pkg/server/catalog"
	"github.com/spiffe/spire/pkg/server/credtemplate"
	"github.com/spiffe/spire/pkg/server/credvalidator"
	"github.com/spiffe/spire/pkg/server/datastore"
	"github.com/spiffe/spire/pkg/server/plugin/keymanager"
	"github.com/spiffe/spire/pkg/server/plugin/notifier"
	"github.com/spiffe/spire/proto/spire/common"
)

const (
	publishJWKTimeout         = 5 * time.Second
	safetyThresholdBundle     = 24 * time.Hour
	safetyThresholdCAJournals = time.Hour * 24 * 14 // Two weeks

	thirtyDays                  = 30 * 24 * time.Hour
	preparationThresholdCap     = thirtyDays
	preparationThresholdDivisor = 2

	sevenDays                  = 7 * 24 * time.Hour
	activationThresholdCap     = sevenDays
	activationThresholdDivisor = 6

	taintBackoffInterval       = 5 * time.Second
	taintBackoffMaxElapsedTime = 1 * time.Minute
)

type ManagedCA interface {
	SetX509CA(*ca.X509CA)
	SetJWTKey(*ca.JWTKey)
	SetWITKey(*ca.WITKey)
	NotifyTaintedX509Authorities([]*x509.Certificate)
}

type JwtKeyPublisher interface {
	PublishJWTKey(ctx context.Context, jwtKey *common.PublicKey) ([]*common.PublicKey, error)
}

type AuthorityManager interface {
	GetCurrentJWTKeySlot() Slot
	GetNextJWTKeySlot() Slot
	PrepareJWTKey(ctx context.Context) error
	RotateJWTKey(ctx context.Context)
	GetCurrentX509CASlot() Slot
	GetNextX509CASlot() Slot
	PrepareX509CA(ctx context.Context) error
	RotateX509CA(ctx context.Context)
	GetCurrentWITKeySlot() Slot
	GetNextWITKeySlot() Slot
	PrepareWITKey(ctx context.Context) error
	RotateWITKey(ctx context.Context)
	IsUpstreamAuthority() bool
	IsJWTSVIDsDisabled() bool
	IsWITSVIDsDisabled() bool
	PublishJWTKey(ctx context.Context, jwtKey *common.PublicKey) ([]*common.PublicKey, error)
	NotifyTaintedX509Authority(ctx context.Context, authorityID string) error
	SubscribeToLocalBundle(ctx context.Context) error
}

type Config struct {
	CredBuilder     *credtemplate.Builder
	CredValidator   *credvalidator.Validator
	CA              ManagedCA
	Catalog         catalog.Catalog
	TrustDomain     spiffeid.TrustDomain
	X509CAKeyType   keymanager.KeyType
	DisableJWTSVIDs bool
	DisableWITSVIDs bool
	JWTKeyType      keymanager.KeyType
	WITKeyType      keymanager.KeyType
	Dir             string
	Log             logrus.FieldLogger
	Metrics         telemetry.Metrics
	Clock           clock.Clock
}

type Manager struct {
	c                            Config
	caTTL                        time.Duration
	bundleUpdatedCh              chan struct{}
	taintedUpstreamAuthoritiesCh chan []*x509.Certificate
	upstreamClient               *ca.UpstreamClient
	upstreamPluginName           string

	currentX509CA *x509CASlot
	nextX509CA    *x509CASlot
	x509CAMutex   sync.RWMutex

	currentJWTKey *jwtKeySlot
	nextJWTKey    *jwtKeySlot
	jwtKeyMutex   sync.RWMutex

	currentWITKey *witKeySlot
	nextWITKey    *witKeySlot
	witKeyMutex   sync.RWMutex

	journal *Journal

	// Used to log a warning only once when the UpstreamAuthority does not support JWT-SVIDs.
	jwtUnimplementedWarnOnce sync.Once

	// Used for testing backoff, must not be set in regular code
	triggerBackOffCh chan error
}

func NewManager(ctx context.Context, c Config) (*Manager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// activate the X509CA immediately if it is set and not within
// activation time of the next X509CA.

// TODO: Activation on journal depends on dates, it will need to be
// refactored to allow to set a status, because when forcing a rotation,
// we are no longer able to depend on a date.

// activate the JWT key immediately if it is set and not within
// activation time of the next JWT key.

// activate the WIT key immediately if it is set and not within
// activation time of the next WIT key.

func (m *Manager) Close() { _ = "STUB: not implemented"; return }

func (m *Manager) NotifyTaintedX509Authority(ctx context.Context, authorityID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) IsJWTSVIDsDisabled() bool { _ = "STUB: not implemented"; return false }

func (m *Manager) IsWITSVIDsDisabled() bool { _ = "STUB: not implemented"; return false }

func (m *Manager) GetCurrentX509CASlot() Slot { _ = "STUB: not implemented"; return *new(Slot) }

func (m *Manager) GetNextX509CASlot() Slot { _ = "STUB: not implemented"; return *new(Slot) }

func (m *Manager) PrepareX509CA(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// If current is not empty, prepare the next.
// If the journal has been started, we will be preparing on next.
// This is only needed when the journal has not been started.

// Set key from new CA, to be able to get it after
// slot moved to old state

func (m *Manager) IsUpstreamAuthority() bool { _ = "STUB: not implemented"; return false }

func (m *Manager) ActivateX509CA(ctx context.Context) { _ = "STUB: not implemented"; return }

func (m *Manager) RotateX509CA(ctx context.Context) { _ = "STUB: not implemented"; return }

func (m *Manager) GetCurrentJWTKeySlot() Slot { _ = "STUB: not implemented"; return *new(Slot) }

func (m *Manager) GetNextJWTKeySlot() Slot { _ = "STUB: not implemented"; return *new(Slot) }

func (m *Manager) PrepareJWTKey(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// If current slot is not empty, use next to prepare

func (m *Manager) ActivateJWTKey(ctx context.Context) { _ = "STUB: not implemented"; return }

func (m *Manager) RotateJWTKey(ctx context.Context) { _ = "STUB: not implemented"; return }

// PublishJWTKey publishes the passed JWK to the upstream server using the configured
// UpstreamAuthority plugin, then appends to the bundle the JWKs returned by the upstream server,
// and finally it returns the updated list of JWT keys contained in the bundle.
//
// The following cases may arise when calling this function:
//
// - The UpstreamAuthority plugin doesn't implement PublishJWTKey, in which case we receive an
// Unimplemented error from the upstream server, and hence we log a one time warning about this,
// append the passed JWK to the bundle, and return the updated list of JWT keys.
//
// - The UpstreamAuthority plugin returned an error, then we return the error.
//
// - There is no UpstreamAuthority plugin configured, then assumes we are the root server and
// just appends the passed JWK to the bundle and returns the updated list of JWT keys.
func (m *Manager) PublishJWTKey(ctx context.Context, jwtKey *common.PublicKey) ([]*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// JWT Key publishing is not supported by the upstream plugin.
// Issue a one-time warning and then fall through to the
// appendBundle call below as if an upstream client was not
// configured so the JWT key gets pushed into the local bundle.

func (m *Manager) GetCurrentWITKeySlot() Slot { _ = "STUB: not implemented"; return *new(Slot) }

func (m *Manager) GetNextWITKeySlot() Slot { _ = "STUB: not implemented"; return *new(Slot) }

func (m *Manager) PrepareWITKey(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// If current slot is not empty, use next to prepare

func (m *Manager) ActivateWITKey(ctx context.Context) { _ = "STUB: not implemented"; return }

func (m *Manager) RotateWITKey(ctx context.Context) { _ = "STUB: not implemented"; return }

func (m *Manager) SubscribeToLocalBundle(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) PruneBundle(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) PruneCAJournals(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// ProcessBundleUpdates Notify any bundle update, or process tainted authorities
func (m *Manager) ProcessBundleUpdates(ctx context.Context) { _ = "STUB: not implemented"; return }

func (m *Manager) NotifyBundleLoaded(ctx context.Context) error {
	_ = "STUB: not implemented"
	// if initialization has triggered a "bundle updated" event (e.g. server CA
	// was rotated), we want to drain it now as we're about to emit the initial
	// bundle loaded event.  otherwise, plugins will get an immediate "bundle
	// updated" event right after "bundle loaded".
	return nil
}

func (m *Manager) activateJWTKey(ctx context.Context) { _ = "STUB: not implemented"; return }

func (m *Manager) activateX509CA(ctx context.Context) { _ = "STUB: not implemented"; return }

func (m *Manager) activateWITKey(ctx context.Context) { _ = "STUB: not implemented"; return }

func (m *Manager) bundleUpdated() { _ = "STUB: not implemented"; return }

func (m *Manager) dropBundleUpdated() { _ = "STUB: not implemented"; return }

func (m *Manager) notifyUpstreamAuthoritiesTainted(taintedAuthorities []*x509.Certificate) {
	_ = "STUB: not implemented"
	return
}

func (m *Manager) fetchRootCAByAuthorityID(ctx context.Context, authorityID string) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Manager) notifyTaintedAuthorities(ctx context.Context, taintedAuthorities []*x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) processTaintedUpstreamAuthorities(ctx context.Context, taintedAuthorities []*x509.Certificate) error {
	_ = "STUB: not implemented"
	// Nothing to rotate if no upstream authority is used
	return nil
}

// No tainted keys found

// Activate the prepared X.509 authority

// Now that we have rotated the intermediate, we can notify about the
// tainted authorities, so agents and downstream servers can start forcing
// the rotation of their SVIDs.

// Intermediate is safe. Notify rotator to force rotation
// of tainted X.509 SVID.

func (m *Manager) notifyBundleUpdated(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) notify(ctx context.Context, event string, advise bool, pre func(context.Context) error, do func(context.Context, notifier.Notifier) error) error {
	_ = "STUB: not implemented"
	return nil
}

// don't select on the ctx here as we can rely on the plugins to
// respond to context cancellation and return an error.

func (m *Manager) fetchRequiredBundle(ctx context.Context) (*common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Manager) fetchOptionalBundle(ctx context.Context) (*common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Manager) upstreamSignX509CA(ctx context.Context, signer crypto.Signer) (*ca.X509CA, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Manager) selfSignX509CA(ctx context.Context, signer crypto.Signer) (*ca.X509CA, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Manager) appendBundle(ctx context.Context, caChain []*x509.Certificate, jwtSigningKeys []*common.PublicKey, witSigningKeys []*common.PublicKey) (*common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Manager) shouldPrepareX509CA(taintedAuthorities []*x509.Certificate) bool {
	_ = "STUB: not implemented"
	return false
}

// MaxSVIDTTL returns the maximum SVID lifetime that can be guaranteed to not
// be cut artificially short by a scheduled rotation.
func MaxSVIDTTL() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// MaxSVIDTTLForCATTL returns the maximum SVID TTL that can be guaranteed given
// a specific CA TTL. In other words, given a CA TTL, what is the largest SVID
// TTL that is guaranteed to not be cut artificially short by a scheduled
// rotation?
func MaxSVIDTTLForCATTL(caTTL time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// MinCATTLForSVIDTTL returns the minimum CA TTL necessary to guarantee an SVID
// TTL of the provided value. In other words, given an SVID TTL, what is the
// minimum CA TTL that will guarantee that the SVIDs lifetime won't be cut
// artificially short by a scheduled rotation?
func MinCATTLForSVIDTTL(svidTTL time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

type bundleUpdater struct {
	log                         logrus.FieldLogger
	trustDomainID               string
	ds                          datastore.DataStore
	updated                     func()
	upstreamAuthoritiesTainted  func([]*x509.Certificate)
	processedTaintedAuthorities map[string]struct{}
}

func (u *bundleUpdater) SyncX509Roots(ctx context.Context, roots []*x509certificate.X509Authority) error {
	_ = "STUB: not implemented"
	return nil
}

// Collect all skIDs

// Verify if new root ca is tainted

// Taint x.509 authority, if required

// Add to the list of new tainted authorities

// Prevent to add tainted keys, since status is updated before

// Notify about tainted authorities to force the rotation of
// intermediates and update the database. This is done in a separate thread
// to prevent agents and downstream servers to start the rotation before the
// current server starts the rotation of the intermediate.

// Only tainted keys can ke revoked

// In case a stored tainted authority is not found,
// from latest bundle update, then revoke it

func (u *bundleUpdater) AppendJWTKeys(ctx context.Context, keys []*common.PublicKey) ([]*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (u *bundleUpdater) LogError(err error, msg string) { _ = "STUB: not implemented"; return }

func (u *bundleUpdater) fetchX509Authorities(ctx context.Context) (map[string]*x509certificate.X509Authority, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Bundle not found

func (u *bundleUpdater) appendBundle(ctx context.Context, bundle *common.Bundle) (*common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newJWTKey(signer crypto.Signer, expiresAt time.Time) (*ca.JWTKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newWITKey(signer crypto.Signer, expiresAt time.Time) (*ca.WITKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newKeyID() (string, error) { _ = "STUB: not implemented"; return "", nil }

func keyIDFromBytes(choices []byte) string { _ = "STUB: not implemented"; return "" }

func publicKeyFromJWTKey(jwtKey *ca.JWTKey) (*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func publicKeyFromWITKey(witKey *ca.WITKey) (*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// isX509AuthorityTainted verifies if the provided X.509 authority is tainted
func isX509AuthorityTainted(x509CA *ca.X509CA, taintedAuthorities []*x509.Certificate) bool {
	_ = "STUB: not implemented"
	return false
}

// Verify certificate chain, using tainted authority as root
