package operations

import (
	"github.com/noresponsehq/go/v16/internal/utils"
	"github.com/noresponsehq/go/v16/models/components"
)

type ListScreeningsRequest struct {
	GroupID string `pathParam:"style=simple,explode=false,name=group_id"`
	// At most this many items.
	Limit *int64 `default:"20" queryParam:"style=form,explode=true,name=limit"`
	// Read screenings older than this one.
	Before *string `queryParam:"style=form,explode=true,name=before"`
}

func (l ListScreeningsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListScreeningsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListScreeningsRequest) GetGroupID() string {
	if l == nil {
		return ""
	}
	return l.GroupID
}

func (l *ListScreeningsRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListScreeningsRequest) GetBefore() *string {
	if l == nil {
		return nil
	}
	return l.Before
}

type ListScreeningsResponse struct {
	Headers map[string][]string
	Result  components.ScreeningList
}

func (l *ListScreeningsResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListScreeningsResponse) GetResult() components.ScreeningList {
	if l == nil {
		return components.ScreeningList{}
	}
	return l.Result
}
