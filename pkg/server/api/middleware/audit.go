package middleware

import (
	"context"

	"github.com/shirou/gopsutil/v4/process"
	"github.com/sirupsen/logrus"
)

func WithAuditLog(localTrackerEnabled bool) Middleware {
	_ = "STUB: not implemented"
	return *new(Middleware)
}

type auditLogMiddleware struct {
	Middleware

	localTrackerEnabled bool
}

func (m auditLogMiddleware) Preprocess(ctx context.Context, _ string, _ any) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (m auditLogMiddleware) Postprocess(ctx context.Context, _ string, _ bool, rpcErr error) {
	_ = "STUB: not implemented"
	return
}

func fieldsFromTracker(ctx context.Context) (logrus.Fields, error) {
	_ = "STUB: not implemented"
	return *new(logrus.Fields), nil
}

// Addr is expected to fail on k8s when "hostPID" is not provided

func getAddr(proc *process.Process) (string, error) { _ = "STUB: not implemented"; return "", nil }
