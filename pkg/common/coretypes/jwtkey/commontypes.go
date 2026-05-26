package jwtkey

import (
	plugintypes "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/types"
	"github.com/spiffe/spire/proto/spire/common"
)

func FromCommonProto(pb *common.PublicKey) (JWTKey, error) {
	_ = "STUB: not implemented"
	return *new(JWTKey), nil
}

func FromCommonProtos(pbs []*common.PublicKey) ([]JWTKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToCommonProto(jwtKey JWTKey) (*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToCommonProtos(jwtKeys []JWTKey) ([]*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToCommonFromPluginProto(pb *plugintypes.JWTKey) (*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToCommonFromPluginProtos(pbs []*plugintypes.JWTKey) ([]*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
