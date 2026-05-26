package tpmutil

import (
	"regexp"
)

var validTPMNames = []*regexp.Regexp{
	regexp.MustCompile(`tpmrm\d+$`),
	regexp.MustCompile(`tpm\d+$`),
}

func AutoDetectTPMPath(baseTPMDir string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Do not return yet, we need to make sure that
// there is only one TPM device.
