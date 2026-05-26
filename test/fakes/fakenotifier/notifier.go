package fakenotifier

import (
	"context"
	"testing"

	notifierv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/notifier/v1"
	"github.com/spiffe/spire/pkg/server/plugin/notifier"
	"github.com/spiffe/spire/proto/spire/common"
)

type Config struct {
	OnNotifyBundleUpdated         func(*common.Bundle) error
	OnNotifyAndAdviseBundleLoaded func(*common.Bundle) error
}

func New(t *testing.T, config Config) notifier.Notifier {
	_ = "STUB: not implemented"
	return *new(notifier.Notifier)
}

type fakeNotifier struct {
	notifierv1.UnimplementedNotifierServer

	config Config
}

func (n *fakeNotifier) Notify(_ context.Context, req *notifierv1.NotifyRequest) (*notifierv1.NotifyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *fakeNotifier) NotifyAndAdvise(_ context.Context, req *notifierv1.NotifyAndAdviseRequest) (*notifierv1.NotifyAndAdviseResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NotifyBundleUpdatedWaiter(t *testing.T) (notifier.Notifier, <-chan *common.Bundle) {
	_ = "STUB: not implemented"
	return *new(notifier.Notifier), nil
}
