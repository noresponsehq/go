package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type CreatePolicyEvaluationRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	// Opaque client key, 16 to 128 characters of [A-Za-z0-9._:-] (a UUID or ULID fits). Scoped to the workspace principal and bound to the canonical body for 24 hours.
	IdempotencyKey string                             `header:"style=simple,explode=false,name=Idempotency-Key"`
	Body           components.PolicyEvaluationRequest `request:"mediaType=application/json"`
}

func (c *CreatePolicyEvaluationRequest) GetWorkspaceID() string {
	if c == nil {
		return ""
	}
	return c.WorkspaceID
}

func (c *CreatePolicyEvaluationRequest) GetEnvironment() components.KeyEnvironment {
	if c == nil {
		return components.KeyEnvironment("")
	}
	return c.Environment
}

func (c *CreatePolicyEvaluationRequest) GetIdempotencyKey() string {
	if c == nil {
		return ""
	}
	return c.IdempotencyKey
}

func (c *CreatePolicyEvaluationRequest) GetBody() components.PolicyEvaluationRequest {
	if c == nil {
		return components.PolicyEvaluationRequest{}
	}
	return c.Body
}

type CreatePolicyEvaluationResponse struct {
	Headers map[string][]string
	Result  components.PolicyEvaluation
}

func (c *CreatePolicyEvaluationResponse) GetHeaders() map[string][]string {
	if c == nil {
		return map[string][]string{}
	}
	return c.Headers
}

func (c *CreatePolicyEvaluationResponse) GetResult() components.PolicyEvaluation {
	if c == nil {
		return components.PolicyEvaluation{}
	}
	return c.Result
}
