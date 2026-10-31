package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type PublishPolicyEvaluationRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment  components.KeyEnvironment                 `pathParam:"style=simple,explode=false,name=environment"`
	EvaluationID string                                    `pathParam:"style=simple,explode=false,name=evaluation_id"`
	Body         components.PolicyEvaluationPublishRequest `request:"mediaType=application/json"`
}

func (p *PublishPolicyEvaluationRequest) GetWorkspaceID() string {
	if p == nil {
		return ""
	}
	return p.WorkspaceID
}

func (p *PublishPolicyEvaluationRequest) GetEnvironment() components.KeyEnvironment {
	if p == nil {
		return components.KeyEnvironment("")
	}
	return p.Environment
}

func (p *PublishPolicyEvaluationRequest) GetEvaluationID() string {
	if p == nil {
		return ""
	}
	return p.EvaluationID
}

func (p *PublishPolicyEvaluationRequest) GetBody() components.PolicyEvaluationPublishRequest {
	if p == nil {
		return components.PolicyEvaluationPublishRequest{}
	}
	return p.Body
}

type PublishPolicyEvaluationResponse struct {
	Headers map[string][]string
	Result  components.Settings
}

func (p *PublishPolicyEvaluationResponse) GetHeaders() map[string][]string {
	if p == nil {
		return map[string][]string{}
	}
	return p.Headers
}

func (p *PublishPolicyEvaluationResponse) GetResult() components.Settings {
	if p == nil {
		return components.Settings{}
	}
	return p.Result
}
