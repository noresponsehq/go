package components

// SearchKind - The kind of the value searched.
type SearchKind string

const (
	SearchKindEmail      SearchKind = "email"
	SearchKindPhone      SearchKind = "phone"
	SearchKindUsername   SearchKind = "username"
	SearchKindPersonName SearchKind = "person_name"
	SearchKindPassword   SearchKind = "password"
	SearchKindHash       SearchKind = "hash"
	SearchKindIP         SearchKind = "ip"
)

func (e SearchKind) ToPointer() *SearchKind {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *SearchKind) IsExact() bool {
	if e != nil {
		switch *e {
		case "email", "phone", "username", "person_name", "password", "hash", "ip":
			return true
		}
	}
	return false
}
