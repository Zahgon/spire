package credtemplate

import (
	"context"
	"crypto"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/go-jose/go-jose/v4"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/common/tlspolicy"
	"github.com/spiffe/spire/pkg/server/plugin/credentialcomposer"
)

const (
	// DefaultX509CATTL is the TTL given to X509 CAs if not overridden by
	// the server config.
	DefaultX509CATTL = time.Hour * 24

	// DefaultX509SVIDTTL is the TTL given to X509 SVIDs if not overridden by
	// the server config.
	DefaultX509SVIDTTL = time.Hour

	// DefaultJWTSVIDTTL is the TTL given to JWT SVIDs if a different TTL is
	// not provided in the signing request.
	DefaultJWTSVIDTTL = time.Minute * 5

	// DefaultWITSVIDTTL is the TTL given to WIT-SVIDs if a different TTL is
	// not provided in the signing request.
	DefaultWITSVIDTTL = time.Hour

	// NotBeforeCushion is how much of a cushion to subtract from the current
	// time when determining the notBefore field of certificates to account
	// for clock skew.
	NotBeforeCushion = 10 * time.Second
)

// DefaultX509CASubject is the default subject set on workload X509SVIDs
// TODO: This is a historic, but poor, default. We should revisit (see issue #3841).
func DefaultX509CASubject() pkix.Name { _ = "STUB: not implemented"; return *new(pkix.Name) }

// DefaultX509SVIDSubject is the default subject set on workload X509SVIDs
// TODO: This is a historic, but poor, default. We should revisit (see issue #3841).
func DefaultX509SVIDSubject() pkix.Name { _ = "STUB: not implemented"; return *new(pkix.Name) }

type SelfSignedX509CAParams struct {
	PublicKey crypto.PublicKey
}

type UpstreamSignedX509CAParams struct {
	PublicKey crypto.PublicKey
}

type DownstreamX509CAParams struct {
	ParentChain []*x509.Certificate
	PublicKey   crypto.PublicKey
	TTL         time.Duration
}

type ServerX509SVIDParams struct {
	ParentChain []*x509.Certificate
	PublicKey   crypto.PublicKey
}

type AgentX509SVIDParams struct {
	ParentChain []*x509.Certificate
	PublicKey   crypto.PublicKey
	SPIFFEID    spiffeid.ID
}

type WorkloadX509SVIDParams struct {
	ParentChain []*x509.Certificate
	PublicKey   crypto.PublicKey
	SPIFFEID    spiffeid.ID
	DNSNames    []string
	TTL         time.Duration
	Subject     pkix.Name
}

type WorkloadJWTSVIDParams struct {
	SPIFFEID      spiffeid.ID
	Audience      []string
	TTL           time.Duration
	ExpirationCap time.Time
}

type WorkloadWITSVIDParams struct {
	SPIFFEID      spiffeid.ID
	PublicKey     jose.JSONWebKey
	TTL           time.Duration
	ExpirationCap time.Time
}

type Config struct {
	TrustDomain         spiffeid.TrustDomain
	Clock               clock.Clock
	X509CASubject       pkix.Name
	X509CATTL           time.Duration
	X509SVIDSubject     pkix.Name
	X509SVIDTTL         time.Duration
	JWTSVIDTTL          time.Duration
	JWTIssuer           string
	WITSVIDTTL          time.Duration
	WITIssuer           string
	AgentSVIDTTL        time.Duration
	CredentialComposers []credentialcomposer.CredentialComposer
	NewSerialNumber     func() (*big.Int, error)
	TLSPolicy           tlspolicy.Policy
}

type Builder struct {
	config Config

	x509CAID spiffeid.ID
	serverID spiffeid.ID
}

func NewBuilder(config Config) (*Builder, error) { _ = "STUB: not implemented"; return nil, nil }

// config.X509SVIDTTL should be initialized by the code above and
// therefore safe to use to initialize the AgentSVIDTTL.

// This check is purely defensive; idutil.ServerID should not fail since the trust domain is valid.

func (b *Builder) Config() Config { _ = "STUB: not implemented"; return *new(Config) }

func (b *Builder) BuildSelfSignedX509CATemplate(ctx context.Context, params SelfSignedX509CAParams) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Builder) BuildUpstreamSignedX509CACSR(ctx context.Context, params UpstreamSignedX509CAParams) (*x509.CertificateRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create the CertificateRequest from the Certificate template. The
// Policies field is ignored since that can be applied by the
// upstream signer and isn't a part of the native CertificateRequest type.
// TODO: maybe revisit this if needed and embed the policy identifiers in
// the extra extensions.

func (b *Builder) BuildDownstreamX509CATemplate(ctx context.Context, params DownstreamX509CAParams) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Builder) BuildServerX509SVIDTemplate(ctx context.Context, params ServerX509SVIDParams) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Builder) BuildAgentX509SVIDTemplate(ctx context.Context, params AgentX509SVIDParams) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Builder) BuildWorkloadX509SVIDTemplate(ctx context.Context, params WorkloadX509SVIDParams) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The first DNS name is also added as the CN by default. This happens
// even if the subject is provided explicitly in the params for backwards
// compatibility. Ideally we wouldn't do override the subject in this
// case. It is still overridable via the credential composers, however.

func (b *Builder) BuildWorkloadJWTSVIDClaims(ctx context.Context, params WorkloadJWTSVIDParams) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AWS will otherwise reject validating timestamps serialized in scientific notation.
// Protobuf serializes large integers as float since Claims are represented as google.protobuf.Struct.

func (b *Builder) BuildWorkloadWITSVIDClaims(ctx context.Context, params WorkloadWITSVIDParams) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Builder) buildX509CATemplate(publicKey crypto.PublicKey, parentChain []*x509.Certificate, ttl time.Duration) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Builder) buildX509SVIDTemplate(spiffeID spiffeid.ID, publicKey crypto.PublicKey, parentChain []*x509.Certificate, subject pkix.Name, ttl time.Duration) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Builder) buildBaseTemplate(spiffeID spiffeid.ID, publicKey crypto.PublicKey, parentChain []*x509.Certificate) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Explicitly set the AKI on the signed certificate, otherwise it won't be
// added if the subject and issuer match (however unlikely).

func (b *Builder) computeX509CALifetime(parentChain []*x509.Certificate, ttl time.Duration) (notBefore, notAfter time.Time) {
	_ = "STUB: not implemented"
	return *new(time.Time), *new(time.Time)
}

func (b *Builder) computeX509SVIDLifetime(parentChain []*x509.Certificate, ttl time.Duration) (notBefore, notAfter time.Time) {
	_ = "STUB: not implemented"
	return *new(time.Time), *new(time.Time)
}

func x509CAAttributesFromTemplate(tmpl *x509.Certificate) credentialcomposer.X509CAAttributes {
	_ = "STUB: not implemented"
	return *new(credentialcomposer.X509CAAttributes)
}

func x509SVIDAttributesFromTemplate(tmpl *x509.Certificate) credentialcomposer.X509SVIDAttributes {
	_ = "STUB: not implemented"
	return *new(credentialcomposer.X509SVIDAttributes)
}

func applyX509CAAttributes(tmpl *x509.Certificate, attribs credentialcomposer.X509CAAttributes) {
	_ = "STUB: not implemented"
	return
}

func applyX509SVIDAttributes(tmpl *x509.Certificate, attribs credentialcomposer.X509SVIDAttributes) {
	_ = "STUB: not implemented"
	return
}

func computeCappedLifetime(clk clock.Clock, ttl time.Duration, expirationCap time.Time) (notBefore, notAfter time.Time) {
	_ = "STUB: not implemented"
	return *new(time.Time), *new(time.Time)
}

func parentChainExpiration(parentChain []*x509.Certificate) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func dropEmptyValues(ss []string) []string { _ = "STUB: not implemented"; return nil }
