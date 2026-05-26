package testkey

import (
	"crypto"
)

type Generator struct{ keys Keys }

func (g *Generator) GenerateRSA2048Key() (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}
func (g *Generator) GenerateRSA4096Key() (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}
func (g *Generator) GenerateEC256Key() (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}
func (g *Generator) GenerateEC384Key() (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}
