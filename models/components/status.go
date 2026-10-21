package components

// Status - complete: final; nothing is still running. A lane that failed or missed its 10 s deadline stays listed in incomplete_sources, and confidence is lower. partial: a recommendation exists and at least one lane is still running or did not settle (see incomplete_sources). pending: nothing usable before the deadline; no recommendation yet, poll or wait for the webhook.
type Status string

const (
	StatusComplete Status = "complete"
	StatusPartial  Status = "partial"
	StatusPending  Status = "pending"
)

func (e Status) ToPointer() *Status {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Status) IsExact() bool {
	if e != nil {
		switch *e {
		case "complete", "partial", "pending":
			return true
		}
	}
	return false
}
