package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// AssessmentList - Body of the 200 answer to POST /v1/assessments/query: one page of the calling workspace's unexpired assessments for one account, newest first, each at its latest assessment_version, as its detail; the account's current risk is GET /v1/accounts/{account_id}. Every item's account_id equals the page's. An assessment past its expires_at is not listed. next_cursor is present only when another page exists.
type AssessmentList struct {
	// Version of the body shape.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	schemaVersion string `const:"1" json:"schema_version"`
	// Your own opaque identifier for the account being checked, scoped to your workspace. Use an id you mint, with no direct identifier (email, phone, name or national id) in it. '|' is accepted so an Auth0 user id such as 'auth0|64abc' passes unchanged; '@', '+', '$' and whitespace are refused by the pattern, as is an exact-shape JWT; a hex digest such as a SHA-256 is allowed, because hashing an internal id is a common pseudonymous scheme. A value such as '07700900123' or 'jane.doe.gmail.com' cannot be told apart from an opaque id. It travels in bodies, and in a URL only as the path segment of GET /v1/accounts/{account_id}, its /assessments and its /linked_accounts. An assessment requested by a provider hook at pre_registration, which has no account yet, carries the server-issued account_id pre: followed by its assessment_id.
	AccountID string             `json:"account_id"`
	Items     []AssessmentDetail `json:"items"`
	// Opaque cursor for the next page, bound to your workspace and the account_id it was issued for. Pass it back unchanged in the next query; a tampered, foreign or expired cursor answers 400 invalid_cursor.
	NextCursor *string `json:"next_cursor,omitzero"`
	// Your current retention period, so a consumer knows how far back history reaches. Every item expires no later than its created_at plus this many days.
	RetentionDays *int64 `default:"365" json:"retention_days"`
}

func (a AssessmentList) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AssessmentList) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AssessmentList) GetSchemaVersion() string {
	return "1"
}

func (a *AssessmentList) GetAccountID() string {
	if a == nil {
		return ""
	}
	return a.AccountID
}

func (a *AssessmentList) GetItems() []AssessmentDetail {
	if a == nil {
		return []AssessmentDetail{}
	}
	return a.Items
}

func (a *AssessmentList) GetNextCursor() *string {
	if a == nil {
		return nil
	}
	return a.NextCursor
}

func (a *AssessmentList) GetRetentionDays() *int64 {
	if a == nil {
		return nil
	}
	return a.RetentionDays
}
