package storage

import (
	"crypto/x509"
	"errors"
	"sync"
	"time"
)

var (
	ErrNotCached = errors.New("not cached")
)

type Storage interface {
	// LoadSVID loads the SVID from storage. Returns ErrNotCached if the SVID
	// does not exist in the cache.
	LoadSVID() ([]*x509.Certificate, bool, error)

	// StoreSVID stores the SVID.
	StoreSVID(certs []*x509.Certificate, reattestable bool) error

	// DeleteSVID deletes the SVID.
	DeleteSVID() error

	// LoadBundle loads the bundle from storage. Returns ErrNotCached if the
	// bundle does not exist in the cache.
	LoadBundle() ([]*x509.Certificate, error)

	// StoreBundle stores the bundle.
	StoreBundle(certs []*x509.Certificate) error

	// LoadBootstrapState returns the Bootstrap state items
	LoadBootstrapState() (use int, start_time time.Time, connectionAttempts int, err error)

	// StoreBootstrapState stores the use and start_time bootstrap states for future use
	StoreBootstrapState(use int, start_time time.Time, connectionAttempts int) error

	// DeleteBootstrapState removes the bootstrap state
	DeleteBootstrapState() error
}

func Open(dir string) (Storage, error) { _ = "STUB: not implemented"; return *new(Storage), nil }

type storage struct {
	dir string

	mtx  sync.RWMutex
	data storageData
}

func (s *storage) LoadBundle() ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *storage) StoreBundle(bundle []*x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *storage) LoadSVID() ([]*x509.Certificate, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (s *storage) StoreSVID(svid []*x509.Certificate, reattestable bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *storage) DeleteSVID() error { _ = "STUB: not implemented"; return nil }

func (s *storage) LoadBootstrapState() (use int, start_time time.Time, connectionAttempts int, err error) {
	_ = "STUB: not implemented"
	return 0, *new(time.Time), 0, nil
}

func (s *storage) StoreBootstrapState(use int, start_time time.Time, connectionAttempts int) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *storage) DeleteBootstrapState() error { _ = "STUB: not implemented"; return nil }

type storageJSON struct {
	SVID               [][]byte  `json:"svid"`
	Bundle             [][]byte  `json:"bundle"`
	Reattestable       bool      `json:"reattestable"`
	BootstrapUse       int       `json:"bootstrap_use"`
	BootstrapStartTime time.Time `json:"bootstrap_start_time"`
	ConnectionAttempts int       `json:"connection_attempts"`
}

type storageData struct {
	SVID               []*x509.Certificate
	Bundle             []*x509.Certificate
	Reattestable       bool
	BootstrapUse       int
	BootstrapStartTime time.Time
	ConnectionAttempts int
}

func (d storageData) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *storageData) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func storeData(dir string, data storageData) error { _ = "STUB: not implemented"; return nil }

func loadData(dir string) (storageData, error) {
	_ = "STUB: not implemented"
	return *new(storageData), nil
}

func parseCertificates(certsPEM [][]byte) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeCertificates(certs []*x509.Certificate) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dataPath(dir string) string { _ = "STUB: not implemented"; return "" }
