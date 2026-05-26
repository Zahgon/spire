package debug

import (
	"context"
	"crypto/x509"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	debugv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/debug/v1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/pkg/server/datastore"
	"github.com/spiffe/spire/pkg/server/svid"
	"github.com/spiffe/spire/test/clock"
	"google.golang.org/grpc"
)

const (
	cacheExpiry = 5 * time.Second
)

// RegisterService registers debug service on provided server
func RegisterService(s grpc.ServiceRegistrar, service *Service) { _ = "STUB: not implemented"; return }

// Config configurations for debug service
type Config struct {
	Clock        clock.Clock
	DataStore    datastore.DataStore
	SVIDObserver svid.Observer
	TrustDomain  spiffeid.TrustDomain
	Uptime       func() time.Duration
}

// New creates a new debug service
func New(config Config) *Service { _ = "STUB: not implemented"; return nil }

// Service implements debug server
type Service struct {
	debugv1.UnsafeDebugServer

	clock  clock.Clock
	ds     datastore.DataStore
	so     svid.Observer
	td     spiffeid.TrustDomain
	uptime func() time.Duration

	getInfoResp getInfoResp
}

type getInfoResp struct {
	mtx  sync.Mutex
	resp *debugv1.GetInfoResponse
	ts   time.Time
}

// GetInfo gets SPIRE Server debug information
func (s *Service) GetInfo(ctx context.Context, _ *debugv1.GetInfoRequest) (*debugv1.GetInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update cache when expired or does not exist

// Reset clock and set current response

func (s *Service) getCertificateChain(ctx context.Context, log logrus.FieldLogger) ([]*debugv1.GetInfoResponse_Cert, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Extract trustdomains bundle and append federated bundles

// Create bundle source using rootCAs

// Verify certificate to extract SVID chain

// Create SVID chain for response

// spiffeIDFromCert gets types SPIFFE ID from certificate, it can be nil
func spiffeIDFromCert(cert *x509.Certificate) *types.SPIFFEID {
	_ = "STUB: not implemented"
	return nil
}
