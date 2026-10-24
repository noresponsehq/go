package components

import (
	"github.com/noresponsehq/go/v7/internal/utils"
)

// BreachSummary - Breach facts for one selector. found is the one required member: false is a definite 'in no breach table'. Every other member is optional, and absent makes no claim. No member carries a breach value, password or hash, and unknown members are refused. truncated means the value matched more tables or records than the lookup keeps, so absence-based claims (a missing name match, for example) are withheld; a truncated result is always found.
type BreachSummary struct {
	Found bool `json:"found"`
	// Distinct breach tables the value appears in.
	BreachCount *int64 `json:"breach_count,omitzero"`
	// Year of the oldest dated table.
	FirstYear *int64 `json:"first_year,omitzero"`
	// Year of the newest dated table.
	LastYear *int64 `json:"last_year,omitzero"`
	// The union of data kinds the matching tables held, without the values.
	Exposed []BreachExposed `json:"exposed,omitzero"`
	// The value is in a credential_list or stealer_log table.
	InCredentialLists *bool `json:"in_credential_lists,omitzero"`
	// A plaintext password sat next to the value in a table dated two years ago or less, or the value is in a stealer_log table.
	RecentPasswordExposure *bool `json:"recent_password_exposure,omitzero"`
	Truncated              *bool `json:"truncated,omitzero"`
}

func (b BreachSummary) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BreachSummary) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (b *BreachSummary) GetFound() bool {
	if b == nil {
		return false
	}
	return b.Found
}

func (b *BreachSummary) GetBreachCount() *int64 {
	if b == nil {
		return nil
	}
	return b.BreachCount
}

func (b *BreachSummary) GetFirstYear() *int64 {
	if b == nil {
		return nil
	}
	return b.FirstYear
}

func (b *BreachSummary) GetLastYear() *int64 {
	if b == nil {
		return nil
	}
	return b.LastYear
}

func (b *BreachSummary) GetExposed() []BreachExposed {
	if b == nil {
		return nil
	}
	return b.Exposed
}

func (b *BreachSummary) GetInCredentialLists() *bool {
	if b == nil {
		return nil
	}
	return b.InCredentialLists
}

func (b *BreachSummary) GetRecentPasswordExposure() *bool {
	if b == nil {
		return nil
	}
	return b.RecentPasswordExposure
}

func (b *BreachSummary) GetTruncated() *bool {
	if b == nil {
		return nil
	}
	return b.Truncated
}
