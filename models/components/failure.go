package components

// Failure - Why the attempt failed: the URL is no longer valid, the host did not resolve, it resolved to an address that is not public, no answer within 5 seconds, the connection failed, a 3xx (redirects are not followed) or another non-2xx status. Absent for a delivered attempt.
type Failure string

const (
	FailureInvalidURL       Failure = "invalid_url"
	FailureResolveFailed    Failure = "resolve_failed"
	FailureAddressRefused   Failure = "address_refused"
	FailureTimeout          Failure = "timeout"
	FailureConnectionFailed Failure = "connection_failed"
	FailureRedirect         Failure = "redirect"
	FailureNon2xx           Failure = "non_2xx"
)

func (e Failure) ToPointer() *Failure {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Failure) IsExact() bool {
	if e != nil {
		switch *e {
		case "invalid_url", "resolve_failed", "address_refused", "timeout", "connection_failed", "redirect", "non_2xx":
			return true
		}
	}
	return false
}
