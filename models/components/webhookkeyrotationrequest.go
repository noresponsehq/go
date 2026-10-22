package components

// WebhookKeyRotationRequest - POST /v1/webhooks/{endpoint_id}/rotate body: an empty object. The old key keeps verifying for 7 days.
type WebhookKeyRotationRequest struct {
}
