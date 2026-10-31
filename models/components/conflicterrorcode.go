package components

// ConflictErrorCode - Machine codes an HTTP 409 answer can carry.
type ConflictErrorCode string

const (
	ConflictErrorCodeIdempotencyKeyReused        ConflictErrorCode = "idempotency_key_reused"
	ConflictErrorCodeKeyRotationInProgress       ConflictErrorCode = "key_rotation_in_progress"
	ConflictErrorCodeWebhookEndpointLimitReached ConflictErrorCode = "webhook_endpoint_limit_reached"
	ConflictErrorCodeWebhookEndpointUnverified   ConflictErrorCode = "webhook_endpoint_unverified"
	ConflictErrorCodeTooManyIntegrations         ConflictErrorCode = "too_many_integrations"
	ConflictErrorCodeRiskVersionAhead            ConflictErrorCode = "risk_version_ahead"
	ConflictErrorCodeKeyLimitReached             ConflictErrorCode = "key_limit_reached"
	ConflictErrorCodeSelfRevocation              ConflictErrorCode = "self_revocation"
	ConflictErrorCodeOutcomeConflict             ConflictErrorCode = "outcome_conflict"
	ConflictErrorCodeCredentialConflict          ConflictErrorCode = "credential_conflict"
	ConflictErrorCodeBindingConflict             ConflictErrorCode = "binding_conflict"
	ConflictErrorCodeListEntryExists             ConflictErrorCode = "list_entry_exists"
	ConflictErrorCodeListLimitReached            ConflictErrorCode = "list_limit_reached"
	ConflictErrorCodeListEntryTooBroad           ConflictErrorCode = "list_entry_too_broad"
	ConflictErrorCodePolicyVersionConflict       ConflictErrorCode = "policy_version_conflict"
	ConflictErrorCodeReviewVersionConflict       ConflictErrorCode = "review_version_conflict"
	ConflictErrorCodeReviewExists                ConflictErrorCode = "review_exists"
	ConflictErrorCodeReviewLimitReached          ConflictErrorCode = "review_limit_reached"
	ConflictErrorCodeExportConflict              ConflictErrorCode = "export_conflict"
	ConflictErrorCodeEvaluationConflict          ConflictErrorCode = "evaluation_conflict"
	ConflictErrorCodeEvaluationNoProposal        ConflictErrorCode = "evaluation_no_proposal"
	ConflictErrorCodeEvaluationLimitReached      ConflictErrorCode = "evaluation_limit_reached"
	ConflictErrorCodeAlreadyMember               ConflictErrorCode = "already_member"
	ConflictErrorCodeMemberLimitReached          ConflictErrorCode = "member_limit_reached"
	ConflictErrorCodeOwnerInvitePending          ConflictErrorCode = "owner_invite_pending"
	ConflictErrorCodeAggregateErasedData         ConflictErrorCode = "aggregate_erased_data"
)

func (e ConflictErrorCode) ToPointer() *ConflictErrorCode {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ConflictErrorCode) IsExact() bool {
	if e != nil {
		switch *e {
		case "idempotency_key_reused", "key_rotation_in_progress", "webhook_endpoint_limit_reached", "webhook_endpoint_unverified", "too_many_integrations", "risk_version_ahead", "key_limit_reached", "self_revocation", "outcome_conflict", "credential_conflict", "binding_conflict", "list_entry_exists", "list_limit_reached", "list_entry_too_broad", "policy_version_conflict", "review_version_conflict", "review_exists", "review_limit_reached", "export_conflict", "evaluation_conflict", "evaluation_no_proposal", "evaluation_limit_reached", "already_member", "member_limit_reached", "owner_invite_pending", "aggregate_erased_data":
			return true
		}
	}
	return false
}
