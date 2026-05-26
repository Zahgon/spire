package cache

import (
	"sync"

	"github.com/spiffe/spire/proto/spire/common"
)

var (
	stringSetPool = sync.Pool{
		New: func() any {
			return make(stringSet)
		},
	}

	selectorSetPool = sync.Pool{
		New: func() any {
			return make(selectorSet)
		},
	}

	lruCacheRecordSetPool = sync.Pool{
		New: func() any {
			return make(lruCacheRecordSet)
		},
	}

	lruCacheSubscriberSetPool = sync.Pool{
		New: func() any {
			return make(lruCacheSubscriberSet)
		},
	}
)

// unique set of strings, allocated from a pool
type stringSet map[string]struct{}

func allocStringSet() (stringSet, func()) { _ = "STUB: not implemented"; return *new(stringSet), nil }

func clearStringSet(set stringSet) { _ = "STUB: not implemented"; return }

func (set stringSet) Merge(ss ...string) { _ = "STUB: not implemented"; return }

// unique set of selectors, allocated from a pool
type selector struct {
	Type  string
	Value string
}

func makeSelector(s *common.Selector) selector { _ = "STUB: not implemented"; return *new(selector) }

type selectorSet map[selector]struct{}

func allocSelectorSet(ss ...*common.Selector) (selectorSet, func()) {
	_ = "STUB: not implemented"
	return *new(selectorSet), nil
}

func clearSelectorSet(set selectorSet) { _ = "STUB: not implemented"; return }

func (set selectorSet) Merge(ss ...*common.Selector) { _ = "STUB: not implemented"; return }

func (set selectorSet) MergeSet(other selectorSet) { _ = "STUB: not implemented"; return }

func (set selectorSet) In(ss ...*common.Selector) bool { _ = "STUB: not implemented"; return false }

func (set selectorSet) SuperSetOf(other selectorSet) bool { _ = "STUB: not implemented"; return false }

// unique set of LRU cache records, allocated from a pool
type lruCacheRecordSet map[*lruCacheRecord]struct{}

func allocLRUCacheRecordSet() (lruCacheRecordSet, func()) {
	_ = "STUB: not implemented"
	return *new(lruCacheRecordSet), nil
}

func clearLRUCacheRecordSet(set lruCacheRecordSet) { _ = "STUB: not implemented"; return }

// unique set of LRU cache subscribers, allocated from a pool
type lruCacheSubscriberSet map[*lruCacheSubscriber]struct{}

func allocLRUCacheSubscriberSet() (lruCacheSubscriberSet, func()) {
	_ = "STUB: not implemented"
	return *new(lruCacheSubscriberSet), nil
}

func clearLRUCacheSubscriberSet(set lruCacheSubscriberSet) { _ = "STUB: not implemented"; return }
