package fakeupstreamauthority

import (
	"testing"

	"github.com/spiffe/spire/pkg/server/plugin/upstreamauthority"
)

func Load(t *testing.T, config Config) (upstreamauthority.UpstreamAuthority, *UpstreamAuthority) {
	_ = "STUB: not implemented"
	return *new(upstreamauthority.UpstreamAuthority), nil
}
