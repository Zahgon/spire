package workloadattestor

import (
	"context"

	workloadattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/agent/workloadattestor/v1"
	"github.com/spiffe/spire/pkg/common/plugin"
	"github.com/spiffe/spire/proto/spire/common"
)

type V1 struct {
	plugin.Facade
	workloadattestorv1.WorkloadAttestorPluginClient
}

func (v1 *V1) Attest(ctx context.Context, pid int) ([]*common.Selector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
