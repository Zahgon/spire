package x509util

import (
	"math/big"
)

var (
	maxUint128 = getMaxUint128()
	one        = big.NewInt(1)
)

// NewSerialNumber creates a random certificate serial number according to CA/Browser forum spec
// Section 7.1:
// "Effective September 30, 2016, CAs SHALL generate non-sequential Certificate serial numbers greater than
// zero (0) containing at least 64 bits of output from a CSPRNG"
func NewSerialNumber() (*big.Int, error) {
	_ = "STUB: not implemented"
	// Creates random integer in range [0,MaxUint128)
	return nil, nil
}

// Adds 1 to return serial number [1,MaxUint128]

func getMaxUint128() *big.Int { _ = "STUB: not implemented"; return nil }

// (2^128 − 1)
