package api

import (
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/pkg/server/datastore"
)

// ProtoToFederationRelationship convert and validate proto to datastore federated relationship
func ProtoToFederationRelationship(f *types.FederationRelationship) (*datastore.FederationRelationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProtoToFederationRelationshipWithMask convert and validate proto to datastore federated relationship, and apply mask
func ProtoToFederationRelationshipWithMask(f *types.FederationRelationship, mask *types.FederationRelationshipMask) (*datastore.FederationRelationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FederationRelationshipToProto converts datastore federation relationship to types proto
func FederationRelationshipToProto(f *datastore.FederationRelationship, mask *types.FederationRelationshipMask) (*types.FederationRelationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
