package components

type ProfileCitation struct {
	// record: a record in lookup.records or profiles[].record. post: an item in posts. check: an item in lookup.checks.
	Kind ProfileCitationKind `json:"kind"`
	// The record or post id. For a check, c followed by its 1-based position in lookup.checks.
	ID string `json:"id"`
}

func (p *ProfileCitation) GetKind() ProfileCitationKind {
	if p == nil {
		return ProfileCitationKind("")
	}
	return p.Kind
}

func (p *ProfileCitation) GetID() string {
	if p == nil {
		return ""
	}
	return p.ID
}
