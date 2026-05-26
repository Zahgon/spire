package agent

import (
	"context"
	"crypto/x509"

	"github.com/andres-erbsen/clock"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	agentv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/agent/v1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/pkg/server/ca"
	"github.com/spiffe/spire/pkg/server/catalog"
	"github.com/spiffe/spire/pkg/server/datastore"
	"github.com/spiffe/spire/pkg/server/plugin/nodeattestor"
	"github.com/spiffe/spire/proto/spire/common"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Config is the service configuration
type Config struct {
	Catalog                 catalog.Catalog
	Clock                   clock.Clock
	DataStore               datastore.DataStore
	ServerCA                ca.ServerCA
	TrustDomain             spiffeid.TrustDomain
	AgentSpiffeIdAsSelector bool
}

// Service implements the v1 agent service
type Service struct {
	agentv1.UnsafeAgentServer

	cat                     catalog.Catalog
	clk                     clock.Clock
	ds                      datastore.DataStore
	ca                      ca.ServerCA
	td                      spiffeid.TrustDomain
	AgentSpiffeIdAsSelector bool
}

// New creates a new agent service
func New(config Config) *Service { _ = "STUB: not implemented"; return nil }

// RegisterService registers the agent service on the gRPC server/
func RegisterService(s grpc.ServiceRegistrar, service *Service) { _ = "STUB: not implemented"; return }

// CountAgents returns the total number of agents.
func (s *Service) CountAgents(ctx context.Context, req *agentv1.CountAgentsRequest) (*agentv1.CountAgentsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse proto filter into datastore request

// ListAgents returns an optionally filtered and/or paginated list of agents.
func (s *Service) ListAgents(ctx context.Context, req *agentv1.ListAgentsRequest) (*agentv1.ListAgentsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse proto filter into datastore request

// Set pagination parameters

// Parse nodes into proto and apply output mask

// GetAgent returns the agent associated with the given SpiffeID.
func (s *Service) GetAgent(ctx context.Context, req *agentv1.GetAgentRequest) (*types.Agent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteAgent removes the agent with the given SpiffeID.
func (s *Service) DeleteAgent(ctx context.Context, req *agentv1.DeleteAgentRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BanAgent sets the agent with the given SpiffeID to the banned state.
func (s *Service) BanAgent(ctx context.Context, req *agentv1.BanAgentRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The agent "Banned" state is pointed out by setting its
// serial numbers (current and new) to empty strings.

// AttestAgent attests the authenticity of the given agent.
func (s *Service) AttestAgent(stream agentv1.Agent_AttestAgentServer) error {
	_ = "STUB: not implemented"
	return nil
}

// validate

// attest

// Ideally we'd do stronger validation that the ID is within the Node
// Attestors scoped area of the reserved agent namespace, but historically
// we haven't been strict here and there are deployments that are emitting
// such IDs.
// Deprecated: enforce that IDs produced by Node Attestors are in the
// reserved namespace for that Node Attestor starting in SPIRE 1.4.

// fetch the agent/node to check if it was already attested or banned

// parse and sign CSR

// dedupe and store node selectors

// create or update attested entry

// build and send response

// RenewAgent renews the SVID of the agent with the given SpiffeID.
func (s *Service) RenewAgent(ctx context.Context, req *agentv1.RenewAgentRequest) (*agentv1.RenewAgentResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Agent attempted to renew when it should've been reattesting

// Send response with new X509 SVID

// PostStatus posts agent status including the agent version
func (s *Service) PostStatus(ctx context.Context, req *agentv1.PostStatusRequest) (*agentv1.PostStatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateJoinToken returns a new JoinToken for an agent.
func (s *Service) CreateJoinToken(ctx context.Context, req *agentv1.CreateJoinTokenRequest) (*types.JoinToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If provided, check that the AgentID is valid BEFORE creating the join token so we can fail early

// Generate a token if one wasn't specified

func (s *Service) createJoinTokenRegistrationEntry(ctx context.Context, token string, agentID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) updateAttestedNode(ctx context.Context, node *common.AttestedNode, mask *common.AttestedNodeMask, log logrus.FieldLogger) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) signSvid(ctx context.Context, agentID spiffeid.ID, csr []byte, log logrus.FieldLogger) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sign a new X509 SVID

func (s *Service) getSelectorsFromAgentID(ctx context.Context, agentID string) ([]*types.Selector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) attestJoinToken(ctx context.Context, token string) (*nodeattestor.AttestResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) attestChallengeResponse(ctx context.Context, agentStream agentv1.Agent_AttestAgentServer, params *agentv1.AttestAgentRequest_Params) (*nodeattestor.AttestResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func applyMask(a *types.Agent, mask *types.AgentMask) { _ = "STUB: not implemented"; return }

func validateAttestAgentParams(params *agentv1.AttestAgentRequest_Params) error {
	_ = "STUB: not implemented"
	return nil
}

func getAttestAgentResponse(spiffeID spiffeid.ID, certificates []*x509.Certificate, canReattest bool) *agentv1.AttestAgentResponse {
	_ = "STUB: not implemented"
	return nil
}

func fieldsFromListAgentsRequest(filter *agentv1.ListAgentsRequest_Filter) logrus.Fields {
	_ = "STUB: not implemented"
	return *new(logrus.Fields)
}

func fieldsFromCountAgentsRequest(filter *agentv1.CountAgentsRequest_Filter) logrus.Fields {
	_ = "STUB: not implemented"
	return *new(logrus.Fields)
}

func joinTokenID(td spiffeid.TrustDomain, token string) (spiffeid.ID, error) {
	_ = "STUB: not implemented"
	return *new(spiffeid.ID), nil
}
