package main

import (
	"context"
	"sync"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/go-jose/go-jose/v4"
	"github.com/sirupsen/logrus"
	bundlev1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/bundle/v1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"google.golang.org/grpc"
)

const (
	DefaultServerAPIPollInterval = time.Second * 10
)

type ServerAPISourceConfig struct {
	Log          logrus.FieldLogger
	GRPCTarget   string
	PollInterval time.Duration
	Clock        clock.Clock
}

type ServerAPISource struct {
	log    logrus.FieldLogger
	clock  clock.Clock
	cancel context.CancelFunc

	mu       sync.RWMutex
	wg       sync.WaitGroup
	bundle   *types.Bundle
	jwks     *jose.JSONWebKeySet
	modTime  time.Time
	pollTime time.Time
}

func NewServerAPISource(config ServerAPISourceConfig) (*ServerAPISource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ServerAPISource) Close() error { _ = "STUB: not implemented"; return nil }

func (s *ServerAPISource) FetchKeySet() (*jose.JSONWebKeySet, time.Time, bool) {
	_ = "STUB: not implemented"
	return nil, *new(time.Time), false
}

func (s *ServerAPISource) LastSuccessfulPoll() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (s *ServerAPISource) pollEvery(ctx context.Context, conn *grpc.ClientConn, interval time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (s *ServerAPISource) pollOnce(ctx context.Context, client bundlev1.BundleClient) {
	_ = "STUB: not implemented"
	// Ensure the stream gets cleaned up
	return
}

func (s *ServerAPISource) parseBundle(bundle *types.Bundle) {
	_ = "STUB: not implemented"
	// If the bundle hasn't changed, don't bother continuing
	return
}
