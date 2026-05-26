package endpoints

import (
	"context"

	attestor "github.com/spiffe/spire/pkg/agent/attestor/workload"
	"github.com/spiffe/spire/proto/spire/common"
)

type PeerTrackerAttestor struct {
	Attestor attestor.Attestor
}

func (a PeerTrackerAttestor) Attest(ctx context.Context) ([]*common.Selector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ensure that the original caller is still alive so that we know we didn't
// attest some other process that happened to be assigned the original PID
