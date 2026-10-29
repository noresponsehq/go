package components

type AccountOutcomeActorType string

const (
	AccountOutcomeActorTypeAPIKey AccountOutcomeActorType = "api_key"
	AccountOutcomeActorTypeUser   AccountOutcomeActorType = "user"
)

func (e AccountOutcomeActorType) ToPointer() *AccountOutcomeActorType {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *AccountOutcomeActorType) IsExact() bool {
	if e != nil {
		switch *e {
		case "api_key", "user":
			return true
		}
	}
	return false
}
