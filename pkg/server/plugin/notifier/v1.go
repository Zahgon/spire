package notifier

import (
	"context"

	notifierv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/notifier/v1"
	"github.com/spiffe/spire/pkg/common/plugin"
	"github.com/spiffe/spire/proto/spire/common"
)

type V1 struct {
	plugin.Facade
	notifierv1.NotifierPluginClient
}

func (v1 *V1) NotifyAndAdviseBundleLoaded(ctx context.Context, b *common.Bundle) error {
	_ = "STUB: not implemented"
	return nil
}

func (v1 *V1) NotifyBundleUpdated(ctx context.Context, b *common.Bundle) error {
	_ = "STUB: not implemented"
	return nil
}
