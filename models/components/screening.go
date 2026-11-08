package components

type ScreeningPlatform string

const (
	ScreeningPlatformInstagram ScreeningPlatform = "instagram"
	ScreeningPlatformX         ScreeningPlatform = "x"
	ScreeningPlatformBluesky   ScreeningPlatform = "bluesky"
	ScreeningPlatformYoutube   ScreeningPlatform = "youtube"
	ScreeningPlatformTiktok    ScreeningPlatform = "tiktok"
	ScreeningPlatformFacebook  ScreeningPlatform = "facebook"
	ScreeningPlatformThreads   ScreeningPlatform = "threads"
	ScreeningPlatformReddit    ScreeningPlatform = "reddit"
	ScreeningPlatformTelegram  ScreeningPlatform = "telegram"
	ScreeningPlatformLinkedin  ScreeningPlatform = "linkedin"
)

func (e ScreeningPlatform) ToPointer() *ScreeningPlatform {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ScreeningPlatform) IsExact() bool {
	if e != nil {
		switch *e {
		case "instagram", "x", "bluesky", "youtube", "tiktok", "facebook", "threads", "reddit", "telegram", "linkedin":
			return true
		}
	}
	return false
}

// Screening - A screening and its progress.
type Screening struct {
	ID          string                   `json:"id"`
	EventID     string                   `json:"event_id"`
	Environment AuditSettingsEnvironment `json:"environment"`
	// partial finished with some attendees not covered; attendees_failed counts them.
	Status      ScreeningStatus     `json:"status"`
	Depth       ScreeningDepth      `json:"depth"`
	Platforms   []ScreeningPlatform `json:"platforms"`
	RequestedBy string              `json:"requested_by"`
	CreatedAt   string              `json:"created_at"`
	StartedAt   *string             `json:"started_at,omitzero"`
	FinishedAt  *string             `json:"finished_at,omitzero"`
	// A screening still running then fails with error_code expired.
	ExpiresAt string `json:"expires_at"`
	// Attendees selected when the screening started.
	AttendeesTotal int64 `json:"attendees_total"`
	// Attendees screened so far, failed ones included.
	AttendeesDone    int64 `json:"attendees_done"`
	AttendeesFlagged int64 `json:"attendees_flagged"`
	AttendeesFailed  int64 `json:"attendees_failed"`
	// Why a failed or cancelled screening stopped.
	ErrorCode *string `json:"error_code,omitzero"`
}

func (s *Screening) GetID() string {
	if s == nil {
		return ""
	}
	return s.ID
}

func (s *Screening) GetEventID() string {
	if s == nil {
		return ""
	}
	return s.EventID
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

func (s *Screening) GetPlatforms() []ScreeningPlatform {
	if s == nil {
		return []ScreeningPlatform{}
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

func (s *Screening) GetAttendeesTotal() int64 {
	if s == nil {
		return 0
	}
	return s.AttendeesTotal
}

func (s *Screening) GetAttendeesDone() int64 {
	if s == nil {
		return 0
	}
	return s.AttendeesDone
}

func (s *Screening) GetAttendeesFlagged() int64 {
	if s == nil {
		return 0
	}
	return s.AttendeesFlagged
}

func (s *Screening) GetAttendeesFailed() int64 {
	if s == nil {
		return 0
	}
	return s.AttendeesFailed
}

func (s *Screening) GetErrorCode() *string {
	if s == nil {
		return nil
	}
	return s.ErrorCode
}
