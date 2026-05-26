package bundleutil

import (
	"time"

	"github.com/spiffe/go-spiffe/v2/bundle/spiffebundle"
)

type marshalConfig struct {
	refreshHint    time.Duration
	sequenceNumber uint64
	noX509SVIDKeys bool
	noJWTSVIDKeys  bool
	standardJWKS   bool
}

type MarshalOption interface {
	configure(*marshalConfig) error
}

type marshalOption func(c *marshalConfig) error

func (o marshalOption) configure(c *marshalConfig) error {
	_ = "STUB: not implemented"

	// OverrideRefreshHint overrides the refresh hint in the bundle
	return nil
}

func OverrideRefreshHint(value time.Duration) MarshalOption {
	_ = "STUB: not implemented"
	return *new(MarshalOption)
}

// OverrideSequenceNumber overrides the sequence number in the bundle
func OverrideSequenceNumber(value uint64) MarshalOption {
	_ = "STUB: not implemented"
	return *new(MarshalOption)
}

// NoX509SVIDKeys skips marshalling X509 SVID keys
func NoX509SVIDKeys() MarshalOption { _ = "STUB: not implemented"; return *new(MarshalOption) }

// NoJWTSVIDKeys skips marshalling JWT SVID keys
func NoJWTSVIDKeys() MarshalOption { _ = "STUB: not implemented"; return *new(MarshalOption) }

// StandardJWKS omits SPIFFE-specific parameters from the marshaled bundle
func StandardJWKS() MarshalOption { _ = "STUB: not implemented"; return *new(MarshalOption) }

func Marshal(bundle *spiffebundle.Bundle, opts ...MarshalOption) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
