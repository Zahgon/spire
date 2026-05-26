//go:build !windows

package diskutil

import (
	"os"
)

const (
	fileModePrivate          = 0600
	fileModePubliclyReadable = 0644
)

// AtomicWritePrivateFile writes data out to a private file.
// It writes to a temp file first, fsyncs that file, then swaps the file in.
// It renames the file using MoveFileEx with  'MOVEFILE_WRITE_THROUGH',
// which waits until the file is synced to disk.
func AtomicWritePrivateFile(path string, data []byte) error { _ = "STUB: not implemented"; return nil }

// AtomicWritePubliclyReadableFile writes data out to a publicly readable file.
// It writes to a temp file first, fsyncs that file, then swaps the file in.
// It renames the file using MoveFileEx with  'MOVEFILE_WRITE_THROUGH',
// which waits until the file is synced to disk.
func AtomicWritePubliclyReadableFile(path string, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func CreateDataDirectory(path string) error { _ = "STUB: not implemented"; return nil }

// WritePrivateFile writes data out to a private file. The file is created if it
// does not exist. If exists, it's overwritten.
func WritePrivateFile(path string, data []byte) error { _ = "STUB: not implemented"; return nil }

// WritePubliclyReadableFile writes data out to a publicly readable file. The
// file is created if it does not exist. If exists, it's overwritten.
func WritePubliclyReadableFile(path string, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func atomicWrite(path string, data []byte, mode os.FileMode) error {
	_ = "STUB: not implemented"
	return nil
}

func rename(tmpPath, path string) error { _ = "STUB: not implemented"; return nil }

// write writes to a file in the specified path with the specified
// security descriptor using the provided data. The sync boolean
// argument is used to indicate whether flushing to disk is required
// or not.
func write(tmpPath string, data []byte, mode os.FileMode, sync bool) error {
	_ = "STUB: not implemented"
	return nil
}
