//go:build windows

package middleware

import (
	"github.com/shirou/gopsutil/v4/process"
	"github.com/sirupsen/logrus"
)

// setFields sets audit log fields specific to the Windows platform.
func setFields(p *process.Process, fields logrus.Fields) error {
	_ = "STUB: not implemented"
	return nil
}

// We don't set group information on Windows. Setting the primary group
// would be confusing, since it is used only by the POSIX subsystem.

func getUserSID(pID int32) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Retrieve an access token to describe the security context of
// the process from which we obtained the handle.
