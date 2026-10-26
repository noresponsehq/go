package components

type Direction string

const (
	DirectionRaises  Direction = "raises"
	DirectionLowers  Direction = "lowers"
	DirectionNeutral Direction = "neutral"
)

func (e Direction) ToPointer() *Direction {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Direction) IsExact() bool {
	if e != nil {
		switch *e {
		case "raises", "lowers", "neutral":
			return true
		}
	}
	return false
}
