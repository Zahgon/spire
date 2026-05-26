package middleware

import (
	"context"

	"google.golang.org/grpc"
)

func Interceptors(middleware Middleware) (grpc.UnaryServerInterceptor, grpc.StreamServerInterceptor) {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor), *new(grpc.StreamServerInterceptor)
}

func UnaryInterceptor(middleware Middleware) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func StreamInterceptor(middleware Middleware) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}

type serverStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (ss serverStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
