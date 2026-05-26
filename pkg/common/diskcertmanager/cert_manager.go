package diskcertmanager

import (
	"context"
	"crypto/tls"
	"os"
	"sync"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/sirupsen/logrus"
)

// DiskCertManager is a certificate manager that loads certificates from disk, and watches for changes.
type DiskCertManager struct {
	certFilePath     string
	keyFilePath      string
	certLastModified time.Time
	keyLastModified  time.Time
	fileSyncInterval time.Duration
	certMtx          sync.RWMutex
	cert             *tls.Certificate
	clk              clock.Clock
	log              logrus.FieldLogger
}

type Config struct {
	CertFilePath     string
	KeyFilePath      string
	FileSyncInterval time.Duration
}

func New(config *Config, clk clock.Clock, log logrus.FieldLogger) (*DiskCertManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TLSConfig returns a TLS configuration that uses the provided certificate stored on disk.
func (m *DiskCertManager) GetTLSConfig() *tls.Config { _ = "STUB: not implemented"; return nil }

// enable HTTP/2

// getCertificate is called by the TLS stack when a new TLS connection is established.
func (m *DiskCertManager) getCertificate(_ *tls.ClientHelloInfo) (*tls.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WatchFileChanges starts a file watcher to watch for changes to the cert and key files.
func (m *DiskCertManager) WatchFileChanges(ctx context.Context) { _ = "STUB: not implemented"; return }

// syncCertificateFiles checks if the cert and key files have been modified, and reloads the certificate if necessary.
func (m *DiskCertManager) syncCertificateFiles() { _ = "STUB: not implemented"; return }

// loadCert read the certificate and key files, and load the x509 certificate to memory.
func (m *DiskCertManager) loadCert() error { _ = "STUB: not implemented"; return nil }

// getFilesInfo returns the file info of the cert and key files, or error if the files are unreadable or do not exist.
func (m *DiskCertManager) getFilesInfo() (os.FileInfo, os.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(os.FileInfo), *new(os.FileInfo), nil
}

// getFileInfo returns the file info of the given path, or error if the file is unreadable or does not exist.
func (m *DiskCertManager) getFileInfo(path string) (os.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(os.FileInfo), nil
}
