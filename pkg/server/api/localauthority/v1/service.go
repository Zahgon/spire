package localauthority

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	localauthorityv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/localauthority/v1"
	"github.com/spiffe/spire/pkg/server/ca/manager"
	"github.com/spiffe/spire/pkg/server/datastore"
	"google.golang.org/grpc"
)

type CAManager interface {
	// JWT
	GetCurrentJWTKeySlot() manager.Slot
	GetNextJWTKeySlot() manager.Slot
	PrepareJWTKey(ctx context.Context) error
	RotateJWTKey(ctx context.Context)
	IsJWTSVIDsDisabled() bool

	// X509
	GetCurrentX509CASlot() manager.Slot
	GetNextX509CASlot() manager.Slot
	PrepareX509CA(ctx context.Context) error
	RotateX509CA(ctx context.Context)

	IsUpstreamAuthority() bool
	NotifyTaintedX509Authority(ctx context.Context, authorityID string) error
}

// RegisterService registers the service on the gRPC server.
func RegisterService(s grpc.ServiceRegistrar, service *Service) { _ = "STUB: not implemented"; return }

// Config is the service configuration
type Config struct {
	TrustDomain spiffeid.TrustDomain
	DataStore   datastore.DataStore
	CAManager   CAManager
}

// New creates a new LocalAuthority service
func New(config Config) *Service { _ = "STUB: not implemented"; return nil }

// Service implements the v1 LocalAuthority service
type Service struct {
	localauthorityv1.UnsafeLocalAuthorityServer

	td spiffeid.TrustDomain
	ds datastore.DataStore
	ca CAManager
}

func (s *Service) GetJWTAuthorityState(ctx context.Context, _ *localauthorityv1.GetJWTAuthorityStateRequest) (*localauthorityv1.GetJWTAuthorityStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// when next has a key indicates that it was initialized

func (s *Service) PrepareJWTAuthority(ctx context.Context, _ *localauthorityv1.PrepareJWTAuthorityRequest) (*localauthorityv1.PrepareJWTAuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) ActivateJWTAuthority(ctx context.Context, req *localauthorityv1.ActivateJWTAuthorityRequest) (*localauthorityv1.ActivateJWTAuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Authority ID is required

/// Only next local authority can be Activated

// Only PREPARED local authorities can be Activated

func (s *Service) TaintJWTAuthority(ctx context.Context, req *localauthorityv1.TaintJWTAuthorityRequest) (*localauthorityv1.TaintJWTAuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Authority ID is required

// It is not possible to taint Active authority

// Only next local authority can be tainted

// Only OLD authorities can be tainted

func (s *Service) RevokeJWTAuthority(ctx context.Context, req *localauthorityv1.RevokeJWTAuthorityRequest) (*localauthorityv1.RevokeJWTAuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) GetX509AuthorityState(ctx context.Context, _ *localauthorityv1.GetX509AuthorityStateRequest) (*localauthorityv1.GetX509AuthorityStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// when next has a key indicates that it was initialized

func (s *Service) PrepareX509Authority(ctx context.Context, _ *localauthorityv1.PrepareX509AuthorityRequest) (*localauthorityv1.PrepareX509AuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) ActivateX509Authority(ctx context.Context, req *localauthorityv1.ActivateX509AuthorityRequest) (*localauthorityv1.ActivateX509AuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Authority ID is required

/// Only next local authority can be Activated

// Only PREPARED local authorities can be Activated

// Move next into current and reset next to clean CA

func (s *Service) TaintX509Authority(ctx context.Context, req *localauthorityv1.TaintX509AuthorityRequest) (*localauthorityv1.TaintX509AuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Authority ID is required

// It is not possible to taint Active authority

// Only next local authority can be tainted

// Only OLD authorities can be tainted

func (s *Service) TaintX509UpstreamAuthority(ctx context.Context, req *localauthorityv1.TaintX509UpstreamAuthorityRequest) (*localauthorityv1.TaintX509UpstreamAuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: may we request in lower case?
// Normalize SKID

func (s *Service) RevokeX509Authority(ctx context.Context, req *localauthorityv1.RevokeX509AuthorityRequest) (*localauthorityv1.RevokeX509AuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) RevokeX509UpstreamAuthority(ctx context.Context, req *localauthorityv1.RevokeX509UpstreamAuthorityRequest) (*localauthorityv1.RevokeX509UpstreamAuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: may we request in lower case?
// Normalize SKID

func (s *Service) GetWITAuthorityState(ctx context.Context, _ *localauthorityv1.GetWITAuthorityStateRequest) (*localauthorityv1.GetWITAuthorityStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) PrepareWITAuthority(ctx context.Context, _ *localauthorityv1.PrepareWITAuthorityRequest) (*localauthorityv1.PrepareWITAuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) ActivateWITAuthority(ctx context.Context, req *localauthorityv1.ActivateWITAuthorityRequest) (*localauthorityv1.ActivateWITAuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) TaintWITAuthority(ctx context.Context, req *localauthorityv1.TaintWITAuthorityRequest) (*localauthorityv1.TaintWITAuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) RevokeWITAuthority(ctx context.Context, req *localauthorityv1.RevokeWITAuthorityRequest) (*localauthorityv1.RevokeWITAuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) isJWTSVIDsDisabled() bool { _ = "STUB: not implemented"; return false }

// validateLocalAuthorityID validates provided authority ID, and return OLD associated public key
func (s *Service) validateLocalAuthorityID(authorityID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) validateUpstreamAuthoritySubjectKey(subjectKeyIDRequest string) error {
	_ = "STUB: not implemented"
	return nil
}

// validateAuthorityID validates provided authority ID
func (s *Service) validateAuthorityID(ctx context.Context, authorityID string) error {
	_ = "STUB: not implemented"
	return nil
}

func buildAuditLogFields(authorityID string) logrus.Fields {
	_ = "STUB: not implemented"
	return *new(logrus.Fields)
}

func buildAuditUpstreamLogFields(authorityID string) logrus.Fields {
	_ = "STUB: not implemented"
	return *new(logrus.Fields)
}

func stateFromSlot(s manager.Slot) *localauthorityv1.AuthorityState {
	_ = "STUB: not implemented"
	return nil
}
