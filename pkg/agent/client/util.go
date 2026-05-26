package client

import (
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/proto/spire/common"
)

func spiffeIDFromProto(protoID *types.SPIFFEID) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func additionalAttributesFromProto(in *types.Entry_AdditionalAttributes) *common.RegistrationEntry_AdditionalAttributes {
	_ = "STUB: not implemented"
	return nil
}

func slicedEntryFromProto(e *types.Entry) (*common.RegistrationEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
