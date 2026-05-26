package nodeattestor

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire/pkg/common/plugin"
)

func JoinToken(log logrus.FieldLogger, token string) NodeAttestor {
	_ = "STUB: not implemented"
	return *new(NodeAttestor)
}

type joinToken struct {
	plugin.Facade
	token string
}

func (plugin joinToken) Attest(ctx context.Context, serverStream ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}
