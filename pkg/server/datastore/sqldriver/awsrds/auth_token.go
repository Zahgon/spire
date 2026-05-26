package awsrds

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
)

const (
	iso8601BasicFormat = "20060102T150405Z"
	clockSkew          = time.Minute // Make sure that the authentication token is valid for one more minute.
)

type authTokenBuilder interface {
	buildAuthToken(ctx context.Context, endpoint string, region string, dbUser string, creds aws.CredentialsProvider, optFns ...func(options *auth.BuildAuthTokenOptions)) (string, error)
}

type authToken struct {
	cachedToken string
	expiresAt   time.Time
}

func (a *authToken) getAuthToken(ctx context.Context, config *Config, tokenBuilder authTokenBuilder) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// X-Amz-Expires is expressed as a duration in seconds.

// shouldRotate returns true if the cached token is either expired or is
// expiring soon. This means that this function will return true also if the
// token is still valid but should be rotated because it's expiring soon. The
// time window that establish when a cached token should be rotated even if it's
// still valid is adjusted by a clock skew, defined in the clockSkew constant.
func (a *authToken) shouldRotate() bool { _ = "STUB: not implemented"; return false }

type awsTokenBuilder struct{}

func (a *awsTokenBuilder) buildAuthToken(ctx context.Context, endpoint string, region string, dbUser string, creds aws.CredentialsProvider, optFns ...func(options *auth.BuildAuthTokenOptions)) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func newAWSClientConfig(ctx context.Context, c *Config) (aws.Config, error) {
	_ = "STUB: not implemented"
	return *new(aws.Config), nil
}
