package components

// AttendeeReviewRequest - Records a decision on an attendee.
type AttendeeReviewRequest struct {
	Decision AttendeeDecision `json:"decision"`
	// Why, for the venue team and later reviewers.
	Note *string `json:"note,omitzero"`
}

func (a *AttendeeReviewRequest) GetDecision() AttendeeDecision {
	if a == nil {
		return AttendeeDecision("")
	}
	return a.Decision
}

func (a *AttendeeReviewRequest) GetNote() *string {
	if a == nil {
		return nil
	}
	return a.Note
}
