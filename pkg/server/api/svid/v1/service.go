package svid

import (
	"context"
	"crypto"
	"crypto/x509"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	svidv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/svid/v1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/pkg/server/api"
	"github.com/spiffe/spire/pkg/server/ca"
	"github.com/spiffe/spire/pkg/server/datastore"
	"google.golang.org/grpc"
)

var (
	supportedRSAWITSigningAlgorithms = []string{"RS256", "RS384", "RS512", "PS256", "PS384", "PS512"}
)

// RegisterService registers the service on the gRPC server.
func RegisterService(s grpc.ServiceRegistrar, service *Service) { _ = "STUB: not implemented"; return }

// Config is the service configuration
type Config struct {
	EntryFetcher api.AuthorizedEntryFetcher
	ServerCA     ca.ServerCA
	TrustDomain  spiffeid.TrustDomain
	DataStore    datastore.DataStore
}

// New creates a new SVID service
func New(config Config) *Service { _ = "STUB: not implemented"; return nil }

// Service implements the v1 SVID service
type Service struct {
	svidv1.UnsafeSVIDServer

	ca                           ca.ServerCA
	ef                           api.AuthorizedEntryFetcher
	td                           spiffeid.TrustDomain
	ds                           datastore.DataStore
	useLegacyDownstreamX509CATTL bool
}

func (s *Service) MintX509SVID(ctx context.Context, req *svidv1.MintX509SVIDRequest) (*svidv1.MintX509SVIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) MintJWTSVID(ctx context.Context, req *svidv1.MintJWTSVIDRequest) (*svidv1.MintJWTSVIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) MintWITSVID(ctx context.Context, req *svidv1.MintWITSVIDRequest) (*svidv1.MintWITSVIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) BatchNewX509SVID(ctx context.Context, req *svidv1.BatchNewX509SVIDRequest) (*svidv1.BatchNewX509SVIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fetch authorized entries

//  Create new SVID

func (s *Service) findEntries(ctx context.Context, log logrus.FieldLogger, entries map[string]struct{}) (map[string]api.ReadOnlyEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newX509SVID creates an X509-SVID using data from registration entry and key from CSR
func (s *Service) newX509SVID(ctx context.Context, param *svidv1.NewX509SVIDParams, entries map[string]api.ReadOnlyEntry) *svidv1.BatchNewX509SVIDResponse_Result {
	_ = "STUB: not implemented"
	return nil
}

// This shouldn't be the case unless there is invalid data in the datastore

func (s *Service) mintJWTSVID(ctx context.Context, protoID *types.SPIFFEID, audience []string, ttl int32) (*types.JWTSVID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) NewJWTSVID(ctx context.Context, req *svidv1.NewJWTSVIDRequest) (resp *svidv1.NewJWTSVIDResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fetch authorized entries

func (s *Service) BatchNewWITSVID(ctx context.Context, req *svidv1.BatchNewWITSVIDRequest) (*svidv1.BatchNewWITSVIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fetch authorized entries

//  Create new SVID

func (s *Service) mintWITSVID(ctx context.Context, protoID *types.SPIFFEID, publicKeyDer []byte, signingAlgorithm string, ttl int32) (*types.WITSVID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newWITSVID creates an WIT-SVID using data from registration entry and public key from input params
func (s *Service) newWITSVID(ctx context.Context, param *svidv1.NewWITSVIDParams, entries map[string]api.ReadOnlyEntry) *svidv1.BatchNewWITSVIDResponse_Result {
	_ = "STUB: not implemented"
	return nil
}

// This shouldn't be the case unless there is invalid data in the datastore

// TODO: add WIT specific TTL (https://github.com/spiffe/spire/issues/6535)

func (s *Service) NewDownstreamX509CA(ctx context.Context, req *svidv1.NewDownstreamX509CARequest) (*svidv1.NewDownstreamX509CAResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use the TTL offered by the downstream server (if any), unless we are
// configured to use the legacy TTL.

// Legacy downstream TTL prefers the downstream workload entry
// TTL (if any) and then the default workload TTL. We'll handle the
// latter inside of the credbuilder package, which already has
// knowledge of the default.

func (s *Service) isJWTSVIDsDisabled() bool { _ = "STUB: not implemented"; return false }

func (s *Service) isWITSVIDsDisabled() bool { _ = "STUB: not implemented"; return false }

func (s Service) fieldsFromJWTSvidParams(ctx context.Context, protoID *types.SPIFFEID, audience []string, ttl int32) logrus.Fields {
	_ = "STUB: not implemented"
	return *new(logrus.Fields)
}

// Don't care about parsing error

func (s Service) fieldsFromWITSvidParams(ctx context.Context, protoID *types.SPIFFEID, ttl int32) logrus.Fields {
	_ = "STUB: not implemented"
	return *new(logrus.Fields)
}

// Don't care about parsing error

func parseAndCheckCSR(ctx context.Context, csrBytes []byte) (*x509.CertificateRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ValidatePublicKeyAndSigningAlgorithm(publicKey crypto.PublicKey, signingAlgorithm string) error {
	_ = "STUB: not implemented"
	return nil
}
