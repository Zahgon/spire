package audit

import (
	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
)

const (
	message = "API accessed"
)

type Logger interface {
	AddFields(logrus.Fields)
	Audit()
	AuditWithFields(logrus.Fields)
	AuditWithTypesStatus(logrus.Fields, *types.Status)
	AuditWithError(error)
}

type logger struct {
	fields logrus.Fields
	log    logrus.FieldLogger
}

func New(l logrus.FieldLogger) Logger { _ = "STUB: not implemented"; return *new(Logger) }

// It is success by default, errors must change it

func (l *logger) AddFields(fields logrus.Fields) { _ = "STUB: not implemented"; return }

func (l *logger) Audit() { _ = "STUB: not implemented"; return }

func (l *logger) AuditWithFields(fields logrus.Fields) { _ = "STUB: not implemented"; return }

func (l *logger) AuditWithError(err error) { _ = "STUB: not implemented"; return }

func (l *logger) AuditWithTypesStatus(fields logrus.Fields, s *types.Status) {
	_ = "STUB: not implemented"
	return
}

func fieldsFromStatus(s *types.Status) logrus.Fields {
	_ = "STUB: not implemented"
	return *new(logrus.Fields)
}

func fieldsFromError(err error) logrus.Fields {
	_ = "STUB: not implemented"
	return *

	// Unknown status is returned for non-proto status
	new(logrus.Fields)
}
