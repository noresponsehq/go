package operations

import (
	"github.com/noresponsehq/go/v10/internal/utils"
	"github.com/noresponsehq/go/v10/models/components"
)

type ListEventsRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	// At most this many items.
	Limit *int64 `default:"20" queryParam:"style=form,explode=true,name=limit"`
	// Read events older than this one.
	Before *string `queryParam:"style=form,explode=true,name=before"`
}

func (l ListEventsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListEventsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListEventsRequest) GetWorkspaceID() string {
	if l == nil {
		return ""
	}
	return l.WorkspaceID
}

func (l *ListEventsRequest) GetEnvironment() components.KeyEnvironment {
	if l == nil {
		return components.KeyEnvironment("")
	}
	return l.Environment
}

func (l *ListEventsRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListEventsRequest) GetBefore() *string {
	if l == nil {
		return nil
	}
	return l.Before
}

type ListEventsResponse struct {
	Headers map[string][]string
	Result  components.EventList
}

func (l *ListEventsResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListEventsResponse) GetResult() components.EventList {
	if l == nil {
		return components.EventList{}
	}
	return l.Result
}
