package main

import (
	"crypto"
	"crypto/x509"
	"crypto/x509/pkix"
	"flag"
	"fmt"
	"math/big"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

type stringArrayFlag []string

func (s *stringArrayFlag) String() string { _ = "STUB: not implemented"; return "" }

func (s *stringArrayFlag) Set(value string) error { _ = "STUB: not implemented"; return nil }

func main() {
	var trustDomain string
	var x509popSans stringArrayFlag
	flag.StringVar(&trustDomain, "trust-domain", "", "Name of the trust domains the certs will be used for")
	flag.Var(&x509popSans, "x509pop-san", "Uri san to set using x509pop:// scheme")

	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Fprintln(os.Stderr, "usage: gencerts SERVERDIR AGENTDIR [AGENTDIR...]")
		os.Exit(1)
	}

	var x509popSanUris []*url.URL
	for _, x509popSan := range x509popSans {
		san, err := url.Parse("x509pop://" + trustDomain + "/" + x509popSan)
		checkErr(err)
		x509popSanUris = append(x509popSanUris, san)
	}

	notAfter := time.Now().Add(time.Hour)

	caKey := generateKey()
	caCert := createRootCertificate(caKey, &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		BasicConstraintsValid: true,
		IsCA:                  true,
		NotAfter:              notAfter,
		Subject:               pkix.Name{CommonName: "Agent CA"},
	})

	writeCerts(filepath.Join(flag.Arg(0), "agent-cacert.pem"), caCert)

	for i, dir := range flag.Args()[1:] {
		agentKey := generateKey()
		agentCert := createCertificate(agentKey, &x509.Certificate{
			SerialNumber: big.NewInt(int64(i)),
			KeyUsage:     x509.KeyUsageDigitalSignature,
			NotAfter:     notAfter,
			Subject:      pkix.Name{CommonName: filepath.Base(dir)},
			URIs:         x509popSanUris,
		}, caKey, caCert)

		writeKey(filepath.Join(dir, "agent.key.pem"), agentKey)
		writeCerts(filepath.Join(dir, "agent.crt.pem"), agentCert)
	}
}

func createRootCertificate(key crypto.Signer, tmpl *x509.Certificate) *x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func createCertificate(key crypto.Signer, tmpl *x509.Certificate, parentKey crypto.Signer, parent *x509.Certificate) *x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func generateKey() crypto.Signer { _ = "STUB: not implemented"; return *new(crypto.Signer) }

func writeKey(path string, key crypto.Signer) { _ = "STUB: not implemented"; return }

// This key is used only for testing purposes.

func writeCerts(path string, certs ...*x509.Certificate) { _ = "STUB: not implemented"; return }

func writeFile(path string, data []byte, mode os.FileMode) { _ = "STUB: not implemented"; return }

func checkErr(err error) { _ = "STUB: not implemented"; return }
