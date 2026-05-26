package cli

import (
	"time"
)

// CommaStringsFlag facilitates parsing flags representing a comma separated list of strings
type CommaStringsFlag []string

func (f CommaStringsFlag) String() string { _ = "STUB: not implemented"; return "" }

func (f *CommaStringsFlag) Set(v string) error { _ = "STUB: not implemented"; return nil }

// DurationFlag facilitates parsing flags representing a time.Duration
type DurationFlag time.Duration

func (f DurationFlag) String() string { _ = "STUB: not implemented"; return "" }

func (f *DurationFlag) Set(v string) error { _ = "STUB: not implemented"; return nil }

// StringsFlag facilitates setting multiple flags
type StringsFlag []string

func (s *StringsFlag) String() string { _ = "STUB: not implemented"; return "" }

func (s *StringsFlag) Set(val string) error { _ = "STUB: not implemented"; return nil }

// BoolFlag is used to define 3 possible states: true, false, or all.
// Take care that false=1, and true=2
type BoolFlag int

const BoolFlagAll = 0
const BoolFlagFalse = 1
const BoolFlagTrue = 2

func (b *BoolFlag) String() string { _ = "STUB: not implemented"; return "" }

func (b *BoolFlag) Set(val string) error { _ = "STUB: not implemented"; return nil }

// if the value received isn't true or false, it will set the default value
