package httpchallenge

import (
	"context"
	"net/http"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
)

const (
	nonceLen = 32

	// PluginName for http based attestor
	PluginName = "http_challenge"
)

type AttestationData struct {
	HostName  string `json:"hostname"`
	AgentName string `json:"agentname"`
	Port      int    `json:"port"`
}

type Challenge struct {
	Nonce string `json:"nonce"`
}

type Response struct {
}

func GenerateChallenge(forceNonce string) (*Challenge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CalculateResponse(_ *Challenge) (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

func VerifyChallenge(ctx context.Context, client *http.Client, attestationData *AttestationData, challenge *Challenge) error {
	_ = "STUB: not implemented"
	return nil
}

// MakeAgentID creates an agent ID
func MakeAgentID(td spiffeid.TrustDomain, hostName string) (spiffeid.ID, error) {
	_ = "STUB: not implemented"
	return *new(spiffeid.ID), nil
}

func generateNonce() (string, error) { _ = "STUB: not implemented"; return "", nil }
