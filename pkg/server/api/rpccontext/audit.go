package rpccontext

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/pkg/server/api/audit"
)

type auditLogKey struct{}

func WithAuditLog(ctx context.Context, auditLog audit.Logger) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func AddRPCAuditFields(ctx context.Context, fields logrus.Fields) {
	_ = "STUB: not implemented"
	return
}

func AuditRPC(ctx context.Context) { _ = "STUB: not implemented"; return }

func AuditRPCWithFields(ctx context.Context, fields logrus.Fields) {
	_ = "STUB: not implemented"
	return
}

func AuditRPCWithError(ctx context.Context, err error) { _ = "STUB: not implemented"; return }

func AuditRPCWithTypesStatus(ctx context.Context, s *types.Status, fieldsFunc func() logrus.Fields) {
	_ = "STUB: not implemented"
	return
}

func AuditLog(ctx context.Context) (audit.Logger, bool) {
	_ = "STUB: not implemented"
	return *new(audit.Logger), false
}
