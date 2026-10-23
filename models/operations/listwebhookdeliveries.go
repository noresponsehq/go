package operations

import (
	"github.com/noresponsehq/go/v3/internal/utils"
	"github.com/noresponsehq/go/v3/models/components"
)

type ListWebhookDeliveriesRequest struct {
	// The endpoint's whe_ id. An id that is unknown, malformed or another tenant's answers 404.
	EndpointID string `pathParam:"style=simple,explode=false,name=endpoint_id"`
	// Attempts per page.
	Limit *int64 `default:"20" queryParam:"style=form,explode=true,name=limit"`
	// next_cursor from the previous page, unchanged.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListWebhookDeliveriesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListWebhookDeliveriesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListWebhookDeliveriesRequest) GetEndpointID() string {
	if l == nil {
		return ""
	}
	return l.EndpointID
}

func (l *ListWebhookDeliveriesRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListWebhookDeliveriesRequest) GetCursor() *string {
	if l == nil {
		return nil
	}
	return l.Cursor
}

type ListWebhookDeliveriesResponse struct {
	Headers map[string][]string
	Result  components.WebhookDeliveryList

	Next func() (*ListWebhookDeliveriesResponse, error)
}

func (l *ListWebhookDeliveriesResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListWebhookDeliveriesResponse) GetResult() components.WebhookDeliveryList {
	if l == nil {
		return components.WebhookDeliveryList{}
	}
	return l.Result
}
