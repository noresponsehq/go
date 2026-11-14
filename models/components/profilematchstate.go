package components

// ProfileMatchState - How sure the screening is that a profile is the person's, from its evidence: confirmed is a handle on the person row, or two or more hard identifiers (email, phone, a supplied handle, or a fact an existing source returned); probable is one identifier beyond the name; possible is the name alone.
type ProfileMatchState string

const (
	ProfileMatchStateConfirmed ProfileMatchState = "confirmed"
	ProfileMatchStateProbable  ProfileMatchState = "probable"
	ProfileMatchStatePossible  ProfileMatchState = "possible"
)

func (e ProfileMatchState) ToPointer() *ProfileMatchState {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ProfileMatchState) IsExact() bool {
	if e != nil {
		switch *e {
		case "confirmed", "probable", "possible":
			return true
		}
	}
	return false
}
