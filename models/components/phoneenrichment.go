package components

import (
	"github.com/noresponsehq/go/v8/internal/utils"
)

// PhoneEnrichment - Facts about the number only: line type, carrier and country. The upstream provider, its scores and its verdicts stay internal.
type PhoneEnrichment struct {
	Fields PhoneEnrichmentFields `json:"fields"`
}

func (p PhoneEnrichment) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PhoneEnrichment) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *PhoneEnrichment) GetFields() PhoneEnrichmentFields {
	if p == nil {
		return PhoneEnrichmentFields{}
	}
	return p.Fields
}
