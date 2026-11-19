package components

// PersonChangeKind - account_added and account_removed: an account or profile appeared or disappeared. breach_added and breach_removed: a breach record appeared or disappeared. identifier_added and identifier_removed: an email, phone, username, name or account id on an account the person already held. field_changed: another field on such an account. posts_added: new posts, counted in count.
type PersonChangeKind string

const (
	PersonChangeKindAccountAdded      PersonChangeKind = "account_added"
	PersonChangeKindAccountRemoved    PersonChangeKind = "account_removed"
	PersonChangeKindBreachAdded       PersonChangeKind = "breach_added"
	PersonChangeKindBreachRemoved     PersonChangeKind = "breach_removed"
	PersonChangeKindIdentifierAdded   PersonChangeKind = "identifier_added"
	PersonChangeKindIdentifierRemoved PersonChangeKind = "identifier_removed"
	PersonChangeKindFieldChanged      PersonChangeKind = "field_changed"
	PersonChangeKindPostsAdded        PersonChangeKind = "posts_added"
)

func (e PersonChangeKind) ToPointer() *PersonChangeKind {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *PersonChangeKind) IsExact() bool {
	if e != nil {
		switch *e {
		case "account_added", "account_removed", "breach_added", "breach_removed", "identifier_added", "identifier_removed", "field_changed", "posts_added":
			return true
		}
	}
	return false
}
