package jwtkey

import (
	apitypes "github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	plugintypes "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/types"
	"github.com/spiffe/spire/proto/spire/common"
)

func FromPluginProto(pb *plugintypes.JWTKey) (JWTKey, error) {
	_ = "STUB: not implemented"
	return *new(JWTKey), nil
}

func FromPluginProtos(pbs []*plugintypes.JWTKey) ([]JWTKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToPluginProto(jwtKey JWTKey) (*plugintypes.JWTKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToPluginProtos(jwtKeys []JWTKey) ([]*plugintypes.JWTKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToPluginFromCommonProto(pb *common.PublicKey) (*plugintypes.JWTKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToPluginFromCommonProtos(pbs []*common.PublicKey) ([]*plugintypes.JWTKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToPluginFromAPIProto(pb *apitypes.JWTKey) (*plugintypes.JWTKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToPluginFromAPIProtos(pbs []*apitypes.JWTKey) ([]*plugintypes.JWTKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
