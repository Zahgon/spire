package credvalidator

import (
	"crypto/x509"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
)

type Config struct {
	Clock       clock.Clock
	TrustDomain spiffeid.TrustDomain
}

type Validator struct {
	clock    clock.Clock
	x509CAID spiffeid.ID
	serverID spiffeid.ID
}

func New(config Config) (*Validator, error) { _ = "STUB: not implemented"; return nil, nil }

// This check is purely defensive; idutil.ServerID should not fail since the trust domain is valid.

func (v *Validator) ValidateX509CA(ca *x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) ValidateServerX509SVID(svid *x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) ValidateX509SVID(svid *x509.Certificate, id spiffeid.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) ValidateWorkloadJWTSVID(rawToken string, id spiffeid.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) ValidateWorkloadWITSVID(rawToken string, id spiffeid.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func checkURISAN(cert *x509.Certificate, isCA bool, id spiffeid.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// A signing certificate should itself be an SVID, but it's not
// mandatory.

// There is at least one URI.
// These validations apply for both CA and non CA certificates.

func checkX509CertificateExpiration(cert *x509.Certificate, now time.Time) error {
	_ = "STUB: not implemented"
	return nil
}
