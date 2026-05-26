package ca

import (
	"context"
	"crypto"
	"crypto/x509"
	"crypto/x509/pkix"
	"sync"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/go-jose/go-jose/v4"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/common/health"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/pkg/server/credtemplate"
	"github.com/spiffe/spire/pkg/server/credvalidator"
)

const (
	backdate = 10 * time.Second
)

// ServerCA is an interface for Server CAs
type ServerCA interface {
	SignDownstreamX509CA(ctx context.Context, params DownstreamX509CAParams) ([]*x509.Certificate, error)
	SignServerX509SVID(ctx context.Context, params ServerX509SVIDParams) ([]*x509.Certificate, error)
	SignAgentX509SVID(ctx context.Context, params AgentX509SVIDParams) ([]*x509.Certificate, error)
	SignWorkloadX509SVID(ctx context.Context, params WorkloadX509SVIDParams) ([]*x509.Certificate, error)
	SignWorkloadJWTSVID(ctx context.Context, params WorkloadJWTSVIDParams) (string, error)
	SignWorkloadWITSVID(ctx context.Context, params WorkloadWITSVIDParams) (string, error)
	TaintedAuthorities() <-chan []*x509.Certificate
	IsJWTSVIDsDisabled() bool
	IsWITSVIDsDisabled() bool
}

// DownstreamX509CAParams are parameters relevant to downstream X.509 CA creation
type DownstreamX509CAParams struct {
	// Public Key
	PublicKey crypto.PublicKey

	// TTL is the desired time-to-live of the SVID. Regardless of the TTL, the
	// lifetime of the certificate will be capped to that of the signing cert.
	TTL time.Duration
}

// ServerX509SVIDParams are parameters relevant to server X509-SVID creation
type ServerX509SVIDParams struct {
	// Public Key
	PublicKey crypto.PublicKey
}

// AgentX509SVIDParams are parameters relevant to agent X509-SVID creation
type AgentX509SVIDParams struct {
	// Public Key
	PublicKey crypto.PublicKey

	// SPIFFE ID of the agent
	SPIFFEID spiffeid.ID
}

// WorkloadX509SVIDParams are parameters relevant to workload X509-SVID creation
type WorkloadX509SVIDParams struct {
	// Public Key
	PublicKey crypto.PublicKey

	// SPIFFE ID of the SVID
	SPIFFEID spiffeid.ID

	// DNSNames is used to add DNS SAN's to the X509 SVID. The first entry
	// is also added as the CN.
	DNSNames []string

	// TTL is the desired time-to-live of the SVID. Regardless of the TTL, the
	// lifetime of the certificate will be capped to that of the signing cert.
	TTL time.Duration

	// Subject of the SVID. Default subject is used if it is empty.
	Subject pkix.Name
}

// WorkloadJWTSVIDParams are parameters relevant to workload JWT-SVID creation
type WorkloadJWTSVIDParams struct {
	// SPIFFE ID of the SVID
	SPIFFEID spiffeid.ID

	// TTL is the desired time-to-live of the SVID. Regardless of the TTL, the
	// lifetime of the token will be capped to that of the signing key.
	TTL time.Duration

	// Audience is used for audience claims
	Audience []string
}

// WorkloadWITSVIDParams are parameters relevant to workload WIT-SVID creation
type WorkloadWITSVIDParams struct {
	// SPIFFE ID of the SVID
	SPIFFEID spiffeid.ID

	// TTL is the desired time-to-live of the SVID. Regardless of the TTL, the
	// lifetime of the token will be capped to that of the signing key.
	TTL time.Duration

	// PublicKey is used for the cnf claim
	PublicKey jose.JSONWebKey
}

type X509CA struct {
	// Signer is used to sign child certificates.
	Signer crypto.Signer

	// Certificate is the CA certificate.
	Certificate *x509.Certificate

	// UpstreamChain contains the CA certificate and intermediates necessary to
	// chain back to the upstream trust bundle. It is only set if the CA is
	// signed by an UpstreamCA.
	UpstreamChain []*x509.Certificate
}

type JWTKey struct {
	// The signer used to sign keys
	Signer crypto.Signer

	// Kid is the JWT key ID (i.e. "kid" claim)
	Kid string

	// NotAfter is the expiration time of the JWT key.
	NotAfter time.Time
}

type WITKey struct {
	// The signer used to sign keys
	Signer crypto.Signer

	// Kid is the WIT key ID (i.e. "kid" claim)
	Kid string

	// NotAfter is the expiration time of the WIT key.
	NotAfter time.Time
}

type Config struct {
	Log             logrus.FieldLogger
	Clock           clock.Clock
	Metrics         telemetry.Metrics
	TrustDomain     spiffeid.TrustDomain
	CredBuilder     *credtemplate.Builder
	CredValidator   *credvalidator.Validator
	HealthChecker   health.Checker
	DisableJWTSVIDs bool
	DisableWITSVIDs bool
}

type CA struct {
	c Config

	mu                   sync.RWMutex
	x509CA               *X509CA
	x509CAChain          []*x509.Certificate
	jwtKey               *JWTKey
	witKey               *WITKey
	taintedAuthoritiesCh chan []*x509.Certificate
}

func NewCA(config Config) *CA { _ = "STUB: not implemented"; return nil }

// Notify caller about any tainted authority

func (ca *CA) X509CA() *X509CA { _ = "STUB: not implemented"; return nil }

func (ca *CA) SetX509CA(x509CA *X509CA) { _ = "STUB: not implemented"; return }

func (ca *CA) JWTKey() *JWTKey { _ = "STUB: not implemented"; return nil }

func (ca *CA) SetJWTKey(jwtKey *JWTKey) { _ = "STUB: not implemented"; return }

func (ca *CA) WITKey() *WITKey { _ = "STUB: not implemented"; return nil }

func (ca *CA) SetWITKey(witKey *WITKey) { _ = "STUB: not implemented"; return }

func (ca *CA) NotifyTaintedX509Authorities(taintedAuthorities []*x509.Certificate) {
	_ = "STUB: not implemented"
	return
}

func (ca *CA) TaintedAuthorities() <-chan []*x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func (ca *CA) SignDownstreamX509CA(ctx context.Context, params DownstreamX509CAParams) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ca *CA) SignServerX509SVID(ctx context.Context, params ServerX509SVIDParams) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ca *CA) SignAgentX509SVID(ctx context.Context, params AgentX509SVIDParams) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ca *CA) SignWorkloadX509SVID(ctx context.Context, params WorkloadX509SVIDParams) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ca *CA) SignWorkloadJWTSVID(ctx context.Context, params WorkloadJWTSVIDParams) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (ca *CA) SignWorkloadWITSVID(ctx context.Context, params WorkloadWITSVIDParams) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// We already validate in other places that a valid algorithm is given

func (ca *CA) getX509CA() (*X509CA, []*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (ca *CA) signX509SVID(x509CA *X509CA, template *x509.Certificate) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ca *CA) signJWTSVID(jwtKey *JWTKey, claims map[string]any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (ca *CA) IsJWTSVIDsDisabled() bool { _ = "STUB: not implemented"; return false }

func (ca *CA) IsWITSVIDsDisabled() bool { _ = "STUB: not implemented"; return false }

func (ca *CA) signWITSVID(witKey *WITKey, claims map[string]any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func makeCertChain(x509CA *X509CA, leaf *x509.Certificate) []*x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}
