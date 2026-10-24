package components

import (
	"github.com/noresponsehq/go/v6/internal/utils"
)

type AssessmentDetailTiming struct {
	// The budget the server applied to the synchronous answer.
	DeadlineMs int64 `json:"deadline_ms"`
	// Server time from HTTP ingress to this revision of the assessment.
	ElapsedMs int64 `json:"elapsed_ms"`
}

func (a AssessmentDetailTiming) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AssessmentDetailTiming) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AssessmentDetailTiming) GetDeadlineMs() int64 {
	if a == nil {
		return 0
	}
	return a.DeadlineMs
}

func (a *AssessmentDetailTiming) GetElapsedMs() int64 {
	if a == nil {
		return 0
	}
	return a.ElapsedMs
}
