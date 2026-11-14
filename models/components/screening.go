package components

// Screening - A screening and its progress.
type Screening struct {
	ID          string                   `json:"id"`
	GroupID     string                   `json:"group_id"`
	Environment AuditSettingsEnvironment `json:"environment"`
	// partial finished with some people or platforms not covered; people_failed and people_partial count them.
	Status      ScreeningStatus  `json:"status"`
	Depth       ScreeningDepth   `json:"depth"`
	Platforms   []HandlePlatform `json:"platforms"`
	RequestedBy string           `json:"requested_by"`
	CreatedAt   string           `json:"created_at"`
	StartedAt   *string          `json:"started_at,omitzero"`
	FinishedAt  *string          `json:"finished_at,omitzero"`
	// A screening still running then fails with error_code expired.
	ExpiresAt string `json:"expires_at"`
	// People selected when the screening started.
	PeopleTotal int64 `json:"people_total"`
	// People screened so far, failed ones included.
	PeopleDone    int64 `json:"people_done"`
	PeopleFlagged int64 `json:"people_flagged"`
	PeopleFailed  int64 `json:"people_failed"`
	// People screened with some requested platform not covered; their coverage names which.
	PeoplePartial int64 `json:"people_partial"`
	// What the screening was estimated to cost when it started, in minor units of USD (cents): its people at its depth and platforms. Until the screening finishes, what it has not yet spent counts against the environment's screening_daily_cap.
	EstimatedSpend int64 `json:"estimated_spend"`
	// What the people it has screened recorded, in minor units of USD. Each counts toward the UTC day it was recorded on; a failed person records nothing.
	Spent int64 `json:"spent"`
	// What screening cost, in US cents, in total and per implementation.
	Spend ScreeningSpend `json:"spend"`
	// Why a failed or cancelled screening stopped.
	ErrorCode *string `json:"error_code,omitzero"`
}

func (s *Screening) GetID() string {
	if s == nil {
		return ""
	}
	return s.ID
}

func (s *Screening) GetGroupID() string {
	if s == nil {
		return ""
	}
	return s.GroupID
}

func (s *Screening) GetEnvironment() AuditSettingsEnvironment {
	if s == nil {
		return AuditSettingsEnvironment("")
	}
	return s.Environment
}

func (s *Screening) GetStatus() ScreeningStatus {
	if s == nil {
		return ScreeningStatus("")
	}
	return s.Status
}

func (s *Screening) GetDepth() ScreeningDepth {
	if s == nil {
		return ScreeningDepth("")
	}
	return s.Depth
}

func (s *Screening) GetPlatforms() []HandlePlatform {
	if s == nil {
		return []HandlePlatform{}
	}
	return s.Platforms
}

func (s *Screening) GetRequestedBy() string {
	if s == nil {
		return ""
	}
	return s.RequestedBy
}

func (s *Screening) GetCreatedAt() string {
	if s == nil {
		return ""
	}
	return s.CreatedAt
}

func (s *Screening) GetStartedAt() *string {
	if s == nil {
		return nil
	}
	return s.StartedAt
}

func (s *Screening) GetFinishedAt() *string {
	if s == nil {
		return nil
	}
	return s.FinishedAt
}

func (s *Screening) GetExpiresAt() string {
	if s == nil {
		return ""
	}
	return s.ExpiresAt
}

func (s *Screening) GetPeopleTotal() int64 {
	if s == nil {
		return 0
	}
	return s.PeopleTotal
}

func (s *Screening) GetPeopleDone() int64 {
	if s == nil {
		return 0
	}
	return s.PeopleDone
}

func (s *Screening) GetPeopleFlagged() int64 {
	if s == nil {
		return 0
	}
	return s.PeopleFlagged
}

func (s *Screening) GetPeopleFailed() int64 {
	if s == nil {
		return 0
	}
	return s.PeopleFailed
}

func (s *Screening) GetPeoplePartial() int64 {
	if s == nil {
		return 0
	}
	return s.PeoplePartial
}

func (s *Screening) GetEstimatedSpend() int64 {
	if s == nil {
		return 0
	}
	return s.EstimatedSpend
}

func (s *Screening) GetSpent() int64 {
	if s == nil {
		return 0
	}
	return s.Spent
}

func (s *Screening) GetSpend() ScreeningSpend {
	if s == nil {
		return ScreeningSpend{}
	}
	return s.Spend
}

func (s *Screening) GetErrorCode() *string {
	if s == nil {
		return nil
	}
	return s.ErrorCode
}
