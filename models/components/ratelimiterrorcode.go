package components

// RateLimitErrorCode - Machine codes an HTTP 429 answer can carry.
type RateLimitErrorCode string

const (
	RateLimitErrorCodeRateLimited               RateLimitErrorCode = "rate_limited"
	RateLimitErrorCodeAssessmentBudgetExhausted RateLimitErrorCode = "assessment_budget_exhausted"
	RateLimitErrorCodeExportLimitReached        RateLimitErrorCode = "export_limit_reached"
)

func (e RateLimitErrorCode) ToPointer() *RateLimitErrorCode {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *RateLimitErrorCode) IsExact() bool {
	if e != nil {
		switch *e {
		case "rate_limited", "assessment_budget_exhausted", "export_limit_reached":
			return true
		}
	}
	return false
}
