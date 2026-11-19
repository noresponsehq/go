package operations

import (
	"github.com/noresponsehq/go/v17/models/components"
)

type RunSearchRequest struct {
	GroupID  string                            `pathParam:"style=simple,explode=false,name=group_id"`
	PersonID string                            `pathParam:"style=simple,explode=false,name=person_id"`
	SearchID string                            `pathParam:"style=simple,explode=false,name=search_id"`
	Body     components.ScreeningCancelRequest `request:"mediaType=application/json"`
}

func (r *RunSearchRequest) GetGroupID() string {
	if r == nil {
		return ""
	}
	return r.GroupID
}

func (r *RunSearchRequest) GetPersonID() string {
	if r == nil {
		return ""
	}
	return r.PersonID
}

func (r *RunSearchRequest) GetSearchID() string {
	if r == nil {
		return ""
	}
	return r.SearchID
}

func (r *RunSearchRequest) GetBody() components.ScreeningCancelRequest {
	if r == nil {
		return components.ScreeningCancelRequest{}
	}
	return r.Body
}

type RunSearchResponse struct {
	Headers map[string][]string
	Result  components.Search
}

func (r *RunSearchResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *RunSearchResponse) GetResult() components.Search {
	if r == nil {
		return components.Search{}
	}
	return r.Result
}
