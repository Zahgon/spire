package azure

import (
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/common/agentpathtemplate"
)

const (
	ImdsPluginName = "azure_imds"
)

var DefaultIMDSAgentPathTemplate = agentpathtemplate.MustParse("/{{ .PluginName }}/{{ .TenantID }}/{{ .SubscriptionID }}/{{ .VMID }}")

// AgentUntrustedMetadata is the untrusted metadata for the IMDS attestation payload.
// Used to help point the server to the correct tenant and VMSS
type AgentUntrustedMetadata struct {
	AgentDomain string  `json:"agentDomain"`
	VMSSName    *string `json:"vmssName"`
}

type IMDSAttestationPayload struct {
	Document AttestedDocument `json:"document"`
	// Nothing in the metadata should ever be trusted, it is used to help point the server to the correct tenant and VMSS
	Metadata AgentUntrustedMetadata `json:"metadata"`
}

type AttestedDocument struct {
	Encoding  string `json:"encoding"`
	Signature string `json:"signature"`
}

type AttestedDocumentContent struct {
	SubscriptionID string `json:"subscriptionId"`
	VMID           string `json:"vmId"`
	Nonce          string `json:"nonce"`
	// TenantID does not actually come from the document, it is added by the server for convenience
	TenantID string `json:"tid"`
}

func FetchAttestedDocument(cl HTTPClient, nonce string) (*AttestedDocument, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type imdsAgentPathTemplateData struct {
	*AttestedDocumentContent
	PluginName string
}

func MakeIMDSAgentID(td spiffeid.TrustDomain, agentPathTemplate *agentpathtemplate.Template, data *AttestedDocumentContent) (spiffeid.ID, error) {
	_ = "STUB: not implemented"
	return *new(spiffeid.ID), nil
}
