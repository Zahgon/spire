//go:build !windows

package log

import (
	"context"
	"os"

	"golang.org/x/sys/unix"
)

const (
	reopenSignal      = unix.SIGUSR2
	failedToReopenMsg = "failed to rotate log after signal"
)

// ReopenOnSignal returns a function compatible with RunTasks.
func ReopenOnSignal(logger *Logger, reopener Reopener) func(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func reopenOnSignal(
	ctx context.Context,
	logger *Logger,
	reopener Reopener,
	signalCh chan os.Signal,
) error {
	_ = "STUB: not implemented"
	return nil
}

// never fail; best effort to log to old file descriptor
