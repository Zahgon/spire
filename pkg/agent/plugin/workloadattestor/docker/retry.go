package docker

import (
	"context"
	"time"

	"github.com/andres-erbsen/clock"
)

const (
	defaultNumRetries     = 3
	defaultInitialBackoff = 100 * time.Millisecond
)

type retryer struct {
	clock          clock.Clock
	disabled       bool
	numRetries     int
	initialBackoff time.Duration
}

func newRetryer() *retryer { _ = "STUB: not implemented"; return nil }

func (r *retryer) Retry(ctx context.Context, fn func() error) error {
	_ = "STUB: not implemented"
	return nil
}

// try once plus the number of retries

// don't wait another backoff cycle if we've already maxed out on retries

func exponentialBackoff(c int) float64 { _ = "STUB: not implemented"; return 0 }
