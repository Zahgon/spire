//go:build !linux && !darwin && !freebsd && !netbsd && !openbsd

package peertracker

func getCallerInfoFromFileDescriptor(uintptr) (CallerInfo, error) {
	_ = "STUB: not implemented"
	return *new(CallerInfo), nil
}
