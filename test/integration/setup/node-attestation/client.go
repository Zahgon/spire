package main

import (
	"context"
	"flag"
	"log"

	agent "github.com/spiffe/spire-api-sdk/proto/spire/api/server/agent/v1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/pkg/common/pemutil"
)

var (
	key, _ = pemutil.ParseSigner([]byte(`-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgs/CcKxAEIyBBEQ9h
ES2kJbWTz79ut45qAb0UgqrGqmOhRANCAARssWdfmS3D4INrpLBdSBxzso5kPPSX
F21JuznwCuYKNV5LnzhUA3nt2+6e18ZIXUDxl+CpkvCYc10MO6SYg6AE
-----END PRIVATE KEY-----`))

	testStep    = flag.String("testStep", "", "jointoken, attest, ban, renew")
	tokenName   = flag.String("tokenName", "tokenName", "token for attestation")
	certificate = flag.String("certificate", "", "certificate for api connection")
	popCert     = flag.String("popCertificate", "/opt/spire/conf/agent/test.crt.pem", "certificate for x509pop attestation")
	popKey      = flag.String("popKey", "/opt/spire/conf/agent/test.key.pem", "key for x509pop attestation")
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Node attestation client failed: %v\n", err)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }

func doJoinTokenStep(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Create a join token using the local socket connection (simulating the CLI running on the spire-server)

// Print the join token so it can be easily used in the subsequent test

func doJoinTokenAttestStep(ctx context.Context, tokenName string) error {
	_ = "STUB: not implemented"
	// Now do agent attestation using the join token and save the resulting SVID to a file. This will give us an SVID
	return nil
}

// Print the SVID so it can easily be used in the next step

func doRenewStep(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Now renew the agent cert

// Print the certificate so it can easily be used in the next step

func doBanStep(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Now ban the agent using the local connection

// doX509popStep tests attestation using x509pop
// Steps:
// - Attest agent
// - Renew agent
// - Delete agent
// - Reattest deleted agent
// - Ban agent
// - Reattest banned agent (must fail because it is banned)
// - Delete agent
// - Reattest deleted agent (must succeed after removing)
func doX509popStep(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Create an admin client to ban/delete agent

// Attest agent

// Reattest agent to "renew"

// Delete agent

// Reattest deleted agent

// Ban agent

// Reattest banned agent, it MUST fail

// Delete banned agent

// Reattest deleted agent, now MUST be successful

// x509popAttest attests agent using x509pop
func x509popAttest(ctx context.Context) (*types.X509SVID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create insecure connection

// deleteAgent delete agent using "admin" connection
func deleteAgent(ctx context.Context, client agent.AgentClient, id *types.SPIFFEID) error {
	_ = "STUB: not implemented"
	return nil
}

// banAgent ban agent using "admin" connection
func banAgent(ctx context.Context, client agent.AgentClient, id *types.SPIFFEID) error {
	_ = "STUB: not implemented"
	return nil
}
