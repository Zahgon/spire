package cli

import (
	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
)

// maxTrustDomainLength is the maximum length of a trust domain according
// to the SPIFFE standard.
const maxTrustDomainLength = 255

// ParseTrustDomain parses a configured trustDomain in a consistent way
// for either the SPIRE agent or server.
func ParseTrustDomain(trustDomain string, logger logrus.FieldLogger) (spiffeid.TrustDomain, error) {
	_ = "STUB: not implemented"
	return *new(spiffeid.TrustDomain), nil
}

func WarnOnLongTrustDomainName(td spiffeid.TrustDomain, logger logrus.FieldLogger) {
	_ = "STUB: not implemented"
	// Warn on a non-conforming trust domain to avoid breaking backwards compatibility
	return
}
