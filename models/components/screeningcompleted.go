package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// ScreeningCompleted - A screening of an event finished. It counts only the attendees this screening set out to reach, not who they are: read the roster with GET /v1/workspaces/{workspace_id}/events/{environment}/{event_id}/attendees, filtered by band, to act on them. attendees_screened + partial_coverage = attendees_total, attendees_flagged is the sum of by_band, no_footprint counts screened attendees with no flag whose accounts were not found, and no_profile_to_check those the screening had nothing to look up for on any requested platform. A cancelled screening sends no event.
type ScreeningCompleted struct {
	// The screening that finished.
	ScreeningID string `json:"screening_id"`
	// The event it screened.
	EventID string `json:"event_id"`
	// completed reached every attendee; partial left some not fully covered (partial_coverage); failed covered nobody.
	Status ScreeningCompletedStatus `json:"status"`
	// Why a partial or failed screening stopped before reaching everyone, as the screening resource says: expired, worker_attempts_exhausted, attempts_exhausted or screening_unavailable.
	ErrorCode *string `json:"error_code,omitzero"`
	// Every attendee the screening set out to reach: those no screening had reached and those on watch or escalated.
	AttendeesTotal int64 `json:"attendees_total"`
	// Attendees the screening covered, flagged or not.
	AttendeesScreened int64 `json:"attendees_screened"`
	// Screened attendees with a finding: the sum of by_band.
	AttendeesFlagged int64 `json:"attendees_flagged"`
	// Screened attendees by the band of their most severe finding.
	ByBand ScreeningCompletedByBand `json:"by_band"`
	// Screened attendees with no finding whose accounts were not found (match_state none_found). A clean result for someone with no visible accounts is not the same as a clean result for someone whose accounts were read.
	NoFootprint int64 `json:"no_footprint"`
	// Screened attendees with no profile to check: the list gave nothing to look up on any requested platform, such as no handle, so nobody looked for a profile (no_profile_to_check on the attendee). Not in no_footprint, and not partial: nothing failed. Their lookups are on their reports.
	NoProfileToCheck *int64 `json:"no_profile_to_check,omitzero"`
	// Attendees the screening did not fully cover. Read each one's coverage on its attendee.
	PartialCoverage int64 `json:"partial_coverage"`
	// When the screening finished; equal to the event's created_at.
	FinishedAt string `json:"finished_at"`
}

func (s ScreeningCompleted) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ScreeningCompleted) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (s *ScreeningCompleted) GetScreeningID() string {
	if s == nil {
		return ""
	}
	return s.ScreeningID
}

func (s *ScreeningCompleted) GetEventID() string {
	if s == nil {
		return ""
	}
	return s.EventID
}

func (s *ScreeningCompleted) GetStatus() ScreeningCompletedStatus {
	if s == nil {
		return ScreeningCompletedStatus("")
	}
	return s.Status
}

func (s *ScreeningCompleted) GetErrorCode() *string {
	if s == nil {
		return nil
	}
	return s.ErrorCode
}

func (s *ScreeningCompleted) GetAttendeesTotal() int64 {
	if s == nil {
		return 0
	}
	return s.AttendeesTotal
}

func (s *ScreeningCompleted) GetAttendeesScreened() int64 {
	if s == nil {
		return 0
	}
	return s.AttendeesScreened
}

func (s *ScreeningCompleted) GetAttendeesFlagged() int64 {
	if s == nil {
		return 0
	}
	return s.AttendeesFlagged
}

func (s *ScreeningCompleted) GetByBand() ScreeningCompletedByBand {
	if s == nil {
		return ScreeningCompletedByBand{}
	}
	return s.ByBand
}

func (s *ScreeningCompleted) GetNoFootprint() int64 {
	if s == nil {
		return 0
	}
	return s.NoFootprint
}

func (s *ScreeningCompleted) GetNoProfileToCheck() *int64 {
	if s == nil {
		return nil
	}
	return s.NoProfileToCheck
}

func (s *ScreeningCompleted) GetPartialCoverage() int64 {
	if s == nil {
		return 0
	}
	return s.PartialCoverage
}

func (s *ScreeningCompleted) GetFinishedAt() string {
	if s == nil {
		return ""
	}
	return s.FinishedAt
}
