package testkey

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"sync"
	"testing"
)

var (
	keys          Keys
	rsa2048Bucket bucket[rsa2048, *rsa.PrivateKey]
	rsa4096Bucket bucket[rsa4096, *rsa.PrivateKey]
	ec256Bucket   bucket[ec256, *ecdsa.PrivateKey]
	ec384Bucket   bucket[ec384, *ecdsa.PrivateKey]
)

func NewRSA2048(tb testing.TB) *rsa.PrivateKey { _ = "STUB: not implemented"; return nil }

func NewRSA2048PKCS1PEM(tb testing.TB) []byte { _ = "STUB: not implemented"; return nil }

func NewRSA2048PKCS8PEM(tb testing.TB) []byte { _ = "STUB: not implemented"; return nil }

func MustRSA2048() *rsa.PrivateKey { _ = "STUB: not implemented"; return nil }

func MustRSA2048PKCS1PEM() []byte { _ = "STUB: not implemented"; return nil }

func MustRSA2048PKCS8PEM() []byte { _ = "STUB: not implemented"; return nil }

func NewRSA4096(tb testing.TB) *rsa.PrivateKey { _ = "STUB: not implemented"; return nil }

func NewRSA4096PKCS1PEM(tb testing.TB) []byte { _ = "STUB: not implemented"; return nil }

func NewRSA4096PKCS8PEM(tb testing.TB) []byte { _ = "STUB: not implemented"; return nil }

func MustRSA4096() *rsa.PrivateKey { _ = "STUB: not implemented"; return nil }

func MustRSA4096PKCS1PEM() []byte { _ = "STUB: not implemented"; return nil }

func MustRSA4096PKCS8PEM() []byte { _ = "STUB: not implemented"; return nil }

func NewEC256(tb testing.TB) *ecdsa.PrivateKey { _ = "STUB: not implemented"; return nil }

func NewEC256PKCS1PEM(tb testing.TB) []byte { _ = "STUB: not implemented"; return nil }

func NewEC256PKCS8PEM(tb testing.TB) []byte { _ = "STUB: not implemented"; return nil }

func MustEC256() *ecdsa.PrivateKey { _ = "STUB: not implemented"; return nil }

func MustEC256PKCS1PEM() []byte { _ = "STUB: not implemented"; return nil }

func MustEC256PKCS8PEM() []byte { _ = "STUB: not implemented"; return nil }

func NewEC384(tb testing.TB) *ecdsa.PrivateKey { _ = "STUB: not implemented"; return nil }

func NewEC384PKCS1PEM(tb testing.TB) []byte { _ = "STUB: not implemented"; return nil }

func NewEC384PKCS8PEM(tb testing.TB) []byte { _ = "STUB: not implemented"; return nil }

func MustEC384() *ecdsa.PrivateKey { _ = "STUB: not implemented"; return nil }

func MustEC384PKCS1PEM() []byte { _ = "STUB: not implemented"; return nil }

func MustEC384PKCS8PEM() []byte { _ = "STUB: not implemented"; return nil }

type Keys struct {
	mtx        sync.Mutex
	rsa2048Idx int
	rsa4096Idx int
	ec256Idx   int
	ec384Idx   int
}

func (ks *Keys) NewRSA2048(tb testing.TB) *rsa.PrivateKey { _ = "STUB: not implemented"; return nil }

func (ks *Keys) MustRSA2048() *rsa.PrivateKey { _ = "STUB: not implemented"; return nil }

func (ks *Keys) NextRSA2048() (*rsa.PrivateKey, error) { _ = "STUB: not implemented"; return nil, nil }

func (ks *Keys) NewRSA4096(tb testing.TB) *rsa.PrivateKey { _ = "STUB: not implemented"; return nil }

func (ks *Keys) MustRSA4096() *rsa.PrivateKey { _ = "STUB: not implemented"; return nil }

func (ks *Keys) NextRSA4096() (*rsa.PrivateKey, error) { _ = "STUB: not implemented"; return nil, nil }

func (ks *Keys) NewEC256(tb testing.TB) *ecdsa.PrivateKey { _ = "STUB: not implemented"; return nil }

func (ks *Keys) MustEC256() *ecdsa.PrivateKey { _ = "STUB: not implemented"; return nil }

func (ks *Keys) NextEC256() (*ecdsa.PrivateKey, error) { _ = "STUB: not implemented"; return nil, nil }

func (ks *Keys) NewEC384(tb testing.TB) *ecdsa.PrivateKey { _ = "STUB: not implemented"; return nil }

func (ks *Keys) MustEC384() *ecdsa.PrivateKey { _ = "STUB: not implemented"; return nil }

func (ks *Keys) NextEC384() (*ecdsa.PrivateKey, error) { _ = "STUB: not implemented"; return nil, nil }

type rsa2048 struct{}

func (rsa2048) Path() string { _ = "STUB: not implemented"; return "" }

func (rsa2048) GenerateKey() (*rsa.PrivateKey, error) { _ = "STUB: not implemented"; return nil, nil }

type rsa4096 struct{}

func (rsa4096) Path() string { _ = "STUB: not implemented"; return "" }

func (rsa4096) GenerateKey() (*rsa.PrivateKey, error) { _ = "STUB: not implemented"; return nil, nil }

type ec256 struct{}

func (ec256) Path() string { _ = "STUB: not implemented"; return "" }

func (ec256) GenerateKey() (*ecdsa.PrivateKey, error) { _ = "STUB: not implemented"; return nil, nil }

type ec384 struct{}

func (ec384) Path() string { _ = "STUB: not implemented"; return "" }

func (ec384) GenerateKey() (*ecdsa.PrivateKey, error) { _ = "STUB: not implemented"; return nil, nil }

func check(err error) { _ = "STUB: not implemented"; return }
