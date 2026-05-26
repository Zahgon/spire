package main

import (
	"context"
	"sync"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/go-jose/go-jose/v4"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/bundle/spiffebundle"
)

const (
	DefaultFilePollInterval = time.Second * 10
)

type FileSourceConfig struct {
	Log          logrus.FieldLogger
	Path         string
	PollInterval time.Duration
	Clock        clock.Clock
}

type FileSource struct {
	log    logrus.FieldLogger
	clock  clock.Clock
	cancel context.CancelFunc

	mu       sync.RWMutex
	wg       sync.WaitGroup
	bundle   *spiffebundle.Bundle
	jwks     *jose.JSONWebKeySet
	modTime  time.Time
	pollTime time.Time
}

func NewFileSource(config FileSourceConfig) *FileSource { _ = "STUB: not implemented"; return nil }

func (s *FileSource) Close() error { _ = "STUB: not implemented"; return nil }

func (s *FileSource) FetchKeySet() (*jose.JSONWebKeySet, time.Time, bool) {
	_ = "STUB: not implemented"
	return nil, *new(time.Time), false
}

func (s *FileSource) LastSuccessfulPoll() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (s *FileSource) pollEvery(ctx context.Context, path string, interval time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (s *FileSource) pollOnce(path string) { _ = "STUB: not implemented"; return }

func (s *FileSource) parseBundle(bundle *spiffebundle.Bundle) {
	_ = "STUB: not implemented"
	// If the bundle hasn't changed, don't bother continuing
	return
}
