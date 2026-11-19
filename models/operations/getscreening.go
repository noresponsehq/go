package operations

import (
	"github.com/noresponsehq/go/v19/models/components"
)

type GetScreeningRequest struct {
	GroupID     string `pathParam:"style=simple,explode=false,name=group_id"`
	ScreeningID string `pathParam:"style=simple,explode=false,name=screening_id"`
}

func (g *GetScreeningRequest) GetGroupID() string {
	if g == nil {
		return ""
	}
	return g.GroupID
}

func (g *GetScreeningRequest) GetScreeningID() string {
	if g == nil {
		return ""
	}
	return g.ScreeningID
}

type GetScreeningResponse struct {
	Headers map[string][]string
	Result  components.Screening
}

func (g *GetScreeningResponse) GetHeaders() map[string][]string {
	if g == nil {
		return map[string][]string{}
	}
	return g.Headers
}

func (g *GetScreeningResponse) GetResult() components.Screening {
	if g == nil {
		return components.Screening{}
	}
	return g.Result
}
