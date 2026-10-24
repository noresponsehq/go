# WebhookEvent

The body of every webhook delivery: {id, type, created_at, data}, where type chooses the shape of data. Deliveries are at least once and unordered: deduplicate on id, and keep an assessment only when its detail.assessment_version is newer than the one you hold. GET /v1/assessments/{assessment_id} stays authoritative.


## Supported Types

### AssessmentUpdatedEvent

```go
webhookEvent := components.CreateWebhookEventAssessmentUpdated(components.AssessmentUpdatedEvent{/* values here */})
```

### EndpointVerificationEvent

```go
webhookEvent := components.CreateWebhookEventEndpointVerification(components.EndpointVerificationEvent{/* values here */})
```

### AccountDeletedEvent

```go
webhookEvent := components.CreateWebhookEventAccountDeleted(components.AccountDeletedEvent{/* values here */})
```

### WorkspaceDeletedEvent

```go
webhookEvent := components.CreateWebhookEventWorkspaceDeleted(components.WorkspaceDeletedEvent{/* values here */})
```

### WebhookEndpointDisabledEvent

```go
webhookEvent := components.CreateWebhookEventWebhookEndpointDisabled(components.WebhookEndpointDisabledEvent{/* values here */})
```

### AccountRiskChangedEvent

```go
webhookEvent := components.CreateWebhookEventAccountRiskChanged(components.AccountRiskChangedEvent{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch webhookEvent.Type {
	case components.WebhookEventTypeAssessmentUpdated:
		// webhookEvent.AssessmentUpdatedEvent is populated
	case components.WebhookEventTypeEndpointVerification:
		// webhookEvent.EndpointVerificationEvent is populated
	case components.WebhookEventTypeAccountDeleted:
		// webhookEvent.AccountDeletedEvent is populated
	case components.WebhookEventTypeWorkspaceDeleted:
		// webhookEvent.WorkspaceDeletedEvent is populated
	case components.WebhookEventTypeWebhookEndpointDisabled:
		// webhookEvent.WebhookEndpointDisabledEvent is populated
	case components.WebhookEventTypeAccountRiskChanged:
		// webhookEvent.AccountRiskChangedEvent is populated
}
```
