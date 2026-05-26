//go:build !windows

package systemd

import (
	"context"
	"sync"

	"github.com/godbus/dbus/v5"
	"github.com/hashicorp/go-hclog"
	workloadattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/agent/workloadattestor/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
)

const (
	systemdDBusInterface      = "org.freedesktop.systemd1"
	systemdDBusPath           = "/org/freedesktop/systemd1"
	systemdGetUnitByPIDMethod = "org.freedesktop.systemd1.Manager.GetUnitByPID"
)

func builtin(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

type DBusUnitInfo struct {
	UnitID           string
	UnitFragmentPath string
}

type Plugin struct {
	workloadattestorv1.UnsafeWorkloadAttestorServer

	log hclog.Logger

	dbusMutex sync.Mutex
	dbusConn  *dbus.Conn

	// hook for tests
	getUnitInfo func(ctx context.Context, p *Plugin, pid uint) (*DBusUnitInfo, error)
}

func New() *Plugin { _ = "STUB: not implemented"; return nil }

func (p *Plugin) SetLogger(log hclog.Logger) { _ = "STUB: not implemented"; return }

func (p *Plugin) Attest(ctx context.Context, req *workloadattestorv1.AttestRequest) (*workloadattestorv1.AttestResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) Close() error { _ = "STUB: not implemented"; return nil }

func (p *Plugin) getDBusConn() (*dbus.Conn, error) { _ = "STUB: not implemented"; return nil, nil }

func getSystemdUnitInfo(ctx context.Context, p *Plugin, pid uint) (*DBusUnitInfo, error) {
	_ = "STUB: not implemented"
	// We are not closing the connection here because it's closed when the Close() function is called as part of unloading the plugin.
	return nil, nil
}

// Get the unit for the given PID from the systemd service.

func getStringProperty(obj dbus.BusObject, prop string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func makeSelectorValue(kind, value string) string { _ = "STUB: not implemented"; return "" }
