package components

// PersonReviewRequest - Records a decision on a person.
type PersonReviewRequest struct {
	Decision PersonDecision `json:"decision"`
	// Why, for the venue team and later reviewers.
	Note *string `json:"note,omitzero"`
}

func (p *PersonReviewRequest) GetDecision() PersonDecision {
	if p == nil {
		return PersonDecision("")
	}
	return p.Decision
}

func (p *PersonReviewRequest) GetNote() *string {
	if p == nil {
		return nil
	}
	return p.Note
}
