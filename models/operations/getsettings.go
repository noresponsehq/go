package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type GetSettingsRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
}

func (g *GetSettingsRequest) GetWorkspaceID() string {
	if g == nil {
		return ""
	}
	return g.WorkspaceID
}

func (g *GetSettingsRequest) GetEnvironment() components.KeyEnvironment {
	if g == nil {
		return components.KeyEnvironment("")
	}
	return g.Environment
}

type GetSettingsResponse struct {
	Headers map[string][]string
	Result  components.Settings
}

func (g *GetSettingsResponse) GetHeaders() map[string][]string {
	if g == nil {
		return map[string][]string{}
	}
	return g.Headers
}

func (g *GetSettingsResponse) GetResult() components.Settings {
	if g == nil {
		return components.Settings{}
	}
	return g.Result
}
