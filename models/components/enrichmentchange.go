package components

import (
	"encoding/json"
	"fmt"
	"github.com/noresponsehq/go/v15/internal/utils"
	"github.com/noresponsehq/go/v15/optionalnullable"
)

type EnrichmentChangeKind string

const (
	EnrichmentChangeKindUsername EnrichmentChangeKind = "username"
	EnrichmentChangeKindEmail    EnrichmentChangeKind = "email"
	EnrichmentChangeKindPhone    EnrichmentChangeKind = "phone"
)

func (e EnrichmentChangeKind) ToPointer() *EnrichmentChangeKind {
	return &e
}
func (e *EnrichmentChangeKind) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "username":
		fallthrough
	case "email":
		fallthrough
	case "phone":
		*e = EnrichmentChangeKind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EnrichmentChangeKind: %v", v)
	}
}

// EnrichmentChange - A change to enrichment. A field left out keeps the value in force; max_spend or deadline null removes the bound.
type EnrichmentChange struct {
	// off: enrichment starts no search. suggest: enrichment starts no search by itself. automatic: after a workspace person's own investigation, and each time a search on that person finishes, enrichment starts searches by itself.
	Mode *EnrichmentMode `json:"mode,omitzero"`
	// Enrichment searches values only from records more likely than this to be the person's, 0 to 0.99.
	MinLikelihood *float64 `json:"min_likelihood,omitzero"`
	// The kinds of value enrichment searches: distinctive usernames, and email addresses and phone numbers, which belong to one person each.
	Kinds []EnrichmentChangeKind `json:"kinds,omitzero"`
	// The most enrichment may spend on a person from one investigation, in minor units of USD (cents): each search it started after that investigation started counts its estimate or what it spent, whichever is more. At the bound enrichment starts nothing more for that investigation. null: only the environment's screening_daily_cap bounds it.
	MaxSpend optionalnullable.OptionalNullable[int64] `json:"max_spend,omitzero"`
	// How long enrichment may start searches after an investigation starts, 60 to 2592000 seconds. null: no deadline.
	Deadline optionalnullable.OptionalNullable[int64] `json:"deadline,omitzero"`
}

func (e EnrichmentChange) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *EnrichmentChange) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (e *EnrichmentChange) GetMode() *EnrichmentMode {
	if e == nil {
		return nil
	}
	return e.Mode
}

func (e *EnrichmentChange) GetMinLikelihood() *float64 {
	if e == nil {
		return nil
	}
	return e.MinLikelihood
}

func (e *EnrichmentChange) GetKinds() []EnrichmentChangeKind {
	if e == nil {
		return nil
	}
	return e.Kinds
}

func (e *EnrichmentChange) GetMaxSpend() optionalnullable.OptionalNullable[int64] {
	if e == nil {
		return nil
	}
	return e.MaxSpend
}

func (e *EnrichmentChange) GetDeadline() optionalnullable.OptionalNullable[int64] {
	if e == nil {
		return nil
	}
	return e.Deadline
}
