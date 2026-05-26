package fakedatastore

import (
	"context"
	"testing"
	"time"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/pkg/server/datastore"
	"github.com/spiffe/spire/proto/spire/common"
)

var (
	ctx = context.Background()
)

type DataStore struct {
	ds   datastore.DataStore
	errs []error
}

var _ datastore.DataStore = (*DataStore)(nil)

func New(tb testing.TB) *DataStore { _ = "STUB: not implemented"; return nil }

func (s *DataStore) CreateBundle(ctx context.Context, bundle *common.Bundle) (*common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) UpdateBundle(ctx context.Context, bundle *common.Bundle, mask *common.BundleMask) (*common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) SetBundle(ctx context.Context, bundle *common.Bundle) (*common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) AppendBundle(ctx context.Context, bundle *common.Bundle) (*common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) CountBundles(ctx context.Context) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *DataStore) DeleteBundle(ctx context.Context, trustDomain string, mode datastore.DeleteMode) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *DataStore) FetchBundle(ctx context.Context, trustDomain string) (*common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) ListBundles(ctx context.Context, req *datastore.ListBundlesRequest) (*datastore.ListBundlesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sorting helps unit-tests have deterministic assertions.

func (s *DataStore) PruneBundle(ctx context.Context, trustDomainID string, expiresBefore time.Time) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *DataStore) CountAttestedNodes(ctx context.Context, req *datastore.CountAttestedNodesRequest) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *DataStore) CreateAttestedNode(ctx context.Context, node *common.AttestedNode) (*common.AttestedNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) FetchAttestedNode(ctx context.Context, spiffeID string) (*common.AttestedNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) ListAttestedNodes(ctx context.Context, req *datastore.ListAttestedNodesRequest) (*datastore.ListAttestedNodesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) UpdateAttestedNode(ctx context.Context, node *common.AttestedNode, mask *common.AttestedNodeMask) (*common.AttestedNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) DeleteAttestedNode(ctx context.Context, spiffeID string) (*common.AttestedNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) PruneAttestedExpiredNodes(ctx context.Context, expiredBefore time.Time, includeNonReattestable bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *DataStore) ListAttestedNodeEvents(ctx context.Context, req *datastore.ListAttestedNodeEventsRequest) (*datastore.ListAttestedNodeEventsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) PruneAttestedNodeEvents(ctx context.Context, olderThan time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *DataStore) CreateAttestedNodeEventForTesting(ctx context.Context, event *datastore.AttestedNodeEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *DataStore) DeleteAttestedNodeEventForTesting(ctx context.Context, eventID uint) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *DataStore) FetchAttestedNodeEvent(ctx context.Context, eventID uint) (*datastore.AttestedNodeEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) TaintX509CA(ctx context.Context, trustDomainID string, subjectKeyIDToTaint string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *DataStore) RevokeX509CA(ctx context.Context, trustDomainID string, subjectKeyIDToRevoke string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *DataStore) TaintJWTKey(ctx context.Context, trustDomainID string, authorityID string) (*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) RevokeJWTKey(ctx context.Context, trustDomainID string, authorityID string) (*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) SetNodeSelectors(ctx context.Context, spiffeID string, selectors []*common.Selector) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *DataStore) ListNodeSelectors(ctx context.Context, req *datastore.ListNodeSelectorsRequest) (*datastore.ListNodeSelectorsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) GetNodeSelectors(ctx context.Context, spiffeID string, dataConsistency datastore.DataConsistency) ([]*common.Selector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sorting helps unit-tests have deterministic assertions.

func (s *DataStore) CountRegistrationEntries(ctx context.Context, req *datastore.CountRegistrationEntriesRequest) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *DataStore) CreateRegistrationEntry(ctx context.Context, entry *common.RegistrationEntry) (*common.RegistrationEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) CreateOrReturnRegistrationEntry(ctx context.Context, entry *common.RegistrationEntry) (*common.RegistrationEntry, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (s *DataStore) FetchRegistrationEntry(ctx context.Context, entryID string) (*common.RegistrationEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) FetchRegistrationEntries(ctx context.Context, entryIDs []string) (map[string]*common.RegistrationEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) ListRegistrationEntries(ctx context.Context, req *datastore.ListRegistrationEntriesRequest) (*datastore.ListRegistrationEntriesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sorting helps unit-tests have deterministic assertions.

func (s *DataStore) UpdateRegistrationEntry(ctx context.Context, entry *common.RegistrationEntry, mask *common.RegistrationEntryMask) (*common.RegistrationEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) DeleteRegistrationEntry(ctx context.Context, entryID string) (*common.RegistrationEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) PruneRegistrationEntries(ctx context.Context, expiresBefore time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *DataStore) ListRegistrationEntryEvents(ctx context.Context, req *datastore.ListRegistrationEntryEventsRequest) (*datastore.ListRegistrationEntryEventsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) PruneRegistrationEntryEvents(ctx context.Context, olderThan time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *DataStore) CreateRegistrationEntryEventForTesting(ctx context.Context, event *datastore.RegistrationEntryEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *DataStore) DeleteRegistrationEntryEventForTesting(ctx context.Context, eventID uint) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *DataStore) FetchRegistrationEntryEvent(ctx context.Context, eventID uint) (*datastore.RegistrationEntryEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) CreateJoinToken(ctx context.Context, token *datastore.JoinToken) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *DataStore) FetchJoinToken(ctx context.Context, token string) (*datastore.JoinToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) DeleteJoinToken(ctx context.Context, token string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *DataStore) PruneJoinTokens(ctx context.Context, expiresBefore time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *DataStore) CreateFederationRelationship(c context.Context, fr *datastore.FederationRelationship) (*datastore.FederationRelationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) DeleteFederationRelationship(c context.Context, trustDomain spiffeid.TrustDomain) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *DataStore) FetchFederationRelationship(c context.Context, trustDomain spiffeid.TrustDomain) (*datastore.FederationRelationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) ListFederationRelationships(ctx context.Context, req *datastore.ListFederationRelationshipsRequest) (*datastore.ListFederationRelationshipsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) UpdateFederationRelationship(ctx context.Context, fr *datastore.FederationRelationship, mask *types.FederationRelationshipMask) (*datastore.FederationRelationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) FetchCAJournal(ctx context.Context, activeX509AuthorityID string) (*datastore.CAJournal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) ListCAJournalsForTesting(ctx context.Context) ([]*datastore.CAJournal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) SetCAJournal(ctx context.Context, caJournal *datastore.CAJournal) (*datastore.CAJournal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DataStore) PruneCAJournals(ctx context.Context, allCAsExpireBefore int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *DataStore) SetNextError(err error) { _ = "STUB: not implemented"; return }

func (s *DataStore) AppendNextError(err error) { _ = "STUB: not implemented"; return }

func (s *DataStore) getNextError() error { _ = "STUB: not implemented"; return nil }
