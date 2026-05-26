//go:build darwin || freebsd || netbsd || openbsd

package peertracker

import (
	"context"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	bsdType = "bsd"
)

var safetyDelay = 250 * time.Millisecond

type bsdTracker struct {
	closer      func()
	ctx         context.Context
	kqfd        int
	mtx         sync.Mutex
	watchedPIDs map[int]chan struct{}
	log         logrus.FieldLogger
}

func newTracker(log logrus.FieldLogger) (*bsdTracker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *bsdTracker) Close() { _ = "STUB: not implemented"; return }

// Be sure to cancel the context before closing the
// kqueue file descriptor so the goroutine watching it
// will know that we are shutting down.

func (b *bsdTracker) NewWatcher(info CallerInfo) (Watcher, error) {
	_ = "STUB: not implemented"
	// If PID == 0, something is wrong...
	return *new(Watcher), nil
}

func (b *bsdTracker) addKeventForWatcher(pid int) error { _ = "STUB: not implemented"; return nil }

func (b *bsdTracker) receiveKevents(kqfd int) { _ = "STUB: not implemented"; return }

// KQUEUE(2) outlines the conditions under which the Kevent call
// can return an error - they are as follows:
//
// EACCESS: The process does not have permission to register a filter.
// EFAULT: There was an error reading or writing the kevent or kevent64_s structure.
// EBADF: The specified descriptor is invalid.
// EINTR: A signal was delivered before the timeout expired and before any events were
//        placed on the kqueue for return.
// EINVAL: The specified time limit or filter is invalid.
// ENOENT: The event could not be found to be modified or deleted.
// ENOMEM: No memory was available to register the event.
// ESRCH: The specified process to attach to does not exist.
//
// Given our usage, the only error that seems possible is EBADF during shutdown.
// If we encounter any other error, we really have no way to recover. This will cause
// all subsequent workload attestations to fail open. After much deliberation, it is
// decided that the safest thing to do is to panic and allow supervision to step in.
// If this is actually encountered in the wild, we can examine the conditions and try
// to do something more intelligent. For now, we will just check to see if we are
// shutting down.

// Don't panic, we're just shutting down

type bsdWatcher struct {
	closed bool
	done   <-chan struct{}
	mtx    sync.Mutex
	pid    int32
	log    logrus.FieldLogger
}

func newBSDWatcher(info CallerInfo, done <-chan struct{}, log logrus.FieldLogger) *bsdWatcher {
	_ = "STUB: not implemented"
	return nil
}

func (b *bsdWatcher) Close() {
	_ = "STUB: not implemented"
	// For simplicity, don't bother cleaning up after ourselves
	// The map entry will be reaped when the process exits
	//
	// Other watchers are unable to track after closed (unlike
	// this one), so to provide consistent behavior, set the closed
	// bit and return an error on subsequent IsAlive() calls
	return
}

func (b *bsdWatcher) IsAlive() error { _ = "STUB: not implemented"; return nil }

// Using kqueue/kevent means we are relying on an asynchronous notification
// system for exit detection. Delays can be incurred on either side: in our
// kevent consumer or in the kernel. Typically, IsAlive() is called following
// workload attestation which can take hundreds of milliseconds, so in practice
// we will probably have been notified of an exit by now if it occurred prior to
// or during the attestation process.
//
// As an extra safety precaution, artificially delay our answer to IsAlive() in
// a blind attempt to allow "enough" time to pass for us to learn of the
// potential exit event.

func (b *bsdWatcher) PID() int32 { _ = "STUB: not implemented"; return 0 }
