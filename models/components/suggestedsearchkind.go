package components

// SuggestedSearchKind - The kind of the value: one of the kinds the environment's enrichment setting names.
type SuggestedSearchKind string

const (
	SuggestedSearchKindUsername SuggestedSearchKind = "username"
	SuggestedSearchKindEmail    SuggestedSearchKind = "email"
	SuggestedSearchKindPhone    SuggestedSearchKind = "phone"
)

func (e SuggestedSearchKind) ToPointer() *SuggestedSearchKind {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *SuggestedSearchKind) IsExact() bool {
	if e != nil {
		switch *e {
		case "username", "email", "phone":
			return true
		}
	}
	return false
}
