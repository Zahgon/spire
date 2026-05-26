package idutil

import (
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
)

// IDProtoString constructs a SPIFFE ID string for the given ID protobuf.
func IDProtoString(id *types.SPIFFEID) (string, error) { _ = "STUB: not implemented"; return "", nil }

// IDProtoFromString parses a SPIFFE ID string into the raw ID proto components.
// It does not attempt to escape/unescape any portion of the ID.
func IDProtoFromString(s string) (*types.SPIFFEID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IDFromProto returns SPIFFE ID from the proto representation
func IDFromProto(id *types.SPIFFEID) (spiffeid.ID, error) {
	_ = "STUB: not implemented"
	return *new(spiffeid.ID), nil
}
