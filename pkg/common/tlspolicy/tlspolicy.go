// Package tlspolicy provides for configuration and enforcement of policies
// relating to TLS.
package tlspolicy

import (
	"crypto/tls"

	"github.com/hashicorp/go-hclog"
)

// Policy describes policy options to be applied to a TLS configuration.
//
// A zero-initialised Policy provides reasonable defaults.
type Policy struct {
	// RequirePQKEM determines if a post-quantum-safe KEM should be required for
	// TLS connections.
	RequirePQKEM bool
}

// LogPolicy logs an informational message reporting the configured policy,
// aiding administrators to determine what policy options have been
// successfully enabled.
func LogPolicy(policy Policy, logger hclog.Logger) { _ = "STUB: not implemented"; return }

// ApplyPolicy applies the policy options in policy to a given tls.Config,
// which is assumed to have already been obtained from the go-spiffe tlsconfig
// package.
func ApplyPolicy(config *tls.Config, policy Policy) error { _ = "STUB: not implemented"; return nil }

// List only known PQ-safe KEMs as valid curves.

// Require TLS 1.3, as all PQ-safe KEMs require it anyway.
