package spiretest

import (
	"testing"
)

// TempDir creates a temporary directory that is cleaned up when the test
// finishes.
// TODO: remove when go1.15 is out, which introduces a new method on
// *testing.T for this purpose.
func TempDir(tb testing.TB) string { _ = "STUB: not implemented"; return "" }
