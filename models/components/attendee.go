package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// Attendee - An attendee and where their screening and review stand.
type Attendee struct {
	ID      string `json:"id"`
	EventID string `json:"event_id"`
	// Your own id for the attendee, unique within the event. A later row with the same external id updates this attendee.
	ExternalID *string `json:"external_id,omitzero"`
	// The attendee's full name.
	Name string `json:"name"`
	// The attendee's email.
	Email *string `json:"email,omitzero"`
	// The attendee's phone in E.164 form.
	Phone *string `json:"phone,omitzero"`
	// Where the attendee works.
	Company *string `json:"company,omitzero"`
	// Where the attendee lives.
	City *string `json:"city,omitzero"`
	// Accounts the attendee is known to hold, at most 20.
	Handles []Handle `json:"handles,omitzero"`
	// The attendee's role at the event: guest (a ticket holder), crew, vendor or press. Leave it out when you do not know it: an attendee without one has no role, not guest. A later row naming the same attendee replaces it, and a change of role alone keeps the attendee's screening results.
	Role *AttendeeRole `json:"role,omitzero"`
	// How sure the latest screening is that it found the attendee's accounts. unscreened until a screening reaches the attendee.
	MatchState MatchState `json:"match_state"`
	// The most severe finding's band, or null when nothing was flagged.
	HighestBand *HighestBand `json:"highest_band"`
	// The latest reviewer decision, or null before any.
	Decision     *AttendeeDecision `json:"decision"`
	DecisionNote *string           `json:"decision_note,omitzero"`
	DecidedBy    *string           `json:"decided_by,omitzero"`
	DecidedAt    *string           `json:"decided_at,omitzero"`
	// The denominator of a screening: accounts found, how many could be read, items read and the period they span, so a clean result is distinct from a person nobody could see.
	Coverage        ScreeningCoverage `json:"coverage"`
	LastScreeningID *string           `json:"last_screening_id,omitzero"`
	ScreenedAt      *string           `json:"screened_at,omitzero"`
	CreatedAt       string            `json:"created_at"`
	UpdatedAt       string            `json:"updated_at"`
	// What screening cost, in US cents, in total and per implementation.
	Spend ScreeningSpend `json:"spend"`
}

func (a Attendee) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *Attendee) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *Attendee) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

func (a *Attendee) GetEventID() string {
	if a == nil {
		return ""
	}
	return a.EventID
}

func (a *Attendee) GetExternalID() *string {
	if a == nil {
		return nil
	}
	return a.ExternalID
}

func (a *Attendee) GetName() string {
	if a == nil {
		return ""
	}
	return a.Name
}

func (a *Attendee) GetEmail() *string {
	if a == nil {
		return nil
	}
	return a.Email
}

func (a *Attendee) GetPhone() *string {
	if a == nil {
		return nil
	}
	return a.Phone
}

func (a *Attendee) GetCompany() *string {
	if a == nil {
		return nil
	}
	return a.Company
}

func (a *Attendee) GetCity() *string {
	if a == nil {
		return nil
	}
	return a.City
}

func (a *Attendee) GetHandles() []Handle {
	if a == nil {
		return nil
	}
	return a.Handles
}

func (a *Attendee) GetRole() *AttendeeRole {
	if a == nil {
		return nil
	}
	return a.Role
}

func (a *Attendee) GetMatchState() MatchState {
	if a == nil {
		return MatchState("")
	}
	return a.MatchState
}

func (a *Attendee) GetHighestBand() *HighestBand {
	if a == nil {
		return nil
	}
	return a.HighestBand
}

func (a *Attendee) GetDecision() *AttendeeDecision {
	if a == nil {
		return nil
	}
	return a.Decision
}

func (a *Attendee) GetDecisionNote() *string {
	if a == nil {
		return nil
	}
	return a.DecisionNote
}

func (a *Attendee) GetDecidedBy() *string {
	if a == nil {
		return nil
	}
	return a.DecidedBy
}

func (a *Attendee) GetDecidedAt() *string {
	if a == nil {
		return nil
	}
	return a.DecidedAt
}

func (a *Attendee) GetCoverage() ScreeningCoverage {
	if a == nil {
		return ScreeningCoverage{}
	}
	return a.Coverage
}

func (a *Attendee) GetLastScreeningID() *string {
	if a == nil {
		return nil
	}
	return a.LastScreeningID
}

func (a *Attendee) GetScreenedAt() *string {
	if a == nil {
		return nil
	}
	return a.ScreenedAt
}

func (a *Attendee) GetCreatedAt() string {
	if a == nil {
		return ""
	}
	return a.CreatedAt
}

func (a *Attendee) GetUpdatedAt() string {
	if a == nil {
		return ""
	}
	return a.UpdatedAt
}

func (a *Attendee) GetSpend() ScreeningSpend {
	if a == nil {
		return ScreeningSpend{}
	}
	return a.Spend
}
