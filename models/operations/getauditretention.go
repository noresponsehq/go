package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type GetAuditRetentionRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string                              `pathParam:"style=simple,explode=false,name=workspace_id"`
	Environment components.AuditSettingsEnvironment `pathParam:"style=simple,explode=false,name=environment"`
}

func (g *GetAuditRetentionRequest) GetWorkspaceID() string {
	if g == nil {
		return ""
	}
	return g.WorkspaceID
}

func (g *GetAuditRetentionRequest) GetEnvironment() components.AuditSettingsEnvironment {
	if g == nil {
		return components.AuditSettingsEnvironment("")
	}
	return g.Environment
}

type GetAuditRetentionResponse struct {
	Headers map[string][]string
	Result  components.AuditSettings
}

func (g *GetAuditRetentionResponse) GetHeaders() map[string][]string {
	if g == nil {
		return map[string][]string{}
	}
	return g.Headers
}

func (g *GetAuditRetentionResponse) GetResult() components.AuditSettings {
	if g == nil {
		return components.AuditSettings{}
	}
	return g.Result
}
