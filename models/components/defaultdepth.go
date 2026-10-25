package components

// DefaultDepth - The depth a request for this flow that names none is served at. Defaults: full for account_creation and purchase, quick for login.
type DefaultDepth string

const (
	DefaultDepthQuick DefaultDepth = "quick"
	DefaultDepthFull  DefaultDepth = "full"
)

func (e DefaultDepth) ToPointer() *DefaultDepth {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *DefaultDepth) IsExact() bool {
	if e != nil {
		switch *e {
		case "quick", "full":
			return true
		}
	}
	return false
}
