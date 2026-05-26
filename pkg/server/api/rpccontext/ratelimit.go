package rpccontext

import (
	"context"

	"github.com/spiffe/spire/pkg/common/api"
)

type rateLimiterKey struct{}

func WithRateLimiter(ctx context.Context, limiter api.RateLimiter) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func RateLimiter(ctx context.Context) (api.RateLimiter, bool) {
	_ = "STUB: not implemented"
	return *new(api.RateLimiter), false
}

func RateLimit(ctx context.Context, count int) error { _ = "STUB: not implemented"; return nil }
