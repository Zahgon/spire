package authoritycommon

import (
	localauthorityv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/localauthority/v1"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
)

func PrettyPrintJWTAuthorityState(env *commoncli.Env, authorityState *localauthorityv1.AuthorityState) {
	_ = "STUB: not implemented"
	return
}

func PrettyPrintX509AuthorityState(env *commoncli.Env, authorityState *localauthorityv1.AuthorityState) {
	_ = "STUB: not implemented"
	return
}

func prettyPrintAuthorityState(env *commoncli.Env, authorityState *localauthorityv1.AuthorityState, includeUpstreamAuthority bool) {
	_ = "STUB: not implemented"
	return
}
