package catalog

import (
	"github.com/spiffe/spire/pkg/agent/plugin/svidstore"
	"github.com/spiffe/spire/pkg/common/catalog"
)

type svidStoreRepository struct {
	svidstore.Repository
}

func (repo *svidStoreRepository) Binder() any { _ = "STUB: not implemented"; return *new(any) }

func (repo *svidStoreRepository) Constraints() catalog.Constraints {
	_ = "STUB: not implemented"
	return *new(catalog.Constraints)
}

func (repo *svidStoreRepository) Versions() []catalog.Version {
	_ = "STUB: not implemented"
	return nil
}

func (repo *svidStoreRepository) BuiltIns() []catalog.BuiltIn {
	_ = "STUB: not implemented"
	return nil
}

type svidStoreV1 struct{}

func (svidStoreV1) New() catalog.Facade { _ = "STUB: not implemented"; return *new(catalog.Facade) }
func (svidStoreV1) Deprecated() bool    { _ = "STUB: not implemented"; return false }
