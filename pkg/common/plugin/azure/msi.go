package azure

import (
	"io"
	"net/http"

	"github.com/go-jose/go-jose/v4/jwt"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/common/agentpathtemplate"
)

const (
	// DefaultMSIResourceID is the default resource ID to use as the intended
	// audience of the MSI token. The current value is the service ID for the
	// Resource Manager API.
	DefaultMSIResourceID = "https://management.azure.com/"
	PluginName           = "azure_msi"
)

// DefaultAgentPathTemplate is the default text/template
var DefaultAgentPathTemplate = agentpathtemplate.MustParse("/{{ .PluginName }}/{{ .TenantID }}/{{ .PrincipalID }}")

type ComputeMetadata struct {
	Name              string `json:"name"`
	SubscriptionID    string `json:"subscriptionId"`
	ResourceGroupName string `json:"resourceGroupName"`
	VMScaleSetName    string `json:"vmScaleSetName"`
}

type InstanceMetadata struct {
	Compute ComputeMetadata `json:"compute"`
}

type MSIAttestationData struct {
	Token string `json:"token"`
}

type MSITokenClaims struct {
	jwt.Claims
	TenantID    string `json:"tid,omitempty"`
	PrincipalID string `json:"sub,omitempty"`
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type HTTPClientFunc func(*http.Request) (*http.Response, error)

func (fn HTTPClientFunc) Do(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FetchMSIToken(cl HTTPClient, resource string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func FetchInstanceMetadata(cl HTTPClient) (*InstanceMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type agentPathTemplateData struct {
	MSITokenClaims
	PluginName string
}

func MakeAgentID(td spiffeid.TrustDomain, agentPathTemplate *agentpathtemplate.Template, claims *MSITokenClaims) (spiffeid.ID, error) {
	_ = "STUB: not implemented"
	return *new(spiffeid.ID), nil
}

func tryRead(r io.Reader) string { _ = "STUB: not implemented"; return "" }
