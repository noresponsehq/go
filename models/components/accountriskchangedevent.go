package components

import (
	"github.com/noresponsehq/go/v8/internal/utils"
)

// AccountRiskChangedEvent - The account.risk_changed event.
type AccountRiskChangedEvent struct {
	// The event id, equal to the NR-Webhook-Id header and across redeliveries and every endpoint the event fans out to.
	ID string `json:"id"`
	// When the event was created. Equal across redeliveries; the signing timestamp header is per delivery.
	CreatedAt string `json:"created_at"`
	// Always account.risk_changed.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	type_ string `const:"account.risk_changed" json:"type"`
	// Present, and true, only on an event sent by POST /v1/webhooks/{endpoint_id}/test: its data is a fixed example assessment, not one of the workspace's, so a receiver acknowledges it and applies nothing. Real events have no test member.
	Test *bool `json:"test,omitzero"`
	// The account's decision for at least one flow changed. account is the account as GET /v1/accounts/{account_id} answers it for the revision that made the change: its new score, decision, reasons and decisions_by_flow, and the assessment the risk comes from. previous_decisions_by_flow are the decisions it replaced, so account.decisions_by_flow differs from them for at least one flow, and previous_risk_version is the risk_version they were announced at. On the account's first event no decision was announced before: previous_risk_version is absent and previous_decisions_by_flow is allow for every flow, the decisions an account has before its first event. Events of one account are not ordered: account.risk_version orders them, so keep the account with the highest risk_version and ignore an event whose account is older than the one held.
	Data AccountRiskChanged `json:"data"`
}

func (a AccountRiskChangedEvent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountRiskChangedEvent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AccountRiskChangedEvent) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

func (a *AccountRiskChangedEvent) GetCreatedAt() string {
	if a == nil {
		return ""
	}
	return a.CreatedAt
}

func (a *AccountRiskChangedEvent) GetType() string {
	return "account.risk_changed"
}

func (a *AccountRiskChangedEvent) GetTest() *bool {
	if a == nil {
		return nil
	}
	return a.Test
}

func (a *AccountRiskChangedEvent) GetData() AccountRiskChanged {
	if a == nil {
		return AccountRiskChanged{}
	}
	return a.Data
}
