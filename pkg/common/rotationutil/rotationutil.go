package rotationutil

import (
	"crypto/x509"
	"time"

	"github.com/spiffe/spire/pkg/agent/client"
)

const (
	gracePeriodThreshold = 12 * time.Hour
)

type RotationStrategy struct {
	x509AvailabilityTarget time.Duration
}

func NewRotationStrategy(x509AvailabilityTarget time.Duration) *RotationStrategy {
	_ = "STUB: not implemented"
	return nil
}

// ShouldFallbackX509DefaultRotation returns true if the availability target is configured but the value is not enough against the SVID lifetime.
func (rs *RotationStrategy) ShouldFallbackX509DefaultRotation(lifetime time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}

// x509AvailabilityTarget is not configured

// ShouldRotateX509 determines if a given SVID should be rotated, based
// on presented current time, and the certificate's expiration.
func (rs *RotationStrategy) ShouldRotateX509(now time.Time, cert *x509.Certificate) bool {
	_ = "STUB: not implemented"
	return false
}

// X509Expired returns true if the given X509 cert has expired
func X509Expired(now time.Time, cert *x509.Certificate) bool {
	_ = "STUB: not implemented"
	return false
}

// JWTSVIDExpiresSoon determines if the given JWT SVID should be rotated
// based on presented current time, the JWT's expiration.
// Also returns true if the JWT is already expired.
func (rs *RotationStrategy) JWTSVIDExpiresSoon(svid *client.JWTSVID, now time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

// if the SVID has less than half of its lifetime left or reaches the availability target,
// consider it as expiring soon

// JWTSVIDExpired returns true if the given SVID is expired.
func JWTSVIDExpired(svid *client.JWTSVID, now time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

func shouldRotateX509(now, beginTime, expiryTime time.Time, availabilityTarget time.Duration) bool {
	_ = "STUB: not implemented"
	return false

	// return true quickly if the expiry is already met.
}

// fall back the default rotation strategy.

func shouldRotateJWT(now, beginTime, expiryTime time.Time) bool {
	_ = "STUB: not implemented"
	return false

	// return true quickly if the expiry is already met.
}

// jitterHalfLifeDelta is a calculated delta centered to the half-life of the SVID.
// It's to spread out the renewal of SVID rotations to avoid spiky renewal requests.
func jitterHalfLifeDelta(halfLife time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func halfLife(lifetime time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *

	// calculateJitteredHalfLife calculates jitter of the half-life of the SVID.
	// The jitter is calculated as ± 10% of the half-life of the SVID.
	new(time.Duration)
}

func calculateJitteredHalfLife(lifetime time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

//nolint // gosec: no need for cryptographic randomness here

// calculateJitteredAvailabilityTarget calculates jitter of the availability target.
// The jitter is calculated as 0 ~ +10min of the availability target.
func calculateJitteredAvailabilityTarget(availabilityTarget time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

//nolint // gosec: no need for cryptographic randomness here

func shouldRotateByAvailabilityTarget(ttl, lifetime, availabilityTarget time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}

func shouldRotateByHalf(ttl, lifetime time.Duration) bool {
	_ = "STUB: not implemented"
	// calculate a jitter delta to spread out rotations
	return false
}

func shouldFallbackX509Default(lifetime, availabilityTarget time.Duration) bool {
	_ = "STUB: not implemented"
	// if the grace period less than the threshold, it should be felt back to the default rotation strategy
	return false
}
