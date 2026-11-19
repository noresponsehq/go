package components

import (
	"github.com/noresponsehq/go/v17/internal/utils"
)

// ScreeningLookup - What the signup check's lookups found for a person: account presence for the email, phone and usernames (the person's handles), the phone lookup, the email domain's registration and the breach search. Context and coverage for the reviewer, not a finding: nothing here sets a band or a match state.
type ScreeningLookup struct {
	// Each lane each identifier was looked up in, identifiers in the order email, phone, username and lanes in the order accounts, phone, breach, domain. Every lane the person's identifiers are looked up in is listed: one that no source answered, or that no source could be asked in, is not_covered. Empty when the person has no email, phone or handle to look up.
	Checks []ScreeningLookupCheck `json:"checks"`
	// Each source asked, with what it answered and when. answered_at is absent while the source has not answered.
	Answers []ScreeningLookupAnswer `json:"answers,omitzero"`
	// What each source answered, whole, each answer once: every field of each account a connector found, and every column of each breach row the person's email or usernames matched, under the source's own names, and what each search a person started found, each match naming its search. Anyone who may read the person reads them: no extra permission applies, breach passwords and hashes included.
	Records []Record `json:"records,omitzero"`
	// For each search and source whose matched rows records leave out, how many: a breach search returns at most 20 rows of a table and 100 in all, from its first 64 tables by name. A search a person starts on one table of its source fetches the rest of that table, up to 100 rows. Absent when records hold every matched row.
	Remaining []RecordRemaining `json:"remaining,omitzero"`
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

func (s *ScreeningLookup) GetAnswers() []ScreeningLookupAnswer {
	if s == nil {
		return nil
	}
	return s.Answers
}

func (s *ScreeningLookup) GetRecords() []Record {
	if s == nil {
		return nil
	}
	return s.Records
}

func (s *ScreeningLookup) GetRemaining() []RecordRemaining {
	if s == nil {
		return nil
	}
	return s.Remaining
}

func (s *ScreeningLookup) GetLookedAt() string {
	if s == nil {
		return ""
	}
	return s.LookedAt
}
