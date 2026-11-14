package components

// MatchState - How sure the latest screening is that it found the person's accounts. unscreened until a screening reaches the person, and after one that had no profile to check (no_profile_to_check).
type MatchState string

const (
	MatchStateUnscreened   MatchState = "unscreened"
	MatchStateConfirmed    MatchState = "confirmed"
	MatchStateProbable     MatchState = "probable"
	MatchStatePossible     MatchState = "possible"
	MatchStateNoneFound    MatchState = "none_found"
	MatchStateInaccessible MatchState = "inaccessible"
)

func (e MatchState) ToPointer() *MatchState {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *MatchState) IsExact() bool {
	if e != nil {
		switch *e {
		case "unscreened", "confirmed", "probable", "possible", "none_found", "inaccessible":
			return true
		}
	}
	return false
}
