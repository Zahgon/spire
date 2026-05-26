package api

import (
	"context"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
)

func TrustDomainMemberIDFromProto(ctx context.Context, td spiffeid.TrustDomain, protoID *types.SPIFFEID) (spiffeid.ID, error) {
	_ = "STUB: not implemented"
	return *new(spiffeid.ID), nil
}

func VerifyTrustDomainMemberID(td spiffeid.TrustDomain, id spiffeid.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func TrustDomainAgentIDFromProto(ctx context.Context, td spiffeid.TrustDomain, protoID *types.SPIFFEID) (spiffeid.ID, error) {
	_ = "STUB: not implemented"
	return *new(spiffeid.ID), nil
}

func VerifyTrustDomainAgentID(td spiffeid.TrustDomain, id spiffeid.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func VerifyTrustDomainAgentIDForNodeAttestor(td spiffeid.TrustDomain, id spiffeid.ID, nodeAttestorName string) error {
	_ = "STUB: not implemented"
	return nil
}

func TrustDomainWorkloadIDFromProto(ctx context.Context, td spiffeid.TrustDomain, protoID *types.SPIFFEID) (spiffeid.ID, error) {
	_ = "STUB: not implemented"
	return *new(spiffeid.ID), nil
}

func VerifyTrustDomainWorkloadID(td spiffeid.TrustDomain, id spiffeid.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// ProtoFromID converts a SPIFFE ID from the given spiffeid.ID to
// types.SPIFFEID
func ProtoFromID(id spiffeid.ID) *types.SPIFFEID { _ = "STUB: not implemented"; return nil }

// IDFromProto converts a SPIFFEID message into an ID type
func IDFromProto(_ context.Context, protoID *types.SPIFFEID) (spiffeid.ID, error) {
	_ = "STUB: not implemented"
	return *new(spiffeid.ID), nil
}
