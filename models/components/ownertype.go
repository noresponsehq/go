package components

// OwnerType - The credential namespace of the creator. An external user ID equal to an API-key label does not grant ownership of that key's private view.
type OwnerType string

const (
	OwnerTypeUser   OwnerType = "user"
	OwnerTypeAPIKey OwnerType = "api_key"
)

func (e OwnerType) ToPointer() *OwnerType {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *OwnerType) IsExact() bool {
	if e != nil {
		switch *e {
		case "user", "api_key":
			return true
		}
	}
	return false
}
