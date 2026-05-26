package witkey

import (
	plugintypes "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/types"
	"github.com/spiffe/spire/proto/spire/common"
)

func RequireFromCommonProto(pb *common.PublicKey) WITKey {
	_ = "STUB: not implemented"
	return *new(WITKey)
}

func RequireFromCommonProtos(pbs []*common.PublicKey) []WITKey {
	_ = "STUB: not implemented"
	return nil
}

func RequireFromPluginProto(pb *plugintypes.WITKey) WITKey {
	_ = "STUB: not implemented"
	return *new(WITKey)
}

func RequireFromPluginProtos(pbs []*plugintypes.WITKey) []WITKey {
	_ = "STUB: not implemented"
	return nil
}

func RequireToCommonFromPluginProto(pb *plugintypes.WITKey) *common.PublicKey {
	_ = "STUB: not implemented"
	return nil
}

func RequireToCommonFromPluginProtos(pbs []*plugintypes.WITKey) []*common.PublicKey {
	_ = "STUB: not implemented"
	return nil
}

func RequireToCommonProto(witKey WITKey) *common.PublicKey { _ = "STUB: not implemented"; return nil }

func RequireToCommonProtos(witKeys []WITKey) []*common.PublicKey {
	_ = "STUB: not implemented"
	return nil
}

func RequireToPluginFromCommonProto(pb *common.PublicKey) *plugintypes.WITKey {
	_ = "STUB: not implemented"
	return nil
}

func RequireToPluginFromCommonProtos(pbs []*common.PublicKey) []*plugintypes.WITKey {
	_ = "STUB: not implemented"
	return nil
}

func RequireToPluginProto(witKey WITKey) *plugintypes.WITKey { _ = "STUB: not implemented"; return nil }

func RequireToPluginProtos(witKeys []WITKey) []*plugintypes.WITKey {
	_ = "STUB: not implemented"
	return nil
}

func panicOnError(err error) { _ = "STUB: not implemented"; return }
