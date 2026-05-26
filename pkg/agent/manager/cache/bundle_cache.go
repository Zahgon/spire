package cache

import (
	"github.com/imkira/go-observer"
	"github.com/spiffe/go-spiffe/v2/bundle/spiffebundle"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
)

type Bundle = spiffebundle.Bundle

type BundleCache struct {
	trustDomain spiffeid.TrustDomain
	bundles     observer.Property
}

func NewBundleCache(trustDomain spiffeid.TrustDomain, bundle *Bundle) *BundleCache {
	_ = "STUB: not implemented"
	return nil
}

func (c *BundleCache) Update(bundles map[spiffeid.TrustDomain]*Bundle) {
	_ = "STUB: not implemented"
	// the bundle map must be copied so that the source can be mutated
	// afterward.
	return
}

func (c *BundleCache) Bundle() *Bundle { _ = "STUB: not implemented"; return nil }

func (c *BundleCache) Bundles() map[spiffeid.TrustDomain]*Bundle {
	_ = "STUB: not implemented"
	return nil
}

func (c *BundleCache) SubscribeToBundleChanges() *BundleStream {
	_ = "STUB: not implemented"
	return nil
}

// Wraps an observer stream to provide a type safe interface
type BundleStream struct {
	stream observer.Stream
}

func NewBundleStream(stream observer.Stream) *BundleStream { _ = "STUB: not implemented"; return nil }

// Value returns the current value for this stream.
func (b *BundleStream) Value() map[spiffeid.TrustDomain]*Bundle {
	_ = "STUB: not implemented"
	return nil
}

// Changes returns the channel that is closed when a new value is available.
func (b *BundleStream) Changes() chan struct{} { _ = "STUB: not implemented"; return nil }

// Next advances this stream to the next state.
// You should never call this unless Changes channel is closed.
func (b *BundleStream) Next() map[spiffeid.TrustDomain]*Bundle {
	_ = "STUB: not implemented"
	return nil
}

// HasNext checks whether there is a new value available.
func (b *BundleStream) HasNext() bool { _ = "STUB: not implemented"; return false }

// WaitNext waits for Changes to be closed, advances the stream and returns
// the current value.
func (b *BundleStream) WaitNext() map[spiffeid.TrustDomain]*Bundle {
	_ = "STUB: not implemented"
	return nil
}

// Clone creates a new independent stream from this one but sharing the same
// Property. Updates to the property will be reflected in both streams, but
// they may have different values depending on when they advance the stream
// with Next.
func (b *BundleStream) Clone() *BundleStream { _ = "STUB: not implemented"; return nil }

// copyBundleMap does a shallow copy of the bundle map.
func copyBundleMap(bundles map[spiffeid.TrustDomain]*Bundle) map[spiffeid.TrustDomain]*Bundle {
	_ = "STUB: not implemented"
	return nil
}
