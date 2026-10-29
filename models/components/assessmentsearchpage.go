package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// AssessmentSearchPage - A bounded page of latest assessment revisions, newest creation first, with how many assessments match the filters. Each item is the revision found and its first answer's decision; no current account risk is attached.
type AssessmentSearchPage struct {
	// Version of the body shape.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	schemaVersion string                 `const:"1" json:"schema_version"`
	Items         []AssessmentSearchItem `json:"items"`
	// How many assessments match the filters over every page, counted when this page was read: exact up to 10000. At 10000 or more it is 10000 and total_capped is true. The cursor does not change it; a later page can count differently, as results are live.
	Total int64 `json:"total"`
	// True when at least 10000 assessments match, so total is 10000 and not the exact count.
	TotalCapped bool `json:"total_capped"`
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

func (a *AssessmentSearchPage) GetItems() []AssessmentSearchItem {
	if a == nil {
		return []AssessmentSearchItem{}
	}
	return a.Items
}

func (a *AssessmentSearchPage) GetTotal() int64 {
	if a == nil {
		return 0
	}
	return a.Total
}

func (a *AssessmentSearchPage) GetTotalCapped() bool {
	if a == nil {
		return false
	}
	return a.TotalCapped
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
