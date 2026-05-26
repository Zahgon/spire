package main

import (
	"context"
	"log"

	"github.com/spiffe/spire/pkg/common/pemutil"
	"github.com/spiffe/spire/test/integration/setup/itclient"
)

var (
	key, _ = pemutil.ParseSigner([]byte(`
-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgs/CcKxAEIyBBEQ9h
ES2kJbWTz79ut45qAb0UgqrGqmOhRANCAARssWdfmS3D4INrpLBdSBxzso5kPPSX
F21JuznwCuYKNV5LnzhUA3nt2+6e18ZIXUDxl+CpkvCYc10MO6SYg6AE
-----END PRIVATE KEY-----
`))
)

func main() {
	// Run all tests cases and if error msg is returned make client fails
	if msg := run(); msg != "" {
		log.Fatal(msg)
	}
	log.Println("Downstream client finished successfully")
}

// run executes all tests cases and return error msg when failing
func run() string { _ = "STUB: not implemented"; return "" }

// Validate call to New Downstream X509 CA

func validateNewDownstreamX509CA(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	// Create csr
	return nil
}

// Create new svid client and new downstream CA

func validatePublishJWTAUthorirty(ctx context.Context, c *itclient.Client) error {
	_ = "STUB: not implemented"
	// Marshal key
	return nil
}

// Authority appended

func validatePermissionError(err error) error { _ = "STUB: not implemented"; return nil }
