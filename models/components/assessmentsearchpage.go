package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// AssessmentSearchPage - A bounded page of latest assessment revisions, newest creation first. Each revision identifies the exact result found; no current account risk is attached.
type AssessmentSearchPage struct {
	// Version of the body shape.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	schemaVersion string               `const:"1" json:"schema_version"`
	Items         []AssessmentRevision `json:"items"`
	// Opaque, expiring cursor bound to this environment, operation and every filter. Keep the filters unchanged.
	NextCursor *string `json:"next_cursor,omitzero"`
	// The calling workspace's configured retention period in days. Default 365 (12 months), which is also the maximum: a workspace may only shorten it, to at least 1 day. Each record expires from its own created_at; see assessment expires_at.
	RetentionDays *int64 `default:"365" json:"retention_days"`
}

func (a AssessmentSearchPage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AssessmentSearchPage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AssessmentSearchPage) GetSchemaVersion() string {
	return "1"
}

func (a *AssessmentSearchPage) GetItems() []AssessmentRevision {
	if a == nil {
		return []AssessmentRevision{}
	}
	return a.Items
}

func (a *AssessmentSearchPage) GetNextCursor() *string {
	if a == nil {
		return nil
	}
	return a.NextCursor
}

func (a *AssessmentSearchPage) GetRetentionDays() *int64 {
	if a == nil {
		return nil
	}
	return a.RetentionDays
}
