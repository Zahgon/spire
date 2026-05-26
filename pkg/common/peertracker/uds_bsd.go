//go:build darwin || freebsd || netbsd || openbsd

package peertracker

func getCallerInfoFromFileDescriptor(fd uintptr) (CallerInfo, error) {
	_ = "STUB: not implemented"
	return *new(CallerInfo), nil
}

// getsockopt(fd, SOL_LOCAL, LOCAL_PEERPID)
