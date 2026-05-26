package catalog

import (
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/server/plugin/nodeattestor"
)

type nodeAttestorRepository struct {
	nodeattestor.Repository
}

func (repo *nodeAttestorRepository) Binder() any { _ = "STUB: not implemented"; return *new(any) }

func (repo *nodeAttestorRepository) Constraints() catalog.Constraints {
	_ = "STUB: not implemented"
	return *new(catalog.Constraints)
}

func (repo *nodeAttestorRepository) Versions() []catalog.Version {
	_ = "STUB: not implemented"
	return nil
}

func (repo *nodeAttestorRepository) BuiltIns() []catalog.BuiltIn {
	_ = "STUB: not implemented"
	return nil
}

type nodeAttestorV1 struct{}

func (nodeAttestorV1) New() catalog.Facade { _ = "STUB: not implemented"; return *new(catalog.Facade) }
func (nodeAttestorV1) Deprecated() bool    { _ = "STUB: not implemented"; return false }
