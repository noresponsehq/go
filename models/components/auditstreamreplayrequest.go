package components

type AuditStreamReplayRequest struct {
	// Opaque, workspace-scoped webhook event id. A redelivery of the same event carries the same event_id, so consumers deduplicate on it.
	EventID string `json:"event_id"`
}

func (a *AuditStreamReplayRequest) GetEventID() string {
	if a == nil {
		return ""
	}
	return a.EventID
}
