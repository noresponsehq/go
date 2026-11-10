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

// OtherEvent - An attendee of another event in the same environment who shares an email or phone with this one, as that event recorded them.
type OtherEvent struct {
	EventID       string `json:"event_id"`
	EventName     string `json:"event_name"`
	EventStartsAt string `json:"event_starts_at"`
	// The attendee on that event's list.
	AttendeeID string `json:"attendee_id"`
	// Which identifiers match, in this order. external_id appears only beside email or phone: the same external id alone is not a match, since ticket numbers may be reused across events.
	MatchedOn []MatchedOn `json:"matched_on"`
	// How sure the latest screening is that it found the attendee's accounts. unscreened until a screening reaches the attendee, and after one that had no profile to check (no_profile_to_check).
	MatchState MatchState `json:"match_state"`
	// true when that attendee's latest screening had no profile to check (see the attendee's no_profile_to_check). Left out otherwise.
	NoProfileToCheck *bool `json:"no_profile_to_check,omitzero"`
	// That attendee's most severe band at that event, or null.
	HighestBand *HighestBand `json:"highest_band"`
	// The latest reviewer decision at that event, or null before any.
	Decision *AttendeeDecision `json:"decision"`
	// Who made that decision.
	DecidedBy *string `json:"decided_by,omitzero"`
	// When that decision was made.
	DecidedAt *string `json:"decided_at,omitzero"`
}

func (o *OtherEvent) GetEventID() string {
	if o == nil {
		return ""
	}
	return o.EventID
}

func (o *OtherEvent) GetEventName() string {
	if o == nil {
		return ""
	}
	return o.EventName
}

func (o *OtherEvent) GetEventStartsAt() string {
	if o == nil {
		return ""
	}
	return o.EventStartsAt
}

func (o *OtherEvent) GetAttendeeID() string {
	if o == nil {
		return ""
	}
	return o.AttendeeID
}

func (o *OtherEvent) GetMatchedOn() []MatchedOn {
	if o == nil {
		return []MatchedOn{}
	}
	return o.MatchedOn
}

func (o *OtherEvent) GetMatchState() MatchState {
	if o == nil {
		return MatchState("")
	}
	return o.MatchState
}

func (o *OtherEvent) GetNoProfileToCheck() *bool {
	if o == nil {
		return nil
	}
	return o.NoProfileToCheck
}

func (o *OtherEvent) GetHighestBand() *HighestBand {
	if o == nil {
		return nil
	}
	return o.HighestBand
}

func (o *OtherEvent) GetDecision() *AttendeeDecision {
	if o == nil {
		return nil
	}
	return o.Decision
}

func (o *OtherEvent) GetDecidedBy() *string {
	if o == nil {
		return nil
	}
	return o.DecidedBy
}

func (o *OtherEvent) GetDecidedAt() *string {
	if o == nil {
		return nil
	}
	return o.DecidedAt
}
