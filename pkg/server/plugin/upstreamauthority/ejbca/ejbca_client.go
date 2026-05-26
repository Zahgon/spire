package ejbca

import (
	"context"

	ejbcaclient "github.com/Keyfactor/ejbca-go-client-sdk/api/ejbca"
)

type ejbcaClient interface {
	EnrollPkcs10Certificate(ctx context.Context) ejbcaclient.ApiEnrollPkcs10CertificateRequest
}

func (p *Plugin) getAuthenticator(config *Config) (ejbcaclient.Authenticator, error) {
	_ = "STUB: not implemented"
	return *new(ejbcaclient.Authenticator), nil
}

// newEjbcaClient generates a new EJBCA client based on the provided configuration.
func (p *Plugin) newEjbcaClient(config *Config, authenticator ejbcaclient.Authenticator) (ejbcaClient, error) {
	_ = "STUB: not implemented"
	return *new(ejbcaClient), nil
}
