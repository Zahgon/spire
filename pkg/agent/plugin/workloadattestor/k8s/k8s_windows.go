//go:build windows

package k8s

import (
	"github.com/hashicorp/go-hclog"
	"github.com/spiffe/spire/pkg/common/container/process"
	"k8s.io/apimachinery/pkg/types"
)

const (
	containerMountPointEnvVar = "CONTAINER_SANDBOX_MOUNT_POINT"
)

func createHelper(*Plugin) ContainerHelper { _ = "STUB: not implemented"; return *new(ContainerHelper) }

type containerHelper struct {
	ph process.Helper
}

func (h *containerHelper) Configure(_ *HCLConfig, _ hclog.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *containerHelper) GetPodUIDAndContainerID(pID int32, log hclog.Logger) (types.UID, string, error) {
	_ = "STUB: not implemented"
	return *new(types.UID), "", nil
}

func (p *Plugin) defaultKubeletCAPath() string { _ = "STUB: not implemented"; return "" }

func (p *Plugin) defaultTokenPath() string { _ = "STUB: not implemented"; return "" }
