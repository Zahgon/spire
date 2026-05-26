package util

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func CheckedCast[To, From Int](v From) (To, error) {
	_ = "STUB: not implemented"

	// Check sign is unchanged. This is violated e.g. by int8(-3) -> uint8.
	// Check converting back gives original value. This is violated e.g. by uint16(300) -> uint8.
	return *new(To), nil
}

// If we got here, then the value can correctly be represented as the 'To' type: success.

func MustCast[To, From Int](v From) To { _ = "STUB: not implemented"; return *new(To) }
