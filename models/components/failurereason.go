package components

// FailureReason - Why a signal or source did not produce a definite found/not_found. timed_out means the source missed the request deadline; a continuation may still settle it. deferred means the (source, identifier) pair is async-only and was not attempted on the synchronous path; a continuation settles it.
type FailureReason string

const (
	FailureReasonTimedOut             FailureReason = "timed_out"
	FailureReasonRateLimited          FailureReason = "rate_limited"
	FailureReasonBlocked              FailureReason = "blocked"
	FailureReasonSourceError          FailureReason = "source_error"
	FailureReasonUnavailable          FailureReason = "unavailable"
	FailureReasonDeferred             FailureReason = "deferred"
	FailureReasonInsufficientEvidence FailureReason = "insufficient_evidence"
)

func (e FailureReason) ToPointer() *FailureReason {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *FailureReason) IsExact() bool {
	if e != nil {
		switch *e {
		case "timed_out", "rate_limited", "blocked", "source_error", "unavailable", "deferred", "insufficient_evidence":
			return true
		}
	}
	return false
}
