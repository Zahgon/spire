package hashicorpvault

import (
	"time"

	"github.com/spiffe/spire/pkg/server/common/vault"
)

// uuidStringLength is the length of a UUID string (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx).
const uuidStringLength = 36

// getKeyEntry converts a Vault transit KeyEntry into a keymanager keyEntry,
// parsing the public key and determining the SPIRE key type from the top-level
// Vault key type. Returns (nil, false, nil) if the key belongs to a different
// server instance (i.e. the key name does not match the expected serverID prefix).
func getKeyEntry(ve *vault.KeyEntry, serverID string) (*keyEntry, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// parseKeyCreationTime extracts and parses the creation_time field from Vault key data.
func parseKeyCreationTime(keyData map[string]any) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// spireKeyIDFromKeyName parses a Vault transit key name to extract the SPIRE Key ID.
// Key names have the format <SERVER-ID>-<UUID>-<SPIRE-KEY-ID>.
// Returns ("", false) if the key name does not match the expected serverID prefix or format.
func spireKeyIDFromKeyName(keyName, serverID string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// rest must be at least "<UUID>-<one-char-spireKeyID>"
