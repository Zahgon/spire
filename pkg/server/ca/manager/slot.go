package manager

import (
	"context"
	"crypto"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/server/ca"
	"github.com/spiffe/spire/pkg/server/catalog"
	"github.com/spiffe/spire/proto/private/server/journal"
	"github.com/spiffe/spire/proto/spire/common"
)

type SlotPosition int

const (
	CurrentX509CASlot SlotPosition = iota
	NextX509CASlot
	CurrentJWTKeySlot
	NextJWTKeySlot
	CurrentWITKeySlot
	NextWITKeySlot
)

type Slot interface {
	KmKeyID() string
	IsEmpty() bool
	Reset()
	ShouldPrepareNext(now time.Time) bool
	ShouldActivateNext(now time.Time) bool
	Status() journal.Status
	UpstreamAuthorityID() string
	AuthorityID() string
	// TODO: This will be removed as part of #5390
	PublicKey() crypto.PublicKey
	NotAfter() time.Time
}

type SlotLoader struct {
	TrustDomain spiffeid.TrustDomain

	Log            logrus.FieldLogger
	Dir            string
	Catalog        catalog.Catalog
	UpstreamClient *ca.UpstreamClient
}

func (s *SlotLoader) load(ctx context.Context) (*Journal, map[SlotPosition]Slot, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Load the journal and see if we can figure out the next and current
// X509CA and JWTKey entries, if any.

// filter out local JwtKeys and X509CAs that do not exist in the database bundle

// getX509CASlots returns X509CA slots based on the status of the slots.
// - If all the statuses are unknown, the two most recent slots are returned.
// - Active entry is returned on current slot if set.
// - The most recent Prepared or Old entry is returned on next slot.
func (s *SlotLoader) getX509CASlots(ctx context.Context, entries []*journal.X509CAEntry) (*x509CASlot, *x509CASlot, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Search from oldest

// Unable to load slot
// TODO: the previous implementation analyzed only the last two entries,
// and if those slots were empty, we created new slots.
// Now we iterate through all the file, to try to get a useful slot.
// Maybe there is room for improvement here, by just verifying if the
// bundle is not expired?

// ACTIVE entry must go into current slot

// Set OLD or PREPARED as next slot
// Get the newest, since Prepared entry must always be located before an Old entry

// If both are set finish iteration

// current is set, complete next if required

// next is set but not current. swap them and initialize next with an empty slot.

// neither are set. initialize them with empty slots.

// getJWTKeysSlots returns JWTKey slots based on the status of the slots.
// - If all status are unknown, choose the two newest on the list
// - Active entry is returned on current if set
// - Newest Prepared or Old entry is returned on next
func (s *SlotLoader) getJWTKeysSlots(ctx context.Context, entries []*journal.JWTKeyEntry) (*jwtKeySlot, *jwtKeySlot, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Search from oldest

// Unable to load slot
// TODO: the previous implementation analyzed only the last two entries,
// and if those slots were empty, we created new slots.
// Now we iterate through all the file, to try to get a useful slot.
// Maybe there is room for improvement here, by just verifying if the
// bundle is not expired?

// ACTIVE entry must go into current slot

// Set OLD or PREPARED as next slot
// Get the newest, since Prepared entry must always be located before an Old entry

// If both are set finish iteration

// current is set, complete next if required

// next is set but not current. swap them and initialize next with an empty slot.

// neither are set. initialize them with empty slots.

// getWITKeysSlots returns WITKey slots based on the status of the slots.
// - If all status are unknown, choose the two newest on the list
// - Active entry is returned on current if set
// - Newest Prepared or Old entry is returned on next
func (s *SlotLoader) getWITKeysSlots(ctx context.Context, entries []*journal.WITKeyEntry) (*witKeySlot, *witKeySlot, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Search from oldest

// Unable to load slot
// TODO: the previous implementation analyzed only the last two entries,
// and if those slots were empty, we created new slots.
// Now we iterate through all the file, to try to get a useful slot.
// Maybe there is room for improvement here, by just verifying if the
// bundle is not expired?

// ACTIVE entry must go into current slot

// Set OLD or PREPARED as next slot
// Get the newest, since Prepared entry must always be located before an Old entry

// If both are set finish iteration

// current is set, complete next if required

// next is set but not current. swap them and initialize next with an empty slot.

// neither are set. initialize them with empty slots.

// filterInvalidEntries takes in a set of journal entries, and removes entries that represent signing keys
// that do not appear in the bundle from the datastore. This prevents SPIRE from entering strange
// and inconsistent states as a result of key mismatch following things like database restore,
// disk/journal manipulation, etc.
//
// If we find such a discrepancy, removing the entry from the journal prior to beginning signing
// operations prevents us from using a signing key that consumers may not be able to validate.
// Instead, we'll rotate into a new one.
func (s *SlotLoader) filterInvalidEntries(ctx context.Context, entries *journal.Entries) ([]*journal.JWTKeyEntry, []*journal.X509CAEntry, []*journal.WITKeyEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// If we have an upstream authority then we're not recovering a root CA, so we do
// not expect to find our CA certificate in the bundle. Simply proceed.

func (s *SlotLoader) fetchOptionalBundle(ctx context.Context) (*common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SlotLoader) tryLoadX509CASlotFromEntry(ctx context.Context, entry *journal.X509CAEntry) (*x509CASlot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SlotLoader) loadX509CASlotFromEntry(ctx context.Context, entry *journal.X509CAEntry) (*x509CASlot, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (s *SlotLoader) tryLoadJWTKeySlotFromEntry(ctx context.Context, entry *journal.JWTKeyEntry) (*jwtKeySlot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SlotLoader) loadJWTKeySlotFromEntry(ctx context.Context, entry *journal.JWTKeyEntry) (*jwtKeySlot, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (s *SlotLoader) tryLoadWITKeySlotFromEntry(ctx context.Context, entry *journal.WITKeyEntry) (*witKeySlot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SlotLoader) loadWITKeySlotFromEntry(ctx context.Context, entry *journal.WITKeyEntry) (*witKeySlot, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (s *SlotLoader) makeSigner(ctx context.Context, keyID string) (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}

func x509CAKmKeyID(id string) string { _ = "STUB: not implemented"; return "" }

func jwtKeyKmKeyID(id string) string { _ = "STUB: not implemented"; return "" }

func witKeyKmKeyID(id string) string { _ = "STUB: not implemented"; return "" }

func containsJwkSigningKeyID(keys []*common.PublicKey, kid string) bool {
	_ = "STUB: not implemented"
	return false
}

func containsX509CA(rootCAs []*common.Certificate, certificate []byte) bool {
	_ = "STUB: not implemented"
	return false
}

func publicKeyEqual(a, b crypto.PublicKey) bool { _ = "STUB: not implemented"; return false }

func otherSlotID(id string) string { _ = "STUB: not implemented"; return "" }

func preparationThreshold(issuedAt, notAfter time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func keyActivationThreshold(issuedAt, notAfter time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

type x509CASlot struct {
	id                  string
	issuedAt            time.Time
	x509CA              *ca.X509CA
	status              journal.Status
	authorityID         string
	publicKey           crypto.PublicKey
	notAfter            time.Time
	upstreamAuthorityID string
}

func newX509CASlot(id string) *x509CASlot { _ = "STUB: not implemented"; return nil }

func (s *x509CASlot) UpstreamAuthorityID() string { _ = "STUB: not implemented"; return "" }

func (s *x509CASlot) KmKeyID() string { _ = "STUB: not implemented"; return "" }

func (s *x509CASlot) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (s *x509CASlot) Reset() { _ = "STUB: not implemented"; return }

func (s *x509CASlot) ShouldPrepareNext(now time.Time) bool { _ = "STUB: not implemented"; return false }

func (s *x509CASlot) ShouldActivateNext(now time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *x509CASlot) Status() journal.Status {
	_ = "STUB: not implemented"
	return *new(journal.Status)
}

func (s *x509CASlot) AuthorityID() string { _ = "STUB: not implemented"; return "" }

func (s *x509CASlot) PublicKey() crypto.PublicKey {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey)
}

func (s *x509CASlot) NotAfter() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

type jwtKeySlot struct {
	id          string
	issuedAt    time.Time
	jwtKey      *ca.JWTKey
	status      journal.Status
	authorityID string
	notAfter    time.Time
}

func newJWTKeySlot(id string) *jwtKeySlot { _ = "STUB: not implemented"; return nil }

func (s *jwtKeySlot) KmKeyID() string { _ = "STUB: not implemented"; return "" }

func (s *jwtKeySlot) Status() journal.Status {
	_ = "STUB: not implemented"
	return *new(journal.Status)
}

func (s *jwtKeySlot) AuthorityID() string { _ = "STUB: not implemented"; return "" }

func (s *jwtKeySlot) UpstreamAuthorityID() string { _ = "STUB: not implemented"; return "" }

func (s *jwtKeySlot) PublicKey() crypto.PublicKey {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey)
}

func (s *jwtKeySlot) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (s *jwtKeySlot) Reset() { _ = "STUB: not implemented"; return }

func (s *jwtKeySlot) ShouldPrepareNext(now time.Time) bool { _ = "STUB: not implemented"; return false }

func (s *jwtKeySlot) ShouldActivateNext(now time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *jwtKeySlot) NotAfter() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

type witKeySlot struct {
	id          string
	issuedAt    time.Time
	witKey      *ca.WITKey
	status      journal.Status
	authorityID string
	notAfter    time.Time
}

func newWITKeySlot(id string) *witKeySlot { _ = "STUB: not implemented"; return nil }

func (s *witKeySlot) KmKeyID() string { _ = "STUB: not implemented"; return "" }

func (s *witKeySlot) Status() journal.Status {
	_ = "STUB: not implemented"
	return *new(journal.Status)
}

func (s *witKeySlot) AuthorityID() string { _ = "STUB: not implemented"; return "" }

func (s *witKeySlot) UpstreamAuthorityID() string { _ = "STUB: not implemented"; return "" }

func (s *witKeySlot) PublicKey() crypto.PublicKey {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey)
}

func (s *witKeySlot) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (s *witKeySlot) Reset() { _ = "STUB: not implemented"; return }

func (s *witKeySlot) ShouldPrepareNext(now time.Time) bool { _ = "STUB: not implemented"; return false }

func (s *witKeySlot) ShouldActivateNext(now time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *witKeySlot) NotAfter() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
