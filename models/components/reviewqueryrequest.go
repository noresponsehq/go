package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

type ReviewQueryRequest struct {
	// Only reviews of this exact account in the authenticated workspace and environment. Applied before pagination.
	AccountID           *string       `json:"account_id,omitzero"`
	Status              *ReviewStatus `json:"status,omitzero"`
	Priority            *Priority     `json:"priority,omitzero"`
	Assignee            *string       `json:"assignee,omitzero"`
	Unassigned          *bool         `json:"unassigned,omitzero"`
	AssigneePrincipalID *string       `json:"assignee_principal_id,omitzero"`
	Limit               *int64        `default:"50" json:"limit"`
	Cursor              *string       `json:"cursor,omitzero"`
}

func (r ReviewQueryRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *ReviewQueryRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (r *ReviewQueryRequest) GetAccountID() *string {
	if r == nil {
		return nil
	}
	return r.AccountID
}

func (r *ReviewQueryRequest) GetStatus() *ReviewStatus {
	if r == nil {
		return nil
	}
	return r.Status
}

func (r *ReviewQueryRequest) GetPriority() *Priority {
	if r == nil {
		return nil
	}
	return r.Priority
}

func (r *ReviewQueryRequest) GetAssignee() *string {
	if r == nil {
		return nil
	}
	return r.Assignee
}

func (r *ReviewQueryRequest) GetUnassigned() *bool {
	if r == nil {
		return nil
	}
	return r.Unassigned
}

func (r *ReviewQueryRequest) GetAssigneePrincipalID() *string {
	if r == nil {
		return nil
	}
	return r.AssigneePrincipalID
}

func (r *ReviewQueryRequest) GetLimit() *int64 {
	if r == nil {
		return nil
	}
	return r.Limit
}

func (r *ReviewQueryRequest) GetCursor() *string {
	if r == nil {
		return nil
	}
	return r.Cursor
}
