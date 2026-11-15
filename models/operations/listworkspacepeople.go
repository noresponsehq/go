package operations

import (
	"github.com/noresponsehq/go/v14/internal/utils"
	"github.com/noresponsehq/go/v14/models/components"
)

type ListWorkspacePeopleRequest struct {
	Limit *int64 `default:"100" queryParam:"style=form,explode=true,name=limit"`
	// Pass next_cursor from the prior page.
	After *string `queryParam:"style=form,explode=true,name=after"`
}

func (l ListWorkspacePeopleRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListWorkspacePeopleRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListWorkspacePeopleRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListWorkspacePeopleRequest) GetAfter() *string {
	if l == nil {
		return nil
	}
	return l.After
}

type ListWorkspacePeopleResponse struct {
	Headers map[string][]string
	Result  components.WorkspacePeopleList
}

func (l *ListWorkspacePeopleResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListWorkspacePeopleResponse) GetResult() components.WorkspacePeopleList {
	if l == nil {
		return components.WorkspacePeopleList{}
	}
	return l.Result
}
