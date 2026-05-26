//go:build !windows

package middleware

import (
	"github.com/shirou/gopsutil/v4/process"
	"github.com/sirupsen/logrus"
)

// setFields sets audit log fields specific to the Unix platforms.
func setFields(p *process.Process, fields logrus.Fields) error {
	_ = "STUB: not implemented"
	return nil
}

func getUID(p *process.Process) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func getGID(p *process.Process) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }
