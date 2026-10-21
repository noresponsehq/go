package components

// IncompleteReason - Why a lane did not settle. deferred means the lane is asynchronous and still running; a continuation settles it. timed_out means the lane missed its deadline. insufficient_evidence is a settled inconclusive signal rather than an incomplete lane.
type IncompleteReason string

const (
	IncompleteReasonTimedOut    IncompleteReason = "timed_out"
	IncompleteReasonRateLimited IncompleteReason = "rate_limited"
	IncompleteReasonBlocked     IncompleteReason = "blocked"
	IncompleteReasonSourceError IncompleteReason = "source_error"
	IncompleteReasonUnavailable IncompleteReason = "unavailable"
	IncompleteReasonDeferred    IncompleteReason = "deferred"
)

func (e IncompleteReason) ToPointer() *IncompleteReason {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *IncompleteReason) IsExact() bool {
	if e != nil {
		switch *e {
		case "timed_out", "rate_limited", "blocked", "source_error", "unavailable", "deferred":
			return true
		}
	}
	return false
}
