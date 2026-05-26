package validate

import (
	"github.com/mitchellh/cli"
	common_cli "github.com/spiffe/spire/pkg/common/cli"
)

const commandName = "validate"

func NewValidateCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newValidateCommand(env *common_cli.Env) *validateCommand {
	_ = "STUB: not implemented"
	return nil
}

type validateCommand struct {
	env *common_cli.Env
}

// Help prints the agent cmd usage
func (c *validateCommand) Help() string { _ = "STUB: not implemented"; return "" }

func (c *validateCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *validateCommand) Run(args []string) int { _ = "STUB: not implemented"; return 0 }

// Ignore error since a failure to write to stderr cannot very well be reported
