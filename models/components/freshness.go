package components

// Freshness class of the evidence behind a signal. unknown whenever no evidence was obtained.
type Freshness string

const (
	FreshnessFresh   Freshness = "fresh"
	FreshnessStale   Freshness = "stale"
	FreshnessUnknown Freshness = "unknown"
)

func (e Freshness) ToPointer() *Freshness {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Freshness) IsExact() bool {
	if e != nil {
		switch *e {
		case "fresh", "stale", "unknown":
			return true
		}
	}
	return false
}
