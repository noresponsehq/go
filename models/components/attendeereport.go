package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// AttendeeReport - Everything known about one attendee, and where else in the environment they appear.
type AttendeeReport struct {
	// An attendee and where their screening and review stand.
	Attendee Attendee `json:"attendee"`
	// The accounts the latest screening matched, with how each was matched.
	Profiles []ScreeningProfile `json:"profiles"`
	// What the latest screening flagged, most severe first, each citing its evidence.
	Findings []ScreeningFinding `json:"findings"`
	// The posts the findings cite.
	Posts []ScreeningPost `json:"posts"`
	// The images and videos the findings cite.
	Media []ScreeningMedia `json:"media"`
	// Every recorded decision, oldest first.
	Reviews []AttendeeReview `json:"reviews"`
	// What the signup check's lookups found for an attendee: account presence for the email, phone and usernames (the attendee's handles), the phone lookup, the email domain's registration and the breach search. Context and coverage for the reviewer, not a finding: nothing here sets a band or a match state.
	Lookup *ScreeningLookup `json:"lookup,omitzero"`
	// Where else in the environment the attendee appears, read live from those events: deleting one removes it here.
	OtherEvents OtherEvents `json:"other_events"`
}

func (a AttendeeReport) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AttendeeReport) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AttendeeReport) GetAttendee() Attendee {
	if a == nil {
		return Attendee{}
	}
	return a.Attendee
}

func (a *AttendeeReport) GetProfiles() []ScreeningProfile {
	if a == nil {
		return []ScreeningProfile{}
	}
	return a.Profiles
}

func (a *AttendeeReport) GetFindings() []ScreeningFinding {
	if a == nil {
		return []ScreeningFinding{}
	}
	return a.Findings
}

func (a *AttendeeReport) GetPosts() []ScreeningPost {
	if a == nil {
		return []ScreeningPost{}
	}
	return a.Posts
}

func (a *AttendeeReport) GetMedia() []ScreeningMedia {
	if a == nil {
		return []ScreeningMedia{}
	}
	return a.Media
}

func (a *AttendeeReport) GetReviews() []AttendeeReview {
	if a == nil {
		return []AttendeeReview{}
	}
	return a.Reviews
}

func (a *AttendeeReport) GetLookup() *ScreeningLookup {
	if a == nil {
		return nil
	}
	return a.Lookup
}

func (a *AttendeeReport) GetOtherEvents() OtherEvents {
	if a == nil {
		return OtherEvents{}
	}
	return a.OtherEvents
}
