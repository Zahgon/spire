package svidstore

import (
	svidstorev1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/agent/svidstore/v1"
)

type Data struct {
	// SPIFFEID is the SPIFFE ID of the SVID
	SPIFFEID string `json:"spiffeID,omitempty"`
	// X509SVID is the PEM encoded certificate chain. MAY include intermediates,
	// the leaf certificate (or SVID itself) MUST come first
	X509SVID string `json:"x509SVID,omitempty"`
	// X509SVIDKey is the PEM encoded PKCS#8 private key.
	X509SVIDKey string `json:"x509SVIDKey,omitempty"`
	// Bundle is the PEM encoded X.509 bundle for the trust domain
	Bundle string `json:"bundle,omitempty"`
	// FederatedBundles is the CA certificate bundles belonging to foreign trust domains that the workload should trust,
	// keyed by trust domain. Bundles are in encoded in PEM format.
	FederatedBundles map[string]string `json:"federatedBundles,omitempty"`
}

func SecretFromProto(req *svidstorev1.PutX509SVIDRequest) (*Data, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseMetadata parses metadata from a slice of strings
// into a map that can be consumed by SVIDStore plugins
func ParseMetadata(metaData []string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func rawKeyToPem(rawKey []byte) (string, error) { _ = "STUB: not implemented"; return "", nil }

func rawCertToPem(rawCerts [][]byte) (string, error) { _ = "STUB: not implemented"; return "", nil }
