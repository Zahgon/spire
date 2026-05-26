package ca

import (
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/common/health"
	"github.com/spiffe/spire/pkg/common/pemutil"
)

var (
	caHealthKey, _ = pemutil.ParsePublicKey([]byte(`-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEzLY1/SRlsMJExTnuvzBO292RjGjU
3L8jFRtmQl0CjBeHdxUlGK1OkNLDYh0b6AW4siWt+y+DcbUAWNb14e5zWg==
-----END PUBLIC KEY-----`))
)

type caHealth struct {
	ca ServerCA
	td spiffeid.TrustDomain
}

func (h *caHealth) CheckHealth() health.State {
	_ = "STUB: not implemented"
	// Prevent a problem with signing the SVID from blocking the health check
	// indefinitely.
	return *new(health.State)
}

// Both liveness and readiness are determined by whether the
// x509 CA was successfully signed.

type caHealthDetails struct {
	SignX509SVIDErr string `json:"sign_x509_svid_err,omitempty"`
}

func errString(err error) string { _ = "STUB: not implemented"; return "" }
