package azuremsi

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources/v3"
)

// apiClient is an interface representing all API methods the resolver
// needs to do its job.
type apiClient interface {
	SubscriptionID() string
	GetVirtualMachineResourceID(ctx context.Context, principalID string) (string, error)
	GetVirtualMachine(ctx context.Context, resourceGroup string, name string) (*armcompute.VirtualMachine, error)
	GetNetworkInterface(ctx context.Context, resourceGroup string, name string) (*armnetwork.Interface, error)
}

// azureClient implements apiClient using Azure SDK client implementations
type azureClient struct {
	subscriptionID string
	r              *armresources.Client
	v              *armcompute.VirtualMachinesClient
	n              *armnetwork.InterfacesClient
}

func newAzureClient(subscriptionID string, cred azcore.TokenCredential) (apiClient, error) {
	_ = "STUB: not implemented"
	return *new(apiClient), nil
}

func (c *azureClient) SubscriptionID() string { _ = "STUB: not implemented"; return "" }

func (c *azureClient) GetVirtualMachineResourceID(ctx context.Context, principalID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *azureClient) GetVirtualMachine(ctx context.Context, resourceGroup string, name string) (*armcompute.VirtualMachine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *azureClient) GetNetworkInterface(ctx context.Context, resourceGroup string, name string) (*armnetwork.Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
