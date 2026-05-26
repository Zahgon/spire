package log

import (
	"io"
	"log"

	"github.com/hashicorp/go-hclog"
	"github.com/sirupsen/logrus"
)

// HCLogAdapter implements the hclog interface, and wraps it
// around a Logrus entry
type HCLogAdapter struct {
	log  logrus.FieldLogger
	name string
	args []any // key/value pairs if this logger was created via With()
}

func NewHCLogAdapter(log logrus.FieldLogger, name string) *HCLogAdapter {
	_ = "STUB: not implemented"
	return nil
}

// HCLog has one more level than we do. As such, we will never
// set trace level.
func (*HCLogAdapter) Trace(_ string, _ ...any) { _ = "STUB: not implemented"; return }

func (a *HCLogAdapter) Debug(msg string, args ...any) { _ = "STUB: not implemented"; return }

func (a *HCLogAdapter) Info(msg string, args ...any) { _ = "STUB: not implemented"; return }

func (a *HCLogAdapter) Warn(msg string, args ...any) { _ = "STUB: not implemented"; return }

func (a *HCLogAdapter) Error(msg string, args ...any) { _ = "STUB: not implemented"; return }

func (a *HCLogAdapter) Log(level hclog.Level, msg string, args ...any) {
	_ = "STUB: not implemented"
	return
}

func (a *HCLogAdapter) IsTrace() bool { _ = "STUB: not implemented"; return false }

func (a *HCLogAdapter) IsDebug() bool { _ = "STUB: not implemented"; return false }

func (a *HCLogAdapter) IsInfo() bool { _ = "STUB: not implemented"; return false }

func (a *HCLogAdapter) IsWarn() bool { _ = "STUB: not implemented"; return false }

func (a *HCLogAdapter) IsError() bool { _ = "STUB: not implemented"; return false }

func (a *HCLogAdapter) SetLevel(hclog.Level) {
	_ = "STUB: not implemented"
	// interface definition says it is ok for this to be a noop if
	// implementations don't need/want to support dynamic level changing, which
	// we don't currently.
	return
}

func (a *HCLogAdapter) GetLevel() hclog.Level {
	_ = "STUB: not implemented"
	// We don't support dynamically setting the level with SetLevel(),
	// so just return a default value here.
	return *new(hclog.Level)
}

func (a *HCLogAdapter) With(args ...any) hclog.Logger {
	_ = "STUB: not implemented"
	return *new(hclog.Logger)
}

// concatFields combines two sets of key/value pairs.
// It allocates a new slice to avoid using append() and
// accidentally overriding the original slice a, e.g.
// when logger.With() is called multiple times to create
// sub-scoped loggers.
func concatFields(a, b []any) []any { _ = "STUB: not implemented"; return nil }

// ImpliedArgs returns With key/value pairs
func (a *HCLogAdapter) ImpliedArgs() []any { _ = "STUB: not implemented"; return nil }

func (a *HCLogAdapter) Name() string { _ = "STUB: not implemented"; return "" }

func (a *HCLogAdapter) Named(name string) hclog.Logger {
	_ = "STUB: not implemented"
	return *new(hclog.Logger)
}

func (a *HCLogAdapter) ResetNamed(name string) hclog.Logger {
	_ = "STUB: not implemented"
	return *new(hclog.Logger)
}

// StandardLogger is meant to return a stdlib Logger type which wraps around
// hclog. It does this by providing an io.Writer and instantiating a new
// Logger. It then tries to interpret the log level by parsing the message.
//
// Since we are not using `hclog` in a generic way, and I cannot find any
// calls to this method from go-plugin, we will poorly support this method.
// Rather than pull in all of hclog writer parsing logic, pass it a Logrus
// writer, and hardcode the level to INFO.
//
// Apologies to those who find themselves here.
func (a *HCLogAdapter) StandardLogger(*hclog.StandardLoggerOptions) *log.Logger {
	_ = "STUB: not implemented"
	return nil
}

func (a *HCLogAdapter) StandardWriter(*hclog.StandardLoggerOptions) io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

func (a *HCLogAdapter) shouldEmit(level logrus.Level) bool { _ = "STUB: not implemented"; return false }

func (a *HCLogAdapter) CreateEntry(args []any) *logrus.Entry { _ = "STUB: not implemented"; return nil }
