package debug

import (
	"context"
	"crypto/x509"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	debugv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/agent/debug/v1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/pkg/agent/manager"
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
	Clock       clock.Clock
	Log         logrus.FieldLogger
	Manager     manager.Manager
	TrustDomain spiffeid.TrustDomain
	Uptime      func() time.Duration
}

// New creates a new debug service
func New(config Config) *Service { _ = "STUB: not implemented"; return nil }

// Service implements debug server
type Service struct {
	debugv1.UnsafeDebugServer

	clock  clock.Clock
	log    logrus.FieldLogger
	m      manager.Manager
	td     spiffeid.TrustDomain
	uptime func() time.Duration

	getInfoResp getInfoResp
}

type getInfoResp struct {
	mtx  sync.Mutex
	resp *debugv1.GetInfoResponse
	ts   time.Time
}

// GetInfo gets SPIRE Agent debug information
func (s *Service) GetInfo(context.Context, *debugv1.GetInfoRequest) (*debugv1.GetInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update cache when expired or does not exist

// Get current agent's credential SVID

// Create SVID chain for response

// Reset clock and set current response

// spiffeIDFromCert gets types SPIFFE ID from certificate, it can be nil
func spiffeIDFromCert(cert *x509.Certificate) *types.SPIFFEID {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) getCertificateChain(svid []*x509.Certificate) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	// Get cached bundle
	return nil, nil
}

// Create bundle source using SVID roots, and verify certificate to extract SVID chain
