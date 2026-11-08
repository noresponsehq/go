package components

type Profile struct {
}

type Finding struct {
}

// AttendeeReport - Everything known about one attendee.
type AttendeeReport struct {
	// An attendee and where their screening and review stand.
	Attendee Attendee `json:"attendee"`
	// The accounts the screening matched, with how each was matched.
	Profiles []Profile `json:"profiles"`
	// What the screening flagged, with its evidence.
	Findings []Finding `json:"findings"`
	// Every recorded decision, oldest first.
	Reviews []AttendeeReview `json:"reviews"`
}

func (a *AttendeeReport) GetAttendee() Attendee {
	if a == nil {
		return Attendee{}
	}
	return a.Attendee
}

func (a *AttendeeReport) GetProfiles() []Profile {
	if a == nil {
		return []Profile{}
	}
	return a.Profiles
}

func (a *AttendeeReport) GetFindings() []Finding {
	if a == nil {
		return []Finding{}
	}
	return a.Findings
}

func (a *AttendeeReport) GetReviews() []AttendeeReview {
	if a == nil {
		return []AttendeeReview{}
	}
	return a.Reviews
}
