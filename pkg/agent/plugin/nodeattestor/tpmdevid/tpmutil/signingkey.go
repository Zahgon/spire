package tpmutil

import (
	"io"

	"github.com/google/go-tpm/legacy/tpm2"
	"github.com/google/go-tpm/tpmutil"
	"github.com/hashicorp/go-hclog"
)

// maxAttempts indicates the max number retries for running TPM commands when
// TPM responds with a tpm2.RCRetry code.
const maxAttempts = 10

// SigningKey represents a TPM loaded key
type SigningKey struct {
	Handle     tpmutil.Handle
	sigHashAlg tpm2.Algorithm
	rw         io.ReadWriter
	log        hclog.Logger
	password   string
}

// Close removes the key from the TPM
func (k *SigningKey) Close() error { _ = "STUB: not implemented"; return nil }

// Sign requests the TPM to sign the given data using this key
func (k *SigningKey) Sign(data []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Certify calls tpm2.Certify using the current key as signer and the provided
// handle as object.
func (k *SigningKey) Certify(object tpmutil.Handle, objectPassword string) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	// For some reason 'tpm2.Certify()' sometimes fails the first attempt and asks for retry.
	// So, we retry in case of getting the RCRetry error.
	// It seems that this issue has been reported: https://github.com/google/go-tpm/issues/59
	return nil, nil, nil
}

// SRKTemplateHighRSA returns the default high range SRK template (called H-1 in the specification).
// https://trustedcomputinggroup.org/wp-content/uploads/TCG_IWG_EKCredentialProfile_v2p3_r2_pub.pdf#page=41
func SRKTemplateHighRSA() tpm2.Public {
	_ = "STUB: not implemented"
	// The client library does not have a function to build the high range template
	// so we build it based on the previous template.
	return *new(tpm2.Public)
}

// SRKTemplateHighECC returns the default high range SRK template (called H-2 in the specification).
// https://trustedcomputinggroup.org/wp-content/uploads/TCG_IWG_EKCredentialProfile_v2p3_r2_pub.pdf#page=42
func SRKTemplateHighECC() tpm2.Public {
	_ = "STUB: not implemented"
	// The client library does not have a function to build the high range template
	// so we build it based on the previous template.
	return *new(tpm2.Public)
}

// isRetry returns true if the given error is a tpm2.Warning that requests retry.
func isRetry(err error) bool { _ = "STUB: not implemented"; return false }

func getSignatureBytes(sig *tpm2.Signature) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
