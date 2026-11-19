package operations

import (
	"github.com/noresponsehq/go/v18/models/components"
)

type RunWorkspacePersonSearchRequest struct {
	PersonID string                            `pathParam:"style=simple,explode=false,name=person_id"`
	SearchID string                            `pathParam:"style=simple,explode=false,name=search_id"`
	Body     components.ScreeningCancelRequest `request:"mediaType=application/json"`
}

func (r *RunWorkspacePersonSearchRequest) GetPersonID() string {
	if r == nil {
		return ""
	}
	return r.PersonID
}

func (r *RunWorkspacePersonSearchRequest) GetSearchID() string {
	if r == nil {
		return ""
	}
	return r.SearchID
}

func (r *RunWorkspacePersonSearchRequest) GetBody() components.ScreeningCancelRequest {
	if r == nil {
		return components.ScreeningCancelRequest{}
	}
	return r.Body
}

type RunWorkspacePersonSearchResponse struct {
	Headers map[string][]string
	Result  components.Search
}

func (r *RunWorkspacePersonSearchResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *RunWorkspacePersonSearchResponse) GetResult() components.Search {
	if r == nil {
		return components.Search{}
	}
	return r.Result
}
