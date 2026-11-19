package components

type RiskDetailsEventIdentifier string

const (
	RiskDetailsEventIdentifierEmail    RiskDetailsEventIdentifier = "email"
	RiskDetailsEventIdentifierUsername RiskDetailsEventIdentifier = "username"
)

func (e RiskDetailsEventIdentifier) ToPointer() *RiskDetailsEventIdentifier {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *RiskDetailsEventIdentifier) IsExact() bool {
	if e != nil {
		switch *e {
		case "email", "username":
			return true
		}
	}
	return false
}
