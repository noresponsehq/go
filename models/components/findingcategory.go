package components

// FindingCategory - What a finding is about: a direct threat to the principal or the event, intent to harm without a named target, fixation on the principal, a grievance against the host, research into the venue, route or security, weapons tied to any of these, disruption or protest planning toward this event, violent extremist content or identification, prior harassment or stalking of staff or principals, or posting the principal's addresses or movements.
type FindingCategory string

const (
	FindingCategoryDirectThreat     FindingCategory = "direct_threat"
	FindingCategoryHarmIntent       FindingCategory = "harm_intent"
	FindingCategoryFixation         FindingCategory = "fixation"
	FindingCategoryGrievance        FindingCategory = "grievance"
	FindingCategorySecurityResearch FindingCategory = "security_research"
	FindingCategoryWeapons          FindingCategory = "weapons"
	FindingCategoryDisruption       FindingCategory = "disruption"
	FindingCategoryViolentExtremism FindingCategory = "violent_extremism"
	FindingCategoryHarassment       FindingCategory = "harassment"
	FindingCategoryMovementExposure FindingCategory = "movement_exposure"
)

func (e FindingCategory) ToPointer() *FindingCategory {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *FindingCategory) IsExact() bool {
	if e != nil {
		switch *e {
		case "direct_threat", "harm_intent", "fixation", "grievance", "security_research", "weapons", "disruption", "violent_extremism", "harassment", "movement_exposure":
			return true
		}
	}
	return false
}
