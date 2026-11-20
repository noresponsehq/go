package components

type Investigation struct {
	// The denominator of a screening: accounts found, how many could be read, items read and the period they span, so a clean result is distinct from a person nobody could see.
	Coverage ScreeningCoverage `json:"coverage"`
	// What screening cost, in US cents, in total and per implementation.
	Spend      ScreeningSpend      `json:"spend"`
	ID         string              `json:"id"`
	PersonID   string              `json:"person_id"`
	GroupID    *string             `json:"group_id"`
	Status     InvestigationStatus `json:"status"`
	Depth      ScreeningDepth      `json:"depth"`
	Platforms  []string            `json:"platforms"`
	MatchState *MatchState         `json:"match_state"`
	ErrorCode  *string             `json:"error_code,omitzero"`
	CreatedAt  string              `json:"created_at"`
	StartedAt  *string             `json:"started_at,omitzero"`
	FinishedAt *string             `json:"finished_at,omitzero"`
	ExpiresAt  string              `json:"expires_at"`
	DeadlineAt string              `json:"deadline_at"`
}

func (i *Investigation) GetCoverage() ScreeningCoverage {
	if i == nil {
		return ScreeningCoverage{}
	}
	return i.Coverage
}

func (i *Investigation) GetSpend() ScreeningSpend {
	if i == nil {
		return ScreeningSpend{}
	}
	return i.Spend
}

func (i *Investigation) GetID() string {
	if i == nil {
		return ""
	}
	return i.ID
}

func (i *Investigation) GetPersonID() string {
	if i == nil {
		return ""
	}
	return i.PersonID
}

func (i *Investigation) GetGroupID() *string {
	if i == nil {
		return nil
	}
	return i.GroupID
}

func (i *Investigation) GetStatus() InvestigationStatus {
	if i == nil {
		return InvestigationStatus("")
	}
	return i.Status
}

func (i *Investigation) GetDepth() ScreeningDepth {
	if i == nil {
		return ScreeningDepth("")
	}
	return i.Depth
}

func (i *Investigation) GetPlatforms() []string {
	if i == nil {
		return []string{}
	}
	return i.Platforms
}

func (i *Investigation) GetMatchState() *MatchState {
	if i == nil {
		return nil
	}
	return i.MatchState
}

func (i *Investigation) GetErrorCode() *string {
	if i == nil {
		return nil
	}
	return i.ErrorCode
}

func (i *Investigation) GetCreatedAt() string {
	if i == nil {
		return ""
	}
	return i.CreatedAt
}

func (i *Investigation) GetStartedAt() *string {
	if i == nil {
		return nil
	}
	return i.StartedAt
}

func (i *Investigation) GetFinishedAt() *string {
	if i == nil {
		return nil
	}
	return i.FinishedAt
}

func (i *Investigation) GetExpiresAt() string {
	if i == nil {
		return ""
	}
	return i.ExpiresAt
}

func (i *Investigation) GetDeadlineAt() string {
	if i == nil {
		return ""
	}
	return i.DeadlineAt
}
