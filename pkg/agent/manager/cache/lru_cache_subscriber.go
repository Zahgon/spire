package cache

import (
	"sync"

	"github.com/spiffe/spire/proto/spire/common"
)

type Subscriber interface {
	Updates() <-chan *WorkloadUpdate
	Finish()
}

type lruCacheSubscriber struct {
	cache   *LRUCache
	set     selectorSet
	setFree func()

	mu   sync.Mutex
	c    chan *WorkloadUpdate
	done bool
}

func newLRUCacheSubscriber(cache *LRUCache, selectors []*common.Selector) *lruCacheSubscriber {
	_ = "STUB: not implemented"
	return nil
}

func (s *lruCacheSubscriber) Updates() <-chan *WorkloadUpdate {
	_ = "STUB: not implemented"
	return nil
}

func (s *lruCacheSubscriber) Finish() { _ = "STUB: not implemented"; return }

func (s *lruCacheSubscriber) notify(update *WorkloadUpdate) { _ = "STUB: not implemented"; return }
