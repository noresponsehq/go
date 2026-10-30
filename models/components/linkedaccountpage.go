package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// LinkedAccountPage - Other accounts in this workspace and environment whose retained assessments submitted an email, phone or customer-namespace username that one of account_id's retained assessments also submitted. Matching uses the identifier search rules: exact, canonical and, for email, mailbox. Only customer-submitted identifiers link accounts; source-discovered identifiers do not. Expired, retention-shortened, deleted and erased assessments link nothing. A shared identifier is not verified ownership. Accounts use stable opaque order, the page and total share one database snapshot, and later pages are live. Cursors bind the workspace, environment, operation and account_id, and expire.
type LinkedAccountPage struct {
	// Version of the body shape.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	schemaVersion string `const:"1" json:"schema_version"`
	// The account whose submitted identifiers were compared. It is excluded from items.
	AccountID string          `json:"account_id"`
	Items     []LinkedAccount `json:"items"`
	// Distinct linked accounts, exact below 10000 and capped at 10000.
	Total int64 `json:"total"`
	// True when at least 10000 accounts are linked.
	TotalCapped bool `json:"total_capped"`
	// Opaque, expiring cursor bound to this environment, operation and account_id. Absent on the last page.
	NextCursor *string `json:"next_cursor,omitzero"`
}

func (l LinkedAccountPage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LinkedAccountPage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *LinkedAccountPage) GetSchemaVersion() string {
	return "1"
}

func (l *LinkedAccountPage) GetAccountID() string {
	if l == nil {
		return ""
	}
	return l.AccountID
}

func (l *LinkedAccountPage) GetItems() []LinkedAccount {
	if l == nil {
		return []LinkedAccount{}
	}
	return l.Items
}

func (l *LinkedAccountPage) GetTotal() int64 {
	if l == nil {
		return 0
	}
	return l.Total
}

func (l *LinkedAccountPage) GetTotalCapped() bool {
	if l == nil {
		return false
	}
	return l.TotalCapped
}

func (l *LinkedAccountPage) GetNextCursor() *string {
	if l == nil {
		return nil
	}
	return l.NextCursor
}
