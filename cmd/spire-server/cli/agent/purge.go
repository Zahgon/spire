package agent

import (
	"context"
	"flag"
	"time"

	"github.com/mitchellh/cli"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

type purgeCommand struct {
	env        *commoncli.Env
	expiredFor time.Duration
	dryRun     bool
	printer    cliprinter.Printer
}

func NewPurgeCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func NewPurgeCommandWithEnv(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

func (*purgeCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*purgeCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *purgeCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *purgeCommand) AppendFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

type expiredAgents struct {
	Agents []*expiredAgent `json:"expired_agents"`
}

type expiredAgent struct {
	AgentID spiffeid.ID `json:"agent_id"`
	Deleted bool        `json:"deleted"`
	Error   string      `json:"error,omitempty"`
}

func (c *purgeCommand) prettyPrintPurgeResult(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *purgeCommand) printAgentsNotPurged(agentsNotPurged []*expiredAgent) {
	_ = "STUB: not implemented"
	return
}

func (c *purgeCommand) printAgentsPurged(agentsPurged []*expiredAgent) {
	_ = "STUB: not implemented"
	return
}
