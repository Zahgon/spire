package middleware

import (
	"context"
)

type PreprocessFunc = func(ctx context.Context, fullMethod string, req any) (context.Context, error)
type PostprocessFunc = func(ctx context.Context, fullMethod string, handlerInvoked bool, rpcErr error)

type Middleware interface {
	// Preprocess is invoked before the gRPC handler is called. It returns a
	// (possibly modified) context that is passed into the handler, which
	// should either be the context passed into the function or one derived
	// from it. If the function returns an error, the gRPC method fails.
	// req passes the request object for unary interceptors and nil for
	// stream interceptors
	Preprocess(ctx context.Context, fullMethod string, req any) (context.Context, error)

	// Postprocess is invoked after the handler is called, or if downstream
	// middleware returns an error from Preprocess. The function is passed an
	// error that was returned from the handler or a downstream middleware
	// during preprocessing. The handlerInvoked boolean, if true, indicates
	// that the handler was executed. If false, then the call failed during
	// preprocessing.
	Postprocess(ctx context.Context, fullMethod string, handlerInvoked bool, rpcErr error)
}

// Preprocess creates a middleware from a function that does pre-processing only.
func Preprocess(fn PreprocessFunc) Middleware { _ = "STUB: not implemented"; return *new(Middleware) }

// Postprocess creates a middleware from a function that does postprocessing only.
func Postprocess(fn PostprocessFunc) Middleware { _ = "STUB: not implemented"; return *new(Middleware) }

// Funcs constructs a Middleware from a pair of functions, one for preprocessing, one for postprocessing.
func Funcs(preprocess PreprocessFunc, postprocess PostprocessFunc) Middleware {
	_ = "STUB: not implemented"
	return *new(Middleware)
}

// Chain chains together a series of middleware. The middleware is called in
// order during preprocessing and in reverse order for postprocessing. The
// context returned by each Middleware during preprocessing is passed into subsequent middlewares
func Chain(middleware ...Middleware) Middleware { _ = "STUB: not implemented"; return *new(Middleware) }

type funcs struct {
	preprocess  PreprocessFunc
	postprocess PostprocessFunc
}

// Preprocess implements the Middleware interface
func (h funcs) Preprocess(ctx context.Context, fullMethod string, req any) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// Preprocess implements the Middleware interface
func (h funcs) Postprocess(ctx context.Context, fullMethod string, handlerInvoked bool, rpcErr error) {
	_ = "STUB: not implemented"
	return
}

type middlewares []Middleware

func (ms middlewares) Preprocess(ctx context.Context, fullMethod string, req any) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// The downstream middleware failed to preprocess. Invoke the
// postprocess step of this middleware layer, passing in the context
// originally set up by this layer.

func (ms middlewares) Postprocess(ctx context.Context, fullMethod string, handlerInvoked bool, rpcErr error) {
	_ = "STUB: not implemented"
	return
}
