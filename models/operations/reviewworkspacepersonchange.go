package operations

import (
	"github.com/noresponsehq/go/v20/models/components"
)

type ReviewWorkspacePersonChangeRequest struct {
	PersonID string                            `pathParam:"style=simple,explode=false,name=person_id"`
	ChangeID string                            `pathParam:"style=simple,explode=false,name=change_id"`
	Body     components.ScreeningCancelRequest `request:"mediaType=application/json"`
}

func (r *ReviewWorkspacePersonChangeRequest) GetPersonID() string {
	if r == nil {
		return ""
	}
	return r.PersonID
}

func (r *ReviewWorkspacePersonChangeRequest) GetChangeID() string {
	if r == nil {
		return ""
	}
	return r.ChangeID
}

func (r *ReviewWorkspacePersonChangeRequest) GetBody() components.ScreeningCancelRequest {
	if r == nil {
		return components.ScreeningCancelRequest{}
	}
	return r.Body
}

type ReviewWorkspacePersonChangeResponse struct {
	Headers map[string][]string
	Result  components.PersonChange
}

func (r *ReviewWorkspacePersonChangeResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *ReviewWorkspacePersonChangeResponse) GetResult() components.PersonChange {
	if r == nil {
		return components.PersonChange{}
	}
	return r.Result
}
