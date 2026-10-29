package components

// Match - exact compares the original submitted value; canonical uses identifier-specific normalisation; mailbox explicitly selects the email alias rules used by lists.
type Match string

const (
	MatchExact     Match = "exact"
	MatchCanonical Match = "canonical"
	MatchMailbox   Match = "mailbox"
)

func (e Match) ToPointer() *Match {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Match) IsExact() bool {
	if e != nil {
		switch *e {
		case "exact", "canonical", "mailbox":
			return true
		}
	}
	return false
}
