package components

// AggregateDimension - decision, recommendation and band: the first answer's. flow: the assessment's. list_match: the action of the list entry that decided the account when it was first answered. reason: a reason the first answer fired, with its impact band. lane: an evidence lane the first answer listed in incomplete_sources. country and asn: the submitted IP's, from the IP data.
type AggregateDimension string

const (
	AggregateDimensionDecision       AggregateDimension = "decision"
	AggregateDimensionFlow           AggregateDimension = "flow"
	AggregateDimensionRecommendation AggregateDimension = "recommendation"
	AggregateDimensionBand           AggregateDimension = "band"
	AggregateDimensionReason         AggregateDimension = "reason"
	AggregateDimensionLane           AggregateDimension = "lane"
	AggregateDimensionCountry        AggregateDimension = "country"
	AggregateDimensionAsn            AggregateDimension = "asn"
	AggregateDimensionListMatch      AggregateDimension = "list_match"
)

func (e AggregateDimension) ToPointer() *AggregateDimension {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *AggregateDimension) IsExact() bool {
	if e != nil {
		switch *e {
		case "decision", "flow", "recommendation", "band", "reason", "lane", "country", "asn", "list_match":
			return true
		}
	}
	return false
}
