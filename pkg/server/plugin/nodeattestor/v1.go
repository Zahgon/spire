package nodeattestor

import (
	"context"

	nodeattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/nodeattestor/v1"
	"github.com/spiffe/spire/pkg/common/plugin"
)

const (
	// This header contains the value of the gRPC :authority pseudo-header as received from the client.
	// Warning: This value is set by the client and is not authenticated or validated by SPIRE.
	// It must not be used for security decisions (such as authentication, authorization, or trust domain selection) in attestor plugins without threat assessment.
	// Valid uses include diagnostics, logging, or configuration side-loading
	XForwardedHostKey = "X-Untrusted-Forwarded-Host"
)

type V1 struct {
	plugin.Facade
	nodeattestorv1.NodeAttestorPluginClient
}

func (v1 *V1) Attest(ctx context.Context, payload []byte, challengeFn func(ctx context.Context, challenge []byte) ([]byte, error)) (*AttestResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// forward original request host to downstream plugins

func (v1 *V1) streamError(err error) error { _ = "STUB: not implemented"; return nil }

func getOriginalHost(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// should be just one in a slice
// example value: spire-server-xyz.spiffe.io:8081
