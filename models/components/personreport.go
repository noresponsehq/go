package components

import (
	"github.com/noresponsehq/go/v11/internal/utils"
)

// PersonReport - Everything known about one person, and where else in the environment they appear.
type PersonReport struct {
	// A person and where their screening and review stand.
	Person Person `json:"person"`
	// The accounts the latest screening matched, with how each was matched.
	Profiles []ScreeningProfile `json:"profiles"`
	// What the latest screening flagged, most severe first, each citing its evidence.
	Findings []ScreeningFinding `json:"findings"`
	// The posts the findings cite.
	Posts []ScreeningPost `json:"posts"`
	// The images and videos the findings cite.
	Media []ScreeningMedia `json:"media"`
	// Every recorded decision, oldest first.
	Reviews []PersonReview `json:"reviews"`
	// What the signup check's lookups found for a person: account presence for the email, phone and usernames (the person's handles), the phone lookup, the email domain's registration and the breach search. Context and coverage for the reviewer, not a finding: nothing here sets a band or a match state.
	Lookup *ScreeningLookup `json:"lookup,omitzero"`
	// Where else in the environment the person appears, read live from those groups: deleting one removes it here.
	OtherGroups OtherGroups `json:"other_groups"`
}

func (p PersonReport) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PersonReport) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *PersonReport) GetPerson() Person {
	if p == nil {
		return Person{}
	}
	return p.Person
}

func (p *PersonReport) GetProfiles() []ScreeningProfile {
	if p == nil {
		return []ScreeningProfile{}
	}
	return p.Profiles
}

func (p *PersonReport) GetFindings() []ScreeningFinding {
	if p == nil {
		return []ScreeningFinding{}
	}
	return p.Findings
}

func (p *PersonReport) GetPosts() []ScreeningPost {
	if p == nil {
		return []ScreeningPost{}
	}
	return p.Posts
}

func (p *PersonReport) GetMedia() []ScreeningMedia {
	if p == nil {
		return []ScreeningMedia{}
	}
	return p.Media
}

func (p *PersonReport) GetReviews() []PersonReview {
	if p == nil {
		return []PersonReview{}
	}
	return p.Reviews
}

func (p *PersonReport) GetLookup() *ScreeningLookup {
	if p == nil {
		return nil
	}
	return p.Lookup
}

func (p *PersonReport) GetOtherGroups() OtherGroups {
	if p == nil {
		return OtherGroups{}
	}
	return p.OtherGroups
}
