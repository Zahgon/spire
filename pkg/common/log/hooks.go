package log

import (
	"github.com/sirupsen/logrus"
)

// LocalTimeHook is a logrus hook that converts all log fields with type time.Time to local time.
type LocalTimeHook struct{}

// Levels defines on which log levels this hook would trigger.
func (l LocalTimeHook) Levels() []logrus.Level { _ = "STUB: not implemented"; return nil }

// Fire is called when one of the log levels defined in Levels() is triggered.
func (l LocalTimeHook) Fire(entry *logrus.Entry) error {
	_ = "STUB: not implemented"
	// Convert all log fields with type time.Time to local time.
	return nil
}
