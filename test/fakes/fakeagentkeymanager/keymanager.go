package fakeagentkeymanager

import (
	"testing"

	"github.com/spiffe/spire/pkg/agent/plugin/keymanager"
)

// New returns a fake key manager
func New(t *testing.T, dir string) keymanager.KeyManager {
	_ = "STUB: not implemented"
	return *new(keymanager.KeyManager)
}
