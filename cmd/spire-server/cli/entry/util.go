package entry

import (
	"io"

	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
)

func printEntry(e *types.Entry, printf func(string, ...any) error) {
	_ = "STUB: not implemented"
	return
}

// admin is rare, so only show admin if true to keep
// from muddying the output.

// idStringToProto converts a SPIFFE ID from the given string to *types.SPIFFEID
func idStringToProto(id string) (*types.SPIFFEID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func printableEntryID(id string) string { _ = "STUB: not implemented"; return "" }

// protoToIDString converts a SPIFFE ID from the given *types.SPIFFEID to string
func protoToIDString(id *types.SPIFFEID) string { _ = "STUB: not implemented"; return "" }

// parseFile parses JSON represented RegistrationEntries
// if path is "-" read JSON from STDIN
func parseFile(path string) ([]*types.Entry, error) { _ = "STUB: not implemented"; return nil, nil }

func parseEntryJSON(in io.Reader, path string) ([]*types.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StringsFlag defines a custom type for string lists. Doing
// this allows us to support repeatable string flags.
type StringsFlag []string

// String returns the string flag.
func (s *StringsFlag) String() string { _ = "STUB: not implemented"; return "" }

// Set appends the string flag.
func (s *StringsFlag) Set(val string) error { _ = "STUB: not implemented"; return nil }
