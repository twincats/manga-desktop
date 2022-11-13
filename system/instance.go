package system

import (
	"syscall"
	"unsafe"
)

var (
	kernel32         = syscall.NewLazyDLL("kernel32.dll")
	ProcCreateMutexW = kernel32.NewProc("CreateMutexW")
	// procCloseHandle  = kernel32.NewProc("CloseHandle")
)

// https://learn.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-createmutexW#return-value
func CreateMutexW(proc *syscall.LazyProc, name string) (uintptr, error) {
	if proc.Name != "CreateMutexW" {
		panic("proc.Name != CreateMutexW")
	}
	lpName, _ := syscall.UTF16PtrFromString(name) // LPCWSTR
	if handleID, _, err := proc.Call(
		0,
		0,
		uintptr(unsafe.Pointer(lpName)),
	); err.(syscall.Errno) == 0 {
		return handleID, nil
	} else {
		return handleID, err
	}
}
