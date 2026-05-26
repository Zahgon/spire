//go:build linux

package peertracker

func getCallerInfoFromFileDescriptor(fd uintptr) (CallerInfo, error) {
	_ = "STUB: not implemented"
	return *new(CallerInfo), nil
}
