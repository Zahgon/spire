package azurekeyvault

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azkeys"
)

type cloudKeyManagementService interface {
	CreateKey(ctx context.Context, name string, parameters azkeys.CreateKeyParameters, options *azkeys.CreateKeyOptions) (azkeys.CreateKeyResponse, error)
	DeleteKey(ctx context.Context, name string, options *azkeys.DeleteKeyOptions) (azkeys.DeleteKeyResponse, error)
	UpdateKey(ctx context.Context, name string, version string, parameters azkeys.UpdateKeyParameters, options *azkeys.UpdateKeyOptions) (azkeys.UpdateKeyResponse, error)
	GetKey(ctx context.Context, name string, version string, options *azkeys.GetKeyOptions) (azkeys.GetKeyResponse, error)
	NewListKeyPropertiesPager(options *azkeys.ListKeyPropertiesOptions) *runtime.Pager[azkeys.ListKeyPropertiesResponse]
	Sign(ctx context.Context, name string, version string, parameters azkeys.SignParameters, options *azkeys.SignOptions) (azkeys.SignResponse, error)
}

type keyVaultClient struct {
	client *azkeys.Client
}

func (c *keyVaultClient) CreateKey(ctx context.Context, name string, parameters azkeys.CreateKeyParameters, options *azkeys.CreateKeyOptions) (azkeys.CreateKeyResponse, error) {
	_ = "STUB: not implemented"
	return *new(azkeys.CreateKeyResponse), nil
}

func (c *keyVaultClient) DeleteKey(ctx context.Context, name string, options *azkeys.DeleteKeyOptions) (azkeys.DeleteKeyResponse, error) {
	_ = "STUB: not implemented"
	return *new(azkeys.DeleteKeyResponse), nil
}

func (c *keyVaultClient) UpdateKey(ctx context.Context, name string, version string, parameters azkeys.UpdateKeyParameters, options *azkeys.UpdateKeyOptions) (azkeys.UpdateKeyResponse, error) {
	_ = "STUB: not implemented"
	return *new(azkeys.UpdateKeyResponse), nil
}

func (c *keyVaultClient) GetKey(ctx context.Context, name string, version string, options *azkeys.GetKeyOptions) (azkeys.GetKeyResponse, error) {
	_ = "STUB: not implemented"
	return *new(azkeys.GetKeyResponse), nil
}

func (c *keyVaultClient) NewListKeyPropertiesPager(options *azkeys.ListKeyPropertiesOptions) *runtime.Pager[azkeys.ListKeyPropertiesResponse] {
	_ = "STUB: not implemented"
	return nil
}

func (c *keyVaultClient) Sign(ctx context.Context, name string, version string, parameters azkeys.SignParameters, options *azkeys.SignOptions) (azkeys.SignResponse, error) {
	_ = "STUB: not implemented"
	return *new(azkeys.SignResponse), nil
}

func newKeyVaultClient(creds azcore.TokenCredential, keyVaultURI string) (cloudKeyManagementService, error) {
	_ = "STUB: not implemented"
	return *new(cloudKeyManagementService), nil
}
