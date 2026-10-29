package components

// AuditEventList - GET /v1/workspaces/{workspace_id}/audit_events: one page of the workspace's audit log, both environments, newest first. Events are kept for as long as the workspace exists and are deleted with it. next_cursor is present only when another page exists.
type AuditEventList struct {
	// The workspace id, which is also its live environment's workspace id; the test environment's is the same id with :test appended. A workspace created through the API is tnt_ and a 26-character ULID.
	WorkspaceID string       `json:"workspace_id"`
	Items       []AuditEvent `json:"items"`
	// Opaque cursor for the next page, MAC'd by the server and bound to the workspace it was issued for. Pass it back unchanged as cursor; a tampered, foreign or expired cursor answers 400 invalid_cursor.
	NextCursor *string `json:"next_cursor,omitzero"`
}

func (a *AuditEventList) GetWorkspaceID() string {
	if a == nil {
		return ""
	}
	return a.WorkspaceID
}

func (a *AuditEventList) GetItems() []AuditEvent {
	if a == nil {
		return []AuditEvent{}
	}
	return a.Items
}

func (a *AuditEventList) GetNextCursor() *string {
	if a == nil {
		return nil
	}
	return a.NextCursor
}
