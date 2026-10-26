package components

import (
	"github.com/noresponsehq/go/v8/internal/utils"
)

type Signal struct {
	// Signal kind, for example account_presence or breach_exposure. New kinds may appear; ignore any you do not know.
	Type string `json:"type"`
	// Which submitted identifier a signal concerns. ip is a selector of breach_exposure only.
	Identifier IdentifierKind `json:"identifier"`
	// Per-signal outcome. A source that timed out, was blocked or errored reports blocked, unavailable or inconclusive with a reason and is listed in incomplete_sources; not_found is a definite absence.
	Outcome SignalOutcome `json:"outcome"`
	// Opaque, workspace-scoped projection of the source that produced a signal. It stays stable for your workspace while the projection key is unchanged and reveals no upstream provider, feed, breach table or source record.
	SourceID string `json:"source_id"`
	// Why a signal or source did not produce a definite found/not_found. timed_out means the source missed the request deadline; a continuation may still settle it. deferred means the (source, identifier) pair is async-only and was not attempted on the synchronous path; a continuation settles it.
	Reason *FailureReason `json:"reason,omitzero"`
	// Opaque pointer to evidence NoResponse holds for this signal. It carries no breach value, credential or page body.
	EvidenceRef *string `json:"evidence_ref,omitzero"`
	// When the evidence was first observed.
	ObservedAt *string `json:"observed_at,omitzero"`
	// When the source was last queried for this identifier.
	LastCheckedAt *string `json:"last_checked_at,omitzero"`
	// Freshness class of the evidence behind a signal. unknown whenever no evidence was obtained.
	Freshness Freshness `json:"freshness"`
	// Facts about the number only: line type, carrier and country. The upstream provider, its scores and its verdicts stay internal.
	Enrichment *PhoneEnrichment `json:"enrichment,omitzero"`
}

func (s Signal) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Signal) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (s *Signal) GetType() string {
	if s == nil {
		return ""
	}
	return s.Type
}

func (s *Signal) GetIdentifier() IdentifierKind {
	if s == nil {
		return IdentifierKind("")
	}
	return s.Identifier
}

func (s *Signal) GetOutcome() SignalOutcome {
	if s == nil {
		return SignalOutcome("")
	}
	return s.Outcome
}

func (s *Signal) GetSourceID() string {
	if s == nil {
		return ""
	}
	return s.SourceID
}

func (s *Signal) GetReason() *FailureReason {
	if s == nil {
		return nil
	}
	return s.Reason
}

func (s *Signal) GetEvidenceRef() *string {
	if s == nil {
		return nil
	}
	return s.EvidenceRef
}

func (s *Signal) GetObservedAt() *string {
	if s == nil {
		return nil
	}
	return s.ObservedAt
}

func (s *Signal) GetLastCheckedAt() *string {
	if s == nil {
		return nil
	}
	return s.LastCheckedAt
}

func (s *Signal) GetFreshness() Freshness {
	if s == nil {
		return Freshness("")
	}
	return s.Freshness
}

func (s *Signal) GetEnrichment() *PhoneEnrichment {
	if s == nil {
		return nil
	}
	return s.Enrichment
}
