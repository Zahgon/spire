package awss3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type simpleStorageService interface {
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

func newAWSConfig(ctx context.Context, c *Config) (aws.Config, error) {
	_ = "STUB: not implemented"
	return *new(aws.Config), nil
}

func newS3Client(c aws.Config) (simpleStorageService, error) {
	_ = "STUB: not implemented"
	return *new(simpleStorageService), nil
}
