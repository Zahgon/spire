package cli

import (
	"context"

	"github.com/spiffe/spire/pkg/common/log"
)

type CLI struct {
	LogOptions         []log.Option
	AllowUnknownConfig bool
}

func (cc *CLI) Run(ctx context.Context, args []string) int { _ = "STUB: not implemented"; return 0 }
