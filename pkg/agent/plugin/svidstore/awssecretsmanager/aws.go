package awssecretsmanager

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/hashicorp/go-hclog"
	svidstorev1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/agent/svidstore/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

const (
	pluginName = "aws_secretsmanager"
)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *SecretsManagerPlugin) catalog.BuiltIn {
	_ = "STUB: not implemented"
	return *new(catalog.BuiltIn)
}

func New() *SecretsManagerPlugin { _ = "STUB: not implemented"; return nil }

func newPlugin(newClient func(ctx context.Context, secretAccessKey, accessKeyID, region string) (SecretsManagerClient, error)) *SecretsManagerPlugin {
	_ = "STUB: not implemented"
	return nil
}

type Configuration struct {
	AccessKeyID     string `hcl:"access_key_id" json:"access_key_id"`
	SecretAccessKey string `hcl:"secret_access_key" json:"secret_access_key"`
	Region          string `hcl:"region" json:"region"`
}

func (p *SecretsManagerPlugin) buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *Configuration {
	_ = "STUB: not implemented"
	return nil
}

type SecretsManagerPlugin struct {
	svidstorev1.UnsafeSVIDStoreServer
	configv1.UnsafeConfigServer

	log      hclog.Logger
	smClient SecretsManagerClient
	mtx      sync.RWMutex

	hooks struct {
		newClient func(ctx context.Context, secretAccessKey, accessKeyID, region string) (SecretsManagerClient, error)
		getenv    func(string) string
	}
}

func (p *SecretsManagerPlugin) SetLogger(log hclog.Logger) {
	_ = "STUB: not implemented"

	// Configure configures the SecretsManagerPlugin.
	return
}

func (p *SecretsManagerPlugin) Configure(ctx context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *SecretsManagerPlugin) Validate(ctx context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PutX509SVID puts the specified X509-SVID in the configured AWS Secrets Manager
func (p *SecretsManagerPlugin) PutX509SVID(ctx context.Context, req *svidstorev1.PutX509SVIDRequest) (*svidstorev1.PutX509SVIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Encode the secret from PutX509SVIDRequest

// Call DescribeSecret to retrieve the details of the secret
// and be able to determine if the secret exists

// Secret not found, creating one with provided `name`

// Purely defensive. This should never happen.

// Validate that the secret has the 'spire-svid' tag. This tag is used to distinguish the secrets
// that have SVID information handled by SPIRE

// If the secret has been scheduled for deletion, restore it

// DeleteX509SVID schedules a deletion to a Secret using AWS secret manager
func (p *SecretsManagerPlugin) DeleteX509SVID(ctx context.Context, req *svidstorev1.DeleteX509SVIDRequest) (*svidstorev1.DeleteX509SVIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Call DescribeSecret to retrieve the details of the secret
// and be able to determine if the secret exists

// Validate that the secret has the 'spire-svid' tag. This tag is used to distinguish the secrets
// that have SVID information handled by SPIRE

type secretOptions struct {
	name     string
	arn      string
	kmsKeyID string
}

// getSecretID gets ARN if it is configured. If not configured, use secret name
func (o *secretOptions) getSecretID() string { _ = "STUB: not implemented"; return "" }

func optionsFromSecretData(metadata []string) (*secretOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createSecret(ctx context.Context, sm SecretsManagerClient, secretBinary []byte, opt *secretOptions) (*secretsmanager.CreateSecretOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validateTag expects that "spire-svid" tag is provided
func validateTag(tags []types.Tag) error { _ = "STUB: not implemented"; return nil }
