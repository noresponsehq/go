package components

import (
	"github.com/noresponsehq/go/v20/internal/utils"
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
	// A username the person uses, with no platform. A search looks it up on every platform that finds accounts by username.
	Username *string `json:"username,omitzero"`
	// Accounts the person is known to hold, at most 20.
	Handles   []Handle `json:"handles,omitzero"`
	ID        string   `json:"id"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	// The time the workspace retention policy removes this person, their evidence and investigation history. Shortening workspace retention can move this time earlier.
	ExpiresAt string `json:"expires_at"`
	// Every group this person belongs to in this environment, most recently added first, with their role in each. A person can belong to any number of groups; removing them from a group or deleting the group removes it here.
	Groups []PersonGroup `json:"groups"`
	// The accounts this result matched, with how each was matched.
	Profiles []ScreeningProfile `json:"profiles"`
	// All posts retained for this result.
	Posts []ScreeningPost `json:"posts"`
	// All images and videos retained for this result.
	Media []ScreeningMedia `json:"media"`
	// What the signup check's lookups found for a person: account presence for the email, phone and usernames (the person's handles), the phone lookup, the email domain's registration and the breach search. Nothing here sets a match state.
	Lookup *ScreeningLookup `json:"lookup,omitzero"`
	// Every search on this person's values, newest first, whether a person started it or enrichment did: each names the value, the record and field it came from, and the search that found that record. Absent when nothing was searched.
	Searches []Search `json:"searches,omitzero"`
	// The searches enrichment would start on this person, likeliest record first, under the environment's enrichment setting for a person's own investigation. In off mode the list is empty. In suggest mode enrichment starts none of them by itself; in automatic mode it starts them as its limits allow. Empty when evidence.read is denied.
	SuggestedSearches []SuggestedSearch `json:"suggested_searches"`
	// True when evidence.read is explicitly denied. Submitted details, coverage, job status and decisions remain available. Collected records, profiles, posts, media, discovered values and free-text decision notes are withheld.
	EvidenceRestricted *bool `json:"evidence_restricted,omitzero"`
	// The most recent retained completed or partial investigation, or a migrated result whose original per-person status is unknown. The evidence below belongs to this investigation. Read investigation history for earlier results and pending jobs.
	EvidenceInvestigation *Investigation `json:"evidence_investigation"`
	// What changed between each settled investigation of this person and the one before it, unreviewed first, then most significant, then newest. The first investigation records no changes. When evidence.read is denied, before and after are withheld.
	Changes []PersonChange `json:"changes"`
	// The person's rerun schedule. Each due run starts a full investigation of the person, as POST /v1/investigations does, and its changes are stored as changes. The schedule lives in the database and the worker runs it, so it survives restarts and releases; each due run starts exactly one investigation.
	Schedule PersonSchedule `json:"schedule"`
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

func (w *WorkspacePersonReport) GetUsername() *string {
	if w == nil {
		return nil
	}
	return w.Username
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

func (w *WorkspacePersonReport) GetGroups() []PersonGroup {
	if w == nil {
		return []PersonGroup{}
	}
	return w.Groups
}

func (w *WorkspacePersonReport) GetProfiles() []ScreeningProfile {
	if w == nil {
		return []ScreeningProfile{}
	}
	return w.Profiles
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

func (w *WorkspacePersonReport) GetSearches() []Search {
	if w == nil {
		return nil
	}
	return w.Searches
}

func (w *WorkspacePersonReport) GetSuggestedSearches() []SuggestedSearch {
	if w == nil {
		return []SuggestedSearch{}
	}
	return w.SuggestedSearches
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

func (w *WorkspacePersonReport) GetChanges() []PersonChange {
	if w == nil {
		return []PersonChange{}
	}
	return w.Changes
}

func (w *WorkspacePersonReport) GetSchedule() PersonSchedule {
	if w == nil {
		return PersonSchedule{}
	}
	return w.Schedule
}
