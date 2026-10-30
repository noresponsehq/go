package components

// AccountType - personal, or business: an account a company or school manages.
type AccountType string

const (
	AccountTypePersonal AccountType = "personal"
	AccountTypeBusiness AccountType = "business"
)

func (e AccountType) ToPointer() *AccountType {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *AccountType) IsExact() bool {
	if e != nil {
		switch *e {
		case "personal", "business":
			return true
		}
	}
	return false
}
