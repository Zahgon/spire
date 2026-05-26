package tpmutil

import (
	"io"

	"github.com/google/go-tpm/tpmutil"
	"github.com/hashicorp/go-hclog"
)

// ekRSACertificateHandle is the default handle for RSA endorsement key according
// to the TCG TPM v2.0 Provisioning Guidance, section 7.8
// https://trustedcomputinggroup.org/resource/tcg-tpm-v2-0-provisioning-guidance/
const EKCertificateHandleRSA = tpmutil.Handle(0x01c00002)

// randomPasswordSize is the number of bytes of generated random passwords
const randomPasswordSize = 32

// Session represents a TPM with loaded DevID credentials and exposes methods
// to perform cryptographic operations relevant to the SPIRE node attestation
// workflow.
type Session struct {
	devID    *SigningKey
	ak       *SigningKey
	ekHandle tpmutil.Handle
	ekPub    []byte
	akPub    []byte

	endorsementHierarchyPassword string
	ownerHierarchyPassword       string

	rwc io.ReadWriteCloser
	log hclog.Logger
}

type TPMPasswords struct {
	EndorsementHierarchy string
	OwnerHierarchy       string
	DevIDKey             string
}

type SessionConfig struct {
	// in future iterations of tpm libraries, TPM will accept a
	// list of device paths (https://github.com/google/go-tpm/pull/256)
	DevicePath string
	DevIDPriv  []byte
	DevIDPub   []byte
	Passwords  TPMPasswords
	Log        hclog.Logger
}

var OpenTPM = openTPM

// NewSession opens a connection to a TPM and configures it to be used for
// node attestation.
func NewSession(scfg *SessionConfig) (*Session, error) { _ = "STUB: not implemented"; return nil, nil }

// Open TPM connection

// Create session

// Close session in case of error

// Create SRK password

// Load DevID

// Create Attestation Key

// Load Attestation Key

// Regenerate Endorsement Key using the default RSA template

// Close unloads TPM loaded objects and closes the connection to the TPM.
func (c *Session) Close() { _ = "STUB: not implemented"; return }

// SolveDevIDChallenge requests the TPM to sign the provided nonce using the loaded
// DevID credentials.
func (c *Session) SolveDevIDChallenge(nonce []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SolveCredActivationChallenge runs credential activation on the TPM. It proves
// that the attestation key resides on the same TPM as the endorsement key.
func (c *Session) SolveCredActivationChallenge(credentialBlob, secret []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Flush only in case of error. If the command executes successfully it
// closes the session. Closing it again produces an error.

// CertifyDevIDKey proves that the DevID Key is in the same TPM than
// Attestation Key.
func (c *Session) CertifyDevIDKey() ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetEKCert returns TPM endorsement certificate.
func (c *Session) GetEKCert() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// In some TPMs, when we read bytes from an NV index, the content read
// includes the DER encoded x.509 certificate + trailing data. We need to
// remove those trailing bytes in order to make the certificate parseable by
// the server that uses x509.ParseCertificate().

// GetEKPublic returns the public part of the Endorsement Key encoded in
// TPM wire format.
func (c *Session) GetEKPublic() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// GetAKPublic returns the public part of the attestation key encoded in
// TPM wire format.
func (c *Session) GetAKPublic() []byte {
	_ = "STUB: not implemented"

	// loadKey loads a key pair into the TPM.
	return nil
}

func (c *Session) loadKey(publicKey, privateKey []byte, parentKeyPassword, keyPassword string) (*SigningKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Session) createAttestationKey(parentKeyPassword, keyPassword string) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// createPolicySessionForEK creates a session-based authorization to access EK.
// We need a session-based authorization to run the activate credential command
// (password-based auth is not enough) because of the attributes of the EK template.
func (c *Session) createPolicySessionForEK() (tpmutil.Handle, error) {
	_ = "STUB: not implemented"
	// The TPM is accessed in a plain session (we assume the bus is trusted) so we use an:
	// un-bounded and un-salted policy session (bindKey = HandleNull, tpmKey = HandleNull, secret = nil,
	// (sym = algNull, nonceCaller = all zeros).
	return *new(tpmutil.Handle), nil
}

// A detailed description of this command and its parameters can be found in TCG spec:
// https://www.trustedcomputinggroup.org/wp-content/uploads/TPM-Rev-2.0-Part-3-Commands-01.38.pdf#page=52

// rw:		TPM channel.
// tpmKey:		Handle to a key to do the decryption of encryptedSalt.
// bindKey:		Handle to a key to bind this session to (concatenates to salt).
// nonceCaller:	Initial nonce from the caller.
// secret:		Encrypted salt.
// se:		Session type.
// sym:		The type of parameter encryption that will be used when the session is set for encrypt or decrypt.
// hashAlg:		The hash algorithm used in computation of the policy digest.

// A detailed description of this command and its parameters can be found in TCG spec:
// https://www.trustedcomputinggroup.org/wp-content/uploads/TPM-Rev-2.0-Part-3-Commands-01.38.pdf#page=228

// 	rw:		TPM channel.
// 	entityHandle:	handle for an entity providing the authorization.
// 		entityAuth:	entity authorization.

// policyHandle:	Handle for the policy session being extended.
// policyNonce:	The policy nonce for the session (can be the Empty Buffer).
// cpHash:		Digest of the command parameters to which this authorization is limited (if it is not limited, the parameter will be the Empty Buffer).
// policyRef:		Reference to a policy relating to the authorization.
// expiry: 		Time when authorization will expire measured in seconds (zero means no expiration).

func (c *Session) flushContext(handle tpmutil.Handle) { _ = "STUB: not implemented"; return }

func newRandomPassword() (string, error) { _ = "STUB: not implemented"; return "", nil }
