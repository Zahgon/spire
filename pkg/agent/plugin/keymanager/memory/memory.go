package memory

import (
	keymanagerbase "github.com/spiffe/spire/pkg/agent/plugin/keymanager/base"
	"github.com/spiffe/spire/pkg/common/catalog"
)

type Generator = keymanagerbase.Generator

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func TestBuiltIn(generator Generator) catalog.BuiltIn {
	_ = "STUB: not implemented"
	return *new(catalog.BuiltIn)
}

func asBuiltIn(p *KeyManager) catalog.BuiltIn {
	_ = "STUB: not implemented"
	return *new(catalog.BuiltIn)
}

type KeyManager struct {
	*keymanagerbase.Base
}

func newKeyManager(generator Generator) *KeyManager { _ = "STUB: not implemented"; return nil }
