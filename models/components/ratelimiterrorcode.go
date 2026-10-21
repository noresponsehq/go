package components

// RateLimitErrorCode - Machine codes an HTTP 429 answer can carry.
type RateLimitErrorCode string

const (
	RateLimitErrorCodeRateLimited RateLimitErrorCode = "rate_limited"
)

func (e RateLimitErrorCode) ToPointer() *RateLimitErrorCode {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *RateLimitErrorCode) IsExact() bool {
	if e != nil {
		switch *e {
		case "rate_limited":
			return true
		}
	}
	return false
}
