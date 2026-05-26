package fakeworkloadattestor

import (
	"context"
	"testing"

	workloadattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/agent/workloadattestor/v1"
	"github.com/spiffe/spire/pkg/agent/plugin/workloadattestor"
)

func New(t *testing.T, name string, pids map[int32][]string) workloadattestor.WorkloadAttestor {
	_ = "STUB: not implemented"
	return *new(workloadattestor.WorkloadAttestor)
}

type workloadAttestor struct {
	workloadattestorv1.UnimplementedWorkloadAttestorServer

	pids map[int32][]string
}

func (p *workloadAttestor) Attest(_ context.Context, req *workloadattestorv1.AttestRequest) (*workloadattestorv1.AttestResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
