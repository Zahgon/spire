package middleware

import (
	"context"
	"sync"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/spiffe/spire/pkg/common/api/middleware"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/pkg/server/api"
	"golang.org/x/time/rate"
)

const (
	// gcInterval is the interval at which per-ip limiters are garbage
	// collected.
	gcInterval = time.Minute
)

var (
	// Used to manipulate time in unit tests
	clk = clock.New()
)

var (
	// newRawRateLimiter is used to create a new ratelimiter. It returns a limiter
	// from the standard rate package by default production.
	newRawRateLimiter = func(limit rate.Limit, burst int) rawRateLimiter {
		return rate.NewLimiter(limit, burst)
	}
)

type noopRateLimiter interface {
	noop()
}

// rawRateLimiter represents the raw limiter functionality.
type rawRateLimiter interface {
	WaitN(ctx context.Context, count int) error
	Limit() rate.Limit
	Burst() int
}

// NoLimit returns a rate limiter that does not rate limit. It is used to
// configure methods that don't do rate limiting.
func NoLimit() api.RateLimiter {
	_ = "STUB: not implemented"

	// DisabledLimit returns a rate limiter that does not rate limit. It is used to
	// configure methods where rate limiting has been disabled by configuration.
	return *new(api.RateLimiter)
}

func DisabledLimit() api.RateLimiter {
	_ = "STUB: not implemented"
	return *

	// PerCallLimit returns a rate limiter that imposes a server-wide limit for
	// calls to the method. It can be shared across methods to enforce a
	// server-wide limit for a group of methods.
	new(api.RateLimiter)
}

func PerCallLimit(limit int) api.RateLimiter {
	_ = "STUB: not implemented"
	return *new(api.RateLimiter)
}

// PerIPLimit returns a rate limiter that imposes a per-ip limit on calls
// to a method. It can be shared across methods to enforce per-ip limits for
// a group of methods.
func PerIPLimit(limit int) api.RateLimiter { _ = "STUB: not implemented"; return *new(api.RateLimiter) }

// WithRateLimits returns a middleware that performs rate limiting for the
// group of methods described by the rateLimits map. It provides the
// configured rate limiter to the method handlers via the request context. If
// the middleware is invoked for a method that is not described in the map, it
// will fail the RPC with an INTERNAL error code, describing the RPC that was
// not configured properly. The middleware also encourages proper rate limiting
// by logging errors if a handler fails to invoke the rate limiter provided on
// the context when a limit has been configured or the handler invokes the rate
// limiter when a no limit has been configured.
//
// WithRateLimits owns the passed rateLimits map and assumes it will not be
// mutated after the method is called.
//
// The WithRateLimits middleware depends on the Logger and Authorization
// middlewares.
func WithRateLimits(rateLimits map[string]api.RateLimiter, metrics telemetry.Metrics) middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}

type noLimit struct{}

func (noLimit) RateLimit(context.Context, int) error { _ = "STUB: not implemented"; return nil }

func (noLimit) noop() { _ = "STUB: not implemented"; return }

type disabledLimit struct{}

func (disabledLimit) RateLimit(context.Context, int) error { _ = "STUB: not implemented"; return nil }

func (disabledLimit) noop() { _ = "STUB: not implemented"; return }

type perCallLimiter struct {
	limiter rawRateLimiter
}

func newPerCallLimiter(limit int) *perCallLimiter { _ = "STUB: not implemented"; return nil }

func (lim *perCallLimiter) RateLimit(ctx context.Context, count int) error {
	_ = "STUB: not implemented"
	return nil
}

type perIPLimiter struct {
	limit int

	mtx sync.RWMutex

	// previous holds all the limiters that were current at the GC
	previous map[string]rawRateLimiter

	// current holds all the limiters that have been created or moved
	// from the previous limiters since the last GC.
	current map[string]rawRateLimiter

	// lastGC is the last GC
	lastGC time.Time
}

func newPerIPLimiter(limit int) *perIPLimiter { _ = "STUB: not implemented"; return nil }

func (lim *perIPLimiter) RateLimit(ctx context.Context, count int) error {
	_ = "STUB: not implemented"
	return nil
}

// Calls not via TCP/IP aren't limited

func (lim *perIPLimiter) getLimiter(ip string) rawRateLimiter {
	_ = "STUB: not implemented"
	return *new(rawRateLimiter)
}

// A limiter does not exist for that address.

// Check the "current" entries in case another goroutine raced on this IP.

// Then check the "previous" entries to see if a limiter exists for this
// IP as of the last GC. If so, move it to current and return it.

// There is no limiter for this IP. Before we create one, we should see
// if we need to do GC.

type rateLimitsMiddleware struct {
	limiters map[string]api.RateLimiter
	metrics  telemetry.Metrics
}

func (i rateLimitsMiddleware) Preprocess(ctx context.Context, fullMethod string, _ any) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (i rateLimitsMiddleware) Postprocess(ctx context.Context, _ string, handlerInvoked bool, rpcErr error) {
	_ = "STUB: not implemented"
	// Handlers are expected to invoke the rate limiter unless they failed to
	// parse parameters. If the handler itself wasn't invoked then there is no
	// need to check if rate limiting was invoked.
	return
}

// This shouldn't be the case unless Preprocess is broken and fails to
// inject the rate limiter into the context.

// This shouldn't be the case unless Preprocess is broken and fails to
// wrap the rate limiter.

func logLimiterMisuse(ctx context.Context, rateLimiter api.RateLimiter, used bool) {
	_ = "STUB: not implemented"
	return
}

// RPC should not invoke the rate limiter, since that would imply a
// misconfiguration. Either the RPC is wrong, or the middleware is
// wrong as to whether the RPC should rate limit.

// RPC should invoke the rate limiter since is an RPC that is normally
// rate limited. The disabled limiter will not actually apply any
// limits but we want to make sure the RPC will be applying limits
// under normal conditions.

// All other rate limiters should definitely be invoked by the RPC or
// it is a bug.

type rateLimiterWrapper struct {
	rateLimiter api.RateLimiter
	used        bool
	metrics     telemetry.Metrics
}

func (w *rateLimiterWrapper) RateLimit(ctx context.Context, count int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w *rateLimiterWrapper) Used() bool { _ = "STUB: not implemented"; return false }

func getNames(ctx context.Context) []string { _ = "STUB: not implemented"; return nil }

func waitN(ctx context.Context, limiter rawRateLimiter, count int) (err error) {
	_ = "STUB: not implemented"
	// limiter.WaitN already provides this check but the error returned is not
	// strongly typed and is a little messy. Lifting this check so we can
	// provide a clean error message.
	return nil
}
