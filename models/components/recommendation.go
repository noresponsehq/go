package components

// Recommendation - The recommended action: allow, step_up or review. Your application applies it.
type Recommendation string

const (
	RecommendationAllow  Recommendation = "allow"
	RecommendationStepUp Recommendation = "step_up"
	RecommendationReview Recommendation = "review"
)

func (e Recommendation) ToPointer() *Recommendation {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Recommendation) IsExact() bool {
	if e != nil {
		switch *e {
		case "allow", "step_up", "review":
			return true
		}
	}
	return false
}
