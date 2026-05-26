package fakeserverkeymanager

import (
	"testing"

	"github.com/spiffe/spire/pkg/server/plugin/keymanager"
	keymanagerbase "github.com/spiffe/spire/pkg/server/plugin/keymanager/base"
)

func New(t *testing.T) keymanager.KeyManager {
	_ = "STUB: not implemented"
	return *new(keymanager.KeyManager)
}

type keyManager struct {
	*keymanagerbase.Base
}
