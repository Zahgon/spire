package catalog

import (
	"github.com/spiffe/spire/pkg/agent/plugin/workloadattestor"
	"github.com/spiffe/spire/pkg/common/catalog"
)

type workloadAttestorRepository struct {
	workloadattestor.Repository
}

func (repo *workloadAttestorRepository) Binder() any { _ = "STUB: not implemented"; return *new(any) }

func (repo *workloadAttestorRepository) Constraints() catalog.Constraints {
	_ = "STUB: not implemented"
	return *new(catalog.Constraints)
}

func (repo *workloadAttestorRepository) Versions() []catalog.Version {
	_ = "STUB: not implemented"
	return nil
}

func (repo *workloadAttestorRepository) BuiltIns() []catalog.BuiltIn {
	_ = "STUB: not implemented"
	return nil
}

type workloadAttestorV1 struct{}

func (workloadAttestorV1) New() catalog.Facade {
	_ = "STUB: not implemented"
	return *new(catalog.Facade)
}
func (workloadAttestorV1) Deprecated() bool { _ = "STUB: not implemented"; return false }
