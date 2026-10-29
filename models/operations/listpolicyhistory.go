package operations

import (
	"github.com/noresponsehq/go/v9/internal/utils"
	"github.com/noresponsehq/go/v9/models/components"
)

type ListPolicyHistoryRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment   components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	BeforeVersion *int64                    `queryParam:"style=form,explode=true,name=before_version"`
	Limit         *int64                    `default:"50" queryParam:"style=form,explode=true,name=limit"`
}

func (l ListPolicyHistoryRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListPolicyHistoryRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListPolicyHistoryRequest) GetWorkspaceID() string {
	if l == nil {
		return ""
	}
	return l.WorkspaceID
}

func (l *ListPolicyHistoryRequest) GetEnvironment() components.KeyEnvironment {
	if l == nil {
		return components.KeyEnvironment("")
	}
	return l.Environment
}

func (l *ListPolicyHistoryRequest) GetBeforeVersion() *int64 {
	if l == nil {
		return nil
	}
	return l.BeforeVersion
}

func (l *ListPolicyHistoryRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

type ListPolicyHistoryResponse struct {
	Headers map[string][]string
	Result  components.PolicyHistory
}

func (l *ListPolicyHistoryResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListPolicyHistoryResponse) GetResult() components.PolicyHistory {
	if l == nil {
		return components.PolicyHistory{}
	}
	return l.Result
}
