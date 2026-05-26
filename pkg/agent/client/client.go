package client

import (
	"context"
	"crypto"
	"crypto/x509"
	"errors"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	agentv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/agent/v1"
	bundlev1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/bundle/v1"
	entryv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/entry/v1"
	svidv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/svid/v1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/pkg/common/tlspolicy"
	"github.com/spiffe/spire/proto/spire/common"
	"google.golang.org/grpc"
)

const (
	rpcTimeout = 30 * time.Second

	// maxBundleWorkers is the maximum number of worker goroutines to use when fetching bundles.
	maxBundleWorkers = 10
)

var (
	ErrUnableToGetStream = errors.New("unable to get a stream")

	entryOutputMask = &types.EntryMask{
		SpiffeId:             true,
		Selectors:            true,
		FederatesWith:        true,
		Admin:                true,
		Downstream:           true,
		RevisionNumber:       true,
		StoreSvid:            true,
		Hint:                 true,
		AdditionalAttributes: true,
		CreatedAt:            true,
	}

	// RPCTimeoutWithCacheHit can be more aggressive with timeouts in cases where a valid SVID
	// exists in the cache but is old enough to try for a new SVID quickly. This is configurable
	// in the Experimental Config of the Agent, and can be set as low as 5 seconds
	RPCTimeoutWithCacheHit = rpcTimeout
)

func SetJWTSVIDCacheHitTimeout(d time.Duration) { _ = "STUB: not implemented"; return }

type X509SVID struct {
	CertChain []byte
	ExpiresAt int64
}

type JWTSVID struct {
	Token     string
	IssuedAt  time.Time
	ExpiresAt time.Time
}

type SyncStats struct {
	Entries SyncEntriesStats
	Bundles SyncBundlesStats
}

type SyncEntriesStats struct {
	Total   int
	Missing int
	Stale   int
	Dropped int
}

type SyncBundlesStats struct {
	Total int
}

type Client interface {
	FetchUpdates(ctx context.Context) (*Update, error)
	SyncUpdates(ctx context.Context, cachedEntries map[string]*common.RegistrationEntry, cachedBundles map[string]*common.Bundle) (SyncStats, error)
	RenewSVID(ctx context.Context, csr []byte) (*X509SVID, error)
	NewX509SVIDs(ctx context.Context, csrs map[string][]byte) (map[string]*X509SVID, error)
	NewJWTSVID(ctx context.Context, entryID string, audience []string, hasCacheHit bool) (*JWTSVID, spiffeid.ID, error)
	PostStatus(ctx context.Context, agentVersion string) error

	// Release releases any resources that were held by this Client, if any.
	Release()
}

// Config holds a client configuration
type Config struct {
	Addr        string
	Log         logrus.FieldLogger
	TrustDomain spiffeid.TrustDomain
	// KeysAndBundle is a callback that must return the keys and bundle used by the client
	// to connect via mTLS to Addr.
	KeysAndBundle func() ([]*x509.Certificate, crypto.Signer, []*x509.Certificate)

	// RotMtx is used to prevent the creation of new connections during SVID rotations
	RotMtx *sync.RWMutex

	// TLSPolicy determines the post-quantum-safe policy to apply to all TLS connections.
	TLSPolicy tlspolicy.Policy
}

type client struct {
	c           *Config
	connections *nodeConn
	m           sync.Mutex

	// dialOpts optionally sets gRPC dial options
	dialOpts []grpc.DialOption
}

// fetchBundleResult contains the result of fetching a federated bundle.
type fetchBundleResult struct {
	bundle *types.Bundle
	err    error
}

// New creates a new client struct with the configuration provided
func New(c *Config) Client { _ = "STUB: not implemented"; return *new(Client) }

func newClient(c *Config) *client { _ = "STUB: not implemented"; return nil }

func (c *client) FetchUpdates(ctx context.Context) (*Update, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get all federated trust domains

func (c *client) SyncUpdates(ctx context.Context, cachedEntries map[string]*common.RegistrationEntry, cachedBundles map[string]*common.Bundle) (SyncStats, error) {
	_ = "STUB: not implemented"
	return *new(SyncStats), nil
}

func (c *client) RenewSVID(ctx context.Context, csr []byte) (*X509SVID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *client) PostStatus(ctx context.Context, agentVersion string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) NewX509SVIDs(ctx context.Context, csrs map[string][]byte) (map[string]*X509SVID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *client) NewJWTSVID(ctx context.Context, entryID string, audience []string, hasCacheHit bool) (*JWTSVID, spiffeid.ID, error) {
	_ = "STUB: not implemented"
	return nil, *new(spiffeid.ID), nil
}

// Release the underlying connection.
func (c *client) Release() { _ = "STUB: not implemented"; return }

func (c *client) release(conn *nodeConn) { _ = "STUB: not implemented"; return }

func (c *client) newServerGRPCClient() (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *client) fetchEntries(ctx context.Context) ([]*types.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *client) syncEntries(ctx context.Context, cachedEntries map[string]*common.RegistrationEntry) (SyncEntriesStats, error) {
	_ = "STUB: not implemented"
	return *new(SyncEntriesStats), nil
}

func entryIsStale(entry *common.RegistrationEntry, revisionNumber, revisionCreatedAt int64) bool {
	_ = "STUB: not implemented"
	return false
}

// TODO: remove in SPIRE 1.14

// Verify that the CreatedAt of the entries match. If they are different, they are
// completely different entries even if the revision number is the same.
// This can happen for example if an entry is deleted and recreated with the
// same entry id.

func (c *client) streamAndSyncEntries(ctx context.Context, entryClient entryv1.EntryClient, cachedEntries map[string]*common.RegistrationEntry) (stats SyncEntriesStats, err error) {
	_ = "STUB: not implemented"
	// Build a set of all the entries to be removed. This set is initialized
	// with all entries currently known. As entries are synced down from the
	// server, they are removed from this set. If the sync is successful,
	// any entry that was not seen during sync, i.e., still remains a member
	// of this set, is removed from the cached entries.
	return *new(SyncEntriesStats), nil
}

// needFull tracks the entry IDs of entries that are either not cached, or
// that have been determined to be stale (based on revision number
// comparison)

// processEntryRevisions determines what needs to be synced down based
// on entry revisions.

// The entry is still authorized for this agent. Don't remove it.

// If entry is either not cached or is stale, record the ID so
// the full entry can be requested after syncing down all
// entry revisions.

// processServerEntries updates the cached entries

// The entry is still authorized for this agent. Don't remove it.

// Update the cached entry

// If the first response does not contain entry revisions then it contains
// the complete list of authorized entries.

// Assume that the page size is the size of the revisions in the first
// response from the server.

// Receive the rest of the entry revisions

// Presort the IDs. The server sorts the requested IDs as an optimization
// for memory and CPU efficient lookups. Even though the server will sort
// them, pre-sorting should reduce server CPU load (Go1.19+ implements
// sorting via the PDQ algorithm, which performs well on pre-sorted data).

// Request the full entries for missing or stale entries one page at a
// time using the assumed page size.

// Request up to a page full of full entries

// Receive the full entries just requested. Even though the entries
// SHOULD come back in a single response (since we matched the page
// size of the server), handle the case where the server decides to
// break them up into multiple pages.

func (c *client) fetchBundles(ctx context.Context, federatedBundles []string) ([]*types.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get bundle

// fetchFederatedBundlesConcurrently fetches federated bundles concurrently.
// This is done to improve sync times when there are many federations. This should ensure that the
// sync does not exceed rpcTimeout.
func (c *client) fetchFederatedBundlesConcurrently(ctx context.Context, bundleClient bundlev1.BundleClient, trustDomains []string, bundles []*types.Bundle) ([]*types.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Start a set of worker goroutines.

// Feed the workers.

// Wait for completion of all jobs.

// Process the results.

// fetchFederatedBundle fetches a single federated bundle from SPIRE server.
func (c *client) fetchFederatedBundle(ctx context.Context, bundleClient bundlev1.BundleClient, trustDomain string) (*types.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *client) fetchSVIDs(ctx context.Context, params []*svidv1.NewX509SVIDParams) ([]*types.X509SVID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *client) newEntryClient() (entryv1.EntryClient, *nodeConn, error) {
	_ = "STUB: not implemented"
	return *new(entryv1.EntryClient), nil, nil
}

func (c *client) newBundleClient() (bundlev1.BundleClient, *nodeConn, error) {
	_ = "STUB: not implemented"
	return *new(bundlev1.BundleClient), nil, nil
}

func (c *client) newSVIDClient() (svidv1.SVIDClient, *nodeConn, error) {
	_ = "STUB: not implemented"
	return *new(svidv1.SVIDClient), nil, nil
}

func (c *client) newAgentClient() (agentv1.AgentClient, *nodeConn, error) {
	_ = "STUB: not implemented"
	return *new(agentv1.AgentClient), nil, nil
}

func (c *client) getOrOpenConn() (*nodeConn, error) { _ = "STUB: not implemented"; return nil, nil }

type stringSet map[string]struct{}

func (ss stringSet) Add(s string) { _ = "STUB: not implemented"; return }

func (ss stringSet) Sorted() []string { _ = "STUB: not implemented"; return nil }

// withErrorFields add fields of gRPC call status in logger
func (c *client) withErrorFields(err error) logrus.FieldLogger {
	_ = "STUB: not implemented"
	return *new(logrus.FieldLogger)
}
