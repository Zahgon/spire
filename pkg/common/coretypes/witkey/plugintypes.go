package witkey

import (
	apitypes "github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	plugintypes "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/types"
	"github.com/spiffe/spire/proto/spire/common"
)

func FromPluginProto(pb *plugintypes.WITKey) (WITKey, error) {
	_ = "STUB: not implemented"
	return *new(WITKey), nil
}

func FromPluginProtos(pbs []*plugintypes.WITKey) ([]WITKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToPluginProto(witKey WITKey) (*plugintypes.WITKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToPluginProtos(witKeys []WITKey) ([]*plugintypes.WITKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToPluginFromCommonProto(pb *common.PublicKey) (*plugintypes.WITKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToPluginFromCommonProtos(pbs []*common.PublicKey) ([]*plugintypes.WITKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToPluginFromAPIProto(pb *apitypes.WITKey) (*plugintypes.WITKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToPluginFromAPIProtos(pbs []*apitypes.WITKey) ([]*plugintypes.WITKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
