package components

// Population - first_answers uses anonymous counters that can include erased detail. retained_first_answers computes exact intersections over retained first answers within 31 days. Requests exceeding the 100000-record work budget are refused.
type Population string

const (
	PopulationFirstAnswers         Population = "first_answers"
	PopulationRetainedFirstAnswers Population = "retained_first_answers"
)

func (e Population) ToPointer() *Population {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Population) IsExact() bool {
	if e != nil {
		switch *e {
		case "first_answers", "retained_first_answers":
			return true
		}
	}
	return false
}
