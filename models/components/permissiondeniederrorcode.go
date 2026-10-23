package components

// PermissionDeniedErrorCode - Machine codes an HTTP 403 answer can carry.
type PermissionDeniedErrorCode string

const (
	PermissionDeniedErrorCodeWorkspaceForbidden PermissionDeniedErrorCode = "workspace_forbidden"
)

func (e PermissionDeniedErrorCode) ToPointer() *PermissionDeniedErrorCode {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *PermissionDeniedErrorCode) IsExact() bool {
	if e != nil {
		switch *e {
		case "workspace_forbidden":
			return true
		}
	}
	return false
}
