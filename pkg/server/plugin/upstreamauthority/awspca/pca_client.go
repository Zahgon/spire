package awspca

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
)

// PCAClient provides an interface which can be mocked to test
// the functionality of the plugin.
type PCAClient interface {
	DescribeCertificateAuthority(context.Context, *acmpca.DescribeCertificateAuthorityInput, ...func(*acmpca.Options)) (*acmpca.DescribeCertificateAuthorityOutput, error)
	IssueCertificate(context.Context, *acmpca.IssueCertificateInput, ...func(*acmpca.Options)) (*acmpca.IssueCertificateOutput, error)
	GetCertificate(context.Context, *acmpca.GetCertificateInput, ...func(*acmpca.Options)) (*acmpca.GetCertificateOutput, error)
}

func newPCAClient(ctx context.Context, cfg *Configuration) (PCAClient, error) {
	_ = "STUB: not implemented"
	return *new(PCAClient), nil
}

func newAWSAssumeRoleConfig(ctx context.Context, region string, awsConf aws.Config, assumeRoleArn string) (aws.Config, error) {
	_ = "STUB: not implemented"
	return *new(aws.Config), nil
}
