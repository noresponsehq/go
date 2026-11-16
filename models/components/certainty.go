package components

// Certainty - confirmed: a linked record or the searched identifier states it directly. likely: a strong inference. possible: a weak inference.
type Certainty string

const (
	CertaintyConfirmed Certainty = "confirmed"
	CertaintyLikely    Certainty = "likely"
	CertaintyPossible  Certainty = "possible"
)

func (e Certainty) ToPointer() *Certainty {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Certainty) IsExact() bool {
	if e != nil {
		switch *e {
		case "confirmed", "likely", "possible":
			return true
		}
	}
	return false
}
