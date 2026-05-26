package azureimds

import (
	"context"
	"regexp"

	"github.com/go-jose/go-jose/v4"
)

var (
	reNetworkSecurityGroupID = regexp.MustCompile(`^/subscriptions/[^/]+/resourceGroups/([^/]+)/providers/Microsoft.Network/networkSecurityGroups/([^/]+)$`)
	reNetworkInterfaceID     = regexp.MustCompile(`^/subscriptions/[^/]+/resourceGroups/([^/]+)/providers/Microsoft.Network/networkInterfaces/([^/]+)$`)
	reVirtualNetworkSubnetID = regexp.MustCompile(`^/subscriptions/[^/]+/resourceGroups/([^/]+)/providers/Microsoft.Network/virtualNetworks/([^/]+)/subnets/([^/]+)$`)
	reTenantId               = regexp.MustCompile(`^https://sts.windows.net/([^/]+)/$`)
	// VMSS name validation: alphanumeric, underscores, periods, and hyphens
	reVMSSNameAllowedChars = regexp.MustCompile(`^[a-zA-Z0-9._-]+$`)
	// Used to make sure token is valid for credential assertion
	allowedJWTSignatureAlgorithms = []jose.SignatureAlgorithm{
		jose.RS256,
		jose.RS384,
		jose.RS512,
		jose.ES256,
		jose.ES384,
		jose.ES512,
		jose.PS256,
		jose.PS384,
		jose.PS512,
	}
)

func selectorValue(parts ...string) string { _ = "STUB: not implemented"; return "" }

func getAzureAssertionFunc(tokenPath string, reader func(name string) ([]byte, error)) func(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return nil
}

func lookupTenantID(domain string) (string, error) {
	_ = "STUB: not implemented"
	// make an http request to https://login.microsoftonline.com/<domain>/.well-known/openid-configuration
	return "", nil
}

func parseNetworkSecurityGroupID(id string) (resourceGroup, name string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func parseNetworkInterfaceID(id string) (resourceGroup, name string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func parseVirtualNetworkSubnetID(id string) (resourceGroup, networkName, subnetName string, err error) {
	_ = "STUB: not implemented"
	return "", "", "", nil
}

func parseIssuer(issuer string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func generateRandomAlphanumeric(length int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// validateVMSSName validates an Azure VM Scale Set name according to Azure naming rules:
// - Length: Must be between 1 and 64 characters long
// - Allowed Characters: Can contain only alphanumeric characters, underscores, periods, and hyphens
// - Start: Must start with an alphanumeric character
// - End: Must end with an alphanumeric character or an underscore
//
// Note: Uniqueness within the resource group is not validated by this function and must be
// checked separately. Case sensitivity is noted for information but not enforced here.
func validateVMSSName(name string) error { _ = "STUB: not implemented"; return nil }

// Check length

// Check allowed characters (alphanumeric, underscores, periods, hyphens)

// Check start: must start with alphanumeric

// Check end: must end with alphanumeric or underscore

// validateUUID validates that a string is a valid UUID using the uuid library.
func validateUUID(s string) error { _ = "STUB: not implemented"; return nil }
