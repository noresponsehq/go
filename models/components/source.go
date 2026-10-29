package components

type Source string

const (
	SourceRequests        Source = "requests"
	SourceWebhookAttempts Source = "webhook_attempts"
)

func (e Source) ToPointer() *Source {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Source) IsExact() bool {
	if e != nil {
		switch *e {
		case "requests", "webhook_attempts":
			return true
		}
	}
	return false
}
