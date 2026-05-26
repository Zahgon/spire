package awspca

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/acmpca"
)

type pcaClientFake struct {
	t testing.TB

	describeCertificateOutput *acmpca.DescribeCertificateAuthorityOutput
	expectedDescribeInput     *acmpca.DescribeCertificateAuthorityInput
	describeCertificateErr    error

	issueCertificateOutput *acmpca.IssueCertificateOutput
	expectedIssueInput     *acmpca.IssueCertificateInput
	issueCertificateErr    error

	expectedGetCertificateInput *acmpca.GetCertificateInput
	getCertificateOutput        *acmpca.GetCertificateOutput
	getCertificateErr           error
}

func (f *pcaClientFake) DescribeCertificateAuthority(_ context.Context, input *acmpca.DescribeCertificateAuthorityInput, _ ...func(*acmpca.Options)) (*acmpca.DescribeCertificateAuthorityOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *pcaClientFake) IssueCertificate(_ context.Context, input *acmpca.IssueCertificateInput, _ ...func(*acmpca.Options)) (*acmpca.IssueCertificateOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *pcaClientFake) GetCertificate(_ context.Context, input *acmpca.GetCertificateInput, _ ...func(*acmpca.Options)) (*acmpca.GetCertificateOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
