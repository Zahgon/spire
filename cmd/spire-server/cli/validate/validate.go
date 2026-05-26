package validate

import (
	"time"

	"github.com/mitchellh/cli"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
)

const (
	commandName       = "validate"
	validationTimeout = 30 * time.Second
)

func NewValidateCommand() cli.Command { _ = "STUB: not implemented"; return *new(cli.Command) }

func newValidateCommand(env *commoncli.Env) *validateCommand { _ = "STUB: not implemented"; return nil }

type validateCommand struct {
	env *commoncli.Env
}

// Help prints the server cmd usage
func (c *validateCommand) Help() string { _ = "STUB: not implemented"; return "" }

func (c *validateCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *validateCommand) Run(args []string) int { _ = "STUB: not implemented"; return 0 }

// Ignore error since a failure to write to stderr cannot very well be reported
