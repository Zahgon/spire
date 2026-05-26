package base

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/spiffe/spire-plugin-sdk/pluginsdk"
	agentstorev1 "github.com/spiffe/spire-plugin-sdk/proto/spire/hostservice/server/agentstore/v1"
)

type Base struct {
	store agentstorev1.AgentStoreServiceClient
}

var _ pluginsdk.NeedsHostServices = (*Base)(nil)

func (p *Base) BrokerHostServices(broker pluginsdk.ServiceBroker) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Base) AssessTOFU(ctx context.Context, agentID string, log hclog.Logger) error {
	_ = "STUB: not implemented"
	return nil
}
