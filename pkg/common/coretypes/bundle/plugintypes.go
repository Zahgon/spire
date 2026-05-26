package bundle

import (
	apitypes "github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	plugintypes "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/types"
	"github.com/spiffe/spire/proto/spire/common"
)

func ToPluginFromAPIProto(pb *apitypes.Bundle) (*plugintypes.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToPluginProtoFromCommon(b *common.Bundle) (*plugintypes.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
