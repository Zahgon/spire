package health

import "context"

type healthCheckKey struct{}

func IsCheck(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func CheckContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
