package operations

import (
	"github.com/noresponsehq/go/v8/models/components"
)

type UpdateSettingsRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment  `pathParam:"style=simple,explode=false,name=environment"`
	Body        components.SettingsRequest `request:"mediaType=application/json"`
}

func (u *UpdateSettingsRequest) GetWorkspaceID() string {
	if u == nil {
		return ""
	}
	return u.WorkspaceID
}

func (u *UpdateSettingsRequest) GetEnvironment() components.KeyEnvironment {
	if u == nil {
		return components.KeyEnvironment("")
	}
	return u.Environment
}

func (u *UpdateSettingsRequest) GetBody() components.SettingsRequest {
	if u == nil {
		return components.SettingsRequest{}
	}
	return u.Body
}

type UpdateSettingsResponse struct {
	Headers map[string][]string
	Result  components.Settings
}

func (u *UpdateSettingsResponse) GetHeaders() map[string][]string {
	if u == nil {
		return map[string][]string{}
	}
	return u.Headers
}

func (u *UpdateSettingsResponse) GetResult() components.Settings {
	if u == nil {
		return components.Settings{}
	}
	return u.Result
}
