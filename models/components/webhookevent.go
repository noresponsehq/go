package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/noresponsehq/go/v10/internal/utils"
)

type WebhookEventType string

const (
	WebhookEventTypeAssessmentUpdated       WebhookEventType = "assessment.updated"
	WebhookEventTypeEndpointVerification    WebhookEventType = "endpoint.verification"
	WebhookEventTypeAccountDeleted          WebhookEventType = "account.deleted"
	WebhookEventTypeWorkspaceDeleted        WebhookEventType = "workspace.deleted"
	WebhookEventTypeWebhookEndpointDisabled WebhookEventType = "webhook_endpoint.disabled"
	WebhookEventTypeAccountRiskChanged      WebhookEventType = "account.risk_changed"
	WebhookEventTypeAuditEvent              WebhookEventType = "audit.event"
)

// WebhookEvent - The body of every webhook delivery: {id, type, created_at, data}, where type chooses the shape of data. Deliveries are at least once and unordered: deduplicate on id, and keep an assessment only when its detail.assessment_version is newer than the one you hold. GET /v1/assessments/{assessment_id} stays authoritative.
type WebhookEvent struct {
	AssessmentUpdatedEvent       *AssessmentUpdatedEvent       `queryParam:"inline" union:"member"`
	EndpointVerificationEvent    *EndpointVerificationEvent    `queryParam:"inline" union:"member"`
	AccountDeletedEvent          *AccountDeletedEvent          `queryParam:"inline" union:"member"`
	WorkspaceDeletedEvent        *WorkspaceDeletedEvent        `queryParam:"inline" union:"member"`
	WebhookEndpointDisabledEvent *WebhookEndpointDisabledEvent `queryParam:"inline" union:"member"`
	AccountRiskChangedEvent      *AccountRiskChangedEvent      `queryParam:"inline" union:"member"`
	AuditEventEvent              *AuditEventEvent              `queryParam:"inline" union:"member"`

	Type WebhookEventType
}

func CreateWebhookEventAssessmentUpdated(assessmentUpdated AssessmentUpdatedEvent) WebhookEvent {
	typ := WebhookEventTypeAssessmentUpdated

	return WebhookEvent{
		AssessmentUpdatedEvent: &assessmentUpdated,
		Type:                   typ,
	}
}

func CreateWebhookEventEndpointVerification(endpointVerification EndpointVerificationEvent) WebhookEvent {
	typ := WebhookEventTypeEndpointVerification

	return WebhookEvent{
		EndpointVerificationEvent: &endpointVerification,
		Type:                      typ,
	}
}

func CreateWebhookEventAccountDeleted(accountDeleted AccountDeletedEvent) WebhookEvent {
	typ := WebhookEventTypeAccountDeleted

	return WebhookEvent{
		AccountDeletedEvent: &accountDeleted,
		Type:                typ,
	}
}

func CreateWebhookEventWorkspaceDeleted(workspaceDeleted WorkspaceDeletedEvent) WebhookEvent {
	typ := WebhookEventTypeWorkspaceDeleted

	return WebhookEvent{
		WorkspaceDeletedEvent: &workspaceDeleted,
		Type:                  typ,
	}
}

func CreateWebhookEventWebhookEndpointDisabled(webhookEndpointDisabled WebhookEndpointDisabledEvent) WebhookEvent {
	typ := WebhookEventTypeWebhookEndpointDisabled

	return WebhookEvent{
		WebhookEndpointDisabledEvent: &webhookEndpointDisabled,
		Type:                         typ,
	}
}

func CreateWebhookEventAccountRiskChanged(accountRiskChanged AccountRiskChangedEvent) WebhookEvent {
	typ := WebhookEventTypeAccountRiskChanged

	return WebhookEvent{
		AccountRiskChangedEvent: &accountRiskChanged,
		Type:                    typ,
	}
}

func CreateWebhookEventAuditEvent(auditEvent AuditEventEvent) WebhookEvent {
	typ := WebhookEventTypeAuditEvent

	return WebhookEvent{
		AuditEventEvent: &auditEvent,
		Type:            typ,
	}
}

func (u *WebhookEvent) UnmarshalJSON(data []byte) (err error) {
	previous := *u
	*u = WebhookEvent{}
	defer func() {
		if err != nil {
			*u = previous
		}
	}()

	type discriminator struct {
		Type string `json:"type"`
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.Type {
	case "assessment.updated":
		assessmentUpdatedEvent := new(AssessmentUpdatedEvent)
		if err := utils.UnmarshalJSON(data, &assessmentUpdatedEvent, "", true, nil); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Type == assessment.updated) type AssessmentUpdatedEvent within WebhookEvent: %w", string(data), err)
		}

		u.AssessmentUpdatedEvent = assessmentUpdatedEvent
		u.Type = WebhookEventTypeAssessmentUpdated
		return nil
	case "endpoint.verification":
		endpointVerificationEvent := new(EndpointVerificationEvent)
		if err := utils.UnmarshalJSON(data, &endpointVerificationEvent, "", true, nil); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Type == endpoint.verification) type EndpointVerificationEvent within WebhookEvent: %w", string(data), err)
		}

		u.EndpointVerificationEvent = endpointVerificationEvent
		u.Type = WebhookEventTypeEndpointVerification
		return nil
	case "account.deleted":
		accountDeletedEvent := new(AccountDeletedEvent)
		if err := utils.UnmarshalJSON(data, &accountDeletedEvent, "", true, nil); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Type == account.deleted) type AccountDeletedEvent within WebhookEvent: %w", string(data), err)
		}

		u.AccountDeletedEvent = accountDeletedEvent
		u.Type = WebhookEventTypeAccountDeleted
		return nil
	case "workspace.deleted":
		workspaceDeletedEvent := new(WorkspaceDeletedEvent)
		if err := utils.UnmarshalJSON(data, &workspaceDeletedEvent, "", true, nil); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Type == workspace.deleted) type WorkspaceDeletedEvent within WebhookEvent: %w", string(data), err)
		}

		u.WorkspaceDeletedEvent = workspaceDeletedEvent
		u.Type = WebhookEventTypeWorkspaceDeleted
		return nil
	case "webhook_endpoint.disabled":
		webhookEndpointDisabledEvent := new(WebhookEndpointDisabledEvent)
		if err := utils.UnmarshalJSON(data, &webhookEndpointDisabledEvent, "", true, nil); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Type == webhook_endpoint.disabled) type WebhookEndpointDisabledEvent within WebhookEvent: %w", string(data), err)
		}

		u.WebhookEndpointDisabledEvent = webhookEndpointDisabledEvent
		u.Type = WebhookEventTypeWebhookEndpointDisabled
		return nil
	case "account.risk_changed":
		accountRiskChangedEvent := new(AccountRiskChangedEvent)
		if err := utils.UnmarshalJSON(data, &accountRiskChangedEvent, "", true, nil); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Type == account.risk_changed) type AccountRiskChangedEvent within WebhookEvent: %w", string(data), err)
		}

		u.AccountRiskChangedEvent = accountRiskChangedEvent
		u.Type = WebhookEventTypeAccountRiskChanged
		return nil
	case "audit.event":
		auditEventEvent := new(AuditEventEvent)
		if err := utils.UnmarshalJSON(data, &auditEventEvent, "", true, nil); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Type == audit.event) type AuditEventEvent within WebhookEvent: %w", string(data), err)
		}

		u.AuditEventEvent = auditEventEvent
		u.Type = WebhookEventTypeAuditEvent
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for WebhookEvent", string(data))
}

func (u WebhookEvent) MarshalJSON() ([]byte, error) {
	if u.AssessmentUpdatedEvent != nil {
		return utils.MarshalJSON(u.AssessmentUpdatedEvent, "", true)
	}

	if u.EndpointVerificationEvent != nil {
		return utils.MarshalJSON(u.EndpointVerificationEvent, "", true)
	}

	if u.AccountDeletedEvent != nil {
		return utils.MarshalJSON(u.AccountDeletedEvent, "", true)
	}

	if u.WorkspaceDeletedEvent != nil {
		return utils.MarshalJSON(u.WorkspaceDeletedEvent, "", true)
	}

	if u.WebhookEndpointDisabledEvent != nil {
		return utils.MarshalJSON(u.WebhookEndpointDisabledEvent, "", true)
	}

	if u.AccountRiskChangedEvent != nil {
		return utils.MarshalJSON(u.AccountRiskChangedEvent, "", true)
	}

	if u.AuditEventEvent != nil {
		return utils.MarshalJSON(u.AuditEventEvent, "", true)
	}

	return nil, errors.New("could not marshal union type WebhookEvent: all fields are null")
}
