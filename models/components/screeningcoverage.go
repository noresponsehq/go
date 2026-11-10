package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// ScreeningCoverage - The denominator of a screening: accounts found, how many could be read, items read and the period they span, so a clean result is distinct from a person nobody could see.
type ScreeningCoverage struct {
	AccountsFound      int64 `json:"accounts_found"`
	AccountsAccessible int64 `json:"accounts_accessible"`
	ItemsRead          int64 `json:"items_read"`
	// Items read that the analyser could not analyse. Above zero, the attendee is partial: a finding may be missing from them.
	ItemsNotAnalysed *int64 `json:"items_not_analysed,omitzero"`
	// The oldest item read on any platform.
	PeriodFrom *string `json:"period_from,omitzero"`
	// The newest item read on any platform.
	PeriodTo  *string                     `json:"period_to,omitzero"`
	Platforms []ScreeningPlatformCoverage `json:"platforms"`
	// Requested platforms where the attendee gave the screening nothing to look up, such as no handle for a platform that reads supplied handles only. Nothing looked and nothing failed: such a platform is not in platforms, is not none_found and does not make the attendee partial.
	NotSearched []HandlePlatform `json:"not_searched,omitzero"`
}

func (s ScreeningCoverage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ScreeningCoverage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (s *ScreeningCoverage) GetAccountsFound() int64 {
	if s == nil {
		return 0
	}
	return s.AccountsFound
}

func (s *ScreeningCoverage) GetAccountsAccessible() int64 {
	if s == nil {
		return 0
	}
	return s.AccountsAccessible
}

func (s *ScreeningCoverage) GetItemsRead() int64 {
	if s == nil {
		return 0
	}
	return s.ItemsRead
}

func (s *ScreeningCoverage) GetItemsNotAnalysed() *int64 {
	if s == nil {
		return nil
	}
	return s.ItemsNotAnalysed
}

func (s *ScreeningCoverage) GetPeriodFrom() *string {
	if s == nil {
		return nil
	}
	return s.PeriodFrom
}

func (s *ScreeningCoverage) GetPeriodTo() *string {
	if s == nil {
		return nil
	}
	return s.PeriodTo
}

func (s *ScreeningCoverage) GetPlatforms() []ScreeningPlatformCoverage {
	if s == nil {
		return []ScreeningPlatformCoverage{}
	}
	return s.Platforms
}

func (s *ScreeningCoverage) GetNotSearched() []HandlePlatform {
	if s == nil {
		return nil
	}
	return s.NotSearched
}
