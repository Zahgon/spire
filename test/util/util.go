package util

import (
	"testing"
	"time"

	"github.com/spiffe/spire/proto/spire/common"
)

// ProjectRoot returns the absolute path to the SPIRE project root
func ProjectRoot() string { _ = "STUB: not implemented"; return "" }

// GetRegistrationEntriesMap gets a map of registration entries from a fixture
func GetRegistrationEntriesMap(fileName string) map[string][]*common.RegistrationEntry {
	_ = "STUB: not implemented"
	return nil
}

// RunWithTimeout runs code within the specified timeout, if execution
// takes longer than that, an error is logged to t with information
// about the caller of this function. Returns how much time it took to
// run the function.
func RunWithTimeout(t *testing.T, timeout time.Duration, code func()) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// make sure the done channel is sent on in the face of panic's or
// other unwinding events (e.g. runtime.Goexit via t.Fatal)
