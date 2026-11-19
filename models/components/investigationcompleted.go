package components

import (
	"github.com/noresponsehq/go/v17/internal/utils"
)

type InvestigationCompleted struct {
	InvestigationID string                       `json:"investigation_id"`
	PersonID        string                       `json:"person_id"`
	Status          InvestigationCompletedStatus `json:"status"`
	ErrorCode       *string                      `json:"error_code,omitzero"`
	FinishedAt      string                       `json:"finished_at"`
}

func (i InvestigationCompleted) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InvestigationCompleted) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (i *InvestigationCompleted) GetInvestigationID() string {
	if i == nil {
		return ""
	}
	return i.InvestigationID
}

func (i *InvestigationCompleted) GetPersonID() string {
	if i == nil {
		return ""
	}
	return i.PersonID
}

func (i *InvestigationCompleted) GetStatus() InvestigationCompletedStatus {
	if i == nil {
		return InvestigationCompletedStatus("")
	}
	return i.Status
}

func (i *InvestigationCompleted) GetErrorCode() *string {
	if i == nil {
		return nil
	}
	return i.ErrorCode
}

func (i *InvestigationCompleted) GetFinishedAt() string {
	if i == nil {
		return ""
	}
	return i.FinishedAt
}
