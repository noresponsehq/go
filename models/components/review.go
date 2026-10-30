package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

type Review struct {
	ID string `json:"id"`
	// Your own opaque identifier for the account being checked, scoped to your workspace. Use an id you mint, with no direct identifier (email, phone, name or national id) in it. '|' is accepted so an Auth0 user id such as 'auth0|64abc' passes unchanged; '@', '+', '$' and whitespace are refused by the pattern, as is an exact-shape JWT; a hex digest such as a SHA-256 is allowed, because hashing an internal id is a common pseudonymous scheme. A value such as '07700900123' or 'jane.doe.gmail.com' cannot be told apart from an opaque id. It travels in bodies, and in a URL only as the path segment of GET /v1/accounts/{account_id}, its /assessments and its /linked_accounts. An assessment requested by a provider hook at pre_registration, which has no account yet, carries the server-issued account_id pre: followed by its assessment_id.
	AccountID          string       `json:"account_id"`
	Status             ReviewStatus `json:"status"`
	Priority           Priority     `json:"priority"`
	Assignee           *string      `json:"assignee,omitzero"`
	Version            int64        `json:"version"`
	OpenedRiskVersion  int64        `json:"opened_risk_version"`
	CurrentRiskVersion int64        `json:"current_risk_version"`
	CreatedBy          string       `json:"created_by"`
	UpdatedBy          string       `json:"updated_by"`
	CreatedAt          string       `json:"created_at"`
	UpdatedAt          string       `json:"updated_at"`
	ExpiresAt          string       `json:"expires_at"`
	Resolution         *Resolution  `json:"resolution,omitzero"`
	// Immutable workspace member ID. assignee is the historical display identity captured at assignment.
	AssigneePrincipalID *string `json:"assignee_principal_id,omitzero"`
}

func (r Review) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *Review) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (r *Review) GetID() string {
	if r == nil {
		return ""
	}
	return r.ID
}

func (r *Review) GetAccountID() string {
	if r == nil {
		return ""
	}
	return r.AccountID
}

func (r *Review) GetStatus() ReviewStatus {
	if r == nil {
		return ReviewStatus("")
	}
	return r.Status
}

func (r *Review) GetPriority() Priority {
	if r == nil {
		return Priority("")
	}
	return r.Priority
}

func (r *Review) GetAssignee() *string {
	if r == nil {
		return nil
	}
	return r.Assignee
}

func (r *Review) GetVersion() int64 {
	if r == nil {
		return 0
	}
	return r.Version
}

func (r *Review) GetOpenedRiskVersion() int64 {
	if r == nil {
		return 0
	}
	return r.OpenedRiskVersion
}

func (r *Review) GetCurrentRiskVersion() int64 {
	if r == nil {
		return 0
	}
	return r.CurrentRiskVersion
}

func (r *Review) GetCreatedBy() string {
	if r == nil {
		return ""
	}
	return r.CreatedBy
}

func (r *Review) GetUpdatedBy() string {
	if r == nil {
		return ""
	}
	return r.UpdatedBy
}

func (r *Review) GetCreatedAt() string {
	if r == nil {
		return ""
	}
	return r.CreatedAt
}

func (r *Review) GetUpdatedAt() string {
	if r == nil {
		return ""
	}
	return r.UpdatedAt
}

func (r *Review) GetExpiresAt() string {
	if r == nil {
		return ""
	}
	return r.ExpiresAt
}

func (r *Review) GetResolution() *Resolution {
	if r == nil {
		return nil
	}
	return r.Resolution
}

func (r *Review) GetAssigneePrincipalID() *string {
	if r == nil {
		return nil
	}
	return r.AssigneePrincipalID
}
