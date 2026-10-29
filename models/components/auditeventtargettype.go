package components

// AuditEventTargetType - The kind of resource the change acted on.
type AuditEventTargetType string

const (
	AuditEventTargetTypeWorkspace       AuditEventTargetType = "workspace"
	AuditEventTargetTypeMember          AuditEventTargetType = "member"
	AuditEventTargetTypeInvite          AuditEventTargetType = "invite"
	AuditEventTargetTypeKey             AuditEventTargetType = "key"
	AuditEventTargetTypeSetting         AuditEventTargetType = "setting"
	AuditEventTargetTypeWebhookEndpoint AuditEventTargetType = "webhook_endpoint"
	AuditEventTargetTypeIntegration     AuditEventTargetType = "integration"
	AuditEventTargetTypeBilling         AuditEventTargetType = "billing"
	AuditEventTargetTypeReview          AuditEventTargetType = "review"
)

func (e AuditEventTargetType) ToPointer() *AuditEventTargetType {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *AuditEventTargetType) IsExact() bool {
	if e != nil {
		switch *e {
		case "workspace", "member", "invite", "key", "setting", "webhook_endpoint", "integration", "billing", "review":
			return true
		}
	}
	return false
}
