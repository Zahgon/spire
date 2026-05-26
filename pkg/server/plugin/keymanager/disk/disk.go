package disk

import (
	"context"
	"sync"

	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
	keymanagerbase "github.com/spiffe/spire/pkg/server/plugin/keymanager/base"
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
	KeysPath string `hcl:"keys_path"`
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *configuration {
	_ = "STUB: not implemented"
	return nil
}

type KeyManager struct {
	*keymanagerbase.Base
	configv1.UnimplementedConfigServer

	mu     sync.Mutex
	config *configuration
}

func newKeyManager(generator Generator) *KeyManager { _ = "STUB: not implemented"; return nil }

func (m *KeyManager) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *KeyManager) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *KeyManager) configure(config *configuration) error {
	_ = "STUB: not implemented"
	// only load entry information on first configure
	return nil
}

func (m *KeyManager) writeEntries(_ context.Context, entries []*keymanagerbase.KeyEntry) error {
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
