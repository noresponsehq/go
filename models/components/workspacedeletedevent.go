package components

import (
	"github.com/noresponsehq/go/v6/internal/utils"
)

// WorkspaceDeletedEvent - The workspace.deleted event.
type WorkspaceDeletedEvent struct {
	// The event id, equal to the NR-Webhook-Id header and across redeliveries and every endpoint the event fans out to.
	ID string `json:"id"`
	// When the event was created. Equal across redeliveries; the signing timestamp header is per delivery.
	CreatedAt string `json:"created_at"`
	// Always workspace.deleted.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	type_ string `const:"workspace.deleted" json:"type"`
	// Present, and true, only on an event sent by POST /v1/webhooks/{endpoint_id}/test: its data is a fixed example assessment, not one of the workspace's, so a receiver acknowledges it and applies nothing. Real events have no test member.
	Test *bool `json:"test,omitzero"`
	// Terminal for the receiving endpoint's environment: its workspace was deleted. Every key of the environment already answers 401, its assessments and accounts are gone, and no further event follows. The workspace id is not reused, so a later event naming it cannot arrive.
	Data WorkspaceDeleted `json:"data"`
}

func (w WorkspaceDeletedEvent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WorkspaceDeletedEvent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (w *WorkspaceDeletedEvent) GetID() string {
	if w == nil {
		return ""
	}
	return w.ID
}

func (w *WorkspaceDeletedEvent) GetCreatedAt() string {
	if w == nil {
		return ""
	}
	return w.CreatedAt
}

func (w *WorkspaceDeletedEvent) GetType() string {
	return "workspace.deleted"
}

func (w *WorkspaceDeletedEvent) GetTest() *bool {
	if w == nil {
		return nil
	}
	return w.Test
}

func (w *WorkspaceDeletedEvent) GetData() WorkspaceDeleted {
	if w == nil {
		return WorkspaceDeleted{}
	}
	return w.Data
}
