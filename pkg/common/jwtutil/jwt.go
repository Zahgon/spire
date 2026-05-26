package jwtutil

import (
	"crypto"

	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
)

// JWTKeysFromProto converts JWT keys from the given []*types.JWTKey to map[string]crypto.PublicKey.
// The key ID of the public key is used as the key in the returned map.
func JWTKeysFromProto(proto []*types.JWTKey) (map[string]crypto.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProtoFromJWTKeys converts JWT keys from the given map[string]crypto.PublicKey to []*types.JWTKey
func ProtoFromJWTKeys(keys map[string]crypto.PublicKey) ([]*types.JWTKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
