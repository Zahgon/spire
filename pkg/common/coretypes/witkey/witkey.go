package witkey

import (
	"crypto"
	"time"
)

type WITKey struct {
	ID        string
	PublicKey crypto.PublicKey
	ExpiresAt time.Time
	Tainted   bool
}

func toProtoFields(witKey WITKey) (id string, publicKey []byte, expiresAt int64, tainted bool, err error) {
	_ = "STUB: not implemented"
	return "", nil, 0, false, nil
}

func fromProtoFields(keyID string, publicKeyPKIX []byte, expiresAtUnix int64, tainted bool) (WITKey, error) {
	_ = "STUB: not implemented"
	return *new(WITKey), nil
}
