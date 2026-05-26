package util

import (
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/proto/spire/common"
)

func DedupRegistrationEntries(entries []*common.RegistrationEntry) []*common.RegistrationEntry {
	_ = "STUB: not implemented"
	return nil
}

func SortRegistrationEntries(entries []*common.RegistrationEntry) {
	_ = "STUB: not implemented"
	// first, sort the selectors for each entry, since the registration
	// entry comparison relies on them being sorted
	return
}

// second, sort the registration entries

func SortSelectors(selectors []*common.Selector) { _ = "STUB: not implemented"; return }

func compareRegistrationEntries(a, b *common.RegistrationEntry) int {
	_ = "STUB: not implemented"
	return 0
}

// The order of this switch clause matters. It ensures that sorting occurs by X509SvidTtl then JwtSvidTtl

func compareSelectors(a, b []*common.Selector) int { _ = "STUB: not implemented"; return 0 }

func compareSelector(a, b *common.Selector) int { _ = "STUB: not implemented"; return 0 }

func SortTypesEntries(entries []*types.Entry) {
	_ = "STUB: not implemented"
	// first, sort the selectors for each entry, since the registration
	// entry comparison relies on them being sorted
	return
}

// second, sort the registration entries

func SortTypesSelectors(selectors []*types.Selector) { _ = "STUB: not implemented"; return }

func compareTypesEntries(a, b *types.Entry) int { _ = "STUB: not implemented"; return 0 }

// The order of this switch clause matters. It ensures that sorting occurs by X509SvidTtl then JwtSvidTtl

func compareTypesSelectors(a, b []*types.Selector) int { _ = "STUB: not implemented"; return 0 }

func compareTypesSelector(a, b *types.Selector) int { _ = "STUB: not implemented"; return 0 }

func cloneRegistrationEntries(entries []*common.RegistrationEntry) []*common.RegistrationEntry {
	_ = "STUB: not implemented"
	return nil
}

func cloneRegistrationEntry(entry *common.RegistrationEntry) *common.RegistrationEntry {
	_ = "STUB: not implemented"
	return nil
}
