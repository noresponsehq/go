package components

// Delay - How long after the assessment's creation the change was made.
type Delay string

const (
	DelayUnder10s Delay = "under_10s"
	DelayUnder1m  Delay = "under_1m"
	DelayUnder10m Delay = "under_10m"
	DelayUnder1h  Delay = "under_1h"
	DelayUnder1d  Delay = "under_1d"
	DelayOver1d   Delay = "over_1d"
)

func (e Delay) ToPointer() *Delay {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Delay) IsExact() bool {
	if e != nil {
		switch *e {
		case "under_10s", "under_1m", "under_10m", "under_1h", "under_1d", "over_1d":
			return true
		}
	}
	return false
}
