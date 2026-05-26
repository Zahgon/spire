package protoutil

import (
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/proto/spire/common"
	"google.golang.org/protobuf/proto"
)

var (
	AllTrueAgentMask                  = MakeAllTrueMask(&types.AgentMask{}).(*types.AgentMask)
	AllTrueBundleMask                 = MakeAllTrueMask(&types.BundleMask{}).(*types.BundleMask)
	AllTrueEntryMask                  = MakeAllTrueMask(&types.EntryMask{}).(*types.EntryMask)
	AllTrueFederationRelationshipMask = MakeAllTrueMask(&types.FederationRelationshipMask{}).(*types.FederationRelationshipMask)

	AllTrueCommonBundleMask = MakeAllTrueMask(&common.BundleMask{}).(*common.BundleMask)
	AllTrueCommonAgentMask  = MakeAllTrueMask(&common.AttestedNodeMask{}).(*common.AttestedNodeMask)
)

func MakeAllTrueMask(m proto.Message) proto.Message {
	_ = "STUB: not implemented"
	return *new(proto.Message)
}

// Skip the protobuf internal fields or those that aren't bools
