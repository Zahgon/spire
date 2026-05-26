//go:build windows

package k8spsat

const (
	containerMountPointEnvVar = "CONTAINER_SANDBOX_MOUNT_POINT"
)

func getDefaultTokenPath() string { _ = "STUB: not implemented"; return "" }
