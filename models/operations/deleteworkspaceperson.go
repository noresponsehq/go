package operations

type DeleteWorkspacePersonRequest struct {
	PersonID string `pathParam:"style=simple,explode=false,name=person_id"`
}

func (d *DeleteWorkspacePersonRequest) GetPersonID() string {
	if d == nil {
		return ""
	}
	return d.PersonID
}

type DeleteWorkspacePersonResponse struct {
	Headers map[string][]string
}

func (d *DeleteWorkspacePersonResponse) GetHeaders() map[string][]string {
	if d == nil {
		return map[string][]string{}
	}
	return d.Headers
}
