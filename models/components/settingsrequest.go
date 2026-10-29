package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// SettingsRequest - Changes only named settings, atomically. Every successful write advances policy_version and retains an immutable settings snapshot. expected_policy_version optionally guards against stale edits; restore always requires it.
type SettingsRequest struct {
	RetentionDays *int64 `json:"retention_days,omitzero"`
	// Replaces the refused sources; [] refuses none. A change applies to new assessments within a minute.
	RefusedSources []SourceID `json:"refused_sources,omitzero"`
	// Use account_creation or its deprecated alias signup, not both.
	Flows *SettingsRequestFlows `json:"flows,omitzero"`
	// Compare-and-set guard. A stale version returns 409 policy_version_conflict before any change; every settings write including retention advances the version.
	ExpectedPolicyVersion *int64  `json:"expected_policy_version,omitzero"`
	ChangeReason          *string `json:"change_reason,omitzero"`
}

func (s SettingsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SettingsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (s *SettingsRequest) GetRetentionDays() *int64 {
	if s == nil {
		return nil
	}
	return s.RetentionDays
}

func (s *SettingsRequest) GetRefusedSources() []SourceID {
	if s == nil {
		return nil
	}
	return s.RefusedSources
}

func (s *SettingsRequest) GetFlows() *SettingsRequestFlows {
	if s == nil {
		return nil
	}
	return s.Flows
}

func (s *SettingsRequest) GetExpectedPolicyVersion() *int64 {
	if s == nil {
		return nil
	}
	return s.ExpectedPolicyVersion
}

func (s *SettingsRequest) GetChangeReason() *string {
	if s == nil {
		return nil
	}
	return s.ChangeReason
}
