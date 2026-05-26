//go:build !windows

package spiretest

const (
	fileNotFound = "no such file or directory"
)

func PathNotFound() string { _ = "STUB: not implemented"; return "" }

func FileNotFound() string { _ = "STUB: not implemented"; return "" }

func SocketFileNotFound() string { _ = "STUB: not implemented"; return "" }
