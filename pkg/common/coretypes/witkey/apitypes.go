package witkey

import (
	apitypes "github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	plugintypes "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/types"
)

func ToAPIProto(witKey WITKey) (*apitypes.WITKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToAPIFromPluginProto(pb *plugintypes.WITKey) (*apitypes.WITKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
