package operations

import (
	"github.com/noresponsehq/go/v14/models/components"
)

type GetInvestigationRequest struct {
	InvestigationID string `pathParam:"style=simple,explode=false,name=investigation_id"`
}

func (g *GetInvestigationRequest) GetInvestigationID() string {
	if g == nil {
		return ""
	}
	return g.InvestigationID
}

type GetInvestigationResponse struct {
	Headers map[string][]string
	Result  components.InvestigationReport
}

func (g *GetInvestigationResponse) GetHeaders() map[string][]string {
	if g == nil {
		return map[string][]string{}
	}
	return g.Headers
}

func (g *GetInvestigationResponse) GetResult() components.InvestigationReport {
	if g == nil {
		return components.InvestigationReport{}
	}
	return g.Result
}
