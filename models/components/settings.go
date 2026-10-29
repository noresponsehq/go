package components

// Settings - One environment's settings: how long its assessments are kept, and for each flow the risk thresholds that decide it and the depth a request that names none is served at, and the sources it refuses. Live and test each have their own. A flow without thresholds of its own uses account_creation's (inherits_account_creation). Every change is recorded in the audit log as settings.updated, with who made it.
type Settings struct {
	// The workspace id, which is also its live environment's workspace id; the test environment's is the same id with :test appended. A workspace created through the API is tnt_ and a 26-character ULID.
	WorkspaceID string `json:"workspace_id"`
	// The environment the key belongs to. A live key assesses with live sources; a test key answers from the deterministic simulator.
	Environment KeyEnvironment `json:"environment"`
	// Days an assessment is kept after it is created, 1 to 365 (default 365). Shortening hides older assessments at once.
	RetentionDays int64 `json:"retention_days"`
	// The monotonic settings configuration epoch, including retention changes. Zero while defaults apply. Historical assessments keep their frozen policy version; the active version reaches new assessments within a minute.
	PolicyVersion int64 `json:"policy_version"`
	// Who last changed the policy: the member's external user id for a change made through the console, the prefix (nrk_<environment>_<key id>) of the management key that made it, or NoResponse. Absent while the defaults apply.
	UpdatedBy *string `json:"updated_by,omitzero"`
	// When the policy last changed. Absent while the defaults apply.
	UpdatedAt *string       `json:"updated_at,omitzero"`
	Flows     SettingsFlows `json:"flows"`
	// The sources this environment refuses, sorted; empty, the default, refuses none. An assessment leaves a refused source out: its coverage is missing (an identifier with no other source lists coverage-unavailable as unavailable) rather than not_found, and a stored result of a refused source is not reused.
	RefusedSources []SourceID `json:"refused_sources"`
}

func (s *Settings) GetWorkspaceID() string {
	if s == nil {
		return ""
	}
	return s.WorkspaceID
}

func (s *Settings) GetEnvironment() KeyEnvironment {
	if s == nil {
		return KeyEnvironment("")
	}
	return s.Environment
}

func (s *Settings) GetRetentionDays() int64 {
	if s == nil {
		return 0
	}
	return s.RetentionDays
}

func (s *Settings) GetPolicyVersion() int64 {
	if s == nil {
		return 0
	}
	return s.PolicyVersion
}

func (s *Settings) GetUpdatedBy() *string {
	if s == nil {
		return nil
	}
	return s.UpdatedBy
}

func (s *Settings) GetUpdatedAt() *string {
	if s == nil {
		return nil
	}
	return s.UpdatedAt
}

func (s *Settings) GetFlows() SettingsFlows {
	if s == nil {
		return SettingsFlows{}
	}
	return s.Flows
}

func (s *Settings) GetRefusedSources() []SourceID {
	if s == nil {
		return []SourceID{}
	}
	return s.RefusedSources
}
