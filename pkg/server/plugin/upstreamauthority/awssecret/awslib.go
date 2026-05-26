package awssecret

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type secretsManagerClient interface {
	GetSecretValue(context.Context, *secretsmanager.GetSecretValueInput, ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error)
}

func readARN(ctx context.Context, sm secretsManagerClient, arn string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// resp is now filled

func newSecretsManagerClient(ctx context.Context, cfg *Configuration, region string) (secretsManagerClient, error) {
	_ = "STUB: not implemented"
	return *new(secretsManagerClient), nil
}

func newAWSAssumeRoleConfig(ctx context.Context, region string, awsConf aws.Config, assumeRoleArn string) (aws.Config, error) {
	_ = "STUB: not implemented"
	return *new(aws.Config), nil
}
