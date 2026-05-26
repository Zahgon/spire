package middleware

import (
	"context"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/pkg/common/api/middleware"
	"github.com/spiffe/spire/pkg/server/authpolicy"
	"google.golang.org/grpc/status"
)

func WithAuthorization(authPolicyEngine *authpolicy.Engine, entryFetcher EntryFetcher, agentAuthorizer AgentAuthorizer, adminIDs []spiffeid.ID) middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}

type authorizationMiddleware struct {
	authPolicyEngine *authpolicy.Engine
	entryFetcher     EntryFetcher
	agentAuthorizer  AgentAuthorizer
	adminIDs         map[spiffeid.ID]struct{}
}

func (m *authorizationMiddleware) Preprocess(ctx context.Context, methodName string, req any) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// Add request ID to logger, it simplifies debugging when calling batch endpoints

func (m *authorizationMiddleware) Postprocess(context.Context, string, bool, error) {
	_ = "STUB: not implemented"
	// Intentionally empty.
	return
}

func adminIDSet(ids []spiffeid.ID) map[spiffeid.ID]struct{} { _ = "STUB: not implemented"; return nil }

func deniedDetailsFromStatus(s *status.Status) *types.PermissionDeniedDetails {
	_ = "STUB: not implemented"
	return nil
}
