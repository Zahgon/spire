package fakeagentstore

import (
	"context"
	"sync"

	agentstorev1 "github.com/spiffe/spire-plugin-sdk/proto/spire/hostservice/server/agentstore/v1"
)

type agentConfig struct {
	info *agentstorev1.AgentInfo
	err  error
}

type AgentStore struct {
	agentstorev1.UnsafeAgentStoreServer

	mu     sync.RWMutex
	agents map[string]agentConfig
}

func New() *AgentStore { _ = "STUB: not implemented"; return nil }

func (s *AgentStore) SetAgentInfo(info *agentstorev1.AgentInfo) { _ = "STUB: not implemented"; return }

func (s *AgentStore) SetAgentErr(agentID string, err error) { _ = "STUB: not implemented"; return }

func (s *AgentStore) GetAgentInfo(_ context.Context, req *agentstorev1.GetAgentInfoRequest) (*agentstorev1.GetAgentInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
