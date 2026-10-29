package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

type RiskDetailsEvent struct {
	Year     int64                `json:"year"`
	Month    *int64               `json:"month,omitzero"`
	Kind     RiskDetailsEventKind `json:"kind"`
	Site     string               `json:"site"`
	Category string               `json:"category"`
	// The source-neutral kind of evidence: account_presence (an account found on a site), auth_provider (the account at your auth provider), breach (breach and credential tables), breach_abuse (abuse and fraud communities' tables) and code_hosting (a code hosting profile).
	Family     FootprintFamily `json:"family"`
	Identifier Identifier      `json:"identifier"`
}

func (r RiskDetailsEvent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RiskDetailsEvent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (r *RiskDetailsEvent) GetYear() int64 {
	if r == nil {
		return 0
	}
	return r.Year
}

func (r *RiskDetailsEvent) GetMonth() *int64 {
	if r == nil {
		return nil
	}
	return r.Month
}

func (r *RiskDetailsEvent) GetKind() RiskDetailsEventKind {
	if r == nil {
		return RiskDetailsEventKind("")
	}
	return r.Kind
}

func (r *RiskDetailsEvent) GetSite() string {
	if r == nil {
		return ""
	}
	return r.Site
}

func (r *RiskDetailsEvent) GetCategory() string {
	if r == nil {
		return ""
	}
	return r.Category
}

func (r *RiskDetailsEvent) GetFamily() FootprintFamily {
	if r == nil {
		return FootprintFamily("")
	}
	return r.Family
}

func (r *RiskDetailsEvent) GetIdentifier() Identifier {
	if r == nil {
		return Identifier("")
	}
	return r.Identifier
}
