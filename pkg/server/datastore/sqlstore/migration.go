package sqlstore

import (
	"github.com/blang/semver/v4"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire/pkg/common/version"
)

// Each time the database requires a migration, the "schema" version is
// increased and the migration code is added to this file. The migration code
// can be opportunistically removed after the following minor version has been
// released, since the supported upgrade path happens on minor version
// boundaries. For example, when 1.2 is released, the migrations that were
// handled by 1.1.x can be removed, since anyone upgrading from 1.0.X to 1.2.X
// will have to upgrade through 1.1.X first, which will apply the proper
// migrations before those done by 1.2.
//
// For convenience, the following table lists the schema versions for each
// SPIRE release, along with what was added in each schema change. SPIRE v0.6.2
// was the first version to introduce migrations.
//
// ================================================================================================
// | SPIRE   | Schema | What changed                                                              |
// ================================================================================================
// | v0.6.2  | 1      | Soft delete support was removed                                           |
// |*********|********|***************************************************************************|
// | v0.7.0  | 2      | Created join table between bundles and entries                            |
// |         |--------|---------------------------------------------------------------------------|
// |         | 3      | Normalized trust domain IDs across all tables                             |
// |         |--------|---------------------------------------------------------------------------|
// |         | 4      | Converted bundle data from DER to protobuf                                |
// |---------|        |                                                                           |
// | v0.7.1  |        |                                                                           |
// |---------|--------|---------------------------------------------------------------------------|
// | v0.7.2  | 5      | Added admin column to Entries                                             |
// |---------|        |                                                                           |
// | v0.7.3  |        |                                                                           |
// |*********|********|***************************************************************************|
// | v0.8.0  | 6      | Added downstream column to entries                                        |
// |         |--------|---------------------------------------------------------------------------|
// |         | 7      | Added expiry column to entries                                            |
// |         |--------|---------------------------------------------------------------------------|
// |         | 8      | Added dns name support for entries                                        |
// |---------|--------|---------------------------------------------------------------------------|
// | v0.8.1  | 9      | Added parent ID, SPIFFE ID and selector indices for entries               |
// |---------|--------|---------------------------------------------------------------------------|
// | v0.8.2  | 10     | Added expiry index for entries                                            |
// |         |--------|---------------------------------------------------------------------------|
// |         | 11     | Added federates with index for entries                                    |
// |---------|        |                                                                           |
// | v0.8.3  |        |                                                                           |
// |---------|        |                                                                           |
// | v0.8.4  |        |                                                                           |
// |---------|        |                                                                           |
// | v0.8.5  |        |                                                                           |
// |*********|********|***************************************************************************|
// | v0.9.0  | 12     | Added support for tracking the code version in the migration table        |
// |         |--------|---------------------------------------------------------------------------|
// |         | 13     | Added "prepared" cert columns to the attested nodes                       |
// |---------|        |                                                                           |
// | v0.9.1  |        |                                                                           |
// |---------|        |                                                                           |
// | v0.9.2  |        |                                                                           |
// |---------|        |                                                                           |
// | v0.9.3  |        |                                                                           |
// |---------|        |                                                                           |
// | v0.9.4  |        |                                                                           |
// |*********|********|***************************************************************************|
// | v0.10.0 | 14     | Added revision number column to entries                                   |
// |---------|        |                                                                           |
// | v0.10.1 |        |                                                                           |
// |---------|        |                                                                           |
// | v0.10.2 |        |                                                                           |
// |*********|        |                                                                           |
// | v0.11.0 |        |                                                                           |
// |---------|        |                                                                           |
// | v0.11.1 |        |                                                                           |
// |---------|        |                                                                           |
// | v0.11.2 |        |                                                                           |
// |---------|        |                                                                           |
// | v0.11.2 |        |                                                                           |
// |*********|********|***************************************************************************|
// | v0.12.0 | 15     | Added expiry index to attested nodes                                      |
// |---------|        |                                                                           |
// | v0.12.1 |        |                                                                           |
// |---------|        |                                                                           |
// | v0.12.2 |        |                                                                           |
// |---------|        |                                                                           |
// | v0.12.3 |        |                                                                           |
// |*********|********|***************************************************************************|
// | v1.0.0  | 16     | Added exportable identity column to entries                               |
// |         |--------|---------------------------------------------------------------------------|
// |         | 17     | Added support for Federated Trust Domains relationships                   |
// |---------|        |---------------------------------------------------------------------------|
// | v1.0.1  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.0.2  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.0.3  |        |                                                                           |
// |*********|        |                                                                           |
// | v1.1.0  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.1.1  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.1.2  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.1.3  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.1.4  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.1.5  |        |                                                                           |
// |*********|********|***************************************************************************|
// | v1.2.0  | 18     | Added hint column to entries and can_reattest column to attested nodes    |
// |---------|        |                                                                           |
// | v1.2.1  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.2.2  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.2.3  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.2.4  |        |                                                                           |
// |*********|        |                                                                           |
// | v1.3.0  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.3.1  |        |                                                                           |
// |---------|--------|---------------------------------------------------------------------------|
// | v1.3.2  | 19     | Added x509_svid_ttl and jwt_svid_ttl columns to entries                   |
// |---------|        |                                                                           |
// | v1.3.3  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.3.4  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.3.5  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.3.6  |        |                                                                           |
// |*********|********|***************************************************************************|
// | v1.4.0  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.4.1  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.4.2  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.4.3  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.4.4  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.4.5  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.4.6  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.4.7  |        |                                                                           |
// |*********|********|***************************************************************************|
// | v1.5.0  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.5.1  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.5.2  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.5.3  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.5.4  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.5.5  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.5.6  |        |                                                                           |
// |*********|********|***************************************************************************|
// | v1.6.0  | 20     | Removed x509_svid_ttl column from registered_entries                      |
// |         |--------|---------------------------------------------------------------------------|
// |         | 21     | Added index in hint column from registered_entries                        |
// |---------|        |                                                                           |
// | v1.6.1  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.6.2  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.6.3  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.6.4  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.6.5  |        |                                                                           |
// |*********|********|***************************************************************************|
// | v1.7.0  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.7.1  |        |                                                                           |
// |---------|--------|---------------------------------------------------------------------------|
// | v1.7.2  | 22     | Added registered_entries_events and attested_node_entries_events tables   |
// |---------|        |                                                                           |
// | v1.7.3  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.7.4  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.7.5  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.7.6  |        |                                                                           |
// |*********|********|***************************************************************************|
// | v1.8.0  | 23     | Added ca_journals table                                                   |
// |---------|        |                                                                           |
// | v1.8.1  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.8.2  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.8.3  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.8.4  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.8.5  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.8.6  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.8.7  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.8.8  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.8.9  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.8.10 |        |                                                                           |
// |---------|        |                                                                           |
// | v1.8.11 |        |                                                                           |
// |*********|********|***************************************************************************|
// | v1.9.0  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.9.1  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.9.2  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.9.3  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.9.4  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.9.5  |        |                                                                           |
// |---------|        |                                                                           |
// | v1.9.6  |        |                                                                           |
// |*********|********|***************************************************************************|
// | v1.10.0 |        |                                                                           |
// |---------|        |                                                                           |
// | v1.10.1 |        |                                                                           |
// |---------|        |                                                                           |
// | v1.10.2 |        |                                                                           |
// |---------|        |                                                                           |
// | v1.10.3 |        |                                                                           |
// |---------|        |                                                                           |
// | v1.10.4 |        |                                                                           |
// |*********|********|***************************************************************************|
// | v1.11.0 |        |                                                                           |
// |---------|        |                                                                           |
// | v1.11.1 |        |                                                                           |
// |---------|        |                                                                           |
// | v1.11.2 |        |                                                                           |
// |---------|        |                                                                           |
// | v1.11.3 |        |                                                                           |
// |*********|********|***************************************************************************|
// | v1.12.0 |        |                                                                           |
// | v1.12.1 |        |                                                                           |
// | v1.12.2 |        |                                                                           |
// | v1.12.3 |        |                                                                           |
// | v1.12.4 |        |                                                                           |
// | v1.12.5 |        |                                                                           |
// | v1.12.6 |        |                                                                           |
// |*********|********|***************************************************************************|
// | v1.13.0 |        |                                                                           |
// | v1.13.1 |        |                                                                           |
// | v1.13.2 |        |                                                                           |
// | v1.13.3 |        |                                                                           |
// | v1.13.4 |        |                                                                           |
// | v1.13.5 |        |                                                                           |
// |*********|********|***************************************************************************|
// | v1.14.0 |        |                                                                           |
// | v1.14.1 |        |                                                                           |
// | v1.14.2 |        |                                                                           |
// | v1.14.3 |        |                                                                           |
// | v1.14.4 |        |                                                                           |
// | v1.14.5 |        |                                                                           |
// ================================================================================================

const (
	// the latest schema version of the database in the code
	latestSchemaVersion = 25

	// lastMinorReleaseSchemaVersion is the schema version supported by the
	// last minor release. When the migrations are opportunistically pruned
	// from the code after a minor release, this number should be updated.
	lastMinorReleaseSchemaVersion = 23
)

// the current code version
var codeVersion = semver.MustParse(version.Version())

func migrateDB(db *gorm.DB, dbType string, disableMigration bool, log logrus.FieldLogger) (err error) {
	_ = "STUB: not implemented"
	// The version comparison logic in this package supports only 0.x and 1.x versioning semantics.
	// It will need to be updated prior to releasing 2.x. Ensure that we're still building a pre-2.0
	// version before continuing, and fail if we're not.
	return nil
}

// ensure migrations table exists so we can check versioning in all cases

// same DB schema; if current code version greater than stored, store newer code version

// The DB schema version can get ahead of us if the cluster is in the middle of
// an upgrade. So long as the version is compatible, log a warning and continue.
// Otherwise, we should bail out. Migration rollbacks are not supported.

// at this point:
// - auto-migration is enabled
// - schema version of DB is behind

func isDisabledMigrationAllowed(thisCodeVersion, dbCodeVersion semver.Version) error {
	_ = "STUB: not implemented"
	// If auto-migrate is disabled, and we are running a compatible version (+/- 1
	// minor from the stored code version) then we are done here
	return nil
}

func getDBCodeVersion(migration Migration) (dbCodeVersion semver.Version, err error) {
	_ = "STUB: not implemented"
	// default to 0.0.0
	return *new(semver.Version), nil
}

// we will have a blank code version from pre-0.9, and fresh, datastores

func isCompatibleCodeVersion(thisCodeVersion, dbCodeVersion semver.Version) bool {
	_ = "STUB: not implemented"
	// If major version is the same and minor version is +/- 1, versions are compatible
	return false
}

func initDB(db *gorm.DB, dbType string, log logrus.FieldLogger) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func tableOptionsForDialect(tx *gorm.DB, dbType string) *gorm.DB {
	_ = "STUB: not implemented"
	// This allows for setting table options for a particular DB type.
	// For MySQL, (for compatibility reasons) we want to make sure that
	// we can support indexes on strings (varchar(255) in the DB).
	return nil
}

func migrateVersion(tx *gorm.DB, currVersion int, log logrus.FieldLogger) (versionOut int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Place all migrations handled by the current minor release here. This
// list can be opportunistically pruned after every minor release but won't
// break things if it isn't.
//
// When adding a supported migration to version XX, add a case and the
// corresponding function. The case in the following switch statement will
// look like this:
//
// case XX:
//   err = migrateToVXX(tx)
//
// And the migrateToVXX function will be like this:
// func migrateToVXX(tx *gorm.DB) error {
//   if err := tx.AutoMigrate(&Foo{}, &Bar{}).Error; err != nil {
//     return sqlError.Wrap(err)
//   }
//   return nil
// }
//

func migrateToV24(tx *gorm.DB) error {
	_ = "STUB: not implemented"
	// Add agent_version column to attested_node_entries table
	return nil
}

func migrateToV25(tx *gorm.DB) error {
	_ = "STUB: not implemented"
	// Add additional_attributes column to registered_entries table
	return nil
}

func addFederatedRegistrationEntriesRegisteredEntryIDIndex(tx *gorm.DB) error {
	_ = "STUB: not implemented"
	// GORM creates the federated_registration_entries implicitly with a primary
	// key tuple (bundle_id, registered_entry_id). Unfortunately, MySQL5 does
	// not use the primary key index efficiently when joining by registered_entry_id
	// during registration entry list operations. We can't use gorm AutoMigrate
	// to introduce the index since there is no explicit struct to add tags to
	// so we have to manually create it.
	return nil
}
