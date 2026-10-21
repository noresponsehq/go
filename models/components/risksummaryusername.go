package components

import (
	"github.com/noresponsehq/go/v2/internal/utils"
)

type RiskSummaryUsername struct {
	// Accounts found for the username.
	AccountsFound *int64 `json:"accounts_found,omitzero"`
	// Breach facts for one selector. found is the one required member: false is a definite 'in no breach table'. Every other member is optional, and absent makes no claim. No member carries a breach value, password or hash, and unknown members are refused. truncated means the value matched more tables or records than the lookup keeps, so absence-based claims (a missing name match, for example) are withheld; a truncated result is always found.
	Breach *BreachSummary `json:"breach,omitzero"`
}

func (r RiskSummaryUsername) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RiskSummaryUsername) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (r *RiskSummaryUsername) GetAccountsFound() *int64 {
	if r == nil {
		return nil
	}
	return r.AccountsFound
}

func (r *RiskSummaryUsername) GetBreach() *BreachSummary {
	if r == nil {
		return nil
	}
	return r.Breach
}
