package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

type AuditStreamRequest struct {
	EndpointID string   `json:"endpoint_id"`
	Actions    []string `json:"actions,omitzero"`
}

func (a AuditStreamRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuditStreamRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AuditStreamRequest) GetEndpointID() string {
	if a == nil {
		return ""
	}
	return a.EndpointID
}

func (a *AuditStreamRequest) GetActions() []string {
	if a == nil {
		return nil
	}
	return a.Actions
}
