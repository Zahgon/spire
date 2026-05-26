package fakehealthchecker

import (
	"github.com/spiffe/spire/pkg/common/health"
)

type Checker struct {
	checkables map[string]health.Checkable
}

var _ health.Checker = (*Checker)(nil)

func New() *Checker { _ = "STUB: not implemented"; return nil }

func (c *Checker) AddCheck(name string, checkable health.Checkable) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Checker) RunChecks() map[string]health.State { _ = "STUB: not implemented"; return nil }
