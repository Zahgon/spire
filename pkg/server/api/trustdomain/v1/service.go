package trustdomain

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/server/datastore"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	trustdomainv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/trustdomain/v1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
)

// BundleRefresher is used by the service to refresh bundles.
type BundleRefresher interface {
	// TriggerConfigReload triggers the refresher to reload it's configuration
	TriggerConfigReload()

	// RefreshBundleFor refreshes the bundle for the given trust domain.
	RefreshBundleFor(ctx context.Context, td spiffeid.TrustDomain) (bool, error)
}

// Config is the service configuration.
type Config struct {
	DataStore       datastore.DataStore
	TrustDomain     spiffeid.TrustDomain
	BundleRefresher BundleRefresher
}

// Service implements the v1 trustdomain service.
type Service struct {
	trustdomainv1.UnsafeTrustDomainServer

	ds datastore.DataStore
	td spiffeid.TrustDomain
	br BundleRefresher
}

// New creates a new trustdomain service.
func New(config Config) *Service { _ = "STUB: not implemented"; return nil }

// RegisterService registers the trustdomain service on the gRPC server.
func RegisterService(s grpc.ServiceRegistrar, service *Service) { _ = "STUB: not implemented"; return }

func (s *Service) ListFederationRelationships(ctx context.Context, req *trustdomainv1.ListFederationRelationshipsRequest) (*trustdomainv1.ListFederationRelationshipsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) GetFederationRelationship(ctx context.Context, req *trustdomainv1.GetFederationRelationshipRequest) (*types.FederationRelationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if the entry is not found, FetchFederationRelationship returns nil, nil

func (s *Service) BatchCreateFederationRelationship(ctx context.Context, req *trustdomainv1.BatchCreateFederationRelationshipRequest) (*trustdomainv1.BatchCreateFederationRelationshipResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) BatchUpdateFederationRelationship(ctx context.Context, req *trustdomainv1.BatchUpdateFederationRelationshipRequest) (*trustdomainv1.BatchUpdateFederationRelationshipResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) BatchDeleteFederationRelationship(ctx context.Context, req *trustdomainv1.BatchDeleteFederationRelationshipRequest) (*trustdomainv1.BatchDeleteFederationRelationshipResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) RefreshBundle(ctx context.Context, req *trustdomainv1.RefreshBundleRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) createFederationRelationship(ctx context.Context, f *types.FederationRelationship, outputMask *types.FederationRelationshipMask) *trustdomainv1.BatchCreateFederationRelationshipResponse_Result {
	_ = "STUB: not implemented"
	return nil
}

// Warning in case of SPIFFE endpoint that does not have a bundle

func (s *Service) updateFederationRelationship(ctx context.Context, fr *types.FederationRelationship, inputMask *types.FederationRelationshipMask, outputMask *types.FederationRelationshipMask) *trustdomainv1.BatchUpdateFederationRelationshipResponse_Result {
	_ = "STUB: not implemented"
	return nil
}

// Warning in case of SPIFFE endpoint that does not have a bundle

func (s *Service) deleteFederationRelationship(ctx context.Context, td string) *trustdomainv1.BatchDeleteFederationRelationshipResponse_Result {
	_ = "STUB: not implemented"
	return nil
}

func fieldsFromRelationshipProto(proto *types.FederationRelationship, mask *types.FederationRelationshipMask) logrus.Fields {
	_ = "STUB: not implemented"
	return *new(logrus.Fields)
}

func validateEndpointBundle(ctx context.Context, ds datastore.DataStore, log logrus.FieldLogger, endpointSPIFFEID spiffeid.ID) {
	_ = "STUB: not implemented"
	return
}

// Bundle is nil when not found
