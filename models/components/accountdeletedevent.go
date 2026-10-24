package components

import (
	"github.com/noresponsehq/go/v6/internal/utils"
)

// AccountDeletedEvent - The account.deleted event.
type AccountDeletedEvent struct {
	// The event id, equal to the NR-Webhook-Id header and across redeliveries and every endpoint the event fans out to.
	ID string `json:"id"`
	// When the event was created. Equal across redeliveries; the signing timestamp header is per delivery.
	CreatedAt string `json:"created_at"`
	// Always account.deleted.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	type_ string `const:"account.deleted" json:"type"`
	// Present, and true, only on an event sent by POST /v1/webhooks/{endpoint_id}/test: its data is a fixed example assessment, not one of the workspace's, so a receiver acknowledges it and applies nothing. Real events have no test member.
	Test *bool `json:"test,omitzero"`
	// Terminal for the account's assessments created at or before deleted_at in the receiving endpoint's environment: delete what you hold from them and ignore later-arriving updates of them. deleted_at is no earlier than the creation of any assessment it erased. An assessment created after it belongs to the account assessed again: keep it, even when this event arrives after it.
	Data AccountDeleted `json:"data"`
}

func (a AccountDeletedEvent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountDeletedEvent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AccountDeletedEvent) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

func (a *AccountDeletedEvent) GetCreatedAt() string {
	if a == nil {
		return ""
	}
	return a.CreatedAt
}

func (a *AccountDeletedEvent) GetType() string {
	return "account.deleted"
}

func (a *AccountDeletedEvent) GetTest() *bool {
	if a == nil {
		return nil
	}
	return a.Test
}

func (a *AccountDeletedEvent) GetData() AccountDeleted {
	if a == nil {
		return AccountDeleted{}
	}
	return a.Data
}
