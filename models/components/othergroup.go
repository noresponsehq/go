package components

type MatchedOn string

const (
	MatchedOnEmail      MatchedOn = "email"
	MatchedOnPhone      MatchedOn = "phone"
	MatchedOnExternalID MatchedOn = "external_id"
)

func (e MatchedOn) ToPointer() *MatchedOn {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *MatchedOn) IsExact() bool {
	if e != nil {
		switch *e {
		case "email", "phone", "external_id":
			return true
		}
	}
	return false
}

// OtherGroup - A person of another group in the same environment who shares an email or phone with this one, as that group recorded them.
type OtherGroup struct {
	GroupID       string `json:"group_id"`
	GroupName     string `json:"group_name"`
	GroupStartsAt string `json:"group_starts_at"`
	// The person on that group's list.
	PersonID string `json:"person_id"`
	// Which identifiers match, in this order. external_id appears only beside email or phone: the same external id alone is not a match, since ticket numbers may be reused across groups.
	MatchedOn []MatchedOn `json:"matched_on"`
	// How sure the latest screening is that it found the person's accounts. unscreened until a screening reaches the person, and after one that had no profile to check (no_profile_to_check).
	MatchState MatchState `json:"match_state"`
	// true when that person's latest screening had no profile to check (see the person's no_profile_to_check). Left out otherwise.
	NoProfileToCheck *bool `json:"no_profile_to_check,omitzero"`
	// That person's most severe band in that group, or null.
	HighestBand *HighestBand `json:"highest_band"`
	// The latest reviewer decision in that group, or null before any.
	Decision *PersonDecision `json:"decision"`
	// Who made that decision.
	DecidedBy *string `json:"decided_by,omitzero"`
	// When that decision was made.
	DecidedAt *string `json:"decided_at,omitzero"`
}

func (o *OtherGroup) GetGroupID() string {
	if o == nil {
		return ""
	}
	return o.GroupID
}

func (o *OtherGroup) GetGroupName() string {
	if o == nil {
		return ""
	}
	return o.GroupName
}

func (o *OtherGroup) GetGroupStartsAt() string {
	if o == nil {
		return ""
	}
	return o.GroupStartsAt
}

func (o *OtherGroup) GetPersonID() string {
	if o == nil {
		return ""
	}
	return o.PersonID
}

func (o *OtherGroup) GetMatchedOn() []MatchedOn {
	if o == nil {
		return []MatchedOn{}
	}
	return o.MatchedOn
}

func (o *OtherGroup) GetMatchState() MatchState {
	if o == nil {
		return MatchState("")
	}
	return o.MatchState
}

func (o *OtherGroup) GetNoProfileToCheck() *bool {
	if o == nil {
		return nil
	}
	return o.NoProfileToCheck
}

func (o *OtherGroup) GetHighestBand() *HighestBand {
	if o == nil {
		return nil
	}
	return o.HighestBand
}

func (o *OtherGroup) GetDecision() *PersonDecision {
	if o == nil {
		return nil
	}
	return o.Decision
}

func (o *OtherGroup) GetDecidedBy() *string {
	if o == nil {
		return nil
	}
	return o.DecidedBy
}

func (o *OtherGroup) GetDecidedAt() *string {
	if o == nil {
		return nil
	}
	return o.DecidedAt
}
