package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// WorkspaceDeleted - Terminal for the receiving endpoint's environment: its workspace was deleted. Every key of the environment already answers 401, its assessments and accounts are gone, and no further event follows. The workspace id is not reused, so a later event naming it cannot arrive.
type WorkspaceDeleted struct {
	// The workspace id, which is also its live environment's workspace id; the test environment's is the same id with :test appended. A workspace created through the API is tnt_ and a 26-character ULID.
	WorkspaceID string                   `json:"workspace_id"`
	Environment AuditSettingsEnvironment `json:"environment"`
	// When the deletion took effect, equal to deleted_at in the DELETE answer.
	DeletedAt string `json:"deleted_at"`
}

func (w WorkspaceDeleted) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WorkspaceDeleted) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (w *WorkspaceDeleted) GetWorkspaceID() string {
	if w == nil {
		return ""
	}
	return w.WorkspaceID
}

func (w *WorkspaceDeleted) GetEnvironment() AuditSettingsEnvironment {
	if w == nil {
		return AuditSettingsEnvironment("")
	}
	return w.Environment
}

func (w *WorkspaceDeleted) GetDeletedAt() string {
	if w == nil {
		return ""
	}
	return w.DeletedAt
}
