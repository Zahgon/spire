package jwtkey

import (
	"crypto"
	"time"
)

type JWTKey struct {
	ID        string
	PublicKey crypto.PublicKey
	ExpiresAt time.Time
	Tainted   bool
}

func toProtoFields(jwtKey JWTKey) (id string, publicKey []byte, expiresAt int64, tainted bool, err error) {
	_ = "STUB: not implemented"
	return "", nil, 0, false, nil
}

func fromProtoFields(keyID string, publicKeyPKIX []byte, expiresAtUnix int64, tainted bool) (JWTKey, error) {
	_ = "STUB: not implemented"
	return *new(JWTKey), nil
}
