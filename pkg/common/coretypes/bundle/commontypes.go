package bundle

import (
	plugintypes "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/types"
	"github.com/spiffe/spire/proto/spire/common"
)

func ToCommonFromPluginProto(pb *plugintypes.Bundle) (*common.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
