package components

// Access - public profiles were read; a private or locked one exists but could not be read.
type Access string

const (
	AccessPublic  Access = "public"
	AccessPrivate Access = "private"
	AccessLocked  Access = "locked"
)

func (e Access) ToPointer() *Access {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Access) IsExact() bool {
	if e != nil {
		switch *e {
		case "public", "private", "locked":
			return true
		}
	}
	return false
}
