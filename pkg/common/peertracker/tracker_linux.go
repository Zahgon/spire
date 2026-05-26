//go:build linux

package peertracker

import (
	"sync"

	"github.com/sirupsen/logrus"
)

const (
	linuxType = "linux"
)

type linuxTracker struct {
	log logrus.FieldLogger
}

func newTracker(log logrus.FieldLogger) (*linuxTracker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *linuxTracker) NewWatcher(info CallerInfo) (Watcher, error) {
	_ = "STUB: not implemented"
	return *new(Watcher), nil
}

func (*linuxTracker) Close() { _ = "STUB: not implemented"; return }

type linuxWatcher struct {
	gid       uint32
	pid       int32
	mtx       sync.Mutex
	procPath  string
	procfd    int
	starttime string
	uid       uint32
	log       logrus.FieldLogger
}

func newLinuxWatcher(info CallerInfo, log logrus.FieldLogger) (*linuxWatcher, error) {
	_ = "STUB: not implemented"
	// If PID == 0, something is wrong...
	return nil, nil
}

// Grab a handle to proc first since that's the fastest thing we can do

func (l *linuxWatcher) Close() { _ = "STUB: not implemented"; return }

func (l *linuxWatcher) IsAlive() error { _ = "STUB: not implemented"; return nil }

// First we will check if we can read from the original directory handle.
// If the process has exited since we opened it, the read should fail (i.e.
// the ReadDirent syscall will return -1)

// A successful fd read should indicate that the original process is still alive, however
// it is not clear if the original inode can be freed by Linux while it is still referenced.
// This _shouldn't_ happen, but if it does, then there might be room for a reused PID to
// collide with the original inode making the read successful. As an extra measure, ensure
// that the current `starttime` matches the one we saw originally.
//
// This is probably overkill.
// TODO: Evaluate the use of `starttime` as the primary exit detection mechanism.

// Finally, read the UID and GID off the proc directory to determine the owner. If
// we got beaten by a PID race when opening the proc handle originally, we can at
// least get to know that the race winner is running as the same user and group as
// the original caller by comparing it to the received CallerInfo.

func (l *linuxWatcher) PID() int32 { _ = "STUB: not implemented"; return 0 }

func parseTaskStat(stat string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func getStarttime(pid int32) (string, error) { _ = "STUB: not implemented"; return "", nil }

// starttime is the 22nd field in the proc stat data
// Field number 38 was introduced in Linux 2.1.22
// Protect against invalid index and reject anything before 2.1.22
