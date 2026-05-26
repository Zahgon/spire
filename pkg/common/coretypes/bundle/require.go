package bundle

import (
	plugintypes "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/types"
	"github.com/spiffe/spire/proto/spire/common"
)

func RequireToCommonFromPluginProto(pb *plugintypes.Bundle) *common.Bundle {
	_ = "STUB: not implemented"
	return nil
}

func panicOnError(err error) { _ = "STUB: not implemented"; return }
