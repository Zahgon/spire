package agentstore

import (
	"context"
	"sync"

	agentstorev1 "github.com/spiffe/spire-plugin-sdk/proto/spire/hostservice/server/agentstore/v1"
	"github.com/spiffe/spire/pkg/server/datastore"
)

type Deps struct {
	// DataStore is used to retrieve agent information. It MUST be set.
	DataStore datastore.DataStore
}

type AgentStore struct {
	mu   sync.RWMutex
	deps *Deps
}

func New() *AgentStore { _ = "STUB: not implemented"; return nil }

func (s *AgentStore) SetDeps(deps Deps) error { _ = "STUB: not implemented"; return nil }

func (s *AgentStore) getDeps() (*Deps, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *AgentStore) V1() agentstorev1.AgentStoreServer {
	_ = "STUB: not implemented"
	return *new(agentstorev1.AgentStoreServer)
}

type agentStoreV1 struct {
	agentstorev1.UnsafeAgentStoreServer

	s *AgentStore
}

func (v1 *agentStoreV1) GetAgentInfo(ctx context.Context, req *agentstorev1.GetAgentInfoRequest) (*agentstorev1.GetAgentInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
