package azureimds

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
)

// apiClient is an interface representing all API methods the resolver
// needs to do its job.
type apiClient interface {
	GetVirtualMachine(ctx context.Context, vmId string, subscriptionId *string) (*VirtualMachine, error)
	GetVMSSInstance(ctx context.Context, vmId, subscriptionID, ssName string) (*VirtualMachine, error)
}

// VirtualMachine is a subset of the fields returned by the Resource Graph API
type VirtualMachine struct {
	ID            string              `json:"id"`
	Name          string              `json:"name"`
	Location      string              `json:"location"`
	Tags          map[string]any      `json:"tags"`
	VMID          string              `json:"vmId"`
	ResourceGroup string              `json:"resourceGroup"`
	Interfaces    []*NetworkInterface `json:"interfaces"`
}
type NetworkInterface struct {
	Name          string        `json:"name"`
	SecurityGroup SecurityGroup `json:"securityGroup"`
	Subnets       []Subnet      `json:"subnets"`
}
type Subnet struct {
	VNet       string `json:"vnet"`
	SubnetName string `json:"name"`
}

type SecurityGroup struct {
	ResourceGroup string `json:"resourceGroup"`
	Name          string `json:"name"`
}

type VMSSInfo struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Location       string `json:"location"`
	ResourceGroup  string `json:"resourceGroup"`
	SubscriptionID string `json:"subscriptionId"`
}

// azureClient implements apiClient using Azure SDK client implementations
type azureClient struct {
	cred azcore.TokenCredential
	g    *armresourcegraph.Client
}

func newAzureClient(cred azcore.TokenCredential) (apiClient, error) {
	_ = "STUB: not implemented"
	return *new(apiClient), nil
}

// A direct scale set VM api client is needed to support VMSS with an orchestration mode of "Uniform".
func (c *azureClient) newScaleSetVMClient(subscriptionID string) (*armcompute.VirtualMachineScaleSetVMsClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *azureClient) GetVirtualMachine(ctx context.Context, vmId string, subscriptionId *string) (*VirtualMachine, error) {
	_ = "STUB: not implemented"
	// For additional fields, see:
	// https://learn.microsoft.com/en-us/azure/templates/microsoft.compute/virtualmachines?pivots=deployment-language-arm-template
	return nil, nil
}

func (c *azureClient) GetVMSSInstance(ctx context.Context, vmId, subscriptionID, ssName string) (*VirtualMachine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *azureClient) getVMSSInfo(ctx context.Context, subscriptionIDs []*string, name string) (*VMSSInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *azureClient) getNetworkInterfaces(ctx context.Context, vmId string, subscriptionId *string) ([]*NetworkInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buildVirtualMachineFromVMSSInstance creates a VirtualMachine struct from a VMSS instance
// with all network interfaces parsed and populated
func buildVirtualMachineFromVMSSInstance(instance *armcompute.VirtualMachineScaleSetVM, resourceGroup string) (*VirtualMachine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parseNetworkInterfaceConfig parses a network interface configuration from a VMSS instance
// and returns a NetworkInterface with parsed security group and subnet information
func parseNetworkInterfaceConfig(interfaceConfig *armcompute.VirtualMachineScaleSetNetworkConfiguration) (*NetworkInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseSecurityGroup(nsg *armcompute.SubResource) (SecurityGroup, error) {
	_ = "STUB: not implemented"
	return *new(SecurityGroup), nil
}

func extractArmResourceGraphItems[T any](resp armresourcegraph.ClientResourcesResponse) ([]*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractArmResourceGraphItem[T any](resp armresourcegraph.ClientResourcesResponse) (*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
