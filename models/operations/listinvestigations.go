package operations

import (
	"github.com/noresponsehq/go/v19/internal/utils"
	"github.com/noresponsehq/go/v19/models/components"
)

type ListInvestigationsRequest struct {
	PersonID *string `queryParam:"style=form,explode=true,name=person_id"`
	Before   *string `queryParam:"style=form,explode=true,name=before"`
	Limit    *int64  `default:"100" queryParam:"style=form,explode=true,name=limit"`
}

func (l ListInvestigationsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListInvestigationsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListInvestigationsRequest) GetPersonID() *string {
	if l == nil {
		return nil
	}
	return l.PersonID
}

func (l *ListInvestigationsRequest) GetBefore() *string {
	if l == nil {
		return nil
	}
	return l.Before
}

func (l *ListInvestigationsRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

type ListInvestigationsResponse struct {
	Headers map[string][]string
	Result  components.InvestigationsList
}

func (l *ListInvestigationsResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListInvestigationsResponse) GetResult() components.InvestigationsList {
	if l == nil {
		return components.InvestigationsList{}
	}
	return l.Result
}
