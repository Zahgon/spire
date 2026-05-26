package cgroups

import (
	"io"
)

// Filesystem abstracts filesystem operations.
type FileSystem interface {
	// Open opens the named file for reading.
	Open(name string) (io.ReadCloser, error)
}

// Cgroup represents a linux cgroup.
type Cgroup struct {
	HierarchyID    string
	ControllerList string
	GroupPath      string
}

// GetCGroups returns a slice of cgroups for pid using fs for filesystem calls.
//
// The expected cgroup format is "hierarchy-ID:controller-list:cgroup-path", and
// this function will return an error if every cgroup does not meet that format.
//
// For more information, see:
//   - http://man7.org/linux/man-pages/man7/cgroups.7.html
//   - https://www.kernel.org/doc/Documentation/cgroup-v2.txt
func GetCgroups(pid int32, fs FileSystem) ([]Cgroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
