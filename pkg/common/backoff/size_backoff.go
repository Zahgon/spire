package backoff

// SizeLimitedBackOff defines interface for implementing a size based backoff for requests which
// contain number of records to be processed by server.
type SizeLimitedBackOff interface {
	// NextBackOff returns the duration to wait before retrying the operation,
	// or backoff.
	NextBackOff() int

	// Success indicates the backoff implementation that previous request succeeded
	// so that it can adjust backoff accordingly for next request.
	Success()

	// Failure indicates the backoff implementation that previous request failed
	// so that it can adjust backoff accordingly for next request.
	Failure()

	// Reset to initial state.
	Reset()
}

type sizeLimitedBackOff struct {
	currentSize int
	maxSize     int
}

var _ SizeLimitedBackOff = (*sizeLimitedBackOff)(nil)

func (r *sizeLimitedBackOff) NextBackOff() int { _ = "STUB: not implemented"; return 0 }

func (r *sizeLimitedBackOff) Success() { _ = "STUB: not implemented"; return }

func (r *sizeLimitedBackOff) Failure() { _ = "STUB: not implemented"; return }

func (r *sizeLimitedBackOff) Reset() { _ = "STUB: not implemented"; return }

// NewSizeLimitedBackOff returns a new SizeLimitedBackOff with provided maxRequestSize and lowest request size of 1.
// On Failure the size gets reduced by half and on Success size gets doubled
func NewSizeLimitedBackOff(maxRequestSize int) SizeLimitedBackOff {
	_ = "STUB: not implemented"
	return *new(SizeLimitedBackOff)
}
