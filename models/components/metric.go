package components

// Metric - assessments: assessments admitted. webhook_deliveries: events sent to endpoints, one per event per endpoint. key_usage_events: key-usage events accepted. screenings: people a screening recorded a final result for, one per person per screening; a failed person counts only when its lookups were kept.
type Metric string

const (
	MetricAssessments       Metric = "assessments"
	MetricWebhookDeliveries Metric = "webhook_deliveries"
	MetricKeyUsageEvents    Metric = "key_usage_events"
	MetricScreenings        Metric = "screenings"
)

func (e Metric) ToPointer() *Metric {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Metric) IsExact() bool {
	if e != nil {
		switch *e {
		case "assessments", "webhook_deliveries", "key_usage_events", "screenings":
			return true
		}
	}
	return false
}
