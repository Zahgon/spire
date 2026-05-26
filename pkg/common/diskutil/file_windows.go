//go:build windows

package diskutil

import (
	"golang.org/x/sys/windows"
)

const (
	movefileReplaceExisting = 0x1
	movefileWriteThrough    = 0x8
)

type fileAttribs struct {
	pathUTF16Ptr *uint16
	sa           *windows.SecurityAttributes
}

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

// MkdirAll is a modified version of os.MkdirAll for use on Windows
// so that it creates the directory with the specified security descriptor.
func MkdirAll(path string, sddl string) error {
	_ = "STUB: not implemented"
	// Fast path: if we can tell whether path is a directory or file, stop with success or error.
	return nil
}

// Slow path: make sure parent exists and then call Mkdir for path.

// Skip trailing path separator.

// Scan backward over element.

// Create parent.

// Parent now exists; invoke Mkdir and use its result.

// Handle arguments like "foo/." by
// double-checking that directory doesn't exist.

// WritePrivateFile writes data out to a private file. The file is created if it
// does not exist. If exists, it's overwritten.
func WritePrivateFile(path string, data []byte) error { _ = "STUB: not implemented"; return nil }

// WritePubliclyReadableFile writes data out to a publicly readable file. The
// file is created if it does not exist. If exists, it's overwritten.
func WritePubliclyReadableFile(path string, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func atomicWrite(path string, data []byte, sddl string) error {
	_ = "STUB: not implemented"
	return nil
}

// write writes to a file in the specified path with the specified
// security descriptor using the provided data. The sync boolean
// argument is used to indicate whether flushing to disk is required
// or not.
func write(path string, data []byte, sddl string, sync bool) error {
	_ = "STUB: not implemented"
	return nil
}

func createFileForWriting(path string, sddl string) (windows.Handle, error) {
	_ = "STUB: not implemented"
	return *new(windows.Handle), nil
}

func atomicRename(oldPath, newPath string) error { _ = "STUB: not implemented"; return nil }

func rename(oldPath, newPath string) error { _ = "STUB: not implemented"; return nil }

// mkdir creates a new directory with a specific security descriptor.
// The security descriptor must be specified using the Security Descriptor
// Definition Language (SDDL).
//
// In the same way as os.MkDir, errors returned are of type *os.PathError.
func mkdir(path string, sddl string) error { _ = "STUB: not implemented"; return nil }

func getFileWithSecurityAttr(path, sddl string) (*fileAttribs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
