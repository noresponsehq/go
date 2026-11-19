package operations

import (
	"github.com/noresponsehq/go/v19/models/components"
)

type ScheduleWorkspacePersonRequest struct {
	PersonID string                           `pathParam:"style=simple,explode=false,name=person_id"`
	Body     components.PersonScheduleRequest `request:"mediaType=application/json"`
}

func (s *ScheduleWorkspacePersonRequest) GetPersonID() string {
	if s == nil {
		return ""
	}
	return s.PersonID
}

func (s *ScheduleWorkspacePersonRequest) GetBody() components.PersonScheduleRequest {
	if s == nil {
		return components.PersonScheduleRequest{}
	}
	return s.Body
}

type ScheduleWorkspacePersonResponse struct {
	Headers map[string][]string
	Result  components.PersonSchedule
}

func (s *ScheduleWorkspacePersonResponse) GetHeaders() map[string][]string {
	if s == nil {
		return map[string][]string{}
	}
	return s.Headers
}

func (s *ScheduleWorkspacePersonResponse) GetResult() components.PersonSchedule {
	if s == nil {
		return components.PersonSchedule{}
	}
	return s.Result
}
