package log

import (
	"io"
	"regexp"
	"runtime"

	"github.com/sirupsen/logrus"
)

const (
	DefaultFormat = ""
	JSONFormat    = "JSON"
	TextFormat    = "TEXT"
)

// An Option can change the Logger to apply desired configuration in NewLogger
type Option func(*Logger) error

// WithOutputFile requires lossy copytruncate directive in logrotate.
func WithOutputFile(file string) Option { _ = "STUB: not implemented"; return *new(Option) }

// If, for some reason, there's another closer set, close it first.

func WithOutputWriter(w io.Writer) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithReopenableOutputFile uses ReopenableFile to support handling a signal
// to rotate log files (e.g. from a logrotate postrotate script).
func WithReopenableOutputFile(reopenableFile *ReopenableFile) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// If, for some reason, there's another closer set, close it first.

func WithFormat(format string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Logrus has a default formatter set up in logrus.New(), so we don't change it

func WithLevel(logLevel string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithSourceLocation() Option { _ = "STUB: not implemented"; return *new(Option) }

// logrus provides a built-in feature that is very close to what we
// want (logger.SetReportCaller). Unfortunately, it always reports the
// immediate caller; but in certain cases, we want to skip over some
// more frames; in particular, this applies to the HCLogAdapter.

type sourceLocHook struct{}

func (sourceLocHook) Levels() []logrus.Level { _ = "STUB: not implemented"; return nil }

func (sourceLocHook) Fire(e *logrus.Entry) error { _ = "STUB: not implemented"; return nil }

func getCaller() *runtime.Frame { _ = "STUB: not implemented"; return nil }

// skip 'runtime.Callers', this function, and its caller

// skip over frames within the logging infrastructure

var loggingFuncRegexp = regexp.MustCompile(
	`^github\.com/(?:sirupsen/logrus|spiffe/spire/pkg/common/log)[./]`)

func isLoggingFunc(funcName string) bool { _ = "STUB: not implemented"; return false }
