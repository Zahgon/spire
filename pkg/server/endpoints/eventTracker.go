package endpoints

import (
	"sync"
	"time"
)

type eventTracker struct {
	pollPeriods uint

	events map[uint]uint

	pool sync.Pool
}

func PollPeriods(pollTime time.Duration, trackTime time.Duration) uint {
	_ = "STUB: not implemented"
	return 0
}

func NewEventTracker(pollPeriods uint) *eventTracker { _ = "STUB: not implemented"; return nil }

// See https://staticcheck.dev/docs/checks#SA6002.

func (et *eventTracker) PollPeriods() uint { _ = "STUB: not implemented"; return 0 }

func (et *eventTracker) Polls() uint { _ = "STUB: not implemented"; return 0 }

func (et *eventTracker) StartTracking(event uint) { _ = "STUB: not implemented"; return }

func (et *eventTracker) StopTracking(event uint) { _ = "STUB: not implemented"; return }

func (et *eventTracker) SelectEvents() []uint { _ = "STUB: not implemented"; return nil }

func (et *eventTracker) FreeEvents(events []uint) { _ = "STUB: not implemented"; return }

func (et *eventTracker) EventCount() int { _ = "STUB: not implemented"; return 0 }
