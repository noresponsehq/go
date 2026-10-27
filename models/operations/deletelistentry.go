package operations

import (
	"github.com/noresponsehq/go/v8/models/components"
)

type DeleteListEntryRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	// The entry's id. An entry of another environment or workspace answers 404.
	ListEntryID string `pathParam:"style=simple,explode=false,name=list_entry_id"`
}

func (d *DeleteListEntryRequest) GetWorkspaceID() string {
	if d == nil {
		return ""
	}
	return d.WorkspaceID
}

func (d *DeleteListEntryRequest) GetEnvironment() components.KeyEnvironment {
	if d == nil {
		return components.KeyEnvironment("")
	}
	return d.Environment
}

func (d *DeleteListEntryRequest) GetListEntryID() string {
	if d == nil {
		return ""
	}
	return d.ListEntryID
}

type DeleteListEntryResponse struct {
	Headers map[string][]string
}

func (d *DeleteListEntryResponse) GetHeaders() map[string][]string {
	if d == nil {
		return map[string][]string{}
	}
	return d.Headers
}
