package components

// PersonReview - One recorded decision.
type PersonReview struct {
	ID        string         `json:"id"`
	Decision  PersonDecision `json:"decision"`
	Note      *string        `json:"note,omitzero"`
	CreatedBy string         `json:"created_by"`
	CreatedAt string         `json:"created_at"`
}

func (p *PersonReview) GetID() string {
	if p == nil {
		return ""
	}
	return p.ID
}

func (p *PersonReview) GetDecision() PersonDecision {
	if p == nil {
		return PersonDecision("")
	}
	return p.Decision
}

func (p *PersonReview) GetNote() *string {
	if p == nil {
		return nil
	}
	return p.Note
}

func (p *PersonReview) GetCreatedBy() string {
	if p == nil {
		return ""
	}
	return p.CreatedBy
}

func (p *PersonReview) GetCreatedAt() string {
	if p == nil {
		return ""
	}
	return p.CreatedAt
}
