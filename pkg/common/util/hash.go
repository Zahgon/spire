package util

// GetSHA256Digest calculates the sha256 digest of a file specified by path. If the size of the file exceeds the provided
// limit, the hash will not be calculated and an error will be returned instead.
func GetSHA256Digest(path string, limit int64) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
