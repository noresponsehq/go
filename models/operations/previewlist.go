package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type PreviewListRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment     `pathParam:"style=simple,explode=false,name=environment"`
	Body        components.ListPreviewRequest `request:"mediaType=application/json"`
}

func (p *PreviewListRequest) GetWorkspaceID() string {
	if p == nil {
		return ""
	}
	return p.WorkspaceID
}

func (p *PreviewListRequest) GetEnvironment() components.KeyEnvironment {
	if p == nil {
		return components.KeyEnvironment("")
	}
	return p.Environment
}

func (p *PreviewListRequest) GetBody() components.ListPreviewRequest {
	if p == nil {
		return components.ListPreviewRequest{}
	}
	return p.Body
}

type PreviewListResponse struct {
	Headers map[string][]string
	Result  components.ListImpact
}

func (p *PreviewListResponse) GetHeaders() map[string][]string {
	if p == nil {
		return map[string][]string{}
	}
	return p.Headers
}

func (p *PreviewListResponse) GetResult() components.ListImpact {
	if p == nil {
		return components.ListImpact{}
	}
	return p.Result
}
