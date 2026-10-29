package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type RestorePolicyRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment       `pathParam:"style=simple,explode=false,name=environment"`
	Body        components.PolicyRestoreRequest `request:"mediaType=application/json"`
}

func (r *RestorePolicyRequest) GetWorkspaceID() string {
	if r == nil {
		return ""
	}
	return r.WorkspaceID
}

func (r *RestorePolicyRequest) GetEnvironment() components.KeyEnvironment {
	if r == nil {
		return components.KeyEnvironment("")
	}
	return r.Environment
}

func (r *RestorePolicyRequest) GetBody() components.PolicyRestoreRequest {
	if r == nil {
		return components.PolicyRestoreRequest{}
	}
	return r.Body
}

type RestorePolicyResponse struct {
	Headers map[string][]string
	Result  components.Settings
}

func (r *RestorePolicyResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *RestorePolicyResponse) GetResult() components.Settings {
	if r == nil {
		return components.Settings{}
	}
	return r.Result
}
