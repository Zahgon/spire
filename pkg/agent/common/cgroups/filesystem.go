package cgroups

import (
	"io"
)

// OSFileSystem implements FileSystem using the local disk
type OSFileSystem struct{}

func (OSFileSystem) Open(name string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}
