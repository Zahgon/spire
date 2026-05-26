package witkey

import (
	plugintypes "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/types"
	"github.com/spiffe/spire/proto/spire/common"
)

func FromCommonProto(pb *common.PublicKey) (WITKey, error) {
	_ = "STUB: not implemented"
	return *new(WITKey), nil
}

func FromCommonProtos(pbs []*common.PublicKey) ([]WITKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToCommonProto(witKey WITKey) (*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToCommonProtos(witKeys []WITKey) ([]*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToCommonFromPluginProto(pb *plugintypes.WITKey) (*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToCommonFromPluginProtos(pbs []*plugintypes.WITKey) ([]*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
