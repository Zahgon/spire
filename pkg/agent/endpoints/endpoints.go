package endpoints

import (
	"context"
	"net"

	secret_v3 "github.com/envoyproxy/go-control-plane/envoy/service/secret/v3"
	"github.com/sirupsen/logrus"
	workload_pb "github.com/spiffe/go-spiffe/v2/proto/spiffe/workload"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/spiffe/spire/pkg/common/telemetry"
)

const (
	readBufferSize = 4096
)

type Server interface {
	ListenAndServe(ctx context.Context) error
	WaitForListening(listening chan struct{})
}

type Endpoints struct {
	addr              net.Addr
	log               logrus.FieldLogger
	metrics           telemetry.Metrics
	workloadAPIServer workload_pb.SpiffeWorkloadAPIServer
	sdsv3Server       secret_v3.SecretDiscoveryServiceServer
	healthServer      grpc_health_v1.HealthServer

	hooks struct {
		listening chan struct{} // Hook to signal when the server starts listening
	}
}

func New(c Config) *Endpoints { _ = "STUB: not implemented"; return nil }

func (e *Endpoints) ListenAndServe(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Update the listening address with the actual address.
// If a TCP address was specified with port 0, this will
// update the address with the actual port that is used
// to listen.

func (e *Endpoints) triggerListeningHook() { _ = "STUB: not implemented"; return }

func (e *Endpoints) WaitForListening(listening chan struct{}) { _ = "STUB: not implemented"; return }
