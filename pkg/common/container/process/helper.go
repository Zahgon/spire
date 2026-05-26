//go:build windows

package process

import (
	"github.com/hashicorp/go-hclog"
	"golang.org/x/sys/windows"
)

const (
	containerPrefix = `\Container_`
)

type Helper interface {
	GetContainerIDByProcess(pID int32, log hclog.Logger) (string, error)
}

func CreateHelper() Helper { _ = "STUB: not implemented"; return *new(Helper) }

type helper struct {
	wapi API
}

// GetContainerIDByProcess gets the container ID from the provided process ID,
// on windows process that are running in a docker containers are grouped by Named Jobs,
// those Jobs has the container ID as name.
// In the format `\Container_${CONTAINER_ID}`
func (h *helper) GetContainerIDByProcess(pID int32, log hclog.Logger) (string, error) {
	_ = "STUB: not implemented"
	// Search all processes that run vmcompute.exe
	return "", nil
}

// Get current process. The handle must not be closed.

// Duplicate the process handle that we want to validate, with limited permissions.

// Verify if process ID is a vmcompute process

// Filter all handles related with vmcompute processes

// searchProcessByExeFile searches all the processes with specified exe file
func (h *helper) searchProcessByExeFile(exeFile string, log hclog.Logger) ([]uint32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *helper) getJobName(handle SystemHandleInformationExItem, currentProcess windows.Handle, childProcessHandle windows.Handle, log hclog.Logger) (string, error) {
	_ = "STUB: not implemented"
	// Open the handle associated with the process ID, with permissions to duplicate the handle
	return "", nil
}

// This is expected when trying to open process as a non admin user

// Duplicate handle to get information

// This is expected when trying to duplicate a process that
// is not managed by docker

// Filter no Jobs handlers

// Jobs created on Windows environments start with "\Container_"
