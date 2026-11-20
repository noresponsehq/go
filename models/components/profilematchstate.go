package components

// ProfileMatchState - How sure the screening is that a profile is the person's, from the likelihood its record scores. Confirmed reaches the join bar of 0.9. Probable reaches 0.85, which takes a shared value beyond the one that found the account. Possible is the rest.
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
