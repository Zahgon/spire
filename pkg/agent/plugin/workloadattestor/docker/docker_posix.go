//go:build !windows

package docker

import (
	"io"
	"regexp"

	"github.com/hashicorp/go-hclog"
	"github.com/spiffe/spire/pkg/agent/common/cgroups"
	"github.com/spiffe/spire/pkg/agent/plugin/workloadattestor/docker/cgroup"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

const (
	defaultPodmanSocketPath         = "unix:///run/podman/podman.sock"
	defaultPodmanSocketPathTemplate = "unix:///run/user/%d/podman/podman.sock"
)

var (
	rePodmanCgroup = regexp.MustCompile(`(?:libpod-|/libpod/)`)
	reUserSliceUID = regexp.MustCompile(`/user-(\d+)\.slice/`)
)

type OSConfig struct {
	// DockerSocketPath is the location of the docker daemon socket, this config can be used only on unix environments (default: "unix:///var/run/docker.sock").
	DockerSocketPath string `hcl:"docker_socket_path" json:"docker_socket_path"`

	// ContainerIDCGroupMatchers is a list of patterns used to discover container IDs from cgroup entries.
	// See the documentation for cgroup.NewContainerIDFinder in the cgroup subpackage for more information. (Unix)
	ContainerIDCGroupMatchers []string `hcl:"container_id_cgroup_matchers" json:"container_id_cgroup_matchers"`

	// UseNewContainerLocator, if true, uses the new container locator
	// mechanism instead of cgroup matchers. Currently defaults to false if
	// unset. This will default to true in a future release. (Unix)
	UseNewContainerLocator *bool `hcl:"use_new_container_locator"`

	// VerboseContainerLocatorLogs, if true, dumps extra information to the log
	// about mountinfo and cgroup information used to locate the container.
	VerboseContainerLocatorLogs bool `hcl:"verbose_container_locator_logs"`

	// PodmanSocketPath is the socket path for rootful Podman (no user namespace).
	// Defaults to "unix:///run/podman/podman.sock".
	PodmanSocketPath string `hcl:"podman_socket_path" json:"podman_socket_path"`

	// PodmanSocketPathTemplate is the socket path template for rootless Podman.
	// The placeholder %d is replaced with the container owner's host UID extracted
	// from the cgroup path. Defaults to "unix:///run/user/%d/podman/podman.sock".
	PodmanSocketPathTemplate string `hcl:"podman_socket_path_template" json:"podman_socket_path_template"`

	// Used by tests to use a fake /proc directory instead of the real one
	rootDir string
}

func (p *Plugin) createHelper(c *dockerPluginConfig, status *pluginconf.Status) *containerHelper {
	_ = "STUB: not implemented"
	return nil
}

type dirFS string

func (d dirFS) Open(p string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

type containerHelper struct {
	rootDir                     string
	containerIDFinder           cgroup.ContainerIDFinder
	verboseContainerLocatorLogs bool
	podmanSocketPath            string
	podmanSocketPathTemplate    string
}

func (h *containerHelper) getContainerIDAndSocket(pID int32, log hclog.Logger) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func (h *containerHelper) detectPodmanSocket(cgroupList []cgroups.Cgroup, log hclog.Logger) string {
	_ = "STUB: not implemented"
	return ""
}

func validatePodmanSocketPathTemplate(template string) error { _ = "STUB: not implemented"; return nil }

func getDockerHost(c *dockerPluginConfig) string { _ = "STUB: not implemented"; return "" }

// getContainerIDFromCGroups returns the container ID from a set of cgroups
// using the given finder. The container ID found on each cgroup path (if any)
// must be consistent. If no container ID is found among the cgroups, i.e.,
// this isn't a docker workload, the function returns an empty string. If more
// than one container ID is found, or the "found" container ID is blank, the
// function will fail.
func getContainerIDFromCGroups(finder cgroup.ContainerIDFinder, cgroups []cgroups.Cgroup) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// This is the first container ID found so far.

// More than one container ID found in the cgroups.

// Not a docker workload. Since it is expected that non-docker workloads will call the
// workload API, it is fine to return a response without any selectors.

// The "finder" found a container ID, but it was blank. This is a
// defensive measure against bad matcher patterns and shouldn't
// be possible with the default finder.
