//go:build windows

package entrypoint

import (
	"context"
	"sync"

	"golang.org/x/sys/windows/svc"
)

const supportedCommand = "run"

type service struct {
	mtx              sync.RWMutex
	executeServiceFn func(ctx context.Context, stop context.CancelFunc, args []string) int
}

func (s *service) Execute(args []string, changeRequest <-chan svc.ChangeRequest, status chan<- svc.Status) (svcSpecificEC bool, exitCode uint32) {
	_ = "STUB: not implemented"
	// Validate that we are executing the "run" command.
	// First argument (args[0]) is always the process name. Command name is
	// expected in the second argument (args[1]).
	return false, 0
}

// Update the status to indicate that SPIRE is running.
// Only Stop and Shutdown commands are accepted (Interrogate is always accepted).

//nolint:gosec // don't care about potential integer conversion overflow
