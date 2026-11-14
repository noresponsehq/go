package components

import (
	"github.com/noresponsehq/go/v11/internal/utils"
)

// Person - A person and where their screening and review stand.
type Person struct {
	ID      string `json:"id"`
	GroupID string `json:"group_id"`
	// Your own id for the person, unique within the group. A later row with the same external id updates this person.
	ExternalID *string `json:"external_id,omitzero"`
	// The person's full name.
	Name string `json:"name"`
	// The person's email.
	Email *string `json:"email,omitzero"`
	// The person's phone in E.164 form.
	Phone *string `json:"phone,omitzero"`
	// Where the person works.
	Company *string `json:"company,omitzero"`
	// Where the person lives.
	City *string `json:"city,omitzero"`
	// Accounts the person is known to hold, at most 20.
	Handles []Handle `json:"handles,omitzero"`
	// The person's role in the group: guest (a ticket holder), crew, vendor or press. Leave it out when you do not know it: a person without one has no role, not guest. A later row naming the same person replaces it, and a change of role alone keeps the person's screening results.
	Role *PersonRole `json:"role,omitzero"`
	// How sure the latest screening is that it found the person's accounts. unscreened until a screening reaches the person, and after one that had no profile to check (no_profile_to_check).
	MatchState MatchState `json:"match_state"`
	// true when the latest screening reached the person but the list gave nothing to look up on any requested platform, such as no handle for a platform that reads supplied handles only: nobody looked for a profile. match_state stays unscreened, every platform is in coverage.not_searched, and the person is neither none_found nor failed and does not make the screening partial. Its lookups are on its report. Left out otherwise.
	NoProfileToCheck *bool `json:"no_profile_to_check,omitzero"`
	// The most severe finding's band, or null when nothing was flagged.
	HighestBand *HighestBand `json:"highest_band"`
	// The latest reviewer decision, or null before any.
	Decision     *PersonDecision `json:"decision"`
	DecisionNote *string         `json:"decision_note,omitzero"`
	DecidedBy    *string         `json:"decided_by,omitzero"`
	DecidedAt    *string         `json:"decided_at,omitzero"`
	// The denominator of a screening: accounts found, how many could be read, items read and the period they span, so a clean result is distinct from a person nobody could see.
	Coverage        ScreeningCoverage `json:"coverage"`
	LastScreeningID *string           `json:"last_screening_id,omitzero"`
	ScreenedAt      *string           `json:"screened_at,omitzero"`
	CreatedAt       string            `json:"created_at"`
	UpdatedAt       string            `json:"updated_at"`
	// What screening cost, in US cents, in total and per implementation.
	Spend ScreeningSpend `json:"spend"`
}

func (p Person) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Person) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *Person) GetID() string {
	if p == nil {
		return ""
	}
	return p.ID
}

func (p *Person) GetGroupID() string {
	if p == nil {
		return ""
	}
	return p.GroupID
}

func (p *Person) GetExternalID() *string {
	if p == nil {
		return nil
	}
	return p.ExternalID
}

func (p *Person) GetName() string {
	if p == nil {
		return ""
	}
	return p.Name
}

func (p *Person) GetEmail() *string {
	if p == nil {
		return nil
	}
	return p.Email
}

func (p *Person) GetPhone() *string {
	if p == nil {
		return nil
	}
	return p.Phone
}

func (p *Person) GetCompany() *string {
	if p == nil {
		return nil
	}
	return p.Company
}

func (p *Person) GetCity() *string {
	if p == nil {
		return nil
	}
	return p.City
}

func (p *Person) GetHandles() []Handle {
	if p == nil {
		return nil
	}
	return p.Handles
}

func (p *Person) GetRole() *PersonRole {
	if p == nil {
		return nil
	}
	return p.Role
}

func (p *Person) GetMatchState() MatchState {
	if p == nil {
		return MatchState("")
	}
	return p.MatchState
}

func (p *Person) GetNoProfileToCheck() *bool {
	if p == nil {
		return nil
	}
	return p.NoProfileToCheck
}

func (p *Person) GetHighestBand() *HighestBand {
	if p == nil {
		return nil
	}
	return p.HighestBand
}

func (p *Person) GetDecision() *PersonDecision {
	if p == nil {
		return nil
	}
	return p.Decision
}

func (p *Person) GetDecisionNote() *string {
	if p == nil {
		return nil
	}
	return p.DecisionNote
}

func (p *Person) GetDecidedBy() *string {
	if p == nil {
		return nil
	}
	return p.DecidedBy
}

func (p *Person) GetDecidedAt() *string {
	if p == nil {
		return nil
	}
	return p.DecidedAt
}

func (p *Person) GetCoverage() ScreeningCoverage {
	if p == nil {
		return ScreeningCoverage{}
	}
	return p.Coverage
}

func (p *Person) GetLastScreeningID() *string {
	if p == nil {
		return nil
	}
	return p.LastScreeningID
}

func (p *Person) GetScreenedAt() *string {
	if p == nil {
		return nil
	}
	return p.ScreenedAt
}

func (p *Person) GetCreatedAt() string {
	if p == nil {
		return ""
	}
	return p.CreatedAt
}

func (p *Person) GetUpdatedAt() string {
	if p == nil {
		return ""
	}
	return p.UpdatedAt
}

func (p *Person) GetSpend() ScreeningSpend {
	if p == nil {
		return ScreeningSpend{}
	}
	return p.Spend
}
