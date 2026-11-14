package components

import (
	"github.com/noresponsehq/go/v11/internal/utils"
)

// ScreeningCompleted - A screening of a group finished. It counts only the people this screening set out to reach, not who they are: read the roster with GET /v1/workspaces/{workspace_id}/groups/{environment}/{group_id}/people, filtered by band, to act on them. people_screened + partial_coverage = people_total, people_flagged is the sum of by_band, no_footprint counts screened people with no flag whose accounts were not found, and no_profile_to_check those the screening had nothing to look up for on any requested platform. A cancelled screening sends no event.
type ScreeningCompleted struct {
	// The screening that finished.
	ScreeningID string `json:"screening_id"`
	// The group it screened.
	GroupID string `json:"group_id"`
	// completed reached every person; partial left some not fully covered (partial_coverage); failed covered nobody.
	Status ScreeningCompletedStatus `json:"status"`
	// Why a partial or failed screening stopped before reaching everyone, as the screening resource says: expired, worker_attempts_exhausted, attempts_exhausted or screening_unavailable.
	ErrorCode *string `json:"error_code,omitzero"`
	// Every person the screening set out to reach: those no screening had reached and those on watch or escalated.
	PeopleTotal int64 `json:"people_total"`
	// People the screening covered, flagged or not.
	PeopleScreened int64 `json:"people_screened"`
	// Screened people with a finding: the sum of by_band.
	PeopleFlagged int64 `json:"people_flagged"`
	// Screened people by the band of their most severe finding.
	ByBand ScreeningCompletedByBand `json:"by_band"`
	// Screened people with no finding whose accounts were not found (match_state none_found). A clean result for someone with no visible accounts is not the same as a clean result for someone whose accounts were read.
	NoFootprint int64 `json:"no_footprint"`
	// Screened people with no profile to check: the list gave nothing to look up on any requested platform, such as no handle, so nobody looked for a profile (no_profile_to_check on the person). Not in no_footprint, and not partial: nothing failed. Their lookups are on their reports.
	NoProfileToCheck *int64 `json:"no_profile_to_check,omitzero"`
	// People the screening did not fully cover. Read each one's coverage on its person.
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

func (s *ScreeningCompleted) GetGroupID() string {
	if s == nil {
		return ""
	}
	return s.GroupID
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

func (s *ScreeningCompleted) GetPeopleTotal() int64 {
	if s == nil {
		return 0
	}
	return s.PeopleTotal
}

func (s *ScreeningCompleted) GetPeopleScreened() int64 {
	if s == nil {
		return 0
	}
	return s.PeopleScreened
}

func (s *ScreeningCompleted) GetPeopleFlagged() int64 {
	if s == nil {
		return 0
	}
	return s.PeopleFlagged
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
