//go:build !windows

package cli

import (
	"github.com/sirupsen/logrus"
)

// The umask for SPIRE processes should not allow write by group, or
// read/write/execute by everyone.
const minimumUmask = 0o027

// SetUmask sets the minimumUmask.
func SetUmask(log logrus.FieldLogger) {
	_ = "STUB: not implemented"
	// Otherwise, make sure the current umask meets the minimum.
	return
}
