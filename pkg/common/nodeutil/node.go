package nodeutil

import (
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/proto/spire/common"
)

var (
	shouldReattest = map[types.PermissionDeniedDetails_Reason]struct{}{
		types.PermissionDeniedDetails_AGENT_EXPIRED:       {},
		types.PermissionDeniedDetails_AGENT_NOT_ACTIVE:    {},
		types.PermissionDeniedDetails_AGENT_NOT_ATTESTED:  {},
		types.PermissionDeniedDetails_AGENT_MUST_REATTEST: {},
	}
	shouldShutDown = map[types.PermissionDeniedDetails_Reason]struct{}{
		types.PermissionDeniedDetails_AGENT_BANNED: {},
	}
)

// IsAgentBanned determines if a given attested node is banned or not.
// An agent is considered as "banned" if its X509 SVID serial number is empty.
func IsAgentBanned(node *common.AttestedNode) bool { _ = "STUB: not implemented"; return false }

// ShouldAgentReattest returns true if the Server returned an error worth rebooting the Agent
func ShouldAgentReattest(err error) bool { _ = "STUB: not implemented"; return false }

// ShouldAgentShutdown returns true if the Server returned an error worth shutting down the Agent
func ShouldAgentShutdown(err error) bool { _ = "STUB: not implemented"; return false }

func isExpectedPermissionDenied(err error, expectedReason map[types.PermissionDeniedDetails_Reason]struct{}) bool {
	_ = "STUB: not implemented"
	return false
}
