package components

import (
	"github.com/noresponsehq/go/v15/internal/utils"
)

// WorkspacePersonReport - Read a workspace person.
type WorkspacePersonReport struct {
	// Your reference for this person within the workspace environment. Only an explicit reference or person_id can reuse an existing person.
	ExternalID *string `json:"external_id,omitzero"`
	Name       *string `json:"name,omitzero"`
	// The person's email.
	Email *string `json:"email,omitzero"`
	// The person's phone in E.164 form.
	Phone *string `json:"phone,omitzero"`
	// Where the person works.
	Company *string `json:"company,omitzero"`
	// Where the person lives.
	City *string `json:"city,omitzero"`
	// Accounts the person is known to hold, at most 20.
	Handles   []Handle `json:"handles,omitzero"`
	ID        string   `json:"id"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	// The time the workspace retention policy removes this person, their evidence and investigation history. Shortening workspace retention can move this time earlier.
	ExpiresAt string `json:"expires_at"`
	// The accounts this result matched, with how each was matched.
	Profiles []ScreeningProfile `json:"profiles"`
	// What this result flagged, most severe first, each citing its evidence.
	Findings []ScreeningFinding `json:"findings"`
	// All posts retained for this result, including posts that no finding cites.
	Posts []ScreeningPost `json:"posts"`
	// All images and videos retained for this result, including media that no finding cites.
	Media []ScreeningMedia `json:"media"`
	// What the signup check's lookups found for a person: account presence for the email, phone and usernames (the person's handles), the phone lookup, the email domain's registration and the breach search. Context and coverage for the reviewer, not a finding: nothing here sets a band or a match state.
	Lookup *ScreeningLookup `json:"lookup,omitzero"`
	// True when evidence.read is explicitly denied. Submitted details, coverage, job status and decisions remain available. Collected records, profiles, findings, posts, media, discovered values and free-text decision notes are withheld.
	EvidenceRestricted *bool `json:"evidence_restricted,omitzero"`
	// The most recent retained completed or partial investigation, or a migrated result whose original per-person status is unknown. The evidence below belongs to this investigation. Read investigation history for earlier results and pending jobs.
	EvidenceInvestigation *Investigation `json:"evidence_investigation"`
}

func (w WorkspacePersonReport) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WorkspacePersonReport) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (w *WorkspacePersonReport) GetExternalID() *string {
	if w == nil {
		return nil
	}
	return w.ExternalID
}

func (w *WorkspacePersonReport) GetName() *string {
	if w == nil {
		return nil
	}
	return w.Name
}

func (w *WorkspacePersonReport) GetEmail() *string {
	if w == nil {
		return nil
	}
	return w.Email
}

func (w *WorkspacePersonReport) GetPhone() *string {
	if w == nil {
		return nil
	}
	return w.Phone
}

func (w *WorkspacePersonReport) GetCompany() *string {
	if w == nil {
		return nil
	}
	return w.Company
}

func (w *WorkspacePersonReport) GetCity() *string {
	if w == nil {
		return nil
	}
	return w.City
}

func (w *WorkspacePersonReport) GetHandles() []Handle {
	if w == nil {
		return nil
	}
	return w.Handles
}

func (w *WorkspacePersonReport) GetID() string {
	if w == nil {
		return ""
	}
	return w.ID
}

func (w *WorkspacePersonReport) GetCreatedAt() string {
	if w == nil {
		return ""
	}
	return w.CreatedAt
}

func (w *WorkspacePersonReport) GetUpdatedAt() string {
	if w == nil {
		return ""
	}
	return w.UpdatedAt
}

func (w *WorkspacePersonReport) GetExpiresAt() string {
	if w == nil {
		return ""
	}
	return w.ExpiresAt
}

func (w *WorkspacePersonReport) GetProfiles() []ScreeningProfile {
	if w == nil {
		return []ScreeningProfile{}
	}
	return w.Profiles
}

func (w *WorkspacePersonReport) GetFindings() []ScreeningFinding {
	if w == nil {
		return []ScreeningFinding{}
	}
	return w.Findings
}

func (w *WorkspacePersonReport) GetPosts() []ScreeningPost {
	if w == nil {
		return []ScreeningPost{}
	}
	return w.Posts
}

func (w *WorkspacePersonReport) GetMedia() []ScreeningMedia {
	if w == nil {
		return []ScreeningMedia{}
	}
	return w.Media
}

func (w *WorkspacePersonReport) GetLookup() *ScreeningLookup {
	if w == nil {
		return nil
	}
	return w.Lookup
}

func (w *WorkspacePersonReport) GetEvidenceRestricted() *bool {
	if w == nil {
		return nil
	}
	return w.EvidenceRestricted
}

func (w *WorkspacePersonReport) GetEvidenceInvestigation() *Investigation {
	if w == nil {
		return nil
	}
	return w.EvidenceInvestigation
}
