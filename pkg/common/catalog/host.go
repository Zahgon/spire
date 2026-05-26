package catalog

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire-plugin-sdk/pluginsdk"
	"google.golang.org/grpc"
)

func newHostServer(log logrus.FieldLogger, pluginName string, hostServices []pluginsdk.ServiceServer) *grpc.Server {
	_ = "STUB: not implemented"
	return nil
}

func streamPluginInterceptor(name string) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}

func unaryPluginInterceptor(name string) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func streamPanicInterceptor(log logrus.FieldLogger) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}

func unaryPanicInterceptor(log logrus.FieldLogger) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func convertPanic(log logrus.FieldLogger, r any) error { _ = "STUB: not implemented"; return nil }

type streamWrapper struct {
	ctx context.Context
	grpc.ServerStream
}

func (w streamWrapper) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
