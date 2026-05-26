package errorutil

import (
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
)

// PermissionDenied formats a PermissionDenied error with an error string.
func PermissionDenied(reason types.PermissionDeniedDetails_Reason, format string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}
