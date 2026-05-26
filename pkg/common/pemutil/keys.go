package pemutil

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
)

func ParsePublicKey(pemBytes []byte) (crypto.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey), nil
}

func LoadPublicKey(path string) (crypto.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey), nil
}

func ParsePrivateKey(pemBytes []byte) (crypto.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PrivateKey), nil
}

func LoadPrivateKey(path string) (crypto.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PrivateKey), nil
}

func ParseSigner(pemBytes []byte) (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}

func LoadSigner(path string) (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}

func ParseRSAPrivateKey(pemBytes []byte) (*rsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LoadRSAPrivateKey(path string) (*rsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncodeRSAPrivateKey(privateKey *rsa.PrivateKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func rsaPrivateKeyFromObject(object any) (*rsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseECPrivateKey(pemBytes []byte) (*ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LoadECPrivateKey(path string) (*ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncodeECPrivateKey(privateKey *ecdsa.PrivateKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncodePKCS8PrivateKey(privateKey any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ecdsaPrivateKeyFromObject(object any) (*ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func signerFromPrivateKey(privateKey crypto.PrivateKey) (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}
