package components

import (
	"github.com/noresponsehq/go/v7/internal/utils"
)

// EndpointVerificationEvent - The endpoint.verification event.
type EndpointVerificationEvent struct {
	// The event id, equal to the NR-Webhook-Id header and across redeliveries and every endpoint the event fans out to.
	ID string `json:"id"`
	// When the event was created. Equal across redeliveries; the signing timestamp header is per delivery.
	CreatedAt string `json:"created_at"`
	// Always endpoint.verification.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	type_ string `const:"endpoint.verification" json:"type"`
	// Present, and true, only on an event sent by POST /v1/webhooks/{endpoint_id}/test: its data is a fixed example assessment, not one of the workspace's, so a receiver acknowledges it and applies nothing. Real events have no test member.
	Test *bool                `json:"test,omitzero"`
	Data EndpointVerification `json:"data"`
}

func (e EndpointVerificationEvent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *EndpointVerificationEvent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (e *EndpointVerificationEvent) GetID() string {
	if e == nil {
		return ""
	}
	return e.ID
}

func (e *EndpointVerificationEvent) GetCreatedAt() string {
	if e == nil {
		return ""
	}
	return e.CreatedAt
}

func (e *EndpointVerificationEvent) GetType() string {
	return "endpoint.verification"
}

func (e *EndpointVerificationEvent) GetTest() *bool {
	if e == nil {
		return nil
	}
	return e.Test
}

func (e *EndpointVerificationEvent) GetData() EndpointVerification {
	if e == nil {
		return EndpointVerification{}
	}
	return e.Data
}
