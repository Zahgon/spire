package selector

// EqualSet determines whether two sets of selectors are equal or not
func EqualSet(a, b *set) bool { _ = "STUB: not implemented"; return false }

// Includes determines whether a given selector is present in a set
func Includes(set *set, item *Selector) bool { _ = "STUB: not implemented"; return false }

// IncludesSet returns true if s2 is included in s1. This is, all the s2 selectors
// are also present in s1.
func IncludesSet(s1, s2 *set) bool {
	_ = "STUB: not implemented"
	// If s2 has more elements than s1, it cannot be included.
	return false
}
