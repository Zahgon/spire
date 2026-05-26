package health

import (
	"context"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/server/datastore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// RegisterService registers the service on the gRPC server.
func RegisterService(s grpc.ServiceRegistrar, service *Service) { _ = "STUB: not implemented"; return }

// Config is the service configuration
type Config struct {
	TrustDomain spiffeid.TrustDomain
	DataStore   datastore.DataStore
}

// New creates a new Health service
func New(config Config) *Service { _ = "STUB: not implemented"; return nil }

// Service implements the v1 Health service
type Service struct {
	grpc_health_v1.UnimplementedHealthServer

	ds datastore.DataStore
	td spiffeid.TrustDomain
}

func (s *Service) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Ensure per-service health is not being requested.
}
