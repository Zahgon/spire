package catalog

import (
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/server/plugin/notifier"
)

type notifierRepository struct {
	notifier.Repository
}

func (repo *notifierRepository) Binder() any { _ = "STUB: not implemented"; return *new(any) }

func (repo *notifierRepository) Constraints() catalog.Constraints {
	_ = "STUB: not implemented"
	return *new(catalog.Constraints)
}

func (repo *notifierRepository) Versions() []catalog.Version { _ = "STUB: not implemented"; return nil }

func (repo *notifierRepository) BuiltIns() []catalog.BuiltIn { _ = "STUB: not implemented"; return nil }

type notifierV1 struct{}

func (notifierV1) New() catalog.Facade { _ = "STUB: not implemented"; return *new(catalog.Facade) }
func (notifierV1) Deprecated() bool    { _ = "STUB: not implemented"; return false }
