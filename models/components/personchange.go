package components

type PersonChange struct {
	ID string `json:"id"`
	// account_added and account_removed: an account or profile appeared or disappeared. breach_added and breach_removed: a breach record appeared or disappeared. identifier_added and identifier_removed: an email, phone, username, name or account id on an account the person already held. field_changed: another field on such an account. posts_added: new posts, counted in count.
	Kind PersonChangeKind `json:"kind"`
	// One rule ranks every change: identity changes (accounts, breaches, identifiers) are high, profile fields are medium, new posts are low.
	Significance Significance `json:"significance"`
	// The connector, site or platform the change came from.
	Source string  `json:"source"`
	Field  *string `json:"field,omitzero"`
	Before *string `json:"before,omitzero"`
	After  *string `json:"after,omitzero"`
	Count  *int64  `json:"count,omitzero"`
	// The investigation that found the change.
	InvestigationID string `json:"investigation_id"`
	// The earlier investigation it was compared with.
	PreviousInvestigationID string  `json:"previous_investigation_id"`
	CreatedAt               string  `json:"created_at"`
	ReviewedAt              *string `json:"reviewed_at"`
	ReviewedBy              *string `json:"reviewed_by,omitzero"`
}

func (p *PersonChange) GetID() string {
	if p == nil {
		return ""
	}
	return p.ID
}

func (p *PersonChange) GetKind() PersonChangeKind {
	if p == nil {
		return PersonChangeKind("")
	}
	return p.Kind
}

func (p *PersonChange) GetSignificance() Significance {
	if p == nil {
		return Significance("")
	}
	return p.Significance
}

func (p *PersonChange) GetSource() string {
	if p == nil {
		return ""
	}
	return p.Source
}

func (p *PersonChange) GetField() *string {
	if p == nil {
		return nil
	}
	return p.Field
}

func (p *PersonChange) GetBefore() *string {
	if p == nil {
		return nil
	}
	return p.Before
}

func (p *PersonChange) GetAfter() *string {
	if p == nil {
		return nil
	}
	return p.After
}

func (p *PersonChange) GetCount() *int64 {
	if p == nil {
		return nil
	}
	return p.Count
}

func (p *PersonChange) GetInvestigationID() string {
	if p == nil {
		return ""
	}
	return p.InvestigationID
}

func (p *PersonChange) GetPreviousInvestigationID() string {
	if p == nil {
		return ""
	}
	return p.PreviousInvestigationID
}

func (p *PersonChange) GetCreatedAt() string {
	if p == nil {
		return ""
	}
	return p.CreatedAt
}

func (p *PersonChange) GetReviewedAt() *string {
	if p == nil {
		return nil
	}
	return p.ReviewedAt
}

func (p *PersonChange) GetReviewedBy() *string {
	if p == nil {
		return nil
	}
	return p.ReviewedBy
}
