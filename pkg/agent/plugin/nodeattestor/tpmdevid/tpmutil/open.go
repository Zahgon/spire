//go:build !windows

package tpmutil

import (
	"io"
)

// openTPM open a channel to the TPM at the given path.
func openTPM(paths ...string) (io.ReadWriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadWriteCloser), nil
}

// closeTPM EmulatorReadWriteCloser type does not need to be closed. It closes
// the connection after each Read() call. Closing it again results in
// an error.
func closeTPM(closer io.ReadWriteCloser) bool { _ = "STUB: not implemented"; return false }
