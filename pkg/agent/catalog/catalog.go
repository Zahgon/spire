package catalog

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/agent/plugin/keymanager"
	"github.com/spiffe/spire/pkg/agent/plugin/nodeattestor"
	"github.com/spiffe/spire/pkg/agent/plugin/svidstore"
	"github.com/spiffe/spire/pkg/agent/plugin/workloadattestor"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/telemetry"
)

const (
	keyManagerType       = "KeyManager"
	nodeAttestorType     = "NodeAttestor"
	svidStoreType        = "SVIDStore"
	workloadattestorType = "WorkloadAttestor"
)

var ReconfigureTask = catalog.ReconfigureTask

type Catalog interface {
	GetKeyManager() keymanager.KeyManager
	GetNodeAttestor() nodeattestor.NodeAttestor
	GetSVIDStoreNamed(name string) (svidstore.SVIDStore, bool)
	GetWorkloadAttestors() []workloadattestor.WorkloadAttestor
}

type PluginConfigs = catalog.PluginConfigs

type PluginConfig = catalog.PluginConfig

type Config struct {
	Log           logrus.FieldLogger
	TrustDomain   spiffeid.TrustDomain
	PluginConfigs PluginConfigs
	Metrics       telemetry.Metrics
}

type Repository struct {
	keyManagerRepository
	nodeAttestorRepository
	svidStoreRepository
	workloadAttestorRepository

	log     logrus.FieldLogger
	catalog *catalog.Catalog
}

func (repo *Repository) Plugins() map[string]catalog.PluginRepo {
	_ = "STUB: not implemented"
	return nil
}

func (repo *Repository) Services() []catalog.ServiceRepo { _ = "STUB: not implemented"; return nil }

func (repo *Repository) Reconfigure(ctx context.Context) { _ = "STUB: not implemented"; return }

func (repo *Repository) Close() { _ = "STUB: not implemented"; return }

func Load(ctx context.Context, config Config) (_ *Repository, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Load the plugins and populate the repository

// Wrap the facades
