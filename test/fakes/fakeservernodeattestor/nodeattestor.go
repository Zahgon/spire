package fakeservernodeattestor

import (
	"testing"

	nodeattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/nodeattestor/v1"
	"github.com/spiffe/spire/pkg/server/plugin/nodeattestor"
)

const (
	defaultTrustDomain = "example.org"
)

type Config struct {
	// TrustDomain is the trust domain for SPIFFE IDs created by the attestor.
	// Defaults to "example.org" if empty.
	TrustDomain string

	// Payloads is a map from attestation payload (as a string) to the
	// associated id produced by the attestor. For example, a mapping from
	// "DATA" ==> "FOO means that an attestation request with the data "DATA"
	// would result in an attestation response with the SPIFFE ID:
	//
	// spiffe://<trustdomain>/spire/agent/<name>/<ID>
	//
	// For example, "spiffe://example.org/spire/agent/foo/bar"
	// In case ReturnLiteral is true value will be returned as base id
	Payloads map[string]string

	// Challenges is a map from ID to a list of echo challenges. The response
	// to each challenge is expected to match the challenge value.
	Challenges map[string][]string

	// Selectors is a map from ID to a list of selector values to return with that id.
	Selectors map[string][]string

	// Return literal from Payloads map
	ReturnLiteral bool
}

func New(t *testing.T, name string, config Config) nodeattestor.NodeAttestor {
	_ = "STUB: not implemented"
	return *new(nodeattestor.NodeAttestor)
}

type nodeAttestor struct {
	nodeattestorv1.UnsafeNodeAttestorServer

	name   string
	config Config
}

func (p *nodeAttestor) Attest(stream nodeattestorv1.NodeAttestor_AttestServer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// challenge/response loop

func (p *nodeAttestor) getAgentID(id string) string { _ = "STUB: not implemented"; return "" }
