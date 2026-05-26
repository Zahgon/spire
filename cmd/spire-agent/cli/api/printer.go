package api

import (
	"crypto/x509"
	"time"

	commoncli "github.com/spiffe/spire/pkg/common/cli"
)

func printX509SVIDResponse(env *commoncli.Env, svids []*X509SVID, respTime time.Duration) {
	_ = "STUB: not implemented"
	return
}

func printX509SVID(env *commoncli.Env, svid *X509SVID) {
	_ = "STUB: not implemented"
	// Print SPIFFE ID first so if we run into a problem, we
	// get to know which record it was
	return
}

func printX509FederatedBundle(env *commoncli.Env, trustDomain string, bundle []*x509.Certificate) {
	_ = "STUB: not implemented"
	return
}
