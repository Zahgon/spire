package authpolicy

import (
	"context"
	_ "embed"
)

var (
	//go:embed policy_data.json
	defaultPolicyData []byte
	//go:embed policy.rego
	defaultPolicyRego string
)

// DefaultAuthPolicy returns the default policy engine
func DefaultAuthPolicy(ctx context.Context) (*Engine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
