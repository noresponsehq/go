package components

type LastError string

const (
	LastErrorScreeningCapExceeded LastError = "screening_cap_exceeded"
	LastErrorScreeningUnavailable LastError = "screening_unavailable"
	LastErrorPersonNotFound       LastError = "person_not_found"
	LastErrorIdempotencyKeyReused LastError = "idempotency_key_reused"
)

func (e LastError) ToPointer() *LastError {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *LastError) IsExact() bool {
	if e != nil {
		switch *e {
		case "screening_cap_exceeded", "screening_unavailable", "person_not_found", "idempotency_key_reused":
			return true
		}
	}
	return false
}
