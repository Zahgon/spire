package main

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/go-jose/go-jose/v4"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/bundle/jwtbundle"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
)

const (
	DefaultWorkloadAPIPollInterval = time.Second * 10
)

type WorkloadAPISourceConfig struct {
	Log          logrus.FieldLogger
	Addr         net.Addr
	TrustDomain  string
	PollInterval time.Duration
	Clock        clock.Clock
}

type WorkloadAPISource struct {
	log         logrus.FieldLogger
	clock       clock.Clock
	trustDomain spiffeid.TrustDomain
	cancel      context.CancelFunc

	mu        sync.RWMutex
	wg        sync.WaitGroup
	rawBundle []byte
	jwks      *jose.JSONWebKeySet
	modTime   time.Time
	pollTime  time.Time
}

func NewWorkloadAPISource(config WorkloadAPISourceConfig) (*WorkloadAPISource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *WorkloadAPISource) Close() error { _ = "STUB: not implemented"; return nil }

func (s *WorkloadAPISource) FetchKeySet() (*jose.JSONWebKeySet, time.Time, bool) {
	_ = "STUB: not implemented"
	return nil, *new(time.Time), false
}

func (s *WorkloadAPISource) LastSuccessfulPoll() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (s *WorkloadAPISource) pollEvery(ctx context.Context, client *workloadapi.Client, interval time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (s *WorkloadAPISource) pollOnce(ctx context.Context, client *workloadapi.Client) {
	_ = "STUB: not implemented"
	return
}

// update pollTime when setJWKS was successful

func (s *WorkloadAPISource) setJWKS(bundle *jwtbundle.Bundle) error {
	_ = "STUB: not implemented"
	return nil
}

// If the bundle hasn't changed, don't bother continuing

// Clean the JWKS
