package testplugin

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/spiffe/spire-plugin-sdk/pluginsdk"
	"github.com/spiffe/spire-plugin-sdk/private/proto/test"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
)

type Plugin struct {
	test.UnimplementedSomePluginServer
	test.UnimplementedSomeServiceServer
	configv1.UnimplementedConfigServer

	log         hclog.Logger
	hostService test.SomeHostServiceServiceClient
}

var _ pluginsdk.NeedsLogger = (*Plugin)(nil)
var _ pluginsdk.NeedsHostServices = (*Plugin)(nil)

func BuiltIn(registerConfig bool) catalog.BuiltIn {
	_ = "STUB: not implemented"
	return *new(catalog.BuiltIn)
}

func (p *Plugin) SetLogger(log hclog.Logger) { _ = "STUB: not implemented"; return }

func (p *Plugin) BrokerHostServices(broker pluginsdk.ServiceBroker) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) PluginEcho(ctx context.Context, req *test.EchoRequest) (*test.EchoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) ServiceEcho(ctx context.Context, req *test.EchoRequest) (*test.EchoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) Close() error { _ = "STUB: not implemented"; return nil }

type SomeHostService struct {
	test.UnimplementedSomeHostServiceServer
}

func (SomeHostService) HostServiceEcho(ctx context.Context, req *test.EchoRequest) (*test.EchoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func wrap(s string, with string) string { _ = "STUB: not implemented"; return "" }
