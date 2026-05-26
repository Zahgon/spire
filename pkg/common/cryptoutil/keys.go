package cryptoutil

import (
	"crypto"

	"github.com/go-jose/go-jose/v4"
)

func PublicKeyEqual(a, b crypto.PublicKey) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func KeyMatches(privateKey crypto.PrivateKey, publicKey crypto.PublicKey) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func JoseAlgFromPublicKey(publicKey any) (jose.SignatureAlgorithm, error) {
	_ = "STUB: not implemented"
	return *new(jose.SignatureAlgorithm), nil
}

// Prevent the use of keys smaller than 2048 bits
