package middleware

import (
	"github.com/sirupsen/logrus"
)

// WithLogger returns logging middleware that provides a per-rpc logger with
// some initial fields set. If unset, it also provides name metadata to the
// handler context.
func WithLogger(log logrus.FieldLogger) Middleware {
	_ = "STUB: not implemented"
	return *new(Middleware)
}
