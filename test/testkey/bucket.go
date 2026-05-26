package testkey

import (
	"crypto"
	"sync"
)

var (
	packageDir string
)

func init() {
	packageDir = initPackageDir()
}

func initPackageDir() string { _ = "STUB: not implemented"; return "" }

type keyType[K crypto.Signer] interface {
	Path() string
	GenerateKey() (K, error)
}

type bucket[KT keyType[K], K crypto.Signer] struct {
	kt KT

	mtx  sync.Mutex
	keys []K
}

func (b *bucket[KT, K]) At(n int) (key K, err error) {
	_ = "STUB: not implemented"
	return *new(K), nil
}

func (b *bucket[KT, K]) load() (err error) { _ = "STUB: not implemented"; return nil }

func (b *bucket[KT, K]) save() error { _ = "STUB: not implemented"; return nil }

func (b *bucket[KT, K]) path() string { _ = "STUB: not implemented"; return "" }
