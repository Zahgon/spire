package disk

import (
	"context"
	"sync"

	"github.com/hashicorp/go-hclog"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
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

type configuration struct {
	Directory string `hcl:"directory"`
}

type KeyManager struct {
	*keymanagerbase.Base
	configv1.UnimplementedConfigServer

	log hclog.Logger

	mu     sync.Mutex
	config *configuration
}

func newKeyManager(generator Generator) *KeyManager { _ = "STUB: not implemented"; return nil }

func (m *KeyManager) SetLogger(log hclog.Logger) { _ = "STUB: not implemented"; return }

func (m *KeyManager) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *KeyManager) configure(config *configuration) error {
	_ = "STUB: not implemented"
	// Only load entry information on first configure
	return nil
}

func (m *KeyManager) verifyDirectory(dir string) error { _ = "STUB: not implemented"; return nil }

func (m *KeyManager) loadEntries(dir string) error {
	_ = "STUB: not implemented"
	// Load the entries from the keys file.
	return nil
}

func (m *KeyManager) writeEntries(_ context.Context, allEntries []*keymanagerbase.KeyEntry, _ *keymanagerbase.KeyEntry) error {
	_ = "STUB: not implemented"
	return nil
}

type entriesData struct {
	Keys map[string][]byte `json:"keys"`
}

func loadEntries(path string) ([]*keymanagerbase.KeyEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func writeEntries(path string, entries []*keymanagerbase.KeyEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func keysPath(dir string) string { _ = "STUB: not implemented"; return "" }
