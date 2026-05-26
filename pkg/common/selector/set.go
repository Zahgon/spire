package selector

import (
	"github.com/spiffe/spire/proto/spire/common"
)

type Set interface {
	Raw() []*common.Selector
	Array() []*Selector
	Equal(otherSet Set) bool
	Includes(selector *Selector) bool
	IncludesSet(s2 Set) bool
	Add(selector *Selector)
	Remove(selector *Selector) *Selector
	String() string
	Size() int
}

type set map[Selector]*Selector

func NewSet(selectors ...*Selector) Set { _ = "STUB: not implemented"; return *new(Set) }

func NewSetFromRaw(c []*common.Selector) Set { _ = "STUB: not implemented"; return *new(Set) }

func (s *set) Raw() []*common.Selector { _ = "STUB: not implemented"; return nil }

// Array returns an array with the elements of the set in any order.
func (s *set) Array() []*Selector { _ = "STUB: not implemented"; return nil }

func (s *set) Equal(otherSet Set) bool { _ = "STUB: not implemented"; return false }

func (s *set) Includes(selector *Selector) bool { _ = "STUB: not implemented"; return false }

func (s *set) IncludesSet(s2 Set) bool { _ = "STUB: not implemented"; return false }

func (s *set) Add(selector *Selector) { _ = "STUB: not implemented"; return }

func (s *set) Remove(selector *Selector) *Selector { _ = "STUB: not implemented"; return nil }

func (s *set) Size() int { _ = "STUB: not implemented"; return 0 }

func (s *set) String() string { _ = "STUB: not implemented"; return "" }
