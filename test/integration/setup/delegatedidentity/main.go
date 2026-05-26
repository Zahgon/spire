package main

import (
	"flag"
	"log"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	agent_delegatedidentityv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/agent/delegatedidentity/v1"
)

var (
	socketPathFlag = flag.String("adminSocketPath", "unix:///opt/admin.sock", "admin agent socket path")
	expectedID     = flag.String("expectedID", "", "expected SPIFFE ID for workload")
	expectedTD     string
)

func main() {
	flag.Parse()

	if *expectedID != "" {
		expectedTD = spiffeid.RequireFromString(*expectedID).TrustDomain().IDString()
	}

	if err := run(); err != nil {
		log.Fatalf("Test for Delegated API failed: %v", err)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }

func validateCode(err error) error { _ = "STUB: not implemented"; return nil }

func validateFetchJWTSVIDsResponse(resp *agent_delegatedidentityv1.FetchJWTSVIDsResponse) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSubscribeToJWTBundlesResponse(resp *agent_delegatedidentityv1.SubscribeToJWTBundlesResponse) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSubscribeToX509BundlesResponse(resp *agent_delegatedidentityv1.SubscribeToX509BundlesResponse) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSubscribeToX509SVIDsResponse(resp *agent_delegatedidentityv1.SubscribeToX509SVIDsResponse) error {
	_ = "STUB: not implemented"
	return nil
}
