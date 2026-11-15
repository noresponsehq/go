package operations

type RemoveGroupPersonRequest struct {
	GroupID  string `pathParam:"style=simple,explode=false,name=group_id"`
	PersonID string `pathParam:"style=simple,explode=false,name=person_id"`
}

func (r *RemoveGroupPersonRequest) GetGroupID() string {
	if r == nil {
		return ""
	}
	return r.GroupID
}

func (r *RemoveGroupPersonRequest) GetPersonID() string {
	if r == nil {
		return ""
	}
	return r.PersonID
}

type RemoveGroupPersonResponse struct {
	Headers map[string][]string
}

func (r *RemoveGroupPersonResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}
