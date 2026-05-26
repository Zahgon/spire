//go:build windows

package docker

import (
	hclog "github.com/hashicorp/go-hclog"
	"github.com/spiffe/spire/pkg/common/container/process"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

type OSConfig struct {
	// DockerHost is the location of the Docker Engine API endpoint on Windows (default: "npipe:////./pipe/docker_engine").
	DockerHost string `hcl:"docker_host" json:"docker_host"`
}

func (p *Plugin) createHelper(*dockerPluginConfig, *pluginconf.Status) *containerHelper {
	_ = "STUB: not implemented"
	return nil
}

type containerHelper struct {
	ph process.Helper
}

func (h *containerHelper) getContainerIDAndSocket(pID int32, log hclog.Logger) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func getDockerHost(c *dockerPluginConfig) string { _ = "STUB: not implemented"; return "" }
