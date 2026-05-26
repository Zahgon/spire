package fakeworkloadattestor

import (
	"context"
	"testing"

	workloadattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/agent/workloadattestor/v1"
	"github.com/spiffe/spire/pkg/agent/plugin/workloadattestor"
)

func NewTimeoutAttestor(t *testing.T, name string, c chan struct{}) workloadattestor.WorkloadAttestor {
	_ = "STUB: not implemented"
	return *new(workloadattestor.WorkloadAttestor)
}

type timeoutWorkloadAttestor struct {
	workloadattestorv1.UnimplementedWorkloadAttestorServer

	c chan struct{}
}

func (twa *timeoutWorkloadAttestor) Attest(_ context.Context, _ *workloadattestorv1.AttestRequest) (*workloadattestorv1.AttestResponse, error) {
	_ = "STUB: not implemented"
	// Block on channel until test sends signal
	return nil, nil
}
