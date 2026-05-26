package cgroup

import (
	"regexp"
)

const (
	// A token to match an entire path component in a "/" delimited path
	wildcardToken = "*"
	// A regex expression that expresses wildcardToken
	regexpWildcard = "[^\\/]*"

	// A token to match, and extract as a container ID, an entire path component in a
	// "/" delimited path
	containerIDToken = "<id>"
	// A regex expression that expresses containerIDToken
	regexpContainerID = "([^\\/]*)"
	// index for slice returned by FindStringSubmatch
	submatchIndex = 1
)

// ContainerIDFinder finds a container id from a cgroup entry.
type ContainerIDFinder interface {
	// FindContainerID returns a container id and true if the known pattern is matched, false otherwise.
	FindContainerID(cgroup string) (containerID string, found bool)
}

func newContainerIDFinder(pattern string) (ContainerIDFinder, error) {
	_ = "STUB: not implemented"
	return *new(ContainerIDFinder), nil
}

// NewContainerIDFinder returns a new ContainerIDFinder.
//
// The patterns provided should use the Tokens defined in this package in order
// to describe how a container id should be extracted from a cgroup entry. The
// given patterns MUST NOT be ambiguous and an error will be returned if multiple
// patterns can match the same input. An example of invalid input:
//
// "/a/b/<id>"
// "/*/b/<id>"
//
// Examples:
//
// "/docker/<id>"
// "/my.slice/*/<id>/*"
//
// Note: The pattern provided is *not* a regular expression. It is a simplified matching
// language that enforces a forward slash-delimited schema.
func NewContainerIDFinder(patterns []string) (ContainerIDFinder, error) {
	_ = "STUB: not implemented"
	return *new(ContainerIDFinder), nil
}

type containerIDFinder struct {
	re *regexp.Regexp
}

func (f *containerIDFinder) FindContainerID(cgroup string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

type containerIDFinders struct {
	finders []ContainerIDFinder
}

func (f *containerIDFinders) FindContainerID(cgroup string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// There must be exactly 0 or 1 pattern that matches a given input. Enforcing
// this at startup, instead of at runtime (e.g. in `FindContainerID`) ensures that
// a bad configuration is found immediately during rollout, rather than once a
// specific cgroup input is encountered.
//
// Given the restricted grammar of wildcardToken and containerIDToken and
// the goal of protecting a user from invalid configuration, detecting ambiguous patterns
// is done as follows:
//
// 1. If the number of path components in two patterns differ, they cannot match identical inputs.
// This assertion follows from the path focused grammar and the fact that the regex
// wildcards (regexpWildcard and regexpContainerID) cannot match "/".
// 2. If the number of path components in two patterns are the same, we test "component
// equivalence" at each index. wildcardToken and containerIDToken are equivalent to
// any other, otherwise, the two components at an index are directly compared.
// From this and the fact the regex wildcards cannot match "/" follows that a single
// non-equivalent path component means the two patterns cannot match the same inputs.
func findAmbiguousPatterns(patterns []string) []string { _ = "STUB: not implemented"; return nil }

// generate all combinations except for equivalent
// index combinations which will always match.

func equivalentPatterns(a, b string) bool { _ = "STUB: not implemented"; return false }
