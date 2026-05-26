package api

import (
	"context"

	"github.com/spiffe/spire/pkg/common/peertracker"

	"google.golang.org/grpc"
)

type Server interface {
	ListenAndServe(ctx context.Context) error
}

type Endpoints struct {
	c        *Config
	listener *peertracker.ListenerFactory
}

func (e *Endpoints) ListenAndServe(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Endpoints) registerDebugAPI(server *grpc.Server) { _ = "STUB: not implemented"; return }

func (e *Endpoints) registerDelegatedIdentityAPI(server *grpc.Server) {
	_ = "STUB: not implemented"
	return
}
