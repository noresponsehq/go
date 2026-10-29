package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type PreviewPolicyRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment       `pathParam:"style=simple,explode=false,name=environment"`
	Body        components.PolicyPreviewRequest `request:"mediaType=application/json"`
}

func (p *PreviewPolicyRequest) GetWorkspaceID() string {
	if p == nil {
		return ""
	}
	return p.WorkspaceID
}

func (p *PreviewPolicyRequest) GetEnvironment() components.KeyEnvironment {
	if p == nil {
		return components.KeyEnvironment("")
	}
	return p.Environment
}

func (p *PreviewPolicyRequest) GetBody() components.PolicyPreviewRequest {
	if p == nil {
		return components.PolicyPreviewRequest{}
	}
	return p.Body
}

type PreviewPolicyResponse struct {
	Headers map[string][]string
	Result  components.PolicyPreview
}

func (p *PreviewPolicyResponse) GetHeaders() map[string][]string {
	if p == nil {
		return map[string][]string{}
	}
	return p.Headers
}

func (p *PreviewPolicyResponse) GetResult() components.PolicyPreview {
	if p == nil {
		return components.PolicyPreview{}
	}
	return p.Result
}
