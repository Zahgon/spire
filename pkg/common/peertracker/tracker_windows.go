//go:build windows

package peertracker

import (
	"sync"

	"github.com/sirupsen/logrus"
	"golang.org/x/sys/windows"
)

const (
	windowsType = "windows"
	stillActive = 259 // https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getexitcodeprocess
)

type windowsTracker struct {
	log logrus.FieldLogger
	sc  systemCaller
}

func newTracker(log logrus.FieldLogger) (*windowsTracker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *windowsTracker) NewWatcher(info CallerInfo) (Watcher, error) {
	_ = "STUB: not implemented"
	return *new(Watcher), nil
}

func (*windowsTracker) Close() { _ = "STUB: not implemented"; return }

type windowsWatcher struct {
	mtx        sync.Mutex
	procHandle windows.Handle

	pid int32
	log logrus.FieldLogger

	sc systemCaller
}

func (t *windowsTracker) newWindowsWatcher(info CallerInfo, log logrus.FieldLogger) (*windowsWatcher, error) {
	_ = "STUB: not implemented"
	// Having an open process handle prevents the process object from being destroyed,
	// keeping the process ID valid, so this is the first thing that we do.
	return nil, nil
}

// Find out if the PID is a well known PID that we don't
// expect from a workload.

// Process ID 0 is the Idle process

// Process ID 4 is the System process

// This is a mitigation for attacks that leverage opening a
// named pipe through the local SMB server that set the PID
// attribute to 0xFEFF (65279). We want to to prevent abusing
// the fact that Windows reuses PID values and an attacker could
// cycle through process creation until it has a suitable process
// meeting the security check requirements from SMB server.
// Note that 65279 is not a valid PID in Windows because is not
// a multiple of 4, but if the SMB server calls OpenProcess on
// 65279 it will round down and open the PID 65276 which could
// be created by the attacker.
// This check makes sure that the process handle obtained from
// the PID discovered through the GetNamedPipeClientProcessId
// call matches the one that is obtained from that process ID.

func (w *windowsWatcher) Close() { _ = "STUB: not implemented"; return }

func (w *windowsWatcher) IsAlive() error { _ = "STUB: not implemented"; return nil }

// The process object remains as long as the process is still running or
// as long as there is a handle to the process object.
// GetExitCodeProcess can be called to retrieve the exit code.

func (w *windowsWatcher) PID() int32 { _ = "STUB: not implemented"; return 0 }

type systemCaller interface {
	// CloseHandle closes an open object handle.
	CloseHandle(windows.Handle) error

	// CompareObjectHandles compares two object handles to determine if they
	// refer to the same underlying kernel object
	CompareObjectHandles(windows.Handle, windows.Handle) error

	// OpenProcess returns an open handle to the specified process id.
	OpenProcess(int32) (windows.Handle, error)

	// GetProcessID retrieves the process identifier corresponding
	// to the specified process handle.
	GetProcessID(windows.Handle) (uint32, error)

	// GetExitCodeProcess retrieves the termination status of the
	// specified process handle.
	GetExitCodeProcess(windows.Handle, *uint32) error

	// IsCompareObjectHandlesFound returns true if the CompareObjectHandles
	// function could be found in this Windows instance
	IsCompareObjectHandlesFound() bool
}

type systemCall struct {
}

func (s *systemCall) CloseHandle(h windows.Handle) error { _ = "STUB: not implemented"; return nil }

func (s *systemCall) IsCompareObjectHandlesFound() bool { _ = "STUB: not implemented"; return false }

func (s *systemCall) CompareObjectHandles(h1, h2 windows.Handle) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *systemCall) GetExitCodeProcess(h windows.Handle, exitCode *uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *systemCall) GetProcessID(h windows.Handle) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *systemCall) OpenProcess(pid int32) (handle windows.Handle, err error) {
	_ = "STUB: not implemented"
	return *new(windows.Handle), nil
}
