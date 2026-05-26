package rpccontext

import (
	"context"
)

type callerPIDKey struct{}

// WithCallerPID returns a context with the given caller PID
func WithCallerPID(ctx context.Context, pid int) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// CallerPID returns the caller pid.
func CallerPID(ctx context.Context) int { _ = "STUB: not implemented"; return 0 }
