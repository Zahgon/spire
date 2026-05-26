package rotator

import (
	"github.com/spiffe/spire/pkg/common/health"
)

// TODO: What would be a good threshold number?
const failedRotationThreshold = 10

type caSyncHealth struct {
	m *Rotator
}

func (h *caSyncHealth) CheckHealth() health.State {
	_ = "STUB: not implemented"
	// Readiness and liveness will be checked by manager's ability to
	// rotate for a certain threshold.
	return *new(health.State)
}

type managerHealthDetails struct {
	RotationErr string `json:"rotation_err,omitempty"`
}

func errString(err error) string { _ = "STUB: not implemented"; return "" }
