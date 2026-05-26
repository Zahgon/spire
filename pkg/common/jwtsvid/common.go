package jwtsvid

import (
	"time"
)

func GetTokenExpiry(token string) (time.Time, time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), *new(time.Time), nil
}
