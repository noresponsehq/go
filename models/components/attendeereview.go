package components

// AttendeeReview - One recorded decision.
type AttendeeReview struct {
	ID        string           `json:"id"`
	Decision  AttendeeDecision `json:"decision"`
	Note      *string          `json:"note,omitzero"`
	CreatedBy string           `json:"created_by"`
	CreatedAt string           `json:"created_at"`
}

func (a *AttendeeReview) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

func (a *AttendeeReview) GetDecision() AttendeeDecision {
	if a == nil {
		return AttendeeDecision("")
	}
	return a.Decision
}

func (a *AttendeeReview) GetNote() *string {
	if a == nil {
		return nil
	}
	return a.Note
}

func (a *AttendeeReview) GetCreatedBy() string {
	if a == nil {
		return ""
	}
	return a.CreatedBy
}

func (a *AttendeeReview) GetCreatedAt() string {
	if a == nil {
		return ""
	}
	return a.CreatedAt
}
