package x509

import (
	"context"
	"flag"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-server/util"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/cliprinter"
)

// NewX509TaintCommand creates a new "x509 taint" subcommand for "localauthority" command.
func NewX509TaintCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

// NewX509TaintCommandWithEnv creates a new "x509 taint" subcommand for "localauthority" command
// using the environment specified
func NewX509TaintCommandWithEnv(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

func newX509TaintCommand(env *commoncli.Env) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

type x509TaintCommand struct {
	authorityID string
	printer     cliprinter.Printer
	env         *commoncli.Env
}

func (c *x509TaintCommand) Name() string { _ = "STUB: not implemented"; return "" }

func (*x509TaintCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *x509TaintCommand) AppendFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

// Run executes all logic associated with a single invocation of the
// `spire-server localauthority x509 taint` CLI command
func (c *x509TaintCommand) Run(ctx context.Context, _ *commoncli.Env, serverClient util.ServerClient) error {
	_ = "STUB: not implemented"
	return nil
}

func prettyPrintX509Taint(env *commoncli.Env, results ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *x509TaintCommand) validate() error { _ = "STUB: not implemented"; return nil }
