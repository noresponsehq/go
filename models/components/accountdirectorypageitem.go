package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

type WaitingFor string

const (
	WaitingForOutcome WaitingFor = "outcome"
	WaitingForReview  WaitingFor = "review"
)

func (e WaitingFor) ToPointer() *WaitingFor {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *WaitingFor) IsExact() bool {
	if e != nil {
		switch *e {
		case "outcome", "review":
			return true
		}
	}
	return false
}

type AccountDirectoryPageItem struct {
	// Present for needs_action searches. outcome means a current review/block decision lacks a retained effective outcome; review means a retained open or in-progress review. Both can apply.
	WaitingFor []WaitingFor `json:"waiting_for,omitzero"`
	// Your own opaque identifier for the account being checked, scoped to your workspace. Use an id you mint, with no direct identifier (email, phone, name or national id) in it. '|' is accepted so an Auth0 user id such as 'auth0|64abc' passes unchanged; '@', '+', '$' and whitespace are refused by the pattern, as is an exact-shape JWT; a hex digest such as a SHA-256 is allowed, because hashing an internal id is a common pseudonymous scheme. A value such as '07700900123' or 'jane.doe.gmail.com' cannot be told apart from an opaque id. It travels in bodies, and in a URL only as the path segment of GET /v1/accounts/{account_id}, its /assessments and its /linked_accounts. An assessment requested by a provider hook at pre_registration, which has no account yet, carries the server-issued account_id pre: followed by its assessment_id.
	AccountID string `json:"account_id"`
	// Retrieve a retained assessment version with its original decision, policy, flow and depth. This response excludes the current account decision. Current retention and response-mode permissions still apply.
	Latest AssessmentRevision `json:"latest"`
	// Body of the 200 answer to GET /v1/accounts/{account_id}: the one current risk of one of the calling workspace's accounts. It is the risk of the account's latest assessment that has a decision (status complete or partial), by created_at and then assessment_id, scored exactly as GET /v1/assessments/{assessment_id} scores that assessment, plus the account's own reason shared_api_key when a key issued to the account was used from many IP addresses or countries (POST /v1/key_usage/batches). Its assessments are the account's history (GET /v1/accounts/{account_id}/assessments).
	Current *Account `json:"current,omitzero"`
	// Count of admitted retained assessments belonging to this account, not just the returned page.
	Assessments int64 `json:"assessments"`
	// Why this result matched a typed identifier search. value is the original value submitted in assessment_id. Provenance records a customer submission, not a source discovery or verified ownership. Equality is deterministic and has no inferred confidence score. An account result names its newest retained matching assessment, which can differ from latest.
	MatchedIdentifier *IdentifierMatch `json:"matched_identifier,omitzero"`
	// Present only for identifier_expression. All satisfied unique terms, ordered by term_index, with original submitted values and retained assessment witnesses. A duplicate normalized term appears once at its first raw leaf index. The legacy single identifier filter instead returns matched_identifier.
	MatchedIdentifiers []IdentifierTermMatch `json:"matched_identifiers,omitzero"`
}

func (a AccountDirectoryPageItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountDirectoryPageItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AccountDirectoryPageItem) GetWaitingFor() []WaitingFor {
	if a == nil {
		return nil
	}
	return a.WaitingFor
}

func (a *AccountDirectoryPageItem) GetAccountID() string {
	if a == nil {
		return ""
	}
	return a.AccountID
}

func (a *AccountDirectoryPageItem) GetLatest() AssessmentRevision {
	if a == nil {
		return AssessmentRevision{}
	}
	return a.Latest
}

func (a *AccountDirectoryPageItem) GetCurrent() *Account {
	if a == nil {
		return nil
	}
	return a.Current
}

func (a *AccountDirectoryPageItem) GetAssessments() int64 {
	if a == nil {
		return 0
	}
	return a.Assessments
}

func (a *AccountDirectoryPageItem) GetMatchedIdentifier() *IdentifierMatch {
	if a == nil {
		return nil
	}
	return a.MatchedIdentifier
}

func (a *AccountDirectoryPageItem) GetMatchedIdentifiers() []IdentifierTermMatch {
	if a == nil {
		return nil
	}
	return a.MatchedIdentifiers
}
