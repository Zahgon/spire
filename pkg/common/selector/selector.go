// The selector package exports functions useful for manipulating and generating
// spire selectors
package selector

import (
	"github.com/spiffe/spire/proto/spire/common"
)

// Type and Value are delimited by a colon (:)
// e.g. "unix:uid:1000"
const Delimiter = ":"

type Selector struct {
	Type  string
	Value string
}

func New(c *common.Selector) *Selector { _ = "STUB: not implemented"; return nil }

func (s *Selector) Raw() *common.Selector { _ = "STUB: not implemented"; return nil }

func Validate(s *common.Selector) error {
	_ = "STUB: not implemented"
	// Validate that the Type does not contain a colon (:) to prevent accidental misconfigurations
	// e.g. type="unix:user" value="root" is the invalid selector
	// and type="unix" value"user:root" is the valid selector
	return nil
}
