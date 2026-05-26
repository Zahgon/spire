package bundleutil

import (
	"io"

	"github.com/spiffe/go-spiffe/v2/bundle/spiffebundle"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
)

func Decode(trustDomain spiffeid.TrustDomain, r io.Reader) (*spiffebundle.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Unmarshal(trustDomain spiffeid.TrustDomain, data []byte) (*spiffebundle.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unmarshal(trustDomain spiffeid.TrustDomain, doc *bundleDoc) (*spiffebundle.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
