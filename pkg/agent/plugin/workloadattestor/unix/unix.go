package unix

import "github.com/spiffe/spire/pkg/common/catalog"

const (
	pluginName = "unix"
)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }
