package components

// AuthenticationErrorCode - Machine codes an HTTP 401 answer can carry.
type AuthenticationErrorCode string

const (
	AuthenticationErrorCodeUnauthorized AuthenticationErrorCode = "unauthorized"
)

func (e AuthenticationErrorCode) ToPointer() *AuthenticationErrorCode {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *AuthenticationErrorCode) IsExact() bool {
	if e != nil {
		switch *e {
		case "unauthorized":
			return true
		}
	}
	return false
}
