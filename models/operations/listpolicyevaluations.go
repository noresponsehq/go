package operations

import (
	"github.com/noresponsehq/go/v9/internal/utils"
	"github.com/noresponsehq/go/v9/models/components"
)

type ListPolicyEvaluationsRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	Limit       *int64                    `default:"20" queryParam:"style=form,explode=true,name=limit"`
	Before      *string                   `queryParam:"style=form,explode=true,name=before"`
}

func (l ListPolicyEvaluationsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListPolicyEvaluationsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListPolicyEvaluationsRequest) GetWorkspaceID() string {
	if l == nil {
		return ""
	}
	return l.WorkspaceID
}

func (l *ListPolicyEvaluationsRequest) GetEnvironment() components.KeyEnvironment {
	if l == nil {
		return components.KeyEnvironment("")
	}
	return l.Environment
}

func (l *ListPolicyEvaluationsRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListPolicyEvaluationsRequest) GetBefore() *string {
	if l == nil {
		return nil
	}
	return l.Before
}

type ListPolicyEvaluationsResponse struct {
	Headers map[string][]string
	Result  components.PolicyEvaluationList

	Next func() (*ListPolicyEvaluationsResponse, error)
}

func (l *ListPolicyEvaluationsResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListPolicyEvaluationsResponse) GetResult() components.PolicyEvaluationList {
	if l == nil {
		return components.PolicyEvaluationList{}
	}
	return l.Result
}
