package operations

import (
	"encoding/json"
	"fmt"
	"github.com/noresponsehq/go/v10/models/components"
)

type Environment string

const (
	EnvironmentTest Environment = "test"
)

func (e Environment) ToPointer() *Environment {
	return &e
}
func (e *Environment) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "test":
		*e = Environment(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Environment: %v", v)
	}
}

type UpdateAuditRetentionRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string                          `pathParam:"style=simple,explode=false,name=workspace_id"`
	Environment Environment                     `pathParam:"style=simple,explode=false,name=environment"`
	Body        components.AuditSettingsRequest `request:"mediaType=application/json"`
}

func (u *UpdateAuditRetentionRequest) GetWorkspaceID() string {
	if u == nil {
		return ""
	}
	return u.WorkspaceID
}

func (u *UpdateAuditRetentionRequest) GetEnvironment() Environment {
	if u == nil {
		return Environment("")
	}
	return u.Environment
}

func (u *UpdateAuditRetentionRequest) GetBody() components.AuditSettingsRequest {
	if u == nil {
		return components.AuditSettingsRequest{}
	}
	return u.Body
}

type UpdateAuditRetentionResponse struct {
	Headers map[string][]string
	Result  components.AuditSettings
}

func (u *UpdateAuditRetentionResponse) GetHeaders() map[string][]string {
	if u == nil {
		return map[string][]string{}
	}
	return u.Headers
}

func (u *UpdateAuditRetentionResponse) GetResult() components.AuditSettings {
	if u == nil {
		return components.AuditSettings{}
	}
	return u.Result
}
