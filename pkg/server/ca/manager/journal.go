package manager

import (
	"context"
	"crypto/x509"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire/pkg/server/ca"
	"github.com/spiffe/spire/pkg/server/catalog"
	"github.com/spiffe/spire/pkg/server/datastore"
	"github.com/spiffe/spire/proto/private/server/journal"
)

const (
	// journalCap is the maximum number of entries per type that we'll
	// hold onto.
	journalCap = 10
)

type journalConfig struct {
	cat catalog.Catalog
	log logrus.FieldLogger
}

// Journal stores X509 CAs, JWT keys, and WIT keys in the datastore as they are
// rotated by the manager.
type Journal struct {
	config *journalConfig

	mu                    sync.RWMutex
	activeX509AuthorityID string
	caJournalID           uint
	entries               *journal.Entries
}

func LoadJournal(ctx context.Context, config *journalConfig) (*Journal, error) {
	_ = "STUB: not implemented"
	// Look for the CA journal of this server in the datastore.
	return nil, nil
}

func (j *Journal) getEntries() *journal.Entries { _ = "STUB: not implemented"; return nil }

func (j *Journal) AppendX509CA(ctx context.Context, slotID string, issuedAt time.Time, x509CA *ca.X509CA) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateX509CAStatus updates a stored X509CA entry to have the given status,
// updating the CA journal.
func (j *Journal) UpdateX509CAStatus(ctx context.Context, authorityID string, status journal.Status) error {
	_ = "STUB: not implemented"
	return nil
}

func (j *Journal) AppendJWTKey(ctx context.Context, slotID string, issuedAt time.Time, jwtKey *ca.JWTKey) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateJWTKeyStatus updates a stored JWTKey entry to have the given status,
// updating the CA journal.
func (j *Journal) UpdateJWTKeyStatus(ctx context.Context, authorityID string, status journal.Status) error {
	_ = "STUB: not implemented"
	return nil
}

func (j *Journal) AppendWITKey(ctx context.Context, slotID string, issuedAt time.Time, witKey *ca.WITKey) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateWITKeyStatus updates a stored WITKey entry to have the given status,
// updating the CA journal.
func (j *Journal) UpdateWITKeyStatus(ctx context.Context, authorityID string, status journal.Status) error {
	_ = "STUB: not implemented"
	return nil
}

func (j *Journal) setEntries(entries *journal.Entries) { _ = "STUB: not implemented"; return }

// saveInDatastore saves the provided marshaled entries in the datastore.
// If caJournalID has not been defined yet (its value is 0), it first finds
// the CA journal record that corresponds to this server. In case there is no
// CA record for this server, it creates one.
// The ID of the CA journal record that was saved is returned, in addition to
// the error (if any) of the operation.
func (j *Journal) saveInDatastore(ctx context.Context, entriesBytes []byte) (caJournalID uint, err error) {
	_ = "STUB: not implemented"
	// Check if we already identified what's the CA journal for this server in
	// the datastore. If not, log that we are creating a new CA journal entry.
	return 0, nil
}

// findCAJournal finds the corresponding CA journal record in the datastore for
// this server. It does that by retrieving all the public keys managed by the
// KeyManager and trying to get a match with a record which last active
// X509 authority ID correspond to one of the keys.
func (j *Journal) findCAJournal(ctx context.Context) (*datastore.CAJournal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get all the public keys managed by the KeyManager.

// There is a CA journal record that has an active X509 authority
// ID that matches with one of the public keys of this server. This
// means that this record belongs to this server.

// save saves the CA journal in the datastore.
func (j *Journal) save(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func chainDER(chain []*x509.Certificate) [][]byte { _ = "STUB: not implemented"; return nil }

// loadJournalFromDS loads the CA journal from the datastore.
// It does that by looking for a CA journal record that matches with one of the
// public keys of this server.
func loadJournalFromDS(ctx context.Context, config *journalConfig) (*Journal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
