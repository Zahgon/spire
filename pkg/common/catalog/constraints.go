package catalog

func ExactlyOne() Constraints { _ = "STUB: not implemented"; return *new(Constraints) }

func MaybeOne() Constraints { _ = "STUB: not implemented"; return *new(Constraints) }

func AtLeastOne() Constraints { _ = "STUB: not implemented"; return *new(Constraints) }

func ZeroOrMore() Constraints { _ = "STUB: not implemented"; return *new(Constraints) }

type Constraints struct {
	// Min is the minimum number of plugins required of a specific type. If
	// zero, there is no lower bound (i.e. the plugin type is optional).
	Min int

	// Max is the maximum number of plugins required of a specific type. If
	// zero, there is no upper bound.
	Max int
}

func (c Constraints) Check(count int) error { _ = "STUB: not implemented"; return nil }
