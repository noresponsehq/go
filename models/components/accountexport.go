package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

type Exclusion string

const (
	ExclusionExpiredOrErasedRecords                 Exclusion = "expired_or_erased_records"
	ExclusionSecretsAndCredentialDigests            Exclusion = "secrets_and_credential_digests"
	ExclusionPrivateSourceAndWorkPayloads           Exclusion = "private_source_and_work_payloads"
	ExclusionCrossCustomerData                      Exclusion = "cross_customer_data"
	ExclusionLegacyOutcomeNotes                     Exclusion = "legacy_outcome_notes"
	ExclusionUncapturedHistoricalDecisions          Exclusion = "uncaptured_historical_decisions"
	ExclusionInvestigatorReviewsRequireReviewExport Exclusion = "investigator_reviews_require_review_export"
)

func (e Exclusion) ToPointer() *Exclusion {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Exclusion) IsExact() bool {
	if e != nil {
		switch *e {
		case "expired_or_erased_records", "secrets_and_credential_digests", "private_source_and_work_payloads", "cross_customer_data", "legacy_outcome_notes", "uncaptured_historical_decisions", "investigator_reviews_require_review_export":
			return true
		}
	}
	return false
}

// AccountExport - Retrieve one consistent snapshot of retained customer-accessible records, including assessment revisions, account decisions, feedback, key associations, enforcements and request diagnostics. Secrets, restricted source information, other customers' data and legacy free-text notes are excluded. The same standard/detailed permissions apply. An export exceeding 10000 records or 8 MiB returns export_too_large without a partial result. Historical decisions that were not captured are identified as an exclusion.
type AccountExport struct {
	// Version of the body shape.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	schemaVersion string `const:"1" json:"schema_version"`
	// Your own opaque identifier for the account being checked, scoped to your workspace. Use an id you mint, with no direct identifier (email, phone, name or national id) in it. '|' is accepted so an Auth0 user id such as 'auth0|64abc' passes unchanged; '@', '+', '$' and whitespace are refused by the pattern, as is an exact-shape JWT; a hex digest such as a SHA-256 is allowed, because hashing an internal id is a common pseudonymous scheme. A value such as '07700900123' or 'jane.doe.gmail.com' cannot be told apart from an opaque id. It travels in bodies, and in a URL only as the path segment of GET /v1/accounts/{account_id}, its /assessments and its /linked_accounts. An assessment requested by a provider hook at pre_registration, which has no account yet, carries the server-issued account_id pre: followed by its assessment_id.
	AccountID  string `json:"account_id"`
	ExportedAt string `json:"exported_at"`
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	complete           bool                      `const:"true" json:"complete"`
	CurrentRiskVersion int64                     `json:"current_risk_version"`
	Assessments        []AssessmentRevision      `json:"assessments"`
	Decisions          []AccountDecisionSnapshot `json:"decisions"`
	Outcomes           []AccountOutcome          `json:"outcomes"`
	AssessmentOutcomes []ExportAssessmentOutcome `json:"assessment_outcomes"`
	Credentials        []CredentialAssociation   `json:"credentials"`
	Bindings           []AccountBinding          `json:"bindings"`
	Enforcements       []Enforcement             `json:"enforcements"`
	Requests           []RequestDiagnostic       `json:"requests"`
	Exclusions         []Exclusion               `json:"exclusions"`
}

func (a AccountExport) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountExport) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AccountExport) GetSchemaVersion() string {
	return "1"
}

func (a *AccountExport) GetAccountID() string {
	if a == nil {
		return ""
	}
	return a.AccountID
}

func (a *AccountExport) GetExportedAt() string {
	if a == nil {
		return ""
	}
	return a.ExportedAt
}

func (a *AccountExport) GetComplete() bool {
	return true
}

func (a *AccountExport) GetCurrentRiskVersion() int64 {
	if a == nil {
		return 0
	}
	return a.CurrentRiskVersion
}

func (a *AccountExport) GetAssessments() []AssessmentRevision {
	if a == nil {
		return []AssessmentRevision{}
	}
	return a.Assessments
}

func (a *AccountExport) GetDecisions() []AccountDecisionSnapshot {
	if a == nil {
		return []AccountDecisionSnapshot{}
	}
	return a.Decisions
}

func (a *AccountExport) GetOutcomes() []AccountOutcome {
	if a == nil {
		return []AccountOutcome{}
	}
	return a.Outcomes
}

func (a *AccountExport) GetAssessmentOutcomes() []ExportAssessmentOutcome {
	if a == nil {
		return []ExportAssessmentOutcome{}
	}
	return a.AssessmentOutcomes
}

func (a *AccountExport) GetCredentials() []CredentialAssociation {
	if a == nil {
		return []CredentialAssociation{}
	}
	return a.Credentials
}

func (a *AccountExport) GetBindings() []AccountBinding {
	if a == nil {
		return []AccountBinding{}
	}
	return a.Bindings
}

func (a *AccountExport) GetEnforcements() []Enforcement {
	if a == nil {
		return []Enforcement{}
	}
	return a.Enforcements
}

func (a *AccountExport) GetRequests() []RequestDiagnostic {
	if a == nil {
		return []RequestDiagnostic{}
	}
	return a.Requests
}

func (a *AccountExport) GetExclusions() []Exclusion {
	if a == nil {
		return []Exclusion{}
	}
	return a.Exclusions
}
