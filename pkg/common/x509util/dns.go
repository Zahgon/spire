package x509util

import (
	"errors"
)

var (
	ErrTooManyWildcards         = errors.New("too many wildcards")
	ErrWildcardMustBeFirstLabel = errors.New("wildcard must be first label")
	ErrEmptyDomain              = errors.New("empty or only whitespace")
	ErrIDNAError                = errors.New("idna error")
	ErrDomainEndsWithDot        = errors.New("domain ends with dot")
	ErrWildcardOverlap          = errors.New("wildcard overlap")
	ErrNameMustBeASCII          = errors.New("name must be ascii")
	ErrLabelMismatchAfterIDNA   = errors.New("label mismatch after idna")
	errNoWildcardAllowed        = errors.New("wildcard not allowed")
)

func ValidateLabel(domain string) error { _ = "STUB: not implemented"; return nil }

func validNonwildcardLabel(domain string) error { _ = "STUB: not implemented"; return nil }

// Defensive check.

func CheckForWildcardOverlap(names []string) error { _ = "STUB: not implemented"; return nil }

// While we're checking, we don't need to care about wildcards

// Let's split this non-wildcard DNS name into its corresponding labels

// Let's now replace the first label with a wildcard
