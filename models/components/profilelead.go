package components

type ProfileLead struct {
	// A value copied exactly from a linked record that the person was not searched by.
	Value    string            `json:"value"`
	Kind     string            `json:"kind"`
	Why      string            `json:"why"`
	Evidence []ProfileCitation `json:"evidence"`
}

func (p *ProfileLead) GetValue() string {
	if p == nil {
		return ""
	}
	return p.Value
}

func (p *ProfileLead) GetKind() string {
	if p == nil {
		return ""
	}
	return p.Kind
}

func (p *ProfileLead) GetWhy() string {
	if p == nil {
		return ""
	}
	return p.Why
}

func (p *ProfileLead) GetEvidence() []ProfileCitation {
	if p == nil {
		return []ProfileCitation{}
	}
	return p.Evidence
}
