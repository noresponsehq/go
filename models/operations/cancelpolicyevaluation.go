package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type CancelPolicyEvaluationRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment  components.KeyEnvironment                `pathParam:"style=simple,explode=false,name=environment"`
	EvaluationID string                                   `pathParam:"style=simple,explode=false,name=evaluation_id"`
	Body         components.PolicyEvaluationCancelRequest `request:"mediaType=application/json"`
}

func (c *CancelPolicyEvaluationRequest) GetWorkspaceID() string {
	if c == nil {
		return ""
	}
	return c.WorkspaceID
}

func (c *CancelPolicyEvaluationRequest) GetEnvironment() components.KeyEnvironment {
	if c == nil {
		return components.KeyEnvironment("")
	}
	return c.Environment
}

func (c *CancelPolicyEvaluationRequest) GetEvaluationID() string {
	if c == nil {
		return ""
	}
	return c.EvaluationID
}

func (c *CancelPolicyEvaluationRequest) GetBody() components.PolicyEvaluationCancelRequest {
	if c == nil {
		return components.PolicyEvaluationCancelRequest{}
	}
	return c.Body
}

type CancelPolicyEvaluationResponse struct {
	Headers map[string][]string
	Result  components.PolicyEvaluation
}

func (c *CancelPolicyEvaluationResponse) GetHeaders() map[string][]string {
	if c == nil {
		return map[string][]string{}
	}
	return c.Headers
}

func (c *CancelPolicyEvaluationResponse) GetResult() components.PolicyEvaluation {
	if c == nil {
		return components.PolicyEvaluation{}
	}
	return c.Result
}
