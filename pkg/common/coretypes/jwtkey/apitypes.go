package jwtkey

import (
	apitypes "github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	plugintypes "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/types"
)

func ToAPIProto(jwtKey JWTKey) (*apitypes.JWTKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToAPIFromPluginProto(pb *plugintypes.JWTKey) (*apitypes.JWTKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
