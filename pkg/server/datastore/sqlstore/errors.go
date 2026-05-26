package sqlstore

const (
	datastoreSQLErrorPrefix        = "datastore-sql"
	datastoreValidationErrorPrefix = "datastore-validation"
)

type sqlError struct {
	err error
	msg string
}

func (s *sqlError) Error() string { _ = "STUB: not implemented"; return "" }

func (s *sqlError) Unwrap() error { _ = "STUB: not implemented"; return nil }

type validationError struct {
	err error
	msg string
}

func (v *validationError) Error() string { _ = "STUB: not implemented"; return "" }

func (v *validationError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func newSQLError(fmtMsg string, args ...any) error { _ = "STUB: not implemented"; return nil }

func newWrappedSQLError(err error) error { _ = "STUB: not implemented"; return nil }

func newValidationError(fmtMsg string, args ...any) error { _ = "STUB: not implemented"; return nil }

func newWrappedValidationError(err error) error { _ = "STUB: not implemented"; return nil }
