package operations

import (
	"github.com/noresponsehq/go/v19/models/components"
)

type ReviewPersonRequest struct {
	GroupID  string                         `pathParam:"style=simple,explode=false,name=group_id"`
	PersonID string                         `pathParam:"style=simple,explode=false,name=person_id"`
	Body     components.PersonReviewRequest `request:"mediaType=application/json"`
}

func (r *ReviewPersonRequest) GetGroupID() string {
	if r == nil {
		return ""
	}
	return r.GroupID
}

func (r *ReviewPersonRequest) GetPersonID() string {
	if r == nil {
		return ""
	}
	return r.PersonID
}

func (r *ReviewPersonRequest) GetBody() components.PersonReviewRequest {
	if r == nil {
		return components.PersonReviewRequest{}
	}
	return r.Body
}

type ReviewPersonResponse struct {
	Headers map[string][]string
	Result  components.Person
}

func (r *ReviewPersonResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *ReviewPersonResponse) GetResult() components.Person {
	if r == nil {
		return components.Person{}
	}
	return r.Result
}
