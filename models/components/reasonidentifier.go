package components

// ReasonIdentifier - The submitted identifier or client context field a reason concerns; absent when it concerns the whole assessment. Client context is a claimed hint, not a device fingerprint.
type ReasonIdentifier string

const (
	ReasonIdentifierEmail    ReasonIdentifier = "email"
	ReasonIdentifierUsername ReasonIdentifier = "username"
	ReasonIdentifierPhone    ReasonIdentifier = "phone"
	ReasonIdentifierName     ReasonIdentifier = "name"
	ReasonIdentifierIP       ReasonIdentifier = "ip"
	ReasonIdentifierContext  ReasonIdentifier = "context"
)

func (e ReasonIdentifier) ToPointer() *ReasonIdentifier {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ReasonIdentifier) IsExact() bool {
	if e != nil {
		switch *e {
		case "email", "username", "phone", "name", "ip", "context":
			return true
		}
	}
	return false
}
