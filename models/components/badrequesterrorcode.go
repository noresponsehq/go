package components

// BadRequestErrorCode - Machine codes an HTTP 400 answer can carry.
type BadRequestErrorCode string

const (
	BadRequestErrorCodeInvalidRequest        BadRequestErrorCode = "invalid_request"
	BadRequestErrorCodeInvalidIdempotencyKey BadRequestErrorCode = "invalid_idempotency_key"
	BadRequestErrorCodeInvalidCursor         BadRequestErrorCode = "invalid_cursor"
)

func (e BadRequestErrorCode) ToPointer() *BadRequestErrorCode {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *BadRequestErrorCode) IsExact() bool {
	if e != nil {
		switch *e {
		case "invalid_request", "invalid_idempotency_key", "invalid_cursor":
			return true
		}
	}
	return false
}
