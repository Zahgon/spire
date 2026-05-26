package log

import (
	"io"
	"os"
	"sync"
)

const (
	fileFlags = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	fileMode  = 0640
)

var _ ReopenableWriteCloser = (*ReopenableFile)(nil)

type (
	// Reopener inspired by https://github.com/client9/reopen
	Reopener interface {
		Reopen() error
	}
	ReopenableWriteCloser interface {
		Reopener
		io.WriteCloser
	}
)

type (
	ReopenableFile struct {
		name      string
		f         *os.File
		closeFunc closeFunc
		mu        sync.Mutex
	}
	// closeFunc must be called while holding the lock. It is intended for
	// injecting errors under test.
	closeFunc func(*os.File) error
)

func NewReopenableFile(name string) (*ReopenableFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ReopenableFile) Reopen() error { _ = "STUB: not implemented"; return nil }

// Ignore errors closing old file descriptor since logger would be using
// file descriptor we fail to close. This could leak file descriptors.

func (r *ReopenableFile) Write(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *ReopenableFile) Close() error { _ = "STUB: not implemented"; return nil }

// Name implements part of os.FileInfo without needing a lock on the
// underlying file.
func (r *ReopenableFile) Name() string { _ = "STUB: not implemented"; return "" }
