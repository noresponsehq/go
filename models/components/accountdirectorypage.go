package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// AccountDirectoryPage - Distinct observed accounts with retained assessments. Typed identifier searches use stable opaque account order, independent of newly submitted assessments; ordinary and legacy email searches use newest-seen order. One repeatable-read database snapshot supplies each page's rows and totals. Pages do not share a snapshot: new matching accounts may appear on a later page or before the cursor and be missed until a fresh search; retention, erasure and changed membership can remove results. Current account decisions are read separately and may advance during a page. Cursors bind the workspace/environment, operation and filters, expire, and cannot cross scopes. Erasure and retention always take precedence over pagination.
type AccountDirectoryPage struct {
	// Version of the body shape.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	schemaVersion string                     `const:"1" json:"schema_version"`
	Items         []AccountDirectoryPageItem `json:"items"`
	// Distinct accounts matching the filters, exact below 10000 and capped at 10000.
	Total int64 `json:"total"`
	// True when at least 10000 accounts match.
	TotalCapped bool `json:"total_capped"`
	// Opaque, expiring cursor bound to this environment, operation and every filter. Keep the filters unchanged.
	NextCursor *string `json:"next_cursor,omitzero"`
}

func (a AccountDirectoryPage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountDirectoryPage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AccountDirectoryPage) GetSchemaVersion() string {
	return "1"
}

func (a *AccountDirectoryPage) GetItems() []AccountDirectoryPageItem {
	if a == nil {
		return []AccountDirectoryPageItem{}
	}
	return a.Items
}

func (a *AccountDirectoryPage) GetTotal() int64 {
	if a == nil {
		return 0
	}
	return a.Total
}

func (a *AccountDirectoryPage) GetTotalCapped() bool {
	if a == nil {
		return false
	}
	return a.TotalCapped
}

func (a *AccountDirectoryPage) GetNextCursor() *string {
	if a == nil {
		return nil
	}
	return a.NextCursor
}
