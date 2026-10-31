package operations

import (
	"github.com/noresponsehq/go/v10/internal/utils"
	"github.com/noresponsehq/go/v10/models/components"
)

type ListReviewsRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	Limit       *int64                    `default:"50" queryParam:"style=form,explode=true,name=limit"`
	Cursor      *string                   `queryParam:"style=form,explode=true,name=cursor"`
	Status      *components.ReviewStatus  `queryParam:"style=form,explode=true,name=status"`
	Priority    *components.Priority      `queryParam:"style=form,explode=true,name=priority"`
	Assignee    *string                   `queryParam:"style=form,explode=true,name=assignee"`
	Unassigned  *bool                     `queryParam:"style=form,explode=true,name=unassigned"`
	// Filter by immutable member ID. Mutually exclusive with assignee and unassigned=true.
	AssigneePrincipalID *string `queryParam:"style=form,explode=true,name=assignee_principal_id"`
}

func (l ListReviewsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListReviewsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListReviewsRequest) GetWorkspaceID() string {
	if l == nil {
		return ""
	}
	return l.WorkspaceID
}

func (l *ListReviewsRequest) GetEnvironment() components.KeyEnvironment {
	if l == nil {
		return components.KeyEnvironment("")
	}
	return l.Environment
}

func (l *ListReviewsRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListReviewsRequest) GetCursor() *string {
	if l == nil {
		return nil
	}
	return l.Cursor
}

func (l *ListReviewsRequest) GetStatus() *components.ReviewStatus {
	if l == nil {
		return nil
	}
	return l.Status
}

func (l *ListReviewsRequest) GetPriority() *components.Priority {
	if l == nil {
		return nil
	}
	return l.Priority
}

func (l *ListReviewsRequest) GetAssignee() *string {
	if l == nil {
		return nil
	}
	return l.Assignee
}

func (l *ListReviewsRequest) GetUnassigned() *bool {
	if l == nil {
		return nil
	}
	return l.Unassigned
}

func (l *ListReviewsRequest) GetAssigneePrincipalID() *string {
	if l == nil {
		return nil
	}
	return l.AssigneePrincipalID
}

type ListReviewsResponse struct {
	Headers map[string][]string
	Result  components.ReviewList

	Next func() (*ListReviewsResponse, error)
}

func (l *ListReviewsResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListReviewsResponse) GetResult() components.ReviewList {
	if l == nil {
		return components.ReviewList{}
	}
	return l.Result
}
