package authorizedentries

import (
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
)

type selectorSet map[Selector]struct{}

func selectorSetFromProto(selectors []*types.Selector) selectorSet {
	_ = "STUB: not implemented"
	return *new(selectorSet)
}

// Returns true if sub is a subset of whole
func isSubset(sub, whole selectorSet) bool { _ = "STUB: not implemented"; return false }
