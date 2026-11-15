package components

// SearchKind - The kind of the value searched. A domain search finds the leaked rows of every email address at the domain.
type SearchKind string

const (
	SearchKindEmail      SearchKind = "email"
	SearchKindPhone      SearchKind = "phone"
	SearchKindUsername   SearchKind = "username"
	SearchKindPersonName SearchKind = "person_name"
	SearchKindPassword   SearchKind = "password"
	SearchKindHash       SearchKind = "hash"
	SearchKindIP         SearchKind = "ip"
	SearchKindDomain     SearchKind = "domain"
)

func (e SearchKind) ToPointer() *SearchKind {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *SearchKind) IsExact() bool {
	if e != nil {
		switch *e {
		case "email", "phone", "username", "person_name", "password", "hash", "ip", "domain":
			return true
		}
	}
	return false
}
