package components

import (
	"github.com/noresponsehq/go/v8/internal/utils"
)

// SettingsRequest - PATCH /v1/workspaces/{workspace_id}/settings/{environment} body. Changes only what it names, in one transaction; a policy change is one new policy_version. Needs the admin or owner role, or an API key with the management scope for its own environment.
type SettingsRequest struct {
	RetentionDays *int64 `json:"retention_days,omitzero"`
	// Replaces the refused sources; [] refuses none. A change applies to new assessments within a minute.
	RefusedSources []SourceID `json:"refused_sources,omitzero"`
	// Use account_creation or its deprecated alias signup, not both.
	Flows *SettingsRequestFlows `json:"flows,omitzero"`
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
