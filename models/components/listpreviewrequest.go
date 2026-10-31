package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// ListPreviewRequest - Choose an addition or a removal. The preview performs no list, account, audit, outcome or outbox write.
type ListPreviewRequest struct {
	// POST /v1/workspaces/{workspace_id}/lists/{environment} body. Adds one entry to the environment's list and, in the same change, re-decides every account it matches: each sends account.risk_changed when a decision changes. Needs the admin or owner role, or a management key of the environment.
	Add    *ListEntryRequest `json:"add,omitzero"`
	Remove *string           `json:"remove,omitzero"`
}

func (l ListPreviewRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListPreviewRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListPreviewRequest) GetAdd() *ListEntryRequest {
	if l == nil {
		return nil
	}
	return l.Add
}

func (l *ListPreviewRequest) GetRemove() *string {
	if l == nil {
		return nil
	}
	return l.Remove
}
