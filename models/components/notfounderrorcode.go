package components

// NotFoundErrorCode - Machine codes an HTTP 404 answer can carry.
type NotFoundErrorCode string

const (
	NotFoundErrorCodeNotFound NotFoundErrorCode = "not_found"
)

func (e NotFoundErrorCode) ToPointer() *NotFoundErrorCode {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *NotFoundErrorCode) IsExact() bool {
	if e != nil {
		switch *e {
		case "not_found":
			return true
		}
	}
	return false
}
