//go:build !windows

package catalog

import (
	"context"

	"github.com/sirupsen/logrus"
)

func ReconfigureOnSignal(ctx context.Context, log logrus.FieldLogger, reconfigurer Reconfigurer) error {
	_ = "STUB: not implemented"
	return nil
}
