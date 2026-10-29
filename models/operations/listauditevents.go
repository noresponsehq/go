package operations

import (
	"github.com/noresponsehq/go/v9/internal/utils"
	"github.com/noresponsehq/go/v9/models/components"
)

type ListAuditEventsRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// Events per page.
	Limit *int64 `default:"20" queryParam:"style=form,explode=true,name=limit"`
	// next_cursor from the previous page, unchanged; keep every filter unchanged.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Affected scope. API keys can request only their own environment. Omit to read the scopes the caller is allowed.
	Environment *components.Environment `queryParam:"style=form,explode=true,name=environment"`
	// Exact action, such as key.created or account.deleted.
	Action *string `queryParam:"style=form,explode=true,name=action"`
	// Exact external user id, public API key id or console credential id.
	Actor *string `queryParam:"style=form,explode=true,name=actor"`
	// Exact target kind, such as key, member, setting or workspace.
	TargetType *string `queryParam:"style=form,explode=true,name=target_type"`
	// Public target identifier. Audit targets exclude subject identifiers.
	TargetID *string `queryParam:"style=form,explode=true,name=target_id"`
	// Exact outcome.
	Outcome *components.CustomerExportRequestFilterOutcome `queryParam:"style=form,explode=true,name=outcome"`
	// Inclusive creation timestamp in RFC3339 format.
	Since *string `queryParam:"style=form,explode=true,name=since"`
	// Exclusive creation timestamp in RFC3339 format; must follow since when both are supplied.
	Until *string `queryParam:"style=form,explode=true,name=until"`
}

func (l ListAuditEventsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAuditEventsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListAuditEventsRequest) GetWorkspaceID() string {
	if l == nil {
		return ""
	}
	return l.WorkspaceID
}

func (l *ListAuditEventsRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListAuditEventsRequest) GetCursor() *string {
	if l == nil {
		return nil
	}
	return l.Cursor
}

func (l *ListAuditEventsRequest) GetEnvironment() *components.Environment {
	if l == nil {
		return nil
	}
	return l.Environment
}

func (l *ListAuditEventsRequest) GetAction() *string {
	if l == nil {
		return nil
	}
	return l.Action
}

func (l *ListAuditEventsRequest) GetActor() *string {
	if l == nil {
		return nil
	}
	return l.Actor
}

func (l *ListAuditEventsRequest) GetTargetType() *string {
	if l == nil {
		return nil
	}
	return l.TargetType
}

func (l *ListAuditEventsRequest) GetTargetID() *string {
	if l == nil {
		return nil
	}
	return l.TargetID
}

func (l *ListAuditEventsRequest) GetOutcome() *components.CustomerExportRequestFilterOutcome {
	if l == nil {
		return nil
	}
	return l.Outcome
}

func (l *ListAuditEventsRequest) GetSince() *string {
	if l == nil {
		return nil
	}
	return l.Since
}

func (l *ListAuditEventsRequest) GetUntil() *string {
	if l == nil {
		return nil
	}
	return l.Until
}

type ListAuditEventsResponse struct {
	Headers map[string][]string
	Result  components.AuditEventList

	Next func() (*ListAuditEventsResponse, error)
}

func (l *ListAuditEventsResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListAuditEventsResponse) GetResult() components.AuditEventList {
	if l == nil {
		return components.AuditEventList{}
	}
	return l.Result
}
