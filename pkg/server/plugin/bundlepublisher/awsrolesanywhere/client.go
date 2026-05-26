package awsrolesanywhere

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rolesanywhere"
)

type rolesAnywhere interface {
	UpdateTrustAnchor(ctx context.Context, params *rolesanywhere.UpdateTrustAnchorInput, optFns ...func(*rolesanywhere.Options)) (*rolesanywhere.UpdateTrustAnchorOutput, error)
}

func newAWSConfig(ctx context.Context, c *Config) (aws.Config, error) {
	_ = "STUB: not implemented"
	return *new(aws.Config), nil
}

func newRolesAnywhereClient(c aws.Config) (rolesAnywhere, error) {
	_ = "STUB: not implemented"
	return *new(rolesAnywhere), nil
}
