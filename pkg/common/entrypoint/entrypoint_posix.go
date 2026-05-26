//go:build !windows

package entrypoint

import (
	"context"
)

type EntryPoint struct {
	runCmdFn func(ctx context.Context, args []string) int
}

func NewEntryPoint(runFn func(ctx context.Context, args []string) int) *EntryPoint {
	_ = "STUB: not implemented"
	return nil
}

func (e *EntryPoint) Main() int { _ = "STUB: not implemented"; return 0 }
