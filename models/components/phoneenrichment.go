package components

// PhoneEnrichment - Facts about the number only: line type, carrier and country. The upstream provider, its scores and its verdicts stay internal.
type PhoneEnrichment struct {
	Fields PhoneEnrichmentFields `json:"fields"`
}

func (p *PhoneEnrichment) GetFields() PhoneEnrichmentFields {
	if p == nil {
		return PhoneEnrichmentFields{}
	}
	return p.Fields
}
