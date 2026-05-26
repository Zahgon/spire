package fakeagentnodeattestor

import (
	"testing"

	nodeattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/agent/nodeattestor/v1"
	"github.com/spiffe/spire/pkg/agent/plugin/nodeattestor"
)

type Config struct {
	// Fail indicates whether fetching attestation data should fail.
	Fail bool

	// Responses are a list of echo responses. The response to each challenge is
	// expected to match the challenge value.
	Responses []string
}

func New(t *testing.T, config Config) nodeattestor.NodeAttestor {
	_ = "STUB: not implemented"
	return *new(nodeattestor.NodeAttestor)
}

type nodeAttestor struct {
	nodeattestorv1.UnimplementedNodeAttestorServer

	config Config
}

func (p *nodeAttestor) AidAttestation(stream nodeattestorv1.NodeAttestor_AidAttestationServer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func makePayload() *nodeattestorv1.PayloadOrChallengeResponse {
	_ = "STUB: not implemented"
	return nil
}

func makeChallengeResponse(challengeResponse []byte) *nodeattestorv1.PayloadOrChallengeResponse {
	_ = "STUB: not implemented"
	return nil
}
