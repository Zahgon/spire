package docker

import (
	"context"
	"sync"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	dockerclient "github.com/docker/docker/client"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/hcl/hcl/token"
	workloadattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/agent/workloadattestor/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/agent/common/sigstore"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

const (
	pluginName                   = "docker"
	subselectorLabel             = "label"
	subselectorImageID           = "image_id"
	subselectorEnv               = "env"
	subselectorImageConfigDigest = "image_config_digest"
)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

// Docker is a subset of the docker client functionality, useful for mocking.
type Docker interface {
	ContainerInspect(ctx context.Context, containerID string) (container.InspectResponse, error)
	ImageInspectWithRaw(ctx context.Context, imageID string) (image.InspectResponse, []byte, error)
}

type podmanDocker interface {
	Docker
	Close() error
}

type Plugin struct {
	workloadattestorv1.UnsafeWorkloadAttestorServer
	configv1.UnsafeConfigServer

	log     hclog.Logger
	retryer *retryer

	mtx                 sync.RWMutex
	docker              Docker
	c                   *containerHelper
	sigstoreVerifier    sigstore.Verifier
	podmanClientFactory func(socketPath string) (podmanDocker, error)
}

func New() *Plugin { _ = "STUB: not implemented"; return nil }

func defaultPodmanClientFactory(socketPath string) (podmanDocker, error) {
	_ = "STUB: not implemented"
	return *new(podmanDocker), nil
}

type dockerPluginConfig struct {
	OSConfig `hcl:",squash"`

	// DockerVersion is the API version of the docker daemon. If not specified, the version is negotiated by the client.
	DockerVersion string `hcl:"docker_version" json:"docker_version"`

	UnusedKeyPositions map[string][]token.Pos `hcl:",unusedKeyPositions"`

	// Sigstore contains sigstore specific configs.
	Sigstore *sigstore.HCLConfig `hcl:"sigstore,omitempty" json:"sigstore"`

	containerHelper *containerHelper
	dockerOpts      []dockerclient.Opt
	sigstoreConfig  *sigstore.Config
}

func (p *Plugin) buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *dockerPluginConfig {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) SetLogger(log hclog.Logger) { _ = "STUB: not implemented"; return }

func (p *Plugin) Attest(ctx context.Context, req *workloadattestorv1.AttestRequest) (*workloadattestorv1.AttestResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Not a docker workload. Nothing more to do.

// Add image_config_digest selector

// RepoDigests is a list of content-addressable digests of locally available
// image manifests that the image is referenced from. Multiple manifests can
// refer to the same image.

func getSelectorValuesFromConfig(cfg *container.Config) []string {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) Configure(ctx context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
