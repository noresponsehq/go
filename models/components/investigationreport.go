package components

import (
	"github.com/noresponsehq/go/v14/internal/utils"
)

// InvestigationReport - Read an investigation and its evidence.
type InvestigationReport struct {
	// The accounts this result matched, with how each was matched.
	Profiles []ScreeningProfile `json:"profiles"`
	// What this result flagged, most severe first, each citing its evidence.
	Findings []ScreeningFinding `json:"findings"`
	// All posts retained for this result, including posts that no finding cites.
	Posts []ScreeningPost `json:"posts"`
	// All images and videos retained for this result, including media that no finding cites.
	Media []ScreeningMedia `json:"media"`
	// Every recorded decision, oldest first.
	Reviews []PersonReview `json:"reviews"`
	// What the signup check's lookups found for a person: account presence for the email, phone and usernames (the person's handles), the phone lookup, the email domain's registration and the breach search. Context and coverage for the reviewer, not a finding: nothing here sets a band or a match state.
	Lookup *ScreeningLookup `json:"lookup,omitzero"`
	// True when evidence.read is explicitly denied. Submitted details, coverage, job status and decisions remain available. Collected records, profiles, findings, posts, media, discovered values and free-text decision notes are withheld.
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

func (i *InvestigationReport) GetFindings() []ScreeningFinding {
	if i == nil {
		return []ScreeningFinding{}
	}
	return i.Findings
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

func (i *InvestigationReport) GetReviews() []PersonReview {
	if i == nil {
		return []PersonReview{}
	}
	return i.Reviews
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
