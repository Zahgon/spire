package api

import (
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/proto/spire/common"
)

// SelectorsFromProto converts a slice of types.Selector to
// a slice of common.Selector
func SelectorsFromProto(proto []*types.Selector) ([]*common.Selector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ProtoFromSelectors(in []*common.Selector) []*types.Selector {
	_ = "STUB: not implemented"
	return nil
}

func SelectorFieldFromProto(proto []*types.Selector) string { _ = "STUB: not implemented"; return "" }
