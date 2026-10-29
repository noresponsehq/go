package components

// Status - Processing state, separate from evidence coverage. complete: final; nothing is still running, but unsupported, failed or unanswered lanes remain in incomplete_sources. partial: a recommendation exists and evidence is incomplete; assessment_url identifies continuing work when present. pending: no recommendation yet; poll or wait for the webhook. Finished processing does not mean all sources answered.
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
