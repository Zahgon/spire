//go:build !windows

package unix

import (
	"context"
	"os/user"
	"sync"

	"github.com/hashicorp/go-hclog"
	"github.com/shirou/gopsutil/v4/process"
	workloadattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/agent/workloadattestor/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

func builtin(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

type processInfo interface {
	Uids() ([]uint32, error)
	Gids() ([]uint32, error)
	Groups() ([]string, error)
	Exe() (string, error)
	NamespacedExe() string
}

type PSProcessInfo struct {
	*process.Process
}

func (ps PSProcessInfo) NamespacedExe() string { _ = "STUB: not implemented"; return "" }

// Groups returns the supplementary group IDs
// This is a custom implementation that only works for linux until the next issue is fixed
// https://github.com/shirou/gopsutil/issues/913
func (ps PSProcessInfo) Groups() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

type Configuration struct {
	DiscoverWorkloadPath bool  `hcl:"discover_workload_path"`
	WorkloadSizeLimit    int64 `hcl:"workload_size_limit"`
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *Configuration {
	_ = "STUB: not implemented"
	return nil
}

type Plugin struct {
	workloadattestorv1.UnsafeWorkloadAttestorServer
	configv1.UnsafeConfigServer

	mu     sync.Mutex
	config *Configuration
	log    hclog.Logger

	// hooks for tests
	hooks struct {
		newProcess      func(pid int32) (processInfo, error)
		lookupUserByID  func(id string) (*user.User, error)
		lookupGroupByID func(id string) (*user.Group, error)
	}
}

func New() *Plugin { _ = "STUB: not implemented"; return nil }

func (p *Plugin) SetLogger(log hclog.Logger) { _ = "STUB: not implemented"; return }

func (p *Plugin) Attest(_ context.Context, req *workloadattestorv1.AttestRequest) (*workloadattestorv1.AttestResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// obtaining the workload process path and digest are behind a config flag
// since it requires the agent to have permissions that might not be
// available.

func (p *Plugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) getConfig() (*Configuration, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *Plugin) getUID(proc processInfo) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *Plugin) getUserName(uid string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (p *Plugin) getGID(proc processInfo) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *Plugin) getGroupName(gid string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (p *Plugin) getPath(proc processInfo) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *Plugin) getNamespacedPath(proc processInfo) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func makeSelectorValue(kind, value string) string { _ = "STUB: not implemented"; return "" }

func getProcPath(pID int32, lastPath string) string { _ = "STUB: not implemented"; return "" }
