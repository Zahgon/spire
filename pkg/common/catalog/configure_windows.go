package catalog

import (
	"context"

	"github.com/sirupsen/logrus"
)

func ReconfigureOnSignal(ctx context.Context, _ logrus.FieldLogger, _ Reconfigurer) error {
	_ = "STUB: not implemented"
	// TODO: maybe drive this using an event?
	return nil
}
