package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// Usage - Body of the 200 answer to GET /v1/workspaces/{workspace_id}/usage: what the workspace's environments were metered for on each UTC day from from through to. Each billable unit is counted once, where it becomes final: an assessment when it is admitted (by flow and depth), a webhook delivery once per event per endpoint (retries and replays add nothing), and a key-usage event once when the ingest accepts it. A count does not change when the assessment, account or workspace is later deleted. A day, environment or metric with nothing counted has no item. No prices or billing. console_lookups is the one limit: the live environment's daily cap on console lookups, as spent today.
type Usage struct {
	// The workspace id, which is also its live environment's workspace id; the test environment's is the same id with :test appended. A workspace created through the API is tnt_ and a 26-character ULID.
	WorkspaceID string `json:"workspace_id"`
	// The first UTC day covered.
	From string `json:"from"`
	// The last UTC day covered, included.
	To string `json:"to"`
	// Ordered by environment (live first), then day, metric, flow and depth.
	Items []UsageCount `json:"items"`
	// The live environment's daily cap on assessments members make through the console (the playground), as spent on the current UTC day. Each console submission with a new Idempotency-Key spends one lookup; a replay spends nothing, and API keys and the test environment spend none. When used reaches limit, the next lookup answers 429 lookup_limit_reached until resets_at. The count is the one the cap enforces.
	ConsoleLookups *ConsoleLookups `json:"console_lookups,omitzero"`
}

func (u Usage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *Usage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (u *Usage) GetWorkspaceID() string {
	if u == nil {
		return ""
	}
	return u.WorkspaceID
}

func (u *Usage) GetFrom() string {
	if u == nil {
		return ""
	}
	return u.From
}

func (u *Usage) GetTo() string {
	if u == nil {
		return ""
	}
	return u.To
}

func (u *Usage) GetItems() []UsageCount {
	if u == nil {
		return []UsageCount{}
	}
	return u.Items
}

func (u *Usage) GetConsoleLookups() *ConsoleLookups {
	if u == nil {
		return nil
	}
	return u.ConsoleLookups
}
