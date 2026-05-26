package catalog

import (
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/server/plugin/credentialcomposer"
)

type credentialComposerRepository struct {
	credentialcomposer.Repository
}

func (repo *credentialComposerRepository) Binder() any { _ = "STUB: not implemented"; return *new(any) }

func (repo *credentialComposerRepository) Constraints() catalog.Constraints {
	_ = "STUB: not implemented"
	return *new(catalog.Constraints)
}

func (repo *credentialComposerRepository) Versions() []catalog.Version {
	_ = "STUB: not implemented"
	return nil
}

func (repo *credentialComposerRepository) BuiltIns() []catalog.BuiltIn {
	_ = "STUB: not implemented"
	return nil
}

type credentialComposerV1 struct{}

func (credentialComposerV1) New() catalog.Facade {
	_ = "STUB: not implemented"
	return *new(catalog.Facade)
}
func (credentialComposerV1) Deprecated() bool { _ = "STUB: not implemented"; return false }
