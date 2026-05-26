package auth

import (
	"context"

	"google.golang.org/grpc"
)

type Authorizer interface {
	AuthorizeCall(ctx context.Context, fullMethod string) (context.Context, error)
}

type AuthorizerFunc func(ctx context.Context, fullMethod string) (context.Context, error)

func (fn AuthorizerFunc) AuthorizeCall(ctx context.Context, fullMethod string) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func UnaryAuthorizeCall(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func StreamAuthorizeCall(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func authorizeCall(ctx context.Context, srv any, fullMethod string) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// used to override the context on a stream
type serverStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (s serverStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
