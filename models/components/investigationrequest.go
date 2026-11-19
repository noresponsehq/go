package components

import (
	"github.com/noresponsehq/go/v16/internal/utils"
)

// InvestigationRequest - Runs a new investigation against one retained workspace person.
type InvestigationRequest struct {
	Depth ScreeningDepth `json:"depth"`
	// The platforms to search. Leave it out for a full search: every platform a connector serves in the environment when the investigation starts, which the investigation then lists.
	Platforms []string `json:"platforms,omitzero"`
	PersonID  string   `json:"person_id"`
	// A change to enrichment. A field left out keeps the value in force; max_spend or deadline null removes the bound.
	Enrichment *EnrichmentChange `json:"enrichment,omitzero"`
}

func (i InvestigationRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InvestigationRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (i *InvestigationRequest) GetDepth() ScreeningDepth {
	if i == nil {
		return ScreeningDepth("")
	}
	return i.Depth
}

func (i *InvestigationRequest) GetPlatforms() []string {
	if i == nil {
		return nil
	}
	return i.Platforms
}

func (i *InvestigationRequest) GetPersonID() string {
	if i == nil {
		return ""
	}
	return i.PersonID
}

func (i *InvestigationRequest) GetEnrichment() *EnrichmentChange {
	if i == nil {
		return nil
	}
	return i.Enrichment
}
