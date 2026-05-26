package sshpop

import (
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/common/agentpathtemplate"
	"golang.org/x/crypto/ssh"
)

type serverHandshakeState int
type clientHandshakeState int

const (
	stateServerInit serverHandshakeState = iota
	stateAttestationDataVerified
	stateChallengeIssued
	stateChallengeVerified
)

const (
	stateClientInit clientHandshakeState = iota
	stateProvidedAttestationData
	stateRespondedToChallenge
)

// ClientHandshake is a single-use object for an agent to do node attestation.
//
// The handshake comprises a state machine that is not goroutine safe.
type ClientHandshake struct {
	c     *Client
	state clientHandshakeState
}

// ServerHandshake is a single-use object for a server to do node attestation.
//
// The handshake comprises a state machine that is not goroutine safe.
type ServerHandshake struct {
	s        *Server
	cert     *ssh.Certificate
	hostname string
	nonce    []byte
	state    serverHandshakeState
}

type attestationData struct {
	Certificate []byte
}

type challengeRequest struct {
	Nonce []byte
}

type challengeResponse struct {
	Nonce     []byte
	Signature *ssh.Signature
}

func (c *ClientHandshake) AttestationData() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClientHandshake) RespondToChallenge(req []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ServerHandshake) VerifyAttestationData(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func decanonicalizeHostname(fqdn, domain string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *ServerHandshake) IssueChallenge() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ServerHandshake) VerifyChallengeResponse(res []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ServerHandshake) AgentID() (spiffeid.ID, error) {
	_ = "STUB: not implemented"
	return *new(spiffeid.ID), nil
}

func newNonce() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func combineNonces(challenge, response []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// write the challenge and response and ignore errors since it will not
// fail writing to the digest

func makeAgentID(td spiffeid.TrustDomain, agentPathTemplate *agentpathtemplate.Template, cert *ssh.Certificate, hostname string) (spiffeid.ID, error) {
	_ = "STUB: not implemented"
	return *new(spiffeid.ID), nil
}

// urlSafeSSHFingerprintSHA256 is a modified version of ssh.FingerprintSHA256
// that returns an unpadded, url-safe version of the fingerprint.
func urlSafeSSHFingerprintSHA256(pubKey ssh.PublicKey) string { _ = "STUB: not implemented"; return "" }
