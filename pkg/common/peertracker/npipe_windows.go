//go:build windows

package peertracker

import (
	"net"

	"golang.org/x/sys/windows"
)

var (
	kernelbase = windows.NewLazyDLL("kernelbase.dll")
	kernel32   = windows.NewLazyDLL("kernel32.dll")

	procCompareObjectHandles           = kernelbase.NewProc("CompareObjectHandles")
	procCompareObjectHandlesErr        = procCompareObjectHandles.Find()
	procGetNamedPipeClientProcessID    = kernel32.NewProc("GetNamedPipeClientProcessId")
	procGetNamedPipeClientProcessIDErr = procGetNamedPipeClientProcessID.Find()
)

func getCallerInfoFromNamedPipeConn(conn net.Conn) (CallerInfo, error) {
	_ = "STUB: not implemented"
	return *new(CallerInfo), nil
}

// getNamedPipeClientProcessID retrieves the client process identifier
// for the specified handle representing a named pipe.
func getNamedPipeClientProcessID(pipe windows.Handle, clientProcessID *int32) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func isCompareObjectHandlesFound() bool { _ = "STUB: not implemented"; return false }

// compareObjectHandles compares two object handles to determine if they
// refer to the same underlying kernel object
func compareObjectHandles(firstHandle, secondHandle windows.Handle) error {
	_ = "STUB: not implemented"
	return nil
}
