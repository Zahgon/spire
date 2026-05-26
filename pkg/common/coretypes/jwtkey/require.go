package jwtkey

import (
	plugintypes "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/types"
	"github.com/spiffe/spire/proto/spire/common"
)

func RequireFromCommonProto(pb *common.PublicKey) JWTKey {
	_ = "STUB: not implemented"
	return *new(JWTKey)
}

func RequireFromCommonProtos(pbs []*common.PublicKey) []JWTKey {
	_ = "STUB: not implemented"
	return nil
}

func RequireFromPluginProto(pb *plugintypes.JWTKey) JWTKey {
	_ = "STUB: not implemented"
	return *new(JWTKey)
}

func RequireFromPluginProtos(pbs []*plugintypes.JWTKey) []JWTKey {
	_ = "STUB: not implemented"
	return nil
}

func RequireToCommonFromPluginProto(pb *plugintypes.JWTKey) *common.PublicKey {
	_ = "STUB: not implemented"
	return nil
}

func RequireToCommonFromPluginProtos(pbs []*plugintypes.JWTKey) []*common.PublicKey {
	_ = "STUB: not implemented"
	return nil
}

func RequireToCommonProto(jwtKey JWTKey) *common.PublicKey { _ = "STUB: not implemented"; return nil }

func RequireToCommonProtos(jwtKeys []JWTKey) []*common.PublicKey {
	_ = "STUB: not implemented"
	return nil
}

func RequireToPluginFromCommonProto(pb *common.PublicKey) *plugintypes.JWTKey {
	_ = "STUB: not implemented"
	return nil
}

func RequireToPluginFromCommonProtos(pbs []*common.PublicKey) []*plugintypes.JWTKey {
	_ = "STUB: not implemented"
	return nil
}

func RequireToPluginProto(jwtKey JWTKey) *plugintypes.JWTKey { _ = "STUB: not implemented"; return nil }

func RequireToPluginProtos(jwtKeys []JWTKey) []*plugintypes.JWTKey {
	_ = "STUB: not implemented"
	return nil
}

func panicOnError(err error) { _ = "STUB: not implemented"; return }
