package components

import (
	"github.com/noresponsehq/go/v10/optionalnullable"
)

type ReviewUpdateRequest struct {
	ExpectedVersion int64                                     `json:"expected_version"`
	Status          *ReviewUpdateRequestStatus                `json:"status,omitzero"`
	Priority        *Priority                                 `json:"priority,omitzero"`
	Assignee        optionalnullable.OptionalNullable[string] `json:"assignee,omitzero"`
}

func (r *ReviewUpdateRequest) GetExpectedVersion() int64 {
	if r == nil {
		return 0
	}
	return r.ExpectedVersion
}

func (r *ReviewUpdateRequest) GetStatus() *ReviewUpdateRequestStatus {
	if r == nil {
		return nil
	}
	return r.Status
}

func (r *ReviewUpdateRequest) GetPriority() *Priority {
	if r == nil {
		return nil
	}
	return r.Priority
}

func (r *ReviewUpdateRequest) GetAssignee() optionalnullable.OptionalNullable[string] {
	if r == nil {
		return nil
	}
	return r.Assignee
}
