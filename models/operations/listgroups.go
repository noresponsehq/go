package operations

import (
	"github.com/noresponsehq/go/v20/internal/utils"
	"github.com/noresponsehq/go/v20/models/components"
)

type ListGroupsRequest struct {
	// At most this many items.
	Limit *int64 `default:"20" queryParam:"style=form,explode=true,name=limit"`
	// Read groups older than this one.
	Before *string `queryParam:"style=form,explode=true,name=before"`
}

func (l ListGroupsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListGroupsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListGroupsRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListGroupsRequest) GetBefore() *string {
	if l == nil {
		return nil
	}
	return l.Before
}

type ListGroupsResponse struct {
	Headers map[string][]string
	Result  components.GroupList
}

func (l *ListGroupsResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListGroupsResponse) GetResult() components.GroupList {
	if l == nil {
		return components.GroupList{}
	}
	return l.Result
}
