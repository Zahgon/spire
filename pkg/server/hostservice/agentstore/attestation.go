package agentstore

import (
	"context"

	agentstorev1 "github.com/spiffe/spire-plugin-sdk/proto/spire/hostservice/server/agentstore/v1"
)

func EnsureNotAttested(ctx context.Context, store agentstorev1.AgentStoreClient, agentID string) error {
	_ = "STUB: not implemented"
	return nil
}

func IsAttested(ctx context.Context, store agentstorev1.AgentStoreClient, agentID string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
