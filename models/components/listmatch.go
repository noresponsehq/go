package components

// ListMatch - The action of the list entry that decided the account.
type ListMatch string

const (
	ListMatchBlock ListMatch = "block"
	ListMatchAllow ListMatch = "allow"
)

func (e ListMatch) ToPointer() *ListMatch {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ListMatch) IsExact() bool {
	if e != nil {
		switch *e {
		case "block", "allow":
			return true
		}
	}
	return false
}
