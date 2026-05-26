package fakeserverca

import (
	"context"
	"crypto/x509"
	"testing"
	"time"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/server/ca"
	"github.com/spiffe/spire/pkg/server/credtemplate"
	"github.com/spiffe/spire/pkg/server/credvalidator"
	"github.com/spiffe/spire/test/clock"
	"github.com/spiffe/spire/test/testkey"
)

var (
	signer = testkey.MustEC256()
)

type Options struct {
	Clock           clock.Clock
	AgentSVIDTTL    time.Duration
	X509SVIDTTL     time.Duration
	JWTSVIDTTL      time.Duration
	WITSVIDTTL      time.Duration
	DisableJWTSVIDs bool
	DisableWITSVIDs bool
}

type CA struct {
	ca              *ca.CA
	credBuilder     *credtemplate.Builder
	credValidator   *credvalidator.Validator
	options         *Options
	bundle          []*x509.Certificate
	err             error
	disableJWTSVIDs bool
	disableWITSVIDs bool
}

func New(t *testing.T, trustDomain spiffeid.TrustDomain, options *Options) *CA {
	_ = "STUB: not implemented"
	return nil
}

func (c *CA) CredBuilder() *credtemplate.Builder { _ = "STUB: not implemented"; return nil }

func (c *CA) CredValidator() *credvalidator.Validator { _ = "STUB: not implemented"; return nil }

func (c *CA) SetX509CA(x509CA *ca.X509CA) { _ = "STUB: not implemented"; return }

func (c *CA) SetJWTKey(jwtKey *ca.JWTKey) { _ = "STUB: not implemented"; return }

func (c *CA) SetWITKey(witKey *ca.WITKey) { _ = "STUB: not implemented"; return }

func (c *CA) NotifyTaintedX509Authorities(taintedAuthorities []*x509.Certificate) {
	_ = "STUB: not implemented"
	return
}

func (c *CA) SignDownstreamX509CA(ctx context.Context, params ca.DownstreamX509CAParams) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CA) SignServerX509SVID(ctx context.Context, params ca.ServerX509SVIDParams) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CA) SignAgentX509SVID(ctx context.Context, params ca.AgentX509SVIDParams) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CA) SignWorkloadX509SVID(ctx context.Context, params ca.WorkloadX509SVIDParams) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CA) SignWorkloadJWTSVID(ctx context.Context, params ca.WorkloadJWTSVIDParams) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *CA) SignWorkloadWITSVID(ctx context.Context, params ca.WorkloadWITSVIDParams) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *CA) TaintedAuthorities() <-chan []*x509.Certificate { _ = "STUB: not implemented"; return nil }

func (c *CA) SetError(err error) { _ = "STUB: not implemented"; return }

func (c *CA) Bundle() []*x509.Certificate { _ = "STUB: not implemented"; return nil }

func (c *CA) Clock() clock.Clock { _ = "STUB: not implemented"; return *new(clock.Clock) }

func (c *CA) X509CATTL() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (c *CA) X509SVIDTTL() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (c *CA) JWTSVIDTTL() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (c *CA) WITSVIDTTL() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (c *CA) IsJWTSVIDsDisabled() bool { _ = "STUB: not implemented"; return false }

func (c *CA) IsWITSVIDsDisabled() bool { _ = "STUB: not implemented"; return false }

func (c *CA) SetDisableJWTSVIDs(disableJWTSVIDs bool) { _ = "STUB: not implemented"; return }

func (c *CA) SetDisableWITSVIDs(disableWITSVIDs bool) { _ = "STUB: not implemented"; return }
