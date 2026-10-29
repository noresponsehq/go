package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

type PolicyHistoryItem struct {
	// One environment's settings: how long its assessments are kept, and for each flow the risk thresholds that decide it and the depth a request that names none is served at, and the sources it refuses. Live and test each have their own. A flow without thresholds of its own uses account_creation's (inherits_account_creation). Every change is recorded in the audit log as settings.updated, with who made it.
	Settings Settings `json:"settings"`
	// Capture time, not a backfilled activation time. Absent only for the current configuration before its first capture.
	RecordedAt   *string `json:"recorded_at,omitzero"`
	ChangeReason *string `json:"change_reason,omitzero"`
	RestoredFrom *int64  `json:"restored_from,omitzero"`
	// Immutable publication provenance remains in policy history after evaluation evidence expires. SHA-256 values identify the frozen manifest, evaluator semantic input, cohort, and canonical result separately.
	Evaluation *Publication `json:"evaluation,omitzero"`
}

func (p PolicyHistoryItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PolicyHistoryItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *PolicyHistoryItem) GetSettings() Settings {
	if p == nil {
		return Settings{}
	}
	return p.Settings
}

func (p *PolicyHistoryItem) GetRecordedAt() *string {
	if p == nil {
		return nil
	}
	return p.RecordedAt
}

func (p *PolicyHistoryItem) GetChangeReason() *string {
	if p == nil {
		return nil
	}
	return p.ChangeReason
}

func (p *PolicyHistoryItem) GetRestoredFrom() *int64 {
	if p == nil {
		return nil
	}
	return p.RestoredFrom
}

func (p *PolicyHistoryItem) GetEvaluation() *Publication {
	if p == nil {
		return nil
	}
	return p.Evaluation
}
