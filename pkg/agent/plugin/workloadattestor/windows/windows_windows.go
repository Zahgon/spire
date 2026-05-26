//go:build windows

package windows

import (
	"context"
	"sync"

	"github.com/hashicorp/go-hclog"
	workloadattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/agent/workloadattestor/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
	"golang.org/x/sys/windows"
)

func builtin(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func New() *Plugin { _ = "STUB: not implemented"; return nil }

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

	log hclog.Logger
	q   processQueryer
}

type processInfo struct {
	pid        int32
	user       string
	userSID    string
	path       string
	groups     []string
	groupsSIDs []string
}

func (p *Plugin) SetLogger(log hclog.Logger) { _ = "STUB: not implemented"; return }

func (p *Plugin) Attest(_ context.Context, req *workloadattestorv1.AttestRequest) (*workloadattestorv1.AttestResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// obtaining the workload process path and digest are behind a config flag
// since it requires the agent to have permissions that might not be
// available.

func (p *Plugin) newProcessInfo(pid int32, queryPath bool) (*processInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve an access token to describe the security context of
// the process from which we obtained the handle.

// Get user information

// Get groups information

// Each group has a set of attributes that control how
// the system uses the SID in an access check.
// We are interested in the SE_GROUP_ENABLED attribute.
// https://docs.microsoft.com/en-us/windows/win32/secauthz/sid-attributes-in-an-access-token

// If the LookupAccount call succeeded, we know that groupAccount is not empty

func (p *Plugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) getConfig() (*Configuration, error) { _ = "STUB: not implemented"; return nil, nil }

type processQueryer interface {
	// OpenProcess returns an open handle to the specified process id.
	OpenProcess(int32) (windows.Handle, error)

	// OpenProcessToken opens the access token associated with a process.
	OpenProcessToken(windows.Handle, *windows.Token) error

	// LookupAccount retrieves the name of the account for the specified
	// SID and the name of the first domain on which that SID is found.
	LookupAccount(sid *windows.SID) (account, domain string, err error)

	// GetTokenUser retrieves user account information of the
	// specified token.
	GetTokenUser(*windows.Token) (*windows.Tokenuser, error)

	// GetTokenGroups retrieves group accounts information of the
	// specified token.
	GetTokenGroups(*windows.Token) (*windows.Tokengroups, error)

	// AllGroups returns a slice that can be used to iterate over
	// the specified Tokengroups.
	AllGroups(*windows.Tokengroups) []windows.SIDAndAttributes

	// CloseHandle closes an open object handle.
	CloseHandle(windows.Handle) error

	// CloseProcessToken releases access to the specified access token.
	CloseProcessToken(windows.Token) error

	// GetProcessExe returns the executable file path relating to the
	// specified process handle.
	GetProcessExe(windows.Handle) (string, error)
}

type processQuery struct{}

func (q *processQuery) OpenProcess(pid int32) (handle windows.Handle, err error) {
	_ = "STUB: not implemented"
	return *new(windows.Handle), nil
}

func (q *processQuery) OpenProcessToken(h windows.Handle, token *windows.Token) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (q *processQuery) LookupAccount(sid *windows.SID) (account, domain string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func (q *processQuery) GetTokenUser(t *windows.Token) (*windows.Tokenuser, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *processQuery) GetTokenGroups(t *windows.Token) (*windows.Tokengroups, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *processQuery) AllGroups(t *windows.Tokengroups) []windows.SIDAndAttributes {
	_ = "STUB: not implemented"
	return nil
}

func (q *processQuery) CloseHandle(h windows.Handle) error { _ = "STUB: not implemented"; return nil }

func (q *processQuery) CloseProcessToken(t windows.Token) error {
	_ = "STUB: not implemented"
	return nil
}

func (q *processQuery) GetProcessExe(h windows.Handle) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func addSelectorValueIfNotEmpty(selectorValues []string, kind, value string) []string {
	_ = "STUB: not implemented"
	return nil
}

func parseAccount(account, domain string) string { _ = "STUB: not implemented"; return "" }

func getGroupEnabledSelector(attributes uint32) string { _ = "STUB: not implemented"; return "" }

func makeSelectorValue(kind, value string) string { _ = "STUB: not implemented"; return "" }
