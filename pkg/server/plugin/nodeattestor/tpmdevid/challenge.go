package tpmdevid

import (
	"crypto/x509"

	"github.com/google/go-tpm/legacy/tpm2"
	devid "github.com/spiffe/spire/pkg/common/plugin/tpmdevid"
)

func newNonce(size int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func VerifyDevIDChallenge(cert *x509.Certificate, challenge, response []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCredActivationChallenge(akPub, ekPub tpm2.Public) (*devid.CredActivation, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func VerifyCredActivationChallenge(expectedNonce, responseNonce []byte) error {
	_ = "STUB: not implemented"
	return nil
}
