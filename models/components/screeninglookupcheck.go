package components

import (
	"github.com/noresponsehq/go/v19/internal/utils"
)

// ScreeningLookupCheck - What came of one lane of the lookups for one identifier.
type ScreeningLookupCheck struct {
	Identifier Identifier `json:"identifier"`
	// An evidence lane: accounts (account presence for email and username), phone, breach, domain (the email domain's registration) and ip.
	Lane Lane `json:"lane"`
	// found: a source found the identifier. not_found: every source answered and none found it. partial: the sources that answered found nothing and the others did not answer. not_covered: no source answered. A lookup that failed reads partial or not_covered, not not_found.
	Status ScreeningLookupCheckStatus `json:"status"`
	// Sources asked: none when no source could be, as when the source is stopped or your workspace refuses it.
	Sources int64 `json:"sources"`
	// Sources that gave a definite answer.
	Answered int64 `json:"answered"`
	// Sources that found the identifier.
	Found int64 `json:"found"`
	// Why a signal or source did not produce a definite found/not_found. timed_out means the source missed the request deadline; a continuation may still settle it. deferred means the (source, identifier) pair is async-only and was not attempted on the synchronous path; a continuation settles it.
	Reason *FailureReason `json:"reason,omitzero"`
	// Each source that did not give a definite answer, by connector, with why: the same connector id a record's source names. A source of the lane that no connector covers is counted in sources and not named here.
	Unanswered []ScreeningLookupCheckUnansweredItem `json:"unanswered,omitzero"`
}

func (s ScreeningLookupCheck) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ScreeningLookupCheck) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (s *ScreeningLookupCheck) GetIdentifier() Identifier {
	if s == nil {
		return Identifier("")
	}
	return s.Identifier
}

func (s *ScreeningLookupCheck) GetLane() Lane {
	if s == nil {
		return Lane("")
	}
	return s.Lane
}

func (s *ScreeningLookupCheck) GetStatus() ScreeningLookupCheckStatus {
	if s == nil {
		return ScreeningLookupCheckStatus("")
	}
	return s.Status
}

func (s *ScreeningLookupCheck) GetSources() int64 {
	if s == nil {
		return 0
	}
	return s.Sources
}

func (s *ScreeningLookupCheck) GetAnswered() int64 {
	if s == nil {
		return 0
	}
	return s.Answered
}

func (s *ScreeningLookupCheck) GetFound() int64 {
	if s == nil {
		return 0
	}
	return s.Found
}

func (s *ScreeningLookupCheck) GetReason() *FailureReason {
	if s == nil {
		return nil
	}
	return s.Reason
}

func (s *ScreeningLookupCheck) GetUnanswered() []ScreeningLookupCheckUnansweredItem {
	if s == nil {
		return nil
	}
	return s.Unanswered
}
