package bundle

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	bundlev1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/bundle/v1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/pkg/server/datastore"
	"github.com/spiffe/spire/proto/spire/common"
	"google.golang.org/grpc"
)

// UpstreamPublisher defines the publisher interface.
type UpstreamPublisher interface {
	PublishJWTKey(ctx context.Context, jwtKey *common.PublicKey) ([]*common.PublicKey, error)
}

// UpstreamPublisherFunc defines the function.
type UpstreamPublisherFunc func(ctx context.Context, jwtKey *common.PublicKey) ([]*common.PublicKey, error)

// PublishJWTKey publishes the JWT key with the given function.
func (fn UpstreamPublisherFunc) PublishJWTKey(ctx context.Context, jwtKey *common.PublicKey) ([]*common.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil,

		// Config defines the bundle service configuration.
		nil
}

type Config struct {
	DataStore         datastore.DataStore
	TrustDomain       spiffeid.TrustDomain
	UpstreamPublisher UpstreamPublisher
}

// Service defines the v1 bundle service properties.
type Service struct {
	bundlev1.UnsafeBundleServer

	ds datastore.DataStore
	td spiffeid.TrustDomain
	up UpstreamPublisher
}

// New creates a new bundle service.
func New(config Config) *Service { _ = "STUB: not implemented"; return nil }

// RegisterService registers the bundle service on the gRPC server.
func RegisterService(s grpc.ServiceRegistrar, service *Service) { _ = "STUB: not implemented"; return }

// CountBundles returns the total number of bundles.
func (s *Service) CountBundles(ctx context.Context, _ *bundlev1.CountBundlesRequest) (*bundlev1.CountBundlesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBundle returns the bundle associated with the given trust domain.
func (s *Service) GetBundle(ctx context.Context, req *bundlev1.GetBundleRequest) (*types.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AppendBundle appends the given authorities to the given bundlev1.
func (s *Service) AppendBundle(ctx context.Context, req *bundlev1.AppendBundleRequest) (*types.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PublishJWTAuthority published the JWT key on the server.
func (s *Service) PublishJWTAuthority(ctx context.Context, req *bundlev1.PublishJWTAuthorityRequest) (*bundlev1.PublishJWTAuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PublishWITAuthority published the WIT key on the server.
func (s *Service) PublishWITAuthority(ctx context.Context, req *bundlev1.PublishWITAuthorityRequest) (*bundlev1.PublishWITAuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListFederatedBundles returns an optionally paginated list of federated bundles.
func (s *Service) ListFederatedBundles(ctx context.Context, req *bundlev1.ListFederatedBundlesRequest) (*bundlev1.ListFederatedBundlesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set pagination parameters

// Filter server bundle

// GetFederatedBundle returns the bundle associated with the given trust domain.
func (s *Service) GetFederatedBundle(ctx context.Context, req *bundlev1.GetFederatedBundleRequest) (*types.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BatchCreateFederatedBundle adds one or more bundles to the server.
func (s *Service) BatchCreateFederatedBundle(ctx context.Context, req *bundlev1.BatchCreateFederatedBundleRequest) (*bundlev1.BatchCreateFederatedBundleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) createFederatedBundle(ctx context.Context, b *types.Bundle, outputMask *types.BundleMask) *bundlev1.BatchCreateFederatedBundleResponse_Result {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) setFederatedBundle(ctx context.Context, b *types.Bundle, outputMask *types.BundleMask) *bundlev1.BatchSetFederatedBundleResponse_Result {
	_ = "STUB: not implemented"
	return nil
}

// BatchUpdateFederatedBundle updates one or more bundles in the server.
func (s *Service) BatchUpdateFederatedBundle(ctx context.Context, req *bundlev1.BatchUpdateFederatedBundleRequest) (*bundlev1.BatchUpdateFederatedBundleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) updateFederatedBundle(ctx context.Context, b *types.Bundle, inputMask, outputMask *types.BundleMask) *bundlev1.BatchUpdateFederatedBundleResponse_Result {
	_ = "STUB: not implemented"
	return nil
}

// BatchSetFederatedBundle upserts one or more bundles in the server.
func (s *Service) BatchSetFederatedBundle(ctx context.Context, req *bundlev1.BatchSetFederatedBundleRequest) (*bundlev1.BatchSetFederatedBundleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BatchDeleteFederatedBundle removes one or more bundles from the server.
func (s *Service) BatchDeleteFederatedBundle(ctx context.Context, req *bundlev1.BatchDeleteFederatedBundleRequest) (*bundlev1.BatchDeleteFederatedBundleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) deleteFederatedBundle(ctx context.Context, log logrus.FieldLogger, trustDomain string, mode datastore.DeleteMode) *bundlev1.BatchDeleteFederatedBundleResponse_Result {
	_ = "STUB: not implemented"
	return nil
}

func parseDeleteMode(mode bundlev1.BatchDeleteFederatedBundleRequest_Mode) (datastore.DeleteMode, error) {
	_ = "STUB: not implemented"
	return *new(datastore.DeleteMode), nil
}

func applyBundleMask(b *types.Bundle, mask *types.BundleMask) { _ = "STUB: not implemented"; return }
