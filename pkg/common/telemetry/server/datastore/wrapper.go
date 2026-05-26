package datastore

import (
	"context"
	"time"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/pkg/server/datastore"
	"github.com/spiffe/spire/proto/spire/common"
)

// WithMetrics wraps a datastore interface and provides per-call metrics. The
// metrics produced include a call counter and elapsed time measurement with
// labels for the status code.
func WithMetrics(ds datastore.DataStore, metrics telemetry.Metrics) datastore.DataStore {
	_ = "STUB: not implemented"
	return *new(datastore.DataStore)
}

type metricsWrapper struct {
	ds datastore.DataStore
	m  telemetry.Metrics
}

func (w metricsWrapper) AppendBundle(ctx context.Context, bundle *common.Bundle) (_ *common.Bundle, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) CreateAttestedNode(ctx context.Context, node *common.AttestedNode) (_ *common.AttestedNode, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) CreateAttestedNodeEventForTesting(ctx context.Context, event *datastore.AttestedNodeEvent) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w metricsWrapper) CreateBundle(ctx context.Context, bundle *common.Bundle) (_ *common.Bundle, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) CreateJoinToken(ctx context.Context, token *datastore.JoinToken) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w metricsWrapper) CreateRegistrationEntry(ctx context.Context, entry *common.RegistrationEntry) (_ *common.RegistrationEntry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) CreateOrReturnRegistrationEntry(ctx context.Context, entry *common.RegistrationEntry) (_ *common.RegistrationEntry, _ bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (w metricsWrapper) CreateRegistrationEntryEventForTesting(ctx context.Context, event *datastore.RegistrationEntryEvent) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w metricsWrapper) CreateFederationRelationship(ctx context.Context, fr *datastore.FederationRelationship) (_ *datastore.FederationRelationship, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) ListFederationRelationships(ctx context.Context, req *datastore.ListFederationRelationshipsRequest) (_ *datastore.ListFederationRelationshipsResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) DeleteAttestedNode(ctx context.Context, spiffeID string) (_ *common.AttestedNode, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) DeleteAttestedNodeEventForTesting(ctx context.Context, eventID uint) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w metricsWrapper) DeleteBundle(ctx context.Context, trustDomain string, mode datastore.DeleteMode) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w metricsWrapper) DeleteFederationRelationship(ctx context.Context, trustDomain spiffeid.TrustDomain) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w metricsWrapper) DeleteJoinToken(ctx context.Context, token string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w metricsWrapper) DeleteRegistrationEntry(ctx context.Context, entryID string) (_ *common.RegistrationEntry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) DeleteRegistrationEntryEventForTesting(ctx context.Context, eventID uint) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w metricsWrapper) FetchAttestedNode(ctx context.Context, spiffeID string) (_ *common.AttestedNode, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) FetchAttestedNodeEvent(ctx context.Context, eventID uint) (_ *datastore.AttestedNodeEvent, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) FetchBundle(ctx context.Context, trustDomain string) (_ *common.Bundle, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) FetchJoinToken(ctx context.Context, token string) (_ *datastore.JoinToken, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) FetchRegistrationEntry(ctx context.Context, entryID string) (_ *common.RegistrationEntry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) FetchRegistrationEntries(ctx context.Context, entryIDs []string) (_ map[string]*common.RegistrationEntry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) FetchRegistrationEntryEvent(ctx context.Context, eventID uint) (_ *datastore.RegistrationEntryEvent, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) FetchFederationRelationship(ctx context.Context, trustDomain spiffeid.TrustDomain) (_ *datastore.FederationRelationship, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) GetNodeSelectors(ctx context.Context, spiffeID string, dataConsistency datastore.DataConsistency) (_ []*common.Selector, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) ListAttestedNodes(ctx context.Context, req *datastore.ListAttestedNodesRequest) (_ *datastore.ListAttestedNodesResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) ListAttestedNodeEvents(ctx context.Context, req *datastore.ListAttestedNodeEventsRequest) (_ *datastore.ListAttestedNodeEventsResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) ListBundles(ctx context.Context, req *datastore.ListBundlesRequest) (_ *datastore.ListBundlesResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) ListNodeSelectors(ctx context.Context, req *datastore.ListNodeSelectorsRequest) (_ *datastore.ListNodeSelectorsResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) ListRegistrationEntries(ctx context.Context, req *datastore.ListRegistrationEntriesRequest) (_ *datastore.ListRegistrationEntriesResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) ListRegistrationEntryEvents(ctx context.Context, req *datastore.ListRegistrationEntryEventsRequest) (_ *datastore.ListRegistrationEntryEventsResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) CountAttestedNodes(ctx context.Context, req *datastore.CountAttestedNodesRequest) (_ int32, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w metricsWrapper) CountBundles(ctx context.Context) (_ int32, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w metricsWrapper) CountRegistrationEntries(ctx context.Context, req *datastore.CountRegistrationEntriesRequest) (_ int32, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w metricsWrapper) PruneAttestedNodeEvents(ctx context.Context, olderThan time.Duration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w metricsWrapper) PruneBundle(ctx context.Context, trustDomainID string, expiresBefore time.Time) (_ bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (w metricsWrapper) PruneJoinTokens(ctx context.Context, expiresBefore time.Time) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w metricsWrapper) PruneRegistrationEntries(ctx context.Context, expiresBefore time.Time) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w metricsWrapper) PruneRegistrationEntryEvents(ctx context.Context, olderThan time.Duration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w metricsWrapper) PruneAttestedExpiredNodes(ctx context.Context, expiredBefore time.Time, includeNonReattestable bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w metricsWrapper) SetBundle(ctx context.Context, bundle *common.Bundle) (_ *common.Bundle, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) TaintX509CA(ctx context.Context, trustDomainID string, subjectKeyIDToTaint string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w metricsWrapper) RevokeX509CA(ctx context.Context, trustDomainID string, subjectKeyIDToRevoke string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w metricsWrapper) TaintJWTKey(ctx context.Context, trustDomainID string, authorityID string) (_ *common.PublicKey, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) RevokeJWTKey(ctx context.Context, trustDomainID string, authorityID string) (_ *common.PublicKey, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) SetNodeSelectors(ctx context.Context, spiffeID string, selectors []*common.Selector) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w metricsWrapper) UpdateAttestedNode(ctx context.Context, node *common.AttestedNode, mask *common.AttestedNodeMask) (_ *common.AttestedNode, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) UpdateBundle(ctx context.Context, bundle *common.Bundle, mask *common.BundleMask) (_ *common.Bundle, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) UpdateRegistrationEntry(ctx context.Context, entry *common.RegistrationEntry, mask *common.RegistrationEntryMask) (_ *common.RegistrationEntry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) UpdateFederationRelationship(ctx context.Context, fr *datastore.FederationRelationship, mask *types.FederationRelationshipMask) (_ *datastore.FederationRelationship, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) SetCAJournal(ctx context.Context, caJournal *datastore.CAJournal) (_ *datastore.CAJournal, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) FetchCAJournal(ctx context.Context, activeX509AuthorityID string) (_ *datastore.CAJournal, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) ListCAJournalsForTesting(ctx context.Context) (_ []*datastore.CAJournal, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w metricsWrapper) PruneCAJournals(ctx context.Context, allCAsExpireBefore int64) (err error) {
	_ = "STUB: not implemented"
	return nil
}
