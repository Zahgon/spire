//go:build windows

package process

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")
	modntdll    = windows.NewLazySystemDLL("ntdll.dll")

	procIsProcessInJob    = modkernel32.NewProc("IsProcessInJob")
	procIsProcessInJobErr = procIsProcessInJob.Find()

	procNtQueryObject               = modntdll.NewProc("NtQueryObject")
	procNtQueryObjectErr            = procNtQueryObject.Find()
	procNtQuerySystemInformation    = modntdll.NewProc("NtQuerySystemInformation")
	procNtQuerySystemInformationErr = procNtQuerySystemInformation.Find()
)

const (
	// ObjectInformationClass values used to call NtQueryObject (https://docs.microsoft.com/en-us/windows/win32/api/winternl/nf-winternl-ntqueryobject)
	ObjectNameInformationClass = 0x1
	ObjectTypeInformationClass = 0x2

	// Includes all processes in the system in the snapshot. (https://docs.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-createtoolhelp32snapshot)
	Th32csSnapProcess uint32 = 0x00000002
)

type API interface {
	// IsProcessInJob determines whether the process is running in the specified job.
	IsProcessInJob(procHandle windows.Handle, jobHandle windows.Handle, result *bool) error

	// GetObjectType gets the object type of the given handle
	GetObjectType(handle windows.Handle) (string, error)

	// GetObjectName gets the object name of the given handle
	GetObjectName(handle windows.Handle) (string, error)

	// QuerySystemExtendedHandleInformation retrieves Extended handle system information.
	QuerySystemExtendedHandleInformation() ([]SystemHandleInformationExItem, error)

	// CurrentProcess returns the handle for the current process.
	// It is a pseudo handle that does not need to be closed.
	CurrentProcess() windows.Handle

	// CloseHandle closes an open object handle.
	CloseHandle(h windows.Handle) error

	// OpenProcess returns an open handle
	OpenProcess(desiredAccess uint32, inheritHandle bool, pID uint32) (windows.Handle, error)

	// DuplicateHandle duplicates an object handle.
	DuplicateHandle(hSourceProcessHandle windows.Handle, hSourceHandle windows.Handle, hTargetProcessHandle windows.Handle, lpTargetHandle *windows.Handle, dwDesiredAccess uint32, bInheritHandle bool, dwOptions uint32) error

	// CreateToolhelp32Snapshot takes a snapshot of the specified processes, as well as the heaps, modules, and threads used by these processes.
	CreateToolhelp32Snapshot(flags uint32, pID uint32) (windows.Handle, error)

	// Process32First retrieves information about the first process encountered in a system snapshot.
	Process32First(snapshot windows.Handle, procEntry *windows.ProcessEntry32) error

	// Process32Next retrieves information about the next process recorded in a system snapshot.
	Process32Next(snapshot windows.Handle, procEntry *windows.ProcessEntry32) error
}

type api struct{}

func (a *api) IsProcessInJob(procHandle windows.Handle, jobHandle windows.Handle, result *bool) error {
	_ = "STUB: not implemented"
	return nil
}

// GetObjectType gets the object type of the given handle
func (a *api) GetObjectType(handle windows.Handle) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetObjectName gets the object name of the given handle
func (a *api) GetObjectName(handle windows.Handle) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (a *api) QuerySystemExtendedHandleInformation() ([]SystemHandleInformationExItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if no error

//nolint:nilerr

func (a *api) OpenProcess(desiredAccess uint32, inheritHandle bool, pID uint32) (windows.Handle, error) {
	_ = "STUB: not implemented"
	return *new(windows.Handle), nil
}

func (a *api) CloseHandle(h windows.Handle) error { _ = "STUB: not implemented"; return nil }

// CurrentProcess returns the handle for the current process.
// It is a pseudo handle that does not need to be closed.
func (a *api) CurrentProcess() windows.Handle {
	_ = "STUB: not implemented"
	return *new(windows.Handle)
}

func (a *api) DuplicateHandle(hSourceProcessHandle windows.Handle, hSourceHandle windows.Handle, hTargetProcessHandle windows.Handle, lpTargetHandle *windows.Handle, dwDesiredAccess uint32, bInheritHandle bool, dwOptions uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *api) CreateToolhelp32Snapshot(flags uint32, pID uint32) (windows.Handle, error) {
	_ = "STUB: not implemented"
	return *new(windows.Handle), nil
}

func (a *api) Process32First(snapshot windows.Handle, procEntry *windows.ProcessEntry32) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *api) Process32Next(snapshot windows.Handle, procEntry *windows.ProcessEntry32) error {
	_ = "STUB: not implemented"
	return nil
}

// System handle extended information item, returned by NtQuerySystemInformation (https://docs.microsoft.com/en-us/windows/win32/api/winternl/nf-winternl-ntquerysysteminformation)
type SystemHandleInformationExItem struct {
	Object                uintptr
	UniqueProcessID       uintptr
	HandleValue           uintptr
	GrantedAccess         uint32
	CreatorBackTraceIndex uint16
	ObjectTypeIndex       uint16
	HandleAttributes      uint32
	Reserved              uint32
}

// System extended handle information summary, returned by NtQuerySystemInformation (https://docs.microsoft.com/en-us/windows/win32/api/winternl/nf-winternl-ntquerysysteminformation)
type SystemExtendedHandleInformation struct {
	NumberOfHandles uintptr
	Reserved        uintptr
	Handles         [1]SystemHandleInformationExItem
}

// Object type returned by calling NtQueryObject function
type ObjectTypeInformation struct {
	TypeName               UnicodeString
	TotalNumberOfObjects   uint32
	TotalNumberOfHandles   uint32
	TotalPagedPoolUsage    uint32
	TotalNonPagedPoolUsage uint32
}

// Unicode string returned by NtQueryObject calls (https://docs.microsoft.com/en-us/windows/win32/api/subauth/ns-subauth-unicode_string)
type UnicodeString struct {
	Length        uint16
	AllocatedSize uint16
	WString       *byte
}

func (u UnicodeString) String() string {
	_ = "STUB: not implemented"

	// TODO: may we recover?
	return ""
}

func ntQueryObject(handle windows.Handle, objectInformationClass uint32, objectInformation *byte, objectInformationLength uint32, returnLength *uint32) windows.NTStatus {
	_ = "STUB: not implemented"
	return *new(windows.NTStatus)
}

func ntQuerySystemInformation(sysInfoClass int32, sysInfo unsafe.Pointer, sysInfoLen uint32, retLen *uint32) windows.NTStatus {
	_ = "STUB: not implemented"
	return *new(windows.NTStatus)
}

func ntStatusFromSyscall(r0 uintptr) windows.NTStatus {
	_ = "STUB: not implemented"
	// NTSTATUS is a 32-bit Windows ABI value even though syscall.SyscallN
	// returns it in a uintptr-sized register. Preserve the low 32 bits instead
	// of treating wider uintptr values as overflow; this keeps statuses with
	// the high bit set intact even if a platform sign-extends the register.
	return *new(windows.NTStatus)
}

//nolint:gosec // G115: intentional ABI conversion from syscall return value.
