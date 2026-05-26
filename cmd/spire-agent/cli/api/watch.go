package api

import (
	"time"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/spiffe/spire/cmd/spire-agent/cli/common"
)

type WatchCLI struct {
	config *common.ConfigOS
}

func (WatchCLI) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (w WatchCLI) Help() string { _ = "STUB: not implemented"; return "" }

func (w *WatchCLI) Run(args []string) int { _ = "STUB: not implemented"; return 0 }

func (w *WatchCLI) parseConfig(args []string) error { _ = "STUB: not implemented"; return nil }

type watcher struct {
	updateTime time.Time
}

func newWatcher() *watcher { _ = "STUB: not implemented"; return nil }

func (w *watcher) OnX509ContextUpdate(x509Context *workloadapi.X509Context) {
	_ = "STUB: not implemented"
	return
}

func (w *watcher) OnX509ContextWatchError(err error) { _ = "STUB: not implemented"; return }
