package operations

import (
	"github.com/noresponsehq/go/v11/internal/utils"
	"github.com/noresponsehq/go/v11/models/components"
)

type ListGroupsRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	// At most this many items.
	Limit *int64 `default:"20" queryParam:"style=form,explode=true,name=limit"`
	// Read groups older than this one.
	Before *string `queryParam:"style=form,explode=true,name=before"`
}

func (l ListGroupsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListGroupsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListGroupsRequest) GetWorkspaceID() string {
	if l == nil {
		return ""
	}
	return l.WorkspaceID
}

func (l *ListGroupsRequest) GetEnvironment() components.KeyEnvironment {
	if l == nil {
		return components.KeyEnvironment("")
	}
	return l.Environment
}

func (l *ListGroupsRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListGroupsRequest) GetBefore() *string {
	if l == nil {
		return nil
	}
	return l.Before
}

type ListGroupsResponse struct {
	Headers map[string][]string
	Result  components.GroupList
}

func (l *ListGroupsResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListGroupsResponse) GetResult() components.GroupList {
	if l == nil {
		return components.GroupList{}
	}
	return l.Result
}
