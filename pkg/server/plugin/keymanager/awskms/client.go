package awskms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type kmsClient interface {
	CreateKey(context.Context, *kms.CreateKeyInput, ...func(*kms.Options)) (*kms.CreateKeyOutput, error)
	DescribeKey(context.Context, *kms.DescribeKeyInput, ...func(*kms.Options)) (*kms.DescribeKeyOutput, error)
	CreateAlias(context.Context, *kms.CreateAliasInput, ...func(*kms.Options)) (*kms.CreateAliasOutput, error)
	UpdateAlias(context.Context, *kms.UpdateAliasInput, ...func(*kms.Options)) (*kms.UpdateAliasOutput, error)
	GetPublicKey(context.Context, *kms.GetPublicKeyInput, ...func(*kms.Options)) (*kms.GetPublicKeyOutput, error)
	ListAliases(context.Context, *kms.ListAliasesInput, ...func(*kms.Options)) (*kms.ListAliasesOutput, error)
	ScheduleKeyDeletion(context.Context, *kms.ScheduleKeyDeletionInput, ...func(*kms.Options)) (*kms.ScheduleKeyDeletionOutput, error)
	Sign(context.Context, *kms.SignInput, ...func(*kms.Options)) (*kms.SignOutput, error)
	ListKeys(context.Context, *kms.ListKeysInput, ...func(*kms.Options)) (*kms.ListKeysOutput, error)
	DeleteAlias(context.Context, *kms.DeleteAliasInput, ...func(*kms.Options)) (*kms.DeleteAliasOutput, error)
}

type stsClient interface {
	GetCallerIdentity(ctx context.Context, params *sts.GetCallerIdentityInput, optFns ...func(*sts.Options)) (*sts.GetCallerIdentityOutput, error)
}

func newKMSClient(c aws.Config) (kmsClient, error) {
	_ = "STUB: not implemented"
	return *new(kmsClient), nil
}

func newSTSClient(c aws.Config) (stsClient, error) {
	_ = "STUB: not implemented"
	return *new(stsClient), nil
}

func newAWSConfig(ctx context.Context, c *Config) (aws.Config, error) {
	_ = "STUB: not implemented"
	return *new(aws.Config), nil
}
