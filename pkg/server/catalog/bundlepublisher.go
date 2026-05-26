package catalog

import (
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/server/plugin/bundlepublisher"
)

type bundlePublisherRepository struct {
	bundlepublisher.Repository
}

func (repo *bundlePublisherRepository) Binder() any { _ = "STUB: not implemented"; return *new(any) }

func (repo *bundlePublisherRepository) Constraints() catalog.Constraints {
	_ = "STUB: not implemented"
	return *new(catalog.Constraints)
}

func (repo *bundlePublisherRepository) Versions() []catalog.Version {
	_ = "STUB: not implemented"
	return nil
}

func (repo *bundlePublisherRepository) BuiltIns() []catalog.BuiltIn {
	_ = "STUB: not implemented"
	return nil
}

type bundlePublisherV1 struct{}

func (bundlePublisherV1) New() catalog.Facade {
	_ = "STUB: not implemented"
	return *new(catalog.Facade)
}
func (bundlePublisherV1) Deprecated() bool { _ = "STUB: not implemented"; return false }
