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
	ConflictErrorCodeListEntryExists             ConflictErrorCode = "list_entry_exists"
	ConflictErrorCodeListLimitReached            ConflictErrorCode = "list_limit_reached"
	ConflictErrorCodeListEntryTooBroad           ConflictErrorCode = "list_entry_too_broad"
)

func (e ConflictErrorCode) ToPointer() *ConflictErrorCode {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ConflictErrorCode) IsExact() bool {
	if e != nil {
		switch *e {
		case "idempotency_key_reused", "key_rotation_in_progress", "webhook_endpoint_limit_reached", "webhook_endpoint_unverified", "too_many_integrations", "risk_version_ahead", "key_limit_reached", "self_revocation", "list_entry_exists", "list_limit_reached", "list_entry_too_broad":
			return true
		}
	}
	return false
}
