package svidstore

import (
	"context"

	svidstorev1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/agent/svidstore/v1"
	"github.com/spiffe/spire/pkg/common/plugin"
)

type V1 struct {
	plugin.Facade

	svidstorev1.SVIDStorePluginClient
}

func (v1 *V1) DeleteX509SVID(ctx context.Context, metadata []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (v1 *V1) PutX509SVID(ctx context.Context, x509SVID *X509SVID) error {
	_ = "STUB: not implemented"
	return nil
}
