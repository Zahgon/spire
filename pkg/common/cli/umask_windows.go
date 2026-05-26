//go:build windows

package cli

import "github.com/sirupsen/logrus"

// SetUmask does nothing on Windows
func SetUmask(logrus.FieldLogger) {
	_ = "STUB: not implemented"
	// Nothing to do in this platform
	return
}
