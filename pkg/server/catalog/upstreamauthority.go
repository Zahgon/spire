package catalog

import (
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/server/plugin/upstreamauthority"
)

type upstreamAuthorityRepository struct {
	upstreamauthority.Repository
}

func (repo *upstreamAuthorityRepository) Binder() any { _ = "STUB: not implemented"; return *new(any) }

func (repo *upstreamAuthorityRepository) Constraints() catalog.Constraints {
	_ = "STUB: not implemented"
	return *new(catalog.Constraints)
}

func (repo *upstreamAuthorityRepository) Versions() []catalog.Version {
	_ = "STUB: not implemented"
	return nil
}

func (repo *upstreamAuthorityRepository) BuiltIns() []catalog.BuiltIn {
	_ = "STUB: not implemented"
	return nil
}

type upstreamAuthorityV1 struct{}

func (upstreamAuthorityV1) New() catalog.Facade {
	_ = "STUB: not implemented"
	return *new(catalog.Facade)
}
func (upstreamAuthorityV1) Deprecated() bool { _ = "STUB: not implemented"; return false }
