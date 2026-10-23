package components

// WebhookDeliveryAttemptStatus - delivered when the endpoint answered 2xx; failed otherwise, with failure naming why.
type WebhookDeliveryAttemptStatus string

const (
	WebhookDeliveryAttemptStatusDelivered WebhookDeliveryAttemptStatus = "delivered"
	WebhookDeliveryAttemptStatusFailed    WebhookDeliveryAttemptStatus = "failed"
)

func (e WebhookDeliveryAttemptStatus) ToPointer() *WebhookDeliveryAttemptStatus {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *WebhookDeliveryAttemptStatus) IsExact() bool {
	if e != nil {
		switch *e {
		case "delivered", "failed":
			return true
		}
	}
	return false
}
