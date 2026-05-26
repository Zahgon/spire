package plugin

import (
	"github.com/sirupsen/logrus"
)

func NullLogger() logrus.FieldLogger { _ = "STUB: not implemented"; return *new(logrus.FieldLogger) }
