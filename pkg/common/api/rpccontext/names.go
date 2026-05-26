package rpccontext

import (
	"context"

	"github.com/spiffe/spire/pkg/common/api"
)

type namesKey struct{}

func WithNames(ctx context.Context, names api.Names) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func Names(ctx context.Context) (api.Names, bool) {
	_ = "STUB: not implemented"
	return *new(api.Names), false
}
