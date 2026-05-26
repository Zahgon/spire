package awsiid

import (
	"context"
	"io"
	"sync"

	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
	"github.com/hashicorp/go-hclog"
	nodeattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/agent/nodeattestor/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	caws "github.com/spiffe/spire/pkg/common/plugin/aws"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

const (
	docPath        = "instance-identity/document"
	sigPath        = "instance-identity/signature"
	sigRSA2048Path = "instance-identity/rsa2048"
)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *IIDAttestorPlugin) catalog.BuiltIn {
	_ = "STUB: not implemented"
	return *new(catalog.BuiltIn)
}

// IIDAttestorConfig configures a IIDAttestorPlugin.
type IIDAttestorConfig struct {
	EC2MetadataEndpoint string `hcl:"ec2_metadata_endpoint"`
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *IIDAttestorConfig {
	_ = "STUB: not implemented"
	return nil
}

// IIDAttestorPlugin implements aws nodeattestation in the agent.
type IIDAttestorPlugin struct {
	nodeattestorv1.UnsafeNodeAttestorServer
	configv1.UnsafeConfigServer

	log    hclog.Logger
	config *IIDAttestorConfig
	mtx    sync.RWMutex
}

// New creates a new IIDAttestorPlugin.
func New() *IIDAttestorPlugin { _ = "STUB: not implemented"; return nil }

func (p *IIDAttestorPlugin) SetLogger(log hclog.Logger) {
	_ = "STUB: not implemented"

	// AidAttestation implements the NodeAttestor interface method of the same name
	return
}

func (p *IIDAttestorPlugin) AidAttestation(stream nodeattestorv1.NodeAttestor_AidAttestationServer) error {
	_ = "STUB: not implemented"
	return nil
}

func fetchMetadata(ctx context.Context, endpoint string) (*caws.IIDAttestationData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Agent sends both RSA-1024 and RSA-2048 signatures. This is for maintaining backwards compatibility, to support
// new SPIRE agents to attest to older SPIRE servers.

func getMetadataDoc(ctx context.Context, client *imds.Client) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getMetadataSig(ctx context.Context, client *imds.Client, signaturePath string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func readStringAndClose(r io.ReadCloser) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Configure implements the Config interface method of the same name
func (p *IIDAttestorPlugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *IIDAttestorPlugin) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *IIDAttestorPlugin) getConfig() (*IIDAttestorConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
