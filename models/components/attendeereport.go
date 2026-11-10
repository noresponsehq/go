package components

// AttendeeReport - Everything known about one attendee.
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
