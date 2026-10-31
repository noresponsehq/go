package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type GetUsageRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The first UTC day, as YYYY-MM-DD.
	From *string `queryParam:"style=form,explode=true,name=from"`
	// The last UTC day, included, as YYYY-MM-DD; no later than today.
	To *string `queryParam:"style=form,explode=true,name=to"`
}

func (g *GetUsageRequest) GetWorkspaceID() string {
	if g == nil {
		return ""
	}
	return g.WorkspaceID
}

func (g *GetUsageRequest) GetFrom() *string {
	if g == nil {
		return nil
	}
	return g.From
}

func (g *GetUsageRequest) GetTo() *string {
	if g == nil {
		return nil
	}
	return g.To
}

type GetUsageResponse struct {
	Headers map[string][]string
	Result  components.Usage
}

func (g *GetUsageResponse) GetHeaders() map[string][]string {
	if g == nil {
		return map[string][]string{}
	}
	return g.Headers
}

func (g *GetUsageResponse) GetResult() components.Usage {
	if g == nil {
		return components.Usage{}
	}
	return g.Result
}
