package components

import (
	"github.com/noresponsehq/go/v20/internal/utils"
)

// InvestigationReport - Read an investigation and its evidence.
type InvestigationReport struct {
	// The accounts this result matched, with how each was matched.
	Profiles []ScreeningProfile `json:"profiles"`
	// All posts retained for this result.
	Posts []ScreeningPost `json:"posts"`
	// All images and videos retained for this result.
	Media []ScreeningMedia `json:"media"`
	// What the signup check's lookups found for a person: account presence for the email, phone and usernames (the person's handles), the phone lookup, the email domain's registration and the breach search. Nothing here sets a match state.
	Lookup *ScreeningLookup `json:"lookup,omitzero"`
	// True when evidence.read is explicitly denied. Submitted details, coverage and job status remain available. Collected records, profiles, posts, media and discovered values are withheld.
	EvidenceRestricted *bool           `json:"evidence_restricted,omitzero"`
	Person             WorkspacePerson `json:"person"`
	Investigation      Investigation   `json:"investigation"`
}

func (i InvestigationReport) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InvestigationReport) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (i *InvestigationReport) GetProfiles() []ScreeningProfile {
	if i == nil {
		return []ScreeningProfile{}
	}
	return i.Profiles
}

func (i *InvestigationReport) GetPosts() []ScreeningPost {
	if i == nil {
		return []ScreeningPost{}
	}
	return i.Posts
}

func (i *InvestigationReport) GetMedia() []ScreeningMedia {
	if i == nil {
		return []ScreeningMedia{}
	}
	return i.Media
}

func (i *InvestigationReport) GetLookup() *ScreeningLookup {
	if i == nil {
		return nil
	}
	return i.Lookup
}

func (i *InvestigationReport) GetEvidenceRestricted() *bool {
	if i == nil {
		return nil
	}
	return i.EvidenceRestricted
}

func (i *InvestigationReport) GetPerson() WorkspacePerson {
	if i == nil {
		return WorkspacePerson{}
	}
	return i.Person
}

func (i *InvestigationReport) GetInvestigation() Investigation {
	if i == nil {
		return Investigation{}
	}
	return i.Investigation
}
