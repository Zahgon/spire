//go:build windows

package tpmutil

import (
	"io"
)

// openTPM open a channel to the TPM, Windows does not receive a path.
func openTPM(paths ...string) (io.ReadWriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadWriteCloser), nil
}

// closeTPM we must close always when running on windows
func closeTPM(io.ReadWriteCloser) bool { _ = "STUB: not implemented"; return false }
