package middleware

import (
	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire/pkg/common/api/middleware"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"google.golang.org/grpc"
)

type Middleware = middleware.Middleware
type PreprocessFunc = middleware.PreprocessFunc
type PostprocessFunc = middleware.PostprocessFunc

func Preprocess(fn PreprocessFunc) Middleware { _ = "STUB: not implemented"; return *new(Middleware) }

func Postprocess(fn PostprocessFunc) Middleware { _ = "STUB: not implemented"; return *new(Middleware) }

func Funcs(preprocess PreprocessFunc, postprocess PostprocessFunc) Middleware {
	_ = "STUB: not implemented"
	return *new(Middleware)
}

func Chain(ms ...Middleware) Middleware { _ = "STUB: not implemented"; return *new(Middleware) }

func WithLogger(log logrus.FieldLogger) Middleware {
	_ = "STUB: not implemented"
	return *new(Middleware)
}

func WithMetrics(metrics telemetry.Metrics) Middleware {
	_ = "STUB: not implemented"
	return *new(Middleware)
}

func Interceptors(m Middleware) (grpc.UnaryServerInterceptor, grpc.StreamServerInterceptor) {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor), *new(grpc.StreamServerInterceptor)
}

func UnaryInterceptor(m Middleware) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func StreamInterceptor(m Middleware) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}
