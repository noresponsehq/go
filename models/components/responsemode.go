package components

// ResponseMode - standard returns the consolidated risk response; detailed adds merged provider-neutral facts when the key has detailed scope. There is no raw mode.
type ResponseMode string

const (
	ResponseModeStandard ResponseMode = "standard"
	ResponseModeDetailed ResponseMode = "detailed"
)

func (e ResponseMode) ToPointer() *ResponseMode {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ResponseMode) IsExact() bool {
	if e != nil {
		switch *e {
		case "standard", "detailed":
			return true
		}
	}
	return false
}
