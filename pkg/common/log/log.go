package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
	io.Closer
}

func NewLogger(options ...Option) (*Logger, error) { _ = "STUB: not implemented"; return nil, nil }

func setHooks(logger *Logger) { _ = "STUB: not implemented"; return }

type nopCloser struct{}

func (nopCloser) Close() error { _ = "STUB: not implemented"; return nil }
