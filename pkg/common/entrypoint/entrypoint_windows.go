//go:build windows

package entrypoint

import (
	"context"

	"golang.org/x/sys/windows/svc"
)

type systemCaller interface {
	IsWindowsService() (bool, error)
	Run(name string, handler svc.Handler) error
}

type systemCall struct {
}

func (s *systemCall) IsWindowsService() (bool, error) {
	_ = "STUB: not implemented"
	// We are using a custom function because the svc.IsWindowsService() one still has an open issue in which it states
	// that it is not working properly in Windows containers: https://github.com/golang/go/issues/56335. Soon as we have
	// a fix for that, we can use the original function.
	return false, nil
}

func (s *systemCall) Run(name string, handler svc.Handler) error {
	_ = "STUB: not implemented"
	return nil
}

type EntryPoint struct {
	handler  svc.Handler
	runCmdFn func(ctx context.Context, args []string) int
	sc       systemCaller
}

func NewEntryPoint(runCmdFn func(ctx context.Context, args []string) int) *EntryPoint {
	_ = "STUB: not implemented"
	return nil
}

func (e *EntryPoint) Main() int {
	_ = "STUB: not implemented"
	// Determining if SPIRE is running as a Windows service is done
	// with a best-effort approach. If there is an error, just fallback
	// to the behavior of not running as a Windows service.
	return 0
}

// Since the service runs in its own process, the service name is ignored.
// https://learn.microsoft.com/en-us/windows/win32/api/winsvc/nf-winsvc-startservicectrldispatcherw

// isWindowsService is a copy of the svc.IsWindowsService() function, but without the parentProcess.SessionID == 0 check
// that is causing the issue in Windows containers, this logic is exactly the same from .NET runtime (>= 6.0.10).
func isWindowsService() (bool, error) {
	_ = "STUB: not implemented"
	// The below technique looks a bit hairy, but it's actually
	// exactly what the .NET runtime (>= 6.0.10) does for the similarly named function:
	// https://github.com/dotnet/runtime/blob/36bf84fc4a89209f4fdbc1fc201e81afd8be49b0/src/libraries/Microsoft.Extensions.Hosting.WindowsServices/src/WindowsServiceHelpers.cs#L20-L33
	// Specifically, it looks up whether the parent process is called "services".
	return false, nil
}
