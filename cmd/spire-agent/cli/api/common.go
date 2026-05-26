package api

import (
	"context"
	"flag"
	"net"
	"time"

	"github.com/spiffe/go-spiffe/v2/proto/spiffe/workload"
	"github.com/spiffe/spire/cmd/spire-agent/cli/common"
	"github.com/spiffe/spire/pkg/common/cli"
	"google.golang.org/grpc"
)

const commandTimeout = 5 * time.Second

type workloadClient struct {
	workload.SpiffeWorkloadAPIClient
	conn    *grpc.ClientConn
	timeout time.Duration
}

func (c *workloadClient) release() { _ = "STUB: not implemented"; return }

type workloadClientMaker func(ctx context.Context, addr net.Addr, timeout time.Duration) (*workloadClient, error)

// newClients is the default client maker
func newWorkloadClient(ctx context.Context, addr net.Addr, timeout time.Duration) (*workloadClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *workloadClient) prepareContext(ctx context.Context) (context.Context, func()) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// command is a common interface for commands in this package. the adapter
// can adapter this interface to the Command interface from github.com/mitchellh/cli.
type command interface {
	name() string
	synopsis() string
	appendFlags(*flag.FlagSet)
	run(context.Context, *cli.Env, *workloadClient) error
}

type adapter struct {
	common.ConfigOS // os specific

	env          *cli.Env
	clientsMaker workloadClientMaker
	cmd          command

	timeout cli.DurationFlag
	flags   *flag.FlagSet
}

// adaptCommand converts a command into one conforming to the Command interface from github.com/mitchellh/cli
func adaptCommand(env *cli.Env, clientsMaker workloadClientMaker, cmd command) *adapter {
	_ = "STUB: not implemented"
	return nil
}

func (a *adapter) Run(args []string) int { _ = "STUB: not implemented"; return 0 }

func (a *adapter) Help() string { _ = "STUB: not implemented"; return "" }

func (a *adapter) Synopsis() string { _ = "STUB: not implemented"; return "" }
