package clock

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/andres-erbsen/clock"
)

// Clock is a clock
type Clock clock.Clock

// Mock is a mock clock that can be precisely controlled
type Mock struct {
	*clock.Mock
	t           testing.TB
	timerC      chan time.Duration
	afterC      chan time.Duration
	tickerC     chan time.Duration
	tickerCount atomic.Int32
	sleepC      chan time.Duration
	afterHook   func(time.Duration) <-chan time.Time
}

// NewMock creates a mock clock which can be precisely controlled
func NewMock(t testing.TB) *Mock { _ = "STUB: not implemented"; return nil }

// NewMockAt creates a mock clock which can be precisely controlled at a specific time.
func NewMockAt(t testing.TB, now time.Time) *Mock { _ = "STUB: not implemented"; return nil }

// TLS verification is being done using a realtime clock so we set the mock clock to
// the current time, truncated to a second which is the granularity available to asn1.
// This ensures that when tests create a certificate with a lifetime of 3 seconds, it
// is exactly 3 seconds (relative to the mock clock).
//
// TODO: plumb the clock into the TLS configs. (Clock).Now should be passed to "crypto/tls".(Config).Time
// and then this can be removed as a clock could be use with a zero value at that point.

func (m *Mock) SetAfterHook(h func(time.Duration) <-chan time.Time) {
	_ = "STUB: not implemented"
	return
}

func (m *Mock) TimerCh() <-chan time.Duration { _ = "STUB: not implemented"; return nil }

func (m *Mock) WaitForAfterCh() <-chan time.Duration {
	_ = "STUB: not implemented"

	// WaitForTimer waits up to the specified timeout for Timer to be called on the clock.
	return nil
}

func (m *Mock) WaitForTimer(timeout time.Duration, format string, args ...any) {
	_ = "STUB: not implemented"
	return
}

// WaitForAfter waits up to the specified timeout for After to be called on the clock.
func (m *Mock) WaitForAfter(timeout time.Duration, format string, args ...any) {
	_ = "STUB: not implemented"
	return
}

// WaitForTicker waits up to the specified timeout for a Ticker to be created from the clock.
func (m *Mock) WaitForTicker(timeout time.Duration, format string, args ...any) {
	_ = "STUB: not implemented"
	return
}

func (m *Mock) WaitForTickerMulti(timeout time.Duration, count int32, format string, args ...any) {
	_ = "STUB: not implemented"
	return
}

// WaitForSleep waits up to the specified timeout for a sleep to begin using the clock.
func (m *Mock) WaitForSleep(timeout time.Duration, format string, args ...any) {
	_ = "STUB: not implemented"
	return
}

// Timer creates a new Timer containing a channel that will send the time with a period specified by the duration argument.
func (m *Mock) Timer(d time.Duration) *clock.Timer { _ = "STUB: not implemented"; return nil }

// After waits for the duration to elapse and then sends the current time on the returned channel.
func (m *Mock) After(d time.Duration) <-chan time.Time { _ = "STUB: not implemented"; return nil }

// Ticker returns a new Ticker containing a channel that will send the time with a period specified by the duration argument.
func (m *Mock) Ticker(d time.Duration) *clock.Ticker { _ = "STUB: not implemented"; return nil }

// Sleep pauses the current goroutine for at least the duration d
func (m *Mock) Sleep(d time.Duration) { _ = "STUB: not implemented"; return }
