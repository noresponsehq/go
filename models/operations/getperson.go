package operations

import (
	"github.com/noresponsehq/go/v14/models/components"
)

type GetPersonRequest struct {
	GroupID  string `pathParam:"style=simple,explode=false,name=group_id"`
	PersonID string `pathParam:"style=simple,explode=false,name=person_id"`
}

func (g *GetPersonRequest) GetGroupID() string {
	if g == nil {
		return ""
	}
	return g.GroupID
}

func (g *GetPersonRequest) GetPersonID() string {
	if g == nil {
		return ""
	}
	return g.PersonID
}

type GetPersonResponse struct {
	Headers map[string][]string
	Result  components.PersonReport
}

func (g *GetPersonResponse) GetHeaders() map[string][]string {
	if g == nil {
		return map[string][]string{}
	}
	return g.Headers
}

func (g *GetPersonResponse) GetResult() components.PersonReport {
	if g == nil {
		return components.PersonReport{}
	}
	return g.Result
}
