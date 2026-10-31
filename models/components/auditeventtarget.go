package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

type AuditEventTarget struct {
	// The kind of resource the change acted on.
	Type AuditEventTargetType `json:"type"`
	// The resource's id, when it has one.
	ID *string `json:"id,omitzero"`
}

func (a AuditEventTarget) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuditEventTarget) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AuditEventTarget) GetType() AuditEventTargetType {
	if a == nil {
		return AuditEventTargetType("")
	}
	return a.Type
}

func (a *AuditEventTarget) GetID() *string {
	if a == nil {
		return nil
	}
	return a.ID
}
