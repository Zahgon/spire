package spiretest

import (
	"testing"

	"github.com/sirupsen/logrus"
)

type LogEntry struct {
	Level   logrus.Level
	Message string
	Data    logrus.Fields
}

func AssertLogs(t *testing.T, entries []*logrus.Entry, expected []LogEntry) {
	_ = "STUB: not implemented"
	return
}

func AssertLogsAnyOrder(t *testing.T, entries []*logrus.Entry, expected []LogEntry) {
	_ = "STUB: not implemented"
	return
}

func AssertLastLogs(t *testing.T, entries []*logrus.Entry, expected []LogEntry) {
	_ = "STUB: not implemented"
	return
}

func AssertLogsContainEntries(t *testing.T, entries []*logrus.Entry, expectedEntries []LogEntry) {
	_ = "STUB: not implemented"
	return
}

func convertLogEntries(entries []*logrus.Entry) (out []LogEntry) {
	_ = "STUB: not implemented"
	return nil
}

func normalizeData(data logrus.Fields) logrus.Fields {
	_ = "STUB: not implemented"
	return *new(logrus.Fields)
}
