package components

import (
	"github.com/noresponsehq/go/v11/internal/utils"
)

// ScreeningLookup - What the signup check's lookups found for a person: account presence for the email, phone and usernames (the person's handles), the phone lookup, the email domain's registration and the breach search. Context and coverage for the reviewer, not a finding: nothing here sets a band or a match state.
type ScreeningLookup struct {
	// Each lane each identifier was looked up in, identifiers in the order email, phone, username and lanes in the order accounts, phone, breach, domain. Every lane the person's identifiers are looked up in is listed: one that no source answered, or that no source could be asked in, is not_covered. Empty when the person has no email, phone or handle to look up.
	Checks []ScreeningLookupCheck `json:"checks"`
	// A few typed facts per identifier: an assessment's risk.summary, and a screened person's lookup summary. Every fact is optional: absent means the evidence did not answer and makes no claim, while false or 0 is a definite answer.
	Summary *RiskSummary `json:"summary,omitzero"`
	// Merged observations behind the summary, served only to a caller who may read evidence: no provider, source, origin, raw breach record or score.
	Details *RiskDetails `json:"details,omitzero"`
	// When the lookups ran.
	LookedAt string `json:"looked_at"`
}

func (s ScreeningLookup) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ScreeningLookup) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (s *ScreeningLookup) GetChecks() []ScreeningLookupCheck {
	if s == nil {
		return []ScreeningLookupCheck{}
	}
	return s.Checks
}

func (s *ScreeningLookup) GetSummary() *RiskSummary {
	if s == nil {
		return nil
	}
	return s.Summary
}

func (s *ScreeningLookup) GetDetails() *RiskDetails {
	if s == nil {
		return nil
	}
	return s.Details
}

func (s *ScreeningLookup) GetLookedAt() string {
	if s == nil {
		return ""
	}
	return s.LookedAt
}
