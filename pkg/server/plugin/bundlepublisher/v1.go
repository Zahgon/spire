package bundlepublisher

import (
	"context"

	bundlepublisherv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/bundlepublisher/v1"
	"github.com/spiffe/spire/pkg/common/plugin"
	"github.com/spiffe/spire/proto/spire/common"
)

type V1 struct {
	plugin.Facade
	bundlepublisherv1.BundlePublisherPluginClient
}

func (v1 *V1) PublishBundle(ctx context.Context, b *common.Bundle) error {
	_ = "STUB: not implemented"
	return nil
}
