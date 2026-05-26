package catalog

import (
	"github.com/spiffe/spire/pkg/common/catalog"

	"github.com/spiffe/spire/pkg/agent/plugin/keymanager"
)

type keyManagerRepository struct {
	keymanager.Repository
}

func (repo *keyManagerRepository) Binder() any { _ = "STUB: not implemented"; return *new(any) }

func (repo *keyManagerRepository) Constraints() catalog.Constraints {
	_ = "STUB: not implemented"
	return *new(catalog.Constraints)
}

func (repo *keyManagerRepository) Versions() []catalog.Version {
	_ = "STUB: not implemented"
	return nil
}

func (repo *keyManagerRepository) BuiltIns() []catalog.BuiltIn {
	_ = "STUB: not implemented"
	return nil
}

type keyManagerV1 struct{}

func (keyManagerV1) New() catalog.Facade { _ = "STUB: not implemented"; return *new(catalog.Facade) }
func (keyManagerV1) Deprecated() bool    { _ = "STUB: not implemented"; return false }
