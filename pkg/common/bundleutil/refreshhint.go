package bundleutil

import (
	"time"

	"github.com/spiffe/go-spiffe/v2/bundle/spiffebundle"
)

const (
	refreshHintLeewayFactor = 10

	// MinimumRefreshHint is the smallest refresh hint the client allows.
	// Anything smaller than the minimum will be reset to the minimum.
	MinimumRefreshHint = time.Minute
)

// CalculateRefreshHint is used to calculate the refresh hint for a given
// bundle. If the bundle already contains a refresh hint, then that is used,
// Otherwise, it looks at the lifetimes of the bundle contents and returns a
// fraction of the smallest. It is fairly aggressive but ensures clients don't
// miss a rotation period and lose their ability to fetch.
// TODO: reevaluate our strategy here when we rework the TTL story inside SPIRE.
func CalculateRefreshHint(bundle *spiffebundle.Bundle) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// TODO: look at JWT key lifetimes... requires us to track issued_at dates
// which we currently do not do.

// Set the refresh hint to a fraction of the smallest lifetime, if found.

func safeRefreshHint(refreshHint time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
