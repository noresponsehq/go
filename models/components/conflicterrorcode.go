package components

// ConflictErrorCode - Machine codes an HTTP 409 answer can carry.
type ConflictErrorCode string

const (
	ConflictErrorCodeIdempotencyKeyReused  ConflictErrorCode = "idempotency_key_reused"
	ConflictErrorCodeKeyRotationInProgress ConflictErrorCode = "key_rotation_in_progress"
	ConflictErrorCodeTooManyIntegrations   ConflictErrorCode = "too_many_integrations"
)

func (e ConflictErrorCode) ToPointer() *ConflictErrorCode {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ConflictErrorCode) IsExact() bool {
	if e != nil {
		switch *e {
		case "idempotency_key_reused", "key_rotation_in_progress", "too_many_integrations":
			return true
		}
	}
	return false
}
