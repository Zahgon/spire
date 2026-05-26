//go:build !windows

package containerinfo

import (
	"io"
	"regexp"
	"strings"

	"github.com/hashicorp/go-hclog"
	"k8s.io/apimachinery/pkg/types"
)

var (
	// This regex covers both the cgroupfs and systemd rendering of the pod
	// UID. The dashes are replaced with underscores in the systemd rendition.
	rePodUID = regexp.MustCompile(`\b(?:pod([[:xdigit:]]{8}[-_][[:xdigit:]]{4}[-_][[:xdigit:]]{4}[-_][[:xdigit:]]{4}[-_][[:xdigit:]]{12}))\b`)

	// The container ID is a 64-character hex string, by convention.
	reContainerID = regexp.MustCompile(`\b([[:xdigit:]]{64})\b`)

	// underToDash replaces underscores with dashes. The systemd cgroups driver
	// doesn't allow dashes so the pod UID component has dashes replaced with
	// underscores by the Kubelet.
	underToDash = strings.NewReplacer("_", "-")
)

type Extractor struct {
	RootDir        string
	VerboseLogging bool
}

func (e *Extractor) GetContainerID(pid int32, log hclog.Logger) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (e *Extractor) GetPodUIDAndContainerID(pid int32, log hclog.Logger) (types.UID, string, error) {
	_ = "STUB: not implemented"
	return *new(types.UID), "", nil
}

func (e *Extractor) extractInfo(pid int32, log hclog.Logger, extractPodUID bool) (types.UID, string, error) {
	_ = "STUB: not implemented"
	// Try to get the information from /proc/pid/mountinfo first. Otherwise,
	// fall back to /proc/pid/cgroup. If it isn't in mountinfo, then the
	// workload being attested likely originates in the same Pod as the agent.
	//
	// It may not be possible to attest a process running in the same container
	// as the agent because, depending on how cgroups are being used,
	// /proc/<pid>/mountinfo or /proc/<pid>/cgroup may not contain any
	// information on the container ID or pod.
	return *new(types.UID), "", nil
}

func (e *Extractor) extractPodUIDAndContainerIDFromMountInfo(pid int32, log hclog.Logger, extractPodUID bool) (types.UID, string, error) {
	_ = "STUB: not implemented"
	return *new(types.UID), "", nil
}

// Scan the cgroup mounts for the pod UID and container ID. The container
// ID is in the last segment, and the pod UID will be in the second to last
// segment, but only when we are attesting a different pod than the agent
// (otherwise, the second to last segment will be "..", since the agent
// exists in the same pod). In the case of cgroup v1 (or a unified
// hierarchy), there may exist multiple cgroup mounts. Out of an abundance
// of caution, all cgroup mounts will be scanned. If a containerID and/or
// pod UID are picked out of a mount, then those extracted from any of the
// remaining mounts will be checked to ensure they match. If not, we'll log
// and fail.

// In addition to cgroup mountInfo with roots at cgroup paths
// containing identifiers, some containers mount the entire
// "/sys/fs/cgroup" with mountInfo.Root = "/" which will not yield
// any pod UID or container ID. Skip to avoid FailedPrecondition
// errors with empty string for container ID and pod UID.

func (e *Extractor) extractPodUIDAndContainerIDFromCGroups(pid int32, log hclog.Logger, extractPodUID bool) (types.UID, string, error) {
	_ = "STUB: not implemented"
	return *new(types.UID), "", nil
}

type dirFS string

func (d dirFS) Open(p string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

type extractor struct {
	podUID        types.UID
	containerID   string
	extractPodUID bool
}

func (e *extractor) PodUID() types.UID { _ = "STUB: not implemented"; return *new(types.UID) }

func (e *extractor) ContainerID() string { _ = "STUB: not implemented"; return "" }

func (e *extractor) Extract(cgroupPathOrMountRoot string, log hclog.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// An entry with a pod UID overrides an entry without. If we currently have
// a pod UID and the new entry does not, then ignore it. If we currently
// don't have a pod UID and the new entry does, then override what we have
// so far.
//
// This helps mitigate situations where there is unified cgroups configured
// while running kind on macOS, which ends up with something like:
//     1:cpuset:/docker/93529524695bb00d91c1f6dba692ea8d3550c3b94fb2463af7bc9ec82f992d26/kubepods/besteffort/poda2830d0d-b0f0-4ff0-81b5-0ee4e299cf80/09bc3d7ade839efec32b6bec4ec79d099027a668ddba043083ec21d3c3b8f1e6
//     0::/docker/93529524695bb00d91c1f6dba692ea8d3550c3b94fb2463af7bc9ec82f992d26/system.slice/containerd.service
// The second entry, with only the container ID of the docker host, should
// be ignored in favor of the first entry which contains the container ID
// and pod UID of the container running in Kind.

// We currently have a pod UID and the new entry does not. Ignore it.

// We currently don't have a pod UID but have found one. Override
// the current values with the new entry.

// Check for conflicting answers for the pod UID or container ID. The safe
// action is to not choose anything.

func (e *extractor) extract(cgroupPathOrMountRoot string) (types.UID, string) {
	_ = "STUB: not implemented"
	// The container ID is typically in the last segment but in some cases
	// there can other path segments that come after. Further, some
	// combinations of kubernetes/cgroups driver/cgroups version/container
	// runtime, etc., use colon separators between the pod UID and containerID,
	// which means they can end up in the same segment together.
	//
	// The basic algorithm is to walk backwards through the path segments until
	// something that looks like a container ID is located. Once located, and
	// if the extractor is configured for it, we'll continue walking backwards
	// (starting with what's left in the segment the container ID was located
	// in) looking for the pod UID.
	return *new(types.UID), ""
}

// Walk backwards through the segments looking for the container ID. If
// found, extract the container ID and truncate the segment so that the
// remainder can (optionally) be searched for the pod UID below.

// If there is no container ID, then don't try to extract the pod UID.

// If the extractor isn't interested in the pod UID, then we're done.

// If the container ID occupied the beginning of the last segment, then
// that segment is consumed, and we should grab the next one.

// Walk backwards through the remaining segments looking for the pod UID.

// For systemd, dashes in pod UIDs are escaped to underscores. Reverse that.
