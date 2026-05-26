package cli

import (
	"context"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/pkg/common/log"
)

// CLI defines the server CLI configuration.
type CLI struct {
	LogOptions         []log.Option
	AllowUnknownConfig bool
}

// Run configures the server CLI commands and subcommands.
func (cc *CLI) Run(ctx context.Context, args []string) int { _ = "STUB: not implemented"; return 0 }

// addCommandsEnabledByFFlags adds commands that are currently available only
// through a feature flag.
// Feature flags support through the fflag package in SPIRE Server is
// designed to work only with the run command and the config file.
// Since feature flags are intended to be used by developers of a specific
// feature only, exposing them through command line arguments is not
// convenient. Instead, we use the SPIRE_SERVER_FFLAGS environment variable
// to read the configured SPIRE Server feature flags from the environment
// when other commands may be enabled through feature flags.
func addCommandsEnabledByFFlags(commands map[string]cli.CommandFactory) {
	_ = "STUB: not implemented"
	return
}
