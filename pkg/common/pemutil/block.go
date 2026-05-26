package pemutil

import (
	"errors"
)

var (
	ErrNoBlocks = errors.New("no PEM blocks")
)

type Block struct {
	Type    string
	Headers map[string]string
	Object  any
}

func LoadBlocks(path string) ([]Block, error) { _ = "STUB: not implemented"; return nil, nil }

func loadBlock(path string, expectedTypes ...string) (*Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadBlocks(path string, expectedCount int, expectedTypes ...string) (blocks []Block, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseBlock(pemBytes []byte, expectedTypes ...string) (*Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseBlocks(pemBytes []byte, expectedCount int, expectedTypes ...string) (blocks []Block, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
