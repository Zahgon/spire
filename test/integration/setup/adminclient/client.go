package main

import (
	"context"
	"encoding/base64"
	"encoding/pem"
	"log"

	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/pkg/common/pemutil"
	"github.com/spiffe/spire/test/integration/setup/itclient"
)

const (
	testBundle = `
-----BEGIN CERTIFICATE-----
MIICOTCCAZqgAwIBAgIBATAKBggqhkjOPQQDBDAeMQswCQYDVQQGEwJVUzEPMA0G
A1UECgwGU1BJRkZFMB4XDTE4MDIxMDAwMzQ0NVoXDTE4MDIxMDAxMzQ1NVowHjEL
MAkGA1UEBhMCVVMxDzANBgNVBAoTBlNQSUZGRTCBmzAQBgcqhkjOPQIBBgUrgQQA
IwOBhgAEAZ6nXrNctKHNjZT7ZkP7xwfpMfvc/DAHc39GdT3qi8mmowY0/XuFQmlJ
cXXwv8ZlOSoGvtuLAEx1lvHNZwv4BuuPALILcIW5tyC8pjcbfqs8PMQYwiC+oFKH
BTxXzolpLeHuFLAD9ccfwWhkT1z/t4pvLkP4FCkkBosG9PVg5JQVJuZJo4GFMIGC
MA4GA1UdDwEB/wQEAwIBhjAPBgNVHRMBAf8EBTADAQH/MB0GA1UdDgQWBBT4RuNt
x6E70yjV0wIvUyrGkMKczzAfBgNVHSMEGDAWgBRGyozl9Mjue0Y3w4c2Q+3u+wVk
CjAfBgNVHREEGDAWhhRzcGlmZmU6Ly9leGFtcGxlLm9yZzAKBggqhkjOPQQDBAOB
jAAwgYgCQgHOtx4sNCioAQnpEx3J/A9M6Lutth/ND/h8D+7luqEkd4tMrBQgnMj4
E0xLGUNtoFNRIrEUlgwksWvKZ3BksIIOMwJCAc8VPA/QYrlJDeQ58FKyQyrOIlPk
Q0qBJEOkL6FrAngY5218TCNUS30YS5HjI2lfyyjB+cSVFXX8Szu019dDBMhV
-----END CERTIFICATE-----
`
)

var (
	blk, _       = pem.Decode([]byte(testBundle))
	pkixBytes, _ = base64.StdEncoding.DecodeString("MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEYSlUVLqTD8DEnA4F1EWMTf5RXc5lnCxw+5WKJwngEL3rPc9i4Tgzz9riR3I/NiSlkgRO1WsxBusqpC284j9dXA==")
	key, _       = pemutil.ParseSigner([]byte(`-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgs/CcKxAEIyBBEQ9h
ES2kJbWTz79ut45qAb0UgqrGqmOhRANCAARssWdfmS3D4INrpLBdSBxzso5kPPSX
F21JuznwCuYKNV5LnzhUA3nt2+6e18ZIXUDxl+CpkvCYc10MO6SYg6AE
-----END PRIVATE KEY-----`))
	// Used between test
	entryID = ""
	agentID = &types.SPIFFEID{}
)

func main() {
	if msg := run(); msg != "" {
		log.Fatal(msg)
	}

	log.Println("Admin client finished successfully")
}

// run execute all test cases return true if all test cases finished successfully
func run() string { _ = "STUB: not implemented"; return "" }

// SVID Client tests

// Bundle Client tests

// Entry client tests

// Agent client tests

// Trustdomain client tests

func mintX509SVID(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

// Create CSR

// Call mint

// Validate error

// Validate certificate

func mintJWTSVID(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

// Parse token

// Validate token

func appendBundle(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func batchCreateFederatedBundle(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate result

func batchUpdateFederatedBundle(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func batchSetFederatedBundle(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate result

func countBundles(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func listFederatedBundles(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func getFederatedBundle(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func batchDeleteFederatedBundle(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func batchCreateEntry(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate result

// Setup entry ID it will be used for another tests

func countEntries(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func listEntries(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func getEntry(ctx context.Context, c *itclient.Client) error { _ = "STUB: not implemented"; return nil }

func batchUpdateEntry(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate result

func batchDeleteEntry(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate result

func createJoinToken(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

// Set agentID that will be used in other tests

// Create CSR

// Attest using generated token

func countAgents(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func listAgents(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate agent

func getAgent(ctx context.Context, c *itclient.Client) error { _ = "STUB: not implemented"; return nil }

func banAgent(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	// Ban agent returns empty as response
	return nil
}

// Validates it is banned

func deleteAgent(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	// Delete agent returns empty as response
	return nil
}

// Validates it is banned

func batchCreateFederationRelationship(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate result

func batchUpdateFederationRelationship(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func listFederationRelationships(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func getFederationRelationship(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func batchDeleteFederationRelationship(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func validatePermissionError(err error) error { _ = "STUB: not implemented"; return nil }

func containsX509Certificate(certs []*types.X509Certificate, b []byte) bool {
	_ = "STUB: not implemented"
	return false
}

func containsJWTKey(keys []*types.JWTKey, key *types.JWTKey) bool {
	_ = "STUB: not implemented"
	return false
}
