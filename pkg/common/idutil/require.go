package idutil

import (
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
)

// RequireIDProtoString constructs a SPIFFE ID string for the given ID proto.
// It panics if the proto is not well-formed.
func RequireIDProtoString(id *types.SPIFFEID) string { _ = "STUB: not implemented"; return "" }

// RequireIDFromProto returns a SPIFFE ID from the proto representation. It
// panics if the proto is not well-formed.
func RequireIDFromProto(id *types.SPIFFEID) spiffeid.ID {
	_ = "STUB: not implemented"
	return *new(spiffeid.ID)
}

// RequireServerID returns the server SPIFFE ID for the given trust domain. It
// panics if the given trust domain isn't valid.
func RequireServerID(td spiffeid.TrustDomain) spiffeid.ID {
	_ = "STUB: not implemented"
	return *new(spiffeid.ID)
}

// RequireAgentID creates an agent SPIFFE ID given a trust domain and a path
// suffix. The path suffix must be an absolute path. The /spire/agent prefix is
// prefixed to the suffix to form the path. It panics if the given trust domain
// isn't valid.
func RequireAgentID(td spiffeid.TrustDomain, suffix string) spiffeid.ID {
	_ = "STUB: not implemented"
	return *new(spiffeid.ID)
}

func panicOnErr(err error) { _ = "STUB: not implemented"; return }
