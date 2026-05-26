package sqlstore

import (
	"context"
	"database/sql"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/server/datastore"
	"github.com/spiffe/spire/proto/spire/common"
)

var validEntryIDChars = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x002d, 0x002e, 1}, // - | .
		{0x0030, 0x0039, 1}, // [0-9]
		{0x0041, 0x005a, 1}, // [A-Z]
		{0x005f, 0x005f, 1}, // _
		{0x0061, 0x007a, 1}, // [a-z]
	},
	LatinOffset: 5,
}

const (
	PluginName = "sql"

	// MySQL database type
	MySQL = "mysql"
	// PostgreSQL database type
	PostgreSQL = "postgres"
	// SQLite database type
	SQLite = "sqlite3"

	// MySQL database provided by an AWS service
	AWSMySQL = "aws_mysql"

	// PostgreSQL database type provided by an AWS service
	AWSPostgreSQL = "aws_postgres"

	// Maximum size for preallocation in a paginated request
	maxResultPreallocation = 1000

	// Maximum size for additional attributes message in a registration entry
	maxAdditionalAttributesSize = 65535
)

// Configuration for the sql datastore implementation.
// Pointer values are used to distinguish between "unset" and "zero" values.
type configuration struct {
	DatabaseTypeNode   ast.Node `hcl:"database_type" json:"database_type"`
	ConnectionString   string   `hcl:"connection_string" json:"connection_string"`
	RoConnectionString string   `hcl:"ro_connection_string" json:"ro_connection_string"`
	RootCAPath         string   `hcl:"root_ca_path" json:"root_ca_path"`
	ClientCertPath     string   `hcl:"client_cert_path" json:"client_cert_path"`
	ClientKeyPath      string   `hcl:"client_key_path" json:"client_key_path"`
	ConnMaxLifetime    *string  `hcl:"conn_max_lifetime" json:"conn_max_lifetime"`
	MaxOpenConns       *int     `hcl:"max_open_conns" json:"max_open_conns"`
	MaxIdleConns       *int     `hcl:"max_idle_conns" json:"max_idle_conns"`
	DisableMigration   bool     `hcl:"disable_migration" json:"disable_migration"`

	databaseTypeConfig *dbTypeConfig
	// Undocumented flags
	LogSQL bool `hcl:"log_sql" json:"log_sql"`
}

type dbTypeConfig struct {
	AWSMySQL     *awsConfig `hcl:"aws_mysql" json:"aws_mysql"`
	AWSPostgres  *awsConfig `hcl:"aws_postgres" json:"aws_postgres"`
	databaseType string
}

type awsConfig struct {
	Region          string `hcl:"region"`
	AccessKeyID     string `hcl:"access_key_id"`
	SecretAccessKey string `hcl:"secret_access_key"`
}

func (a *awsConfig) validate() error { _ = "STUB: not implemented"; return nil }

type sqlDB struct {
	databaseType     string
	connectionString string
	raw              *sql.DB
	*gorm.DB

	dialect     dialect
	stmtCache   *stmtCache
	supportsCTE bool

	// this lock is only required for synchronized writes with "sqlite3". see
	// the withTx() implementation for details.
	opMu sync.Mutex
}

func (db *sqlDB) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Plugin is a DataStore plugin implemented via a SQL database
type Plugin struct {
	mu                  sync.Mutex
	db                  *sqlDB
	roDb                *sqlDB
	log                 logrus.FieldLogger
	useServerTimestamps bool
}

// New creates a new sql plugin struct. Configure must be called
// in order to start the db.
func New(log logrus.FieldLogger) *Plugin { _ = "STUB: not implemented"; return nil }

// CreateBundle stores the given bundle
func (ds *Plugin) CreateBundle(ctx context.Context, b *common.Bundle) (bundle *common.Bundle, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateBundle updates an existing bundle with the given CAs. Overwrites any
// existing certificates.
func (ds *Plugin) UpdateBundle(ctx context.Context, b *common.Bundle, mask *common.BundleMask) (bundle *common.Bundle, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetBundle sets bundle contents. If no bundle exists for the trust domain, it is created.
func (ds *Plugin) SetBundle(ctx context.Context, b *common.Bundle) (bundle *common.Bundle, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AppendBundle append bundle contents to the existing bundle (by trust domain). If no existing one is present, create it.
func (ds *Plugin) AppendBundle(ctx context.Context, b *common.Bundle) (bundle *common.Bundle, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteBundle deletes the bundle with the matching TrustDomain. Any CACert data passed is ignored.
func (ds *Plugin) DeleteBundle(ctx context.Context, trustDomainID string, mode datastore.DeleteMode) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// FetchBundle returns the bundle matching the specified Trust Domain.
func (ds *Plugin) FetchBundle(ctx context.Context, trustDomainID string) (resp *common.Bundle, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CountBundles can be used to count all existing bundles.
func (ds *Plugin) CountBundles(ctx context.Context) (count int32, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ListBundles can be used to fetch all existing bundles.
func (ds *Plugin) ListBundles(ctx context.Context, req *datastore.ListBundlesRequest) (resp *datastore.ListBundlesResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PruneBundle removes expired certs and keys from a bundle
func (ds *Plugin) PruneBundle(ctx context.Context, trustDomainID string, expiresBefore time.Time) (changed bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// TaintX509CAByKey taints an X.509 CA signed using the provided public key
func (ds *Plugin) TaintX509CA(ctx context.Context, trustDoaminID string, subjectKeyIDToTaint string) error {
	_ = "STUB: not implemented"
	return nil
}

// RevokeX509CA removes a Root CA from the bundle
func (ds *Plugin) RevokeX509CA(ctx context.Context, trustDoaminID string, subjectKeyIDToRevoke string) error {
	_ = "STUB: not implemented"
	return nil
}

// TaintJWTKey taints a JWT Authority key
func (ds *Plugin) TaintJWTKey(ctx context.Context, trustDoaminID string, authorityID string) (*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RevokeJWTAuthority removes JWT key from the bundle
func (ds *Plugin) RevokeJWTKey(ctx context.Context, trustDoaminID string, authorityID string) (*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateAttestedNode stores the given attested node
func (ds *Plugin) CreateAttestedNode(ctx context.Context, node *common.AttestedNode) (attestedNode *common.AttestedNode, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FetchAttestedNode fetches an existing attested node by SPIFFE ID
func (ds *Plugin) FetchAttestedNode(ctx context.Context, spiffeID string) (attestedNode *common.AttestedNode, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CountAttestedNodes counts all attested nodes
func (ds *Plugin) CountAttestedNodes(ctx context.Context, req *datastore.CountAttestedNodesRequest) (count int32, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ListAttestedNodes lists all attested nodes (pagination available)
func (ds *Plugin) ListAttestedNodes(ctx context.Context,
	req *datastore.ListAttestedNodesRequest,
) (resp *datastore.ListAttestedNodesResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateAttestedNode updates the given node's cert serial and expiration.
func (ds *Plugin) UpdateAttestedNode(ctx context.Context, n *common.AttestedNode, mask *common.AttestedNodeMask) (node *common.AttestedNode, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteAttestedNode deletes the given attested node and the associated node selectors.
func (ds *Plugin) DeleteAttestedNode(ctx context.Context, spiffeID string) (attestedNode *common.AttestedNode, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PruneAttestedExpiredNodes deletes attested nodes with expiration time further than a given duration in the past.
// Non-reattestable nodes are not deleted by default, and have to be included explicitly by setting
// includeNonReattestable = true. Banned nodes are not deleted.
func (ds *Plugin) PruneAttestedExpiredNodes(ctx context.Context, expiredBefore time.Time, includeNonReattestable bool) error {
	_ = "STUB: not implemented"
	return nil
}

// ListAttestedNodeEvents lists all attested node events
func (ds *Plugin) ListAttestedNodeEvents(ctx context.Context, req *datastore.ListAttestedNodeEventsRequest) (resp *datastore.ListAttestedNodeEventsResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PruneAttestedNodeEvents deletes all attested node events older than a specified duration (i.e. more than 24 hours old)
func (ds *Plugin) PruneAttestedNodeEvents(ctx context.Context, olderThan time.Duration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// CreateRegistrationEntryEventForTestingForTesting creates an attested node event. Used for unit testing.
func (ds *Plugin) CreateAttestedNodeEventForTesting(ctx context.Context, event *datastore.AttestedNodeEvent) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteAttestedNodeEventForTesting deletes an attested node event by event ID. Used for unit testing.
func (ds *Plugin) DeleteAttestedNodeEventForTesting(ctx context.Context, eventID uint) error {
	_ = "STUB: not implemented"
	return nil
}

// FetchAttestedNodeEvent fetches an existing attested node event by event ID
func (ds *Plugin) FetchAttestedNodeEvent(ctx context.Context, eventID uint) (event *datastore.AttestedNodeEvent, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetNodeSelectors sets node (agent) selectors by SPIFFE ID, deleting old selectors first
func (ds *Plugin) SetNodeSelectors(ctx context.Context, spiffeID string, selectors []*common.Selector) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// GetNodeSelectors gets node (agent) selectors by SPIFFE ID
func (ds *Plugin) GetNodeSelectors(ctx context.Context, spiffeID string,
	dataConsistency datastore.DataConsistency,
) (selectors []*common.Selector, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListNodeSelectors gets node (agent) selectors by SPIFFE ID
func (ds *Plugin) ListNodeSelectors(ctx context.Context,
	req *datastore.ListNodeSelectorsRequest,
) (resp *datastore.ListNodeSelectorsResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateRegistrationEntry stores the given registration entry
func (ds *Plugin) CreateRegistrationEntry(ctx context.Context,
	entry *common.RegistrationEntry,
) (registrationEntry *common.RegistrationEntry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateOrReturnRegistrationEntry stores the given registration entry. If an
// entry already exists with the same (parentID, spiffeID, selector) tuple,
// that entry is returned instead.
func (ds *Plugin) CreateOrReturnRegistrationEntry(ctx context.Context,
	entry *common.RegistrationEntry,
) (registrationEntry *common.RegistrationEntry, existing bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (ds *Plugin) createOrReturnRegistrationEntry(ctx context.Context,
	entry *common.RegistrationEntry,
) (registrationEntry *common.RegistrationEntry, existing bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// FetchRegistrationEntry fetches an existing registration by entry ID
func (ds *Plugin) FetchRegistrationEntry(ctx context.Context,
	entryID string,
) (*common.RegistrationEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return the last element in the list

// FetchRegistrationEntries fetches existing registrations by entry IDs
func (ds *Plugin) FetchRegistrationEntries(ctx context.Context,
	entryIDs []string,
) (map[string]*common.RegistrationEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CountRegistrationEntries counts all registrations (pagination available)
func (ds *Plugin) CountRegistrationEntries(ctx context.Context, req *datastore.CountRegistrationEntriesRequest) (count int32, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ListRegistrationEntries lists all registrations (pagination available)
func (ds *Plugin) ListRegistrationEntries(ctx context.Context,
	req *datastore.ListRegistrationEntriesRequest,
) (resp *datastore.ListRegistrationEntriesResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateRegistrationEntry updates an existing registration entry
func (ds *Plugin) UpdateRegistrationEntry(ctx context.Context, e *common.RegistrationEntry, mask *common.RegistrationEntryMask) (entry *common.RegistrationEntry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteRegistrationEntry deletes the given registration
func (ds *Plugin) DeleteRegistrationEntry(ctx context.Context,
	entryID string,
) (registrationEntry *common.RegistrationEntry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PruneRegistrationEntries takes a registration entry message, and deletes all entries which have expired
// before the date in the message
func (ds *Plugin) PruneRegistrationEntries(ctx context.Context, expiresBefore time.Time) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// ListRegistrationEntryEvents lists all registration entry events
func (ds *Plugin) ListRegistrationEntryEvents(ctx context.Context, req *datastore.ListRegistrationEntryEventsRequest) (resp *datastore.ListRegistrationEntryEventsResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PruneRegistrationEntryEvents deletes all registration entry events older than a specified duration (i.e. more than 24 hours old)
func (ds *Plugin) PruneRegistrationEntryEvents(ctx context.Context, olderThan time.Duration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// CreateRegistrationEntryEventForTesting creates a registration entry event. Used for unit testing.
func (ds *Plugin) CreateRegistrationEntryEventForTesting(ctx context.Context, event *datastore.RegistrationEntryEvent) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteRegistrationEntryEventForTesting deletes the given registration entry event. Used for unit testing.
func (ds *Plugin) DeleteRegistrationEntryEventForTesting(ctx context.Context, eventID uint) error {
	_ = "STUB: not implemented"
	return nil
}

// FetchRegistrationEntryEvent fetches an existing registration entry event by event ID
func (ds *Plugin) FetchRegistrationEntryEvent(ctx context.Context, eventID uint) (event *datastore.RegistrationEntryEvent, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateJoinToken takes a Token message and stores it
func (ds *Plugin) CreateJoinToken(ctx context.Context, token *datastore.JoinToken) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// FetchJoinToken takes a Token message and returns one, populating the fields
// we have knowledge of
func (ds *Plugin) FetchJoinToken(ctx context.Context, token string) (resp *datastore.JoinToken, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteJoinToken deletes the given join token
func (ds *Plugin) DeleteJoinToken(ctx context.Context, token string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// PruneJoinTokens takes a Token message, and deletes all tokens which have expired
// before the date in the message
func (ds *Plugin) PruneJoinTokens(ctx context.Context, expiry time.Time) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// CreateFederationRelationship creates a new federation relationship. If the bundle endpoint
// profile is 'https_spiffe' and the given federation relationship contains a bundle, the current
// stored bundle is overridden.
// If no bundle is provided and there is not a previously stored bundle in the datastore, the
// federation relationship is not created.
func (ds *Plugin) CreateFederationRelationship(ctx context.Context, fr *datastore.FederationRelationship) (newFr *datastore.FederationRelationship, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteFederationRelationship deletes the federation relationship to the
// given trust domain. The associated trust bundle is not deleted.
func (ds *Plugin) DeleteFederationRelationship(ctx context.Context, trustDomain spiffeid.TrustDomain) error {
	_ = "STUB: not implemented"
	return nil
}

// FetchFederationRelationship fetches the federation relationship that matches
// the given trust domain. If the federation relationship is not found, nil is returned.
func (ds *Plugin) FetchFederationRelationship(ctx context.Context, trustDomain spiffeid.TrustDomain) (fr *datastore.FederationRelationship, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListFederationRelationships can be used to list all existing federation relationships
func (ds *Plugin) ListFederationRelationships(ctx context.Context, req *datastore.ListFederationRelationshipsRequest) (resp *datastore.ListFederationRelationshipsResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateFederationRelationship updates the given federation relationship.
// Attributes are only updated if the correspondent mask value is set to true.
func (ds *Plugin) UpdateFederationRelationship(ctx context.Context, fr *datastore.FederationRelationship, mask *types.FederationRelationshipMask) (newFr *datastore.FederationRelationship, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetUseServerTimestamps controls whether server-generated timestamps should be used in the database.
// This is only intended to be used by tests in order to produce deterministic timestamp data,
// since some databases round off timestamp data with lower precision.
func (ds *Plugin) SetUseServerTimestamps(useServerTimestamps bool) {
	_ = "STUB: not implemented"
	return
}

// FetchCAJournal fetches the CA journal that has the given active X509
// authority domain. If the CA journal is not found, nil is returned.
func (ds *Plugin) FetchCAJournal(ctx context.Context, activeX509AuthorityID string) (caJournal *datastore.CAJournal, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListCAJournalsForTesting returns all the CA journal records, and is meant to
// be used in tests.
func (ds *Plugin) ListCAJournalsForTesting(ctx context.Context) (caJournals []*datastore.CAJournal, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetCAJournal sets the content for the specified CA journal. If the CA journal
// does not exist, it is created.
func (ds *Plugin) SetCAJournal(ctx context.Context, caJournal *datastore.CAJournal) (caj *datastore.CAJournal, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The CA journal already exists, update it.

// PruneCAJournals prunes the CA journals that have all of their authorities
// expired.
func (ds *Plugin) PruneCAJournals(ctx context.Context, allAuthoritiesExpireBefore int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *Plugin) pruneCAJournals(tx *gorm.DB, allAuthoritiesExpireBefore int64) error {
	_ = "STUB: not implemented"
	return nil
}

// Configure parses HCL config payload into config struct, opens new DB based on the result, and
// prunes all orphaned records
func (ds *Plugin) Configure(ctx context.Context, hclConfiguration string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *Plugin) Validate(ctx context.Context, coreConfig catalog.CoreConfig, configuration string) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildConfig(hclConfiguration string) (*configuration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *Plugin) openConnections(ctx context.Context, config *configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *Plugin) openConnection(ctx context.Context, config *configuration, isReadOnly bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *Plugin) Close() error { _ = "STUB: not implemented"; return nil }

// withReadModifyWriteTx wraps the operation in a transaction appropriate for
// operations that will read one or more rows, change one or more columns in
// those rows, and then set them back. This requires a stronger level of
// consistency that prevents two transactions from doing read-modify-write
// concurrently.
func (ds *Plugin) withReadModifyWriteTx(ctx context.Context, op func(tx *gorm.DB) error) error {
	_ = "STUB: not implemented"
	return nil
}

// MySQL REPEATABLE READ is weaker than that of PostgreSQL. Namely,
// PostgreSQL, beyond providing the minimum consistency guarantees
// mandated for REPEATABLE READ in the standard, automatically fails
// concurrent transactions that try to update the same target row.
//
// To get the same consistency guarantees, have the queries do a
// `SELECT .. FOR UPDATE` which will implicitly lock queried rows
// from update by other transactions. This is preferred to a stronger
// isolation level, like SERIALIZABLE, which is not supported by
// some MySQL-compatible databases (i.e. Percona XtraDB cluster)

// `SELECT .. FOR UPDATE`is also required when PostgreSQL is in
// hot standby mode for this operation to work properly (see issue #3039).

// withWriteTx wraps the operation in a transaction appropriate for operations
// that unconditionally create/update rows, without reading them first. If two
// transactions try and update at the same time, last writer wins.
func (ds *Plugin) withWriteTx(ctx context.Context, op func(tx *gorm.DB) error) error {
	_ = "STUB: not implemented"
	return nil
}

// withReadTx wraps the operation in a transaction appropriate for operations
// that only read rows.
func (ds *Plugin) withReadTx(ctx context.Context, op func(tx *gorm.DB) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *Plugin) withTx(ctx context.Context, op func(tx *gorm.DB) error, readOnly bool) error {
	_ = "STUB: not implemented"
	return nil
}

// sqlite3 can only have one writer at a time. since we're in WAL mode,
// there can be concurrent reads and writes, so no lock is necessary
// over the read operations.

// rolling back makes sure that functions that are invoked with
// withReadTx, and then do writes, will not pass unit tests, since the
// writes won't be committed.

// gormToGRPCStatus takes an error, and converts it to a GRPC error.  If the
// error is already a gRPC status , it will be returned unmodified. Otherwise
// if the error is a gorm error type with a known mapping to a GRPC status,
// that code will be set, otherwise the code will be set to Unknown.
func (ds *Plugin) gormToGRPCStatus(err error) error { _ = "STUB: not implemented"; return nil }

func (ds *Plugin) openDB(ctx context.Context, cfg *configuration, isReadOnly bool) (*gorm.DB, string, bool, dialect, error) {
	_ = "STUB: not implemented"
	return nil, "", false, *new(dialect), nil
}

// Round to nearest second to be consistent with how timestamps are rounded in CreateRegistrationEntry calls

type gormLogger struct {
	log logrus.FieldLogger
}

func (logger gormLogger) Print(v ...any) { _ = "STUB: not implemented"; return }

func createBundle(tx *gorm.DB, bundle *common.Bundle) (*common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func updateBundle(tx *gorm.DB, newBundle *common.Bundle, mask *common.BundleMask) (*common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func applyBundleMask(model *Bundle, newBundle *common.Bundle, inputMask *common.BundleMask) ([]byte, *common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func setBundle(tx *gorm.DB, b *common.Bundle) (*common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fetch existing or create new

func appendBundle(tx *gorm.DB, b *common.Bundle) (*common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fetch existing or create new

// parse the bundle data and add missing elements

func deleteBundle(tx *gorm.DB, trustDomainID string, mode datastore.DeleteMode) error {
	_ = "STUB: not implemented"
	return nil
}

// Get a count of associated registration entries

// TODO: figure out how to do this gracefully with GORM.

// fetchBundle returns the bundle matching the specified Trust Domain.
func fetchBundle(tx *gorm.DB, trustDomainID string) (*common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// countBundles can be used to count existing bundles
func countBundles(tx *gorm.DB) (int32, error) { _ = "STUB: not implemented"; return 0, nil }

// listBundles can be used to fetch all existing bundles.
func listBundles(tx *gorm.DB, req *datastore.ListBundlesRequest) (*datastore.ListBundlesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set token only if page size is the same than bundles len

func pruneBundle(tx *gorm.DB, trustDomainID string, expiry time.Time, log logrus.FieldLogger) (bool, error) {
	_ = "STUB: not implemented"
	// Get current bundle
	return false, nil
}

// No bundle to prune

// Prune

// Update only if bundle was modified

func taintX509CA(tx *gorm.DB, trustDomainID string, subjectKeyIDToTaint string) error {
	_ = "STUB: not implemented"
	return nil
}

func revokeX509CA(tx *gorm.DB, trustDomainID string, subjectKeyIDToRevoke string) error {
	_ = "STUB: not implemented"
	return nil
}

func taintJWTKey(tx *gorm.DB, trustDomainID string, authorityID string) (*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if a JWT Key with the provided keyID was already
// tainted in this loop. This is purely defensive since we do not
// allow to have repeated key IDs.

func revokeJWTKey(tx *gorm.DB, trustDomainID string, authorityID string) (*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if a JWT Key with the provided keyID was already
// found in this loop. This is purely defensive since we do not
// allow to have repeated key IDs.

func getBundle(tx *gorm.DB, trustDomainID string) (*common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createAttestedNode(tx *gorm.DB, node *common.AttestedNode) (*common.AttestedNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fetchAttestedNode(tx *gorm.DB, spiffeID string) (*common.AttestedNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func countAttestedNodes(tx *gorm.DB) (int32, error) { _ = "STUB: not implemented"; return 0, nil }

func countAttestedNodesHasFilters(req *datastore.CountAttestedNodesRequest) bool {
	_ = "STUB: not implemented"
	return false
}

func listAttestedNodes(ctx context.Context, db *sqlDB, log logrus.FieldLogger, req *datastore.ListAttestedNodesRequest) (*datastore.ListAttestedNodesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Now that we've filtered the nodes based on selectors, prune off
// selectors from the response if they were not requested.

// This check is purely defensive. Assuming the pagination code is
// correct, a request with a given token should never yield that
// same token. Just in case, we don't want the server to loop
// indefinitely.

func countAttestedNodesWithFilters(ctx context.Context, db *sqlDB, _ logrus.FieldLogger, req *datastore.CountAttestedNodesRequest) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func createAttestedNodeEvent(tx *gorm.DB, event *datastore.AttestedNodeEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func listAttestedNodeEvents(db *sqlDB, req *datastore.ListAttestedNodeEventsRequest) (*datastore.ListAttestedNodeEventsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func pruneAttestedNodeEvents(tx *gorm.DB, olderThan time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func notBanned(tx *gorm.DB) *gorm.DB { _ = "STUB: not implemented"; return nil }

func expiredForDuration(expiredBefore time.Time) func(db *gorm.DB) *gorm.DB {
	_ = "STUB: not implemented"
	return nil
}

func includeNonReattestable(include bool) func(db *gorm.DB) *gorm.DB {
	_ = "STUB: not implemented"
	return nil
}

func pruneAttestedExpiredNodes(tx *gorm.DB, expiredBefore time.Time, include bool, logger logrus.FieldLogger) error {
	_ = "STUB: not implemented"
	return nil
}

func fetchAttestedNodeEvent(db *sqlDB, eventID uint) (*datastore.AttestedNodeEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deleteAttestedNodeEvent(tx *gorm.DB, eventID uint) error {
	_ = "STUB: not implemented"
	return nil
}

// filterNodesBySelectorSet filters nodes based on provided selectors
func filterNodesBySelectorSet(nodes []*common.AttestedNode, selectors []*common.Selector) []*common.AttestedNode {
	_ = "STUB: not implemented"
	return nil
}

func listAttestedNodesOnce(ctx context.Context, db *sqlDB, req *datastore.ListAttestedNodesRequest) (*datastore.ListAttestedNodesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildListAttestedNodesQuery(dbType string, supportsCTE bool, req *datastore.ListAttestedNodesRequest) (string, []any, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// The PostgreSQL queries unconditionally leverage CTE since all versions
// of PostgreSQL supported by the plugin support CTE.

func buildListAttestedNodesQueryCTE(req *datastore.ListAttestedNodesRequest, dbType string) (string, []any, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Selectors will be fetched only when `FetchSelectors` or BySelectorMatch are in request

// Creates filtered nodes, `true` is added to simplify code, all filters will start with `AND`

// Filter by pagination token

// Filter by expiration

// Filter by Attestation type

// Filter by banned, an Attestation Node is banned when serial number is empty.
// This filter allows 3 outputs:
// - nil:  returns all
// - true: returns banned entries
// - false: returns no banned entries

// Filter by canReattest,
// This filter allows 3 outputs:
//  - nil:  returns all
// - true: returns nodes with canReattest=true
// - false: returns nodes with canReattest=false

// Fetch all selectors from filtered entries

// Add expected fields

// Add "optional" fields for selectors

// Choose what table will be used

// MySQL requires a subquery in order to apply pagination

// Add filter by selectors

// Select IDs, that will be used to fetch "paged" entrieSelect IDs, that will be used to fetch "paged" entries

// Subset needs a union, so we need to group them and add the group
// as a child to the root

// MySQL does not support INTERSECT, so use INNER JOIN instead

// First subquery does not need USING(ID)

// Last query does not need INNER JOIN

// Add all selectors as arguments

// Prevent duplicate IDs when fetching selectors

// Add workaround for limit

func buildListAttestedNodesQueryMySQL(req *datastore.ListAttestedNodesRequest) (string, []any, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Selectors will be fetched only when `FetchSelectors` or `BySelectorMatch` are in request

// Add expected fields

// Add "optional" fields for selectors

// Filter by pagination token

// Filter by expiration

// Filter by valid_at

// Filter by Attestation type

// Filter by banned, an Attestation Node is banned when serial number is empty.
// This filter allows 3 outputs:
// - nil:  returns all
// - true: returns banned entries
// - false: returns no banned entries

// Filter by CanReattest. This is similar to ByBanned

// Add filter by selectors

// subset needs a union, so we need to group them and add the group
// as a child to the root.

func updateAttestedNode(tx *gorm.DB, n *common.AttestedNode, mask *common.AttestedNodeMask) (*common.AttestedNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deleteAttestedNodeAndSelectors(tx *gorm.DB, spiffeID string) (*common.AttestedNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// batch delete all associated node selectors

func setNodeSelectors(tx *gorm.DB, spiffeID string, selectors []*common.Selector) error {
	_ = "STUB: not implemented"
	// Previously the deletion of the previous set of node selectors was
	// implemented via query like DELETE FROM node_resolver_map_entries WHERE
	// spiffe_id = ?, but unfortunately this triggered some pessimistic gap
	// locks on the index even when there were no rows matching the WHERE
	// clause (i.e. rows for that spiffe_id). The gap locks caused MySQL
	// deadlocks when SetNodeSelectors was being called concurrently. Changing
	// the transaction isolation level fixed the deadlocks but only when there
	// were no existing rows; the deadlocks still occurred when existing rows
	// existed (i.e. re-attestation). Instead, gather all the IDs to be
	// deleted and delete them from separate queries, which does not trigger
	// gap locks on the index.
	return nil
}

func getNodeSelectors(ctx context.Context, db *sqlDB, spiffeID string) ([]*common.Selector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func listNodeSelectors(ctx context.Context, db *sqlDB, req *datastore.ListNodeSelectorsRequest) (*datastore.ListNodeSelectorsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildListNodeSelectorsQuery(req *datastore.ListNodeSelectorsRequest) (query string, args []any) {
	_ = "STUB: not implemented"
	return "", nil
}

// This ordering is required to make listNodeSelectors efficient but not
// needed for correctness. Since the query can be wholly satisfied using
// the node_resolver_map_entries unique index over (spiffe_id,type,value)
// it is unlikely to impact database performance as that index is already
// ordered primarily by spiffe_id.

func createRegistrationEntry(tx *gorm.DB, entry *common.RegistrationEntry) (*common.RegistrationEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fetchRegistrationEntries(ctx context.Context, db *sqlDB, entryIDs []string) (map[string]*common.RegistrationEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert array to map

func buildFetchRegistrationEntriesQuery(dbType string, supportsCTE bool, entryIDs []string) (string, []any, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// The SQLite3 queries unconditionally leverage CTE since the
// embedded version of SQLite3 supports CTE.

// The PostgreSQL queries unconditionally leverage CTE since all versions
// of PostgreSQL supported by the plugin support CTE.

func buildFetchRegistrationEntriesQuerySQLite3(entryIDs []string) (string, []any, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func buildFetchRegistrationEntriesQueryPostgreSQL(entryIDs []string) (string, []any, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func buildFetchRegistrationEntriesQueryMySQL(entryIDs []string) (string, []any, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func buildFetchRegistrationEntriesQueryMySQLCTE(entryIDs []string) (string, []any, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func listRegistrationEntries(ctx context.Context, db *sqlDB, log logrus.FieldLogger, req *datastore.ListRegistrationEntriesRequest) (*datastore.ListRegistrationEntriesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Exact/subset selector matching requires filtering out all registration
// entries returned by the query whose selectors are not fully represented
// in the request selectors. For this reason, it's possible that a paged
// query returns rows that are completely filtered out. If that happens,
// keep querying until a page gets at least one result.

// This check is purely defensive. Assuming the pagination code is
// correct, a request with a given token should never yield that
// same token. Just in case, we don't want the server to loop
// indefinitely.

func filterEntriesBySelectorSet(entries []*common.RegistrationEntry, selectors []*common.Selector) []*common.RegistrationEntry {
	_ = "STUB: not implemented"
	// Nothing to filter
	return nil
}

type queryContext interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

func listRegistrationEntriesOnce(ctx context.Context, db queryContext, databaseType string, supportsCTE bool, req *datastore.ListRegistrationEntriesRequest) (*datastore.ListRegistrationEntriesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildListRegistrationEntriesQuery(dbType string, supportsCTE bool, req *datastore.ListRegistrationEntriesRequest) (string, []any, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// The SQLite3 queries unconditionally leverage CTE since the
// embedded version of SQLite3 supports CTE.

// The PostgreSQL queries unconditionally leverage CTE since all versions
// of PostgreSQL supported by the plugin support CTE.

func buildListRegistrationEntriesQuerySQLite3(req *datastore.ListRegistrationEntriesRequest) (string, []any, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func buildListRegistrationEntriesQueryPostgreSQL(req *datastore.ListRegistrationEntriesRequest) (string, []any, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func maybeRebind(dbType, query string) string { _ = "STUB: not implemented"; return "" }

func postgreSQLRebind(s string) string { _ = "STUB: not implemented"; return "" }

func buildListRegistrationEntriesQueryMySQL(req *datastore.ListRegistrationEntriesRequest) (string, []any, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func buildListRegistrationEntriesQueryMySQLCTE(req *datastore.ListRegistrationEntriesRequest) (string, []any, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Count Registration Entries
func countRegistrationEntries(ctx context.Context, db *sqlDB, _ logrus.FieldLogger, req *datastore.CountRegistrationEntriesRequest) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type idFilterNode struct {
	idColumn string

	// mutually exclusive with children
	// supports multiline query
	query []string

	// mutually exclusive with query
	children []idFilterNode
	union    bool
	name     string

	fixed bool
}

func (n idFilterNode) Render(builder *strings.Builder, dbType string, indentation int, eol bool) {
	_ = "STUB: not implemented"
	return
}

func (n idFilterNode) render(builder *strings.Builder, dbType string, sibling int, indentation int, bol, eol bool) {
	_ = "STUB: not implemented"
	return
}

func indent(builder *strings.Builder, indentation int) { _ = "STUB: not implemented"; return }

func appendListRegistrationEntriesFilterQuery(filterExp string, builder *strings.Builder, dbType string, req *datastore.ListRegistrationEntriesRequest) (bool, []any, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// subset needs a union, so we need to group them and add the group
// as a child to the root.

// exact match does use an intersection, so we can just add these
// directly to the root idFilterNode, since it is already an intersection

// Take the trust domains from the request without duplicates

// Exact/subset federates-with matching requires filtering out all registration
// entries whose federated trust domains are not fully represented in the request

// Subset federates-with matching requires filtering out all registration
// entries that don't federate with even one trust domain in the request

// Exact federates-with matching requires filtering out all registration
// entries that don't federate with all the trust domains in the request

// MatchAny federates-with matching requires filtering out all registration
// entries that has at least one trust domain in the request

// SuperSet federates-with matching requires filtering out all registration
// entries has all trustdomains

func buildSliceArg(length int) string { _ = "STUB: not implemented"; return "" }

type nodeRow struct {
	EId             uint64
	SpiffeID        string
	DataType        sql.NullString
	SerialNumber    sql.NullString
	ExpiresAt       sql.NullTime
	NewSerialNumber sql.NullString
	NewExpiresAt    sql.NullTime
	CanReattest     sql.NullBool
	AgentVersion    sql.NullString
	SelectorType    sql.NullString
	SelectorValue   sql.NullString
}

func scanNodeRow(rs *sql.Rows, r *nodeRow) error { _ = "STUB: not implemented"; return nil }

func fillNodeFromRow(node *common.AttestedNode, r *nodeRow) error {
	_ = "STUB: not implemented"
	return nil
}

type nodeSelectorRow struct {
	SpiffeID sql.NullString
	Type     sql.NullString
	Value    sql.NullString
}

func scanNodeSelectorRow(rs *sql.Rows, r *nodeSelectorRow) error {
	_ = "STUB: not implemented"
	return nil
}

func fillNodeSelectorFromRow(nodeSelector *common.Selector, r *nodeSelectorRow) {
	_ = "STUB: not implemented"
	return
}

type entryRow struct {
	EId                  uint64
	EntryID              sql.NullString
	SpiffeID             sql.NullString
	ParentID             sql.NullString
	RegTTL               sql.NullInt64
	Admin                sql.NullBool
	Downstream           sql.NullBool
	Expiry               sql.NullInt64
	SelectorID           sql.NullInt64
	SelectorType         sql.NullString
	SelectorValue        sql.NullString
	StoreSvid            sql.NullBool
	Hint                 sql.NullString
	CreatedAt            sql.NullTime
	TrustDomain          sql.NullString
	DNSNameID            sql.NullInt64
	DNSName              sql.NullString
	RevisionNumber       sql.NullInt64
	RegJwtSvidTTL        sql.NullInt64
	AdditionalAttributes sql.Null[[]byte]
}

func scanEntryRow(rs *sql.Rows, r *entryRow) error { _ = "STUB: not implemented"; return nil }

func fillEntryFromRow(entry *common.RegistrationEntry, r *entryRow) error {
	_ = "STUB: not implemented"
	return nil
}

// applyPagination  add order limit and token to current query
func applyPagination(p *datastore.Pagination, entryTx *gorm.DB) (*gorm.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func updateRegistrationEntry(tx *gorm.DB, e *common.RegistrationEntry, mask *common.RegistrationEntryMask) (*common.RegistrationEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get the existing entry

// Delete existing selectors - we will write new ones

// Verify that final selectors contains the same 'type' when entry is used for store SVIDs

// Delete existing DNSs - we will write new ones

// Revision number is increased by 1 on every update call

// The FederatesWith field in entry is filled in by the call to modelToEntry below

func deleteRegistrationEntry(tx *gorm.DB, entryID string) (*common.RegistrationEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deleteRegistrationEntrySupport(tx *gorm.DB, entry RegisteredEntry) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete existing selectors

// Delete existing dns_names

func pruneRegistrationEntries(tx *gorm.DB, expiresBefore time.Time, logger logrus.FieldLogger) error {
	_ = "STUB: not implemented"
	return nil
}

func createRegistrationEntryEvent(tx *gorm.DB, event *datastore.RegistrationEntryEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func fetchRegistrationEntryEvent(db *sqlDB, eventID uint) (*datastore.RegistrationEntryEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deleteRegistrationEntryEvent(tx *gorm.DB, eventID uint) error {
	_ = "STUB: not implemented"
	return nil
}

func listRegistrationEntryEvents(db *sqlDB, req *datastore.ListRegistrationEntryEventsRequest) (*datastore.ListRegistrationEntryEventsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func pruneRegistrationEntryEvents(tx *gorm.DB, olderThan time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func buildListEventsQueryString(greaterThanEventID, lessThanEventID uint) (*strings.Builder, uint, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func createJoinToken(tx *gorm.DB, token *datastore.JoinToken) error {
	_ = "STUB: not implemented"
	return nil
}

func fetchJoinToken(tx *gorm.DB, token string) (*datastore.JoinToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deleteJoinToken(tx *gorm.DB, token string) error { _ = "STUB: not implemented"; return nil }

func pruneJoinTokens(tx *gorm.DB, expiresBefore time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func createFederationRelationship(tx *gorm.DB, fr *datastore.FederationRelationship) (*datastore.FederationRelationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// overwrite current bundle

func deleteFederationRelationship(tx *gorm.DB, trustDomain spiffeid.TrustDomain) error {
	_ = "STUB: not implemented"
	return nil
}

func fetchFederationRelationship(tx *gorm.DB, trustDomain spiffeid.TrustDomain) (*datastore.FederationRelationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// listFederationRelationships can be used to fetch all existing federation relationships.
func listFederationRelationships(tx *gorm.DB, req *datastore.ListFederationRelationshipsRequest) (*datastore.ListFederationRelationshipsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set token only if page size is the same as federationRelationships len

func updateFederationRelationship(tx *gorm.DB, fr *datastore.FederationRelationship, mask *types.FederationRelationshipMask) (*datastore.FederationRelationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// overwrite current bundle

func validateFederationRelationship(fr *datastore.FederationRelationship, mask *types.FederationRelationshipMask) error {
	_ = "STUB: not implemented"
	return nil
}

func modelToFederationRelationship(tx *gorm.DB, model *FederatedTrustDomain) (*datastore.FederationRelationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// modelToBundle converts the given bundle model to a Protobuf bundle message. It will also
// include any embedded CACert models.
func modelToBundle(model *Bundle) (*common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func marshalAndValidateAdditionalAttributes(additionalAttributes *common.RegistrationEntry_AdditionalAttributes) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateRegistrationEntry(entry *common.RegistrationEntry) error {
	_ = "STUB: not implemented"
	return nil
}

// In case of StoreSvid is set, all entries 'must' be the same type,
// it is done to avoid users to mix selectors from different platforms in
// entries with storable SVIDs

// Selectors must never be empty

// equalSelectorTypes validates that all selectors has the same type,
func equalSelectorTypes(selectors []Selector) bool { _ = "STUB: not implemented"; return false }

func validateRegistrationEntryForUpdate(entry *common.RegistrationEntry, mask *common.RegistrationEntryMask) error {
	_ = "STUB: not implemented"
	return nil
}

// bundleToModel converts the given Protobuf bundle message to a database model. It
// performs validation, and fully parses certificates to form CACert embedded models.
func bundleToModel(pb *common.Bundle) (*Bundle, error) { _ = "STUB: not implemented"; return nil, nil }

func modelToEntry(tx *gorm.DB, model RegisteredEntry) (*common.RegistrationEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createOrReturnEntryID(entry *common.RegistrationEntry) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func newRegistrationEntryID() (string, error) { _ = "STUB: not implemented"; return "", nil }

func modelToAttestedNode(model AttestedNode) *common.AttestedNode {
	_ = "STUB: not implemented"
	return nil
}

func modelToJoinToken(model JoinToken) *datastore.JoinToken { _ = "STUB: not implemented"; return nil }

func modelToCAJournal(model CAJournal) *datastore.CAJournal { _ = "STUB: not implemented"; return nil }

func makeFederatesWith(tx *gorm.DB, ids []string) ([]*Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// make sure all the ids were found

func bindVars(db *gorm.DB, query string) string { _ = "STUB: not implemented"; return "" }

func bindVarsFn(fn func(int) string, query string) string { _ = "STUB: not implemented"; return "" }

func (cfg *configuration) Validate() error { _ = "STUB: not implemented"; return nil }

// getConnectionString returns the connection string corresponding to the database connection.
func getConnectionString(cfg *configuration, isReadOnly bool) string {
	_ = "STUB: not implemented"
	return ""
}

func queryVersion(ctx context.Context, gormDB *gorm.DB, query string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func nullableDBTimeToUnixTime(dbTime *time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func nullableUnixTimeToDBTime(unixTime int64) *time.Time { _ = "STUB: not implemented"; return nil }

func lookupSimilarEntry(ctx context.Context, db *sqlDB, tx *gorm.DB, entry *common.RegistrationEntry) (*common.RegistrationEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// listRegistrationEntriesOnce returns both exact and superset matches.
// Filter out the superset matches to get an exact match

func rowsToCommonRegistrationEntries(rows *sql.Rows, entries []*common.RegistrationEntry) ([]*common.RegistrationEntry, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Due to previous bugs (i.e. #1191), there can be cruft rows related
// to a deleted registration entries that are fetched with the list
// query. To avoid hydrating partial entries, append only entries that
// have data from the registered_entries table (i.e. those with an
// entry id).

// roundedInSecondsUnix rounds the time to the nearest second, and return the time in seconds since the
// unix epoch. This function is used to avoid issues with databases versions that do not support sub-second precision.
func roundedInSecondsUnix(t time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func createCAJournal(tx *gorm.DB, caJournal *datastore.CAJournal) (*datastore.CAJournal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fetchCAJournal(tx *gorm.DB, activeX509AuthorityID string) (*datastore.CAJournal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func listCAJournalsForTesting(tx *gorm.DB) (caJournals []*datastore.CAJournal, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func updateCAJournal(tx *gorm.DB, caJournal *datastore.CAJournal) (*datastore.CAJournal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateCAJournal(caJournal *datastore.CAJournal) error { _ = "STUB: not implemented"; return nil }

func deleteCAJournal(tx *gorm.DB, caJournalID uint) error { _ = "STUB: not implemented"; return nil }

func parseDatabaseTypeASTNode(node ast.Node) (*dbTypeConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We expect the node to be *ast.ObjectList.

func isMySQLDbType(dbType string) bool { _ = "STUB: not implemented"; return false }

func isPostgresDbType(dbType string) bool { _ = "STUB: not implemented"; return false }

func isSQLiteDbType(dbType string) bool { _ = "STUB: not implemented"; return false }

func calculateResultPreallocation(pagination *datastore.Pagination) int32 {
	_ = "STUB: not implemented"
	return 0
}

// buildQuestions build list of question marks, one for each arg
// Used to build a list of args to match for in a sql IN clause in MySQl and sqlite
func buildQuestions(args []string) string { _ = "STUB: not implemented"; return "" }

// Add last question mark without trailing comma

// buildPlaceholders builds a list like $1, $2, $3...
// For use in parameterized postgres queries
func buildPlaceholders(args []string) string { _ = "STUB: not implemented"; return "" }

// buildArgs convert as slice of strings to a slice of any
func buildArgs(args []string) []any { _ = "STUB: not implemented"; return nil }
