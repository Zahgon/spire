//go:build windows

package spiretest

const (
	pathNotFound       = "The system cannot find the path specified."
	fileNotFound       = "The system cannot find the file specified."
	socketFileNotFound = "No connection could be made because the target machine actively refused it."
)

func FileNotFound() string { _ = "STUB: not implemented"; return "" }

func PathNotFound() string { _ = "STUB: not implemented"; return "" }

func SocketFileNotFound() string { _ = "STUB: not implemented"; return "" }
