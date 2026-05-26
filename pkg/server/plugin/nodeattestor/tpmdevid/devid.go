package tpmdevid

import (
	"context"
	"crypto/x509"
	"sync"

	"github.com/google/go-tpm/legacy/tpm2"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	nodeattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/nodeattestor/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	common_devid "github.com/spiffe/spire/pkg/common/plugin/tpmdevid"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

// We use a 32 bytes nonce to provide enough cryptographical randomness and to be
// consistent with other nonces sizes around the project.
const devIDChallengeNonceSize = 32

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

type Config struct {
	DevIDBundlePath       string `hcl:"devid_ca_path"`
	EndorsementBundlePath string `hcl:"endorsement_ca_path"`
}

type config struct {
	trustDomain spiffeid.TrustDomain

	devIDRoots *x509.CertPool
	ekRoots    *x509.CertPool
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *config {
	_ = "STUB: not implemented"
	return nil
}

// Create initial internal configuration

// Load DevID bundle

// Load endorsement bundle if configured

type Plugin struct {
	nodeattestorv1.UnsafeNodeAttestorServer
	configv1.UnsafeConfigServer

	m sync.Mutex
	c *config
}

func New() *Plugin { _ = "STUB: not implemented"; return nil }

func (p *Plugin) Attest(stream nodeattestorv1.NodeAttestor_AttestServer) error {
	_ = "STUB: not implemented"
	// Receive attestation request
	return nil
}

// Unmarshall received attestation data

// Decode attestation data

// Verify DevID certificate chain of trust

// Issue a DevID challenge (to prove the possession of the DevID private key).

// Verify DevID residency

// Marshal challenges

// Send challenges to the agent

// Receive challenges response

// Unmarshal challenges response

// Verify DevID challenge

// Verify credential activation challenge

// Create SPIFFE ID and selectors

func (p *Plugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) Validate(ctx context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) getConfiguration() *config { _ = "STUB: not implemented"; return nil }

func verifyDevIDSignature(cert *x509.Certificate, intermediates *x509.CertPool, roots *x509.CertPool) ([][]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// verifyDevIDResidency verifies that the DevID resides on the same TPM as EK.
// This is done in two steps:
// (1) Verify that the DevID resides in the same TPM as the AK
// (2) Verify that the AK is in the same TPM as the EK.
// The verification is complete once the agent solves the challenge that this
// function generates.
func verifyDevIDResidency(attData *common_devid.AttestationRequest, ekRoots *x509.CertPool) (*common_devid.CredActivation, []byte, error) {
	_ = "STUB: not implemented"
	// Check that request contains all the information required to validate DevID residency
	return nil, nil, nil
}

// Decode attestation data

// Verify the public part of the EK generated from the template is the same
// as the one in the EK certificate.

// Verify EK chain of trust using the provided manufacturer roots.

// Verify DevID resides in the same TPM than AK

// Issue a credential activation challenge (to verify AK is in the same TPM as EK)

func isDevIDResidencyInfoComplete(attReq *common_devid.AttestationRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func verifyEKSignature(ekCert *x509.Certificate, roots *x509.CertPool) error {
	_ = "STUB: not implemented"
	// Check UnhandledCriticalExtensions for OIDs that we know what to do about
	// it (e.g. it's safe to ignore)
	return nil
}

// Endorsement certificate's SAN is not fully processed by x509 package

// verifyEKsMatch checks that the public key generated using the EK template
// matches the public key included in the Endorsement Certificate.
func verifyEKsMatch(ekCert *x509.Certificate, ekPub tpm2.Public) error {
	_ = "STUB: not implemented"
	return nil
}

func VerifyDevIDCertification(pubAK, pubDevID *tpm2.Public, attestData, attestSig []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func checkSignature(pub *tpm2.Public, data, sigRaw []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func getSignatureScheme(pub tpm2.Public) (*tpm2.SigScheme, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
