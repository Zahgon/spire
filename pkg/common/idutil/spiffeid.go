package idutil

import (
	"github.com/spiffe/go-spiffe/v2/spiffeid"
)

const (
	ServerIDPath = "/spire/server"
)

func MemberFromString(td spiffeid.TrustDomain, s string) (spiffeid.ID, error) {
	_ = "STUB: not implemented"
	return *new(spiffeid.ID), nil
}

// IsAgentPath returns true if the given string is an
// SPIRE agent ID path. SPIRE agent IDs are prefixed
// with "/spire/agent/".
func IsAgentPath(path string) bool { _ = "STUB: not implemented"; return false }

// IsAgentPathForNodeAttestor returns if the path lives under the agent
// namesepace for the given node attestor
func IsAgentPathForNodeAttestor(path string, nodeAttestor string) bool {
	_ = "STUB: not implemented"
	return false
}

func IsReservedPath(path string) bool { _ = "STUB: not implemented"; return false }

// AgentID creates an agent SPIFFE ID given a trust domain and a path suffix.
// The path suffix must be an absolute path. The /spire/agent prefix is
// prefixed to the suffix to form the path.
func AgentID(td spiffeid.TrustDomain, suffix string) (spiffeid.ID, error) {
	_ = "STUB: not implemented"
	return *new(spiffeid.ID), nil
}

// ServerID creates a server SPIFFE ID string given a trust domain.
func ServerID(td spiffeid.TrustDomain) (spiffeid.ID, error) {
	_ = "STUB: not implemented"
	return *new(spiffeid.ID), nil
}
