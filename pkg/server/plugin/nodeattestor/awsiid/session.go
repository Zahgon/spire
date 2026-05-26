package awsiid

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// SessionConfig is a common config for AWS session config.
type SessionConfig struct {
	AccessKeyID     string `hcl:"access_key_id"`
	SecretAccessKey string `hcl:"secret_access_key"`
	AssumeRole      string `hcl:"assume_role"`
	Partition       string `hcl:"partition"`
}

func (cfg *SessionConfig) Validate(defaultAccessKeyID, defaultSecretAccessKey string) error {
	_ = "STUB: not implemented"
	return nil
}

// newAWSSession create an AWS config from the credentials and given region
func newAWSConfig(ctx context.Context, accessKeyID, secretAccessKey, region, assumeRoleArn string) (aws.Config, error) {
	_ = "STUB: not implemented"
	return *new(aws.Config), nil
}

func newAWSAssumeRoleConfig(ctx context.Context, region string, stsConf aws.Config, assumeRoleArn string) (aws.Config, error) {
	_ = "STUB: not implemented"
	return *new(aws.Config), nil
}
