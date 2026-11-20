package components

import (
	"github.com/noresponsehq/go/v20/internal/utils"
)

type PersonBatchItem struct {
	// The row's position in the request.
	Index int64  `json:"index"`
	ID    string `json:"id"`
	// unchanged: the row matched a person with the same details.
	Outcome          PersonBatchItemOutcome `json:"outcome"`
	SuggestedMatches []PersonMatch          `json:"suggested_matches,omitzero"`
}

func (p PersonBatchItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PersonBatchItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *PersonBatchItem) GetIndex() int64 {
	if p == nil {
		return 0
	}
	return p.Index
}

func (p *PersonBatchItem) GetID() string {
	if p == nil {
		return ""
	}
	return p.ID
}

func (p *PersonBatchItem) GetOutcome() PersonBatchItemOutcome {
	if p == nil {
		return PersonBatchItemOutcome("")
	}
	return p.Outcome
}

func (p *PersonBatchItem) GetSuggestedMatches() []PersonMatch {
	if p == nil {
		return nil
	}
	return p.SuggestedMatches
}
