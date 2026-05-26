package awsiid

import (
	"crypto/x509"
	"sync"
)

// PublicKeyType is the type of public key used to verify the AWS signature.
type PublicKeyType int

const (
	KeyTypeUnset PublicKeyType = iota
	RSA1024
	RSA2048
)

var certCache sync.Map

func getAWSCACertificate(region string, keyType PublicKeyType) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fall back to the default cert
