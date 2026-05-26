package authorizedentries

import "sync"

var (
	stringSetPool = sync.Pool{
		New: func() any {
			return make(stringSet)
		},
	}
)

type stringSet map[string]struct{}

func allocStringSet() stringSet { _ = "STUB: not implemented"; return *new(stringSet) }

func freeStringSet(set stringSet) { _ = "STUB: not implemented"; return }
