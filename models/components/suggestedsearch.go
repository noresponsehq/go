package components

// SuggestedSearch - A search enrichment would start: a distinctive or unique value of a likely record, not yet searched for the person. A person starts it with POST .../searches, naming the record and the field.
type SuggestedSearch struct {
	// The lookup record whose value enrichment would search. It is among the report's lookup records.
	RecordID string `json:"record_id"`
	// The field of that record that holds the value. Send record_id and field to POST .../searches to start the search.
	Field string `json:"field"`
	// The kind of the value: one of the kinds the environment's enrichment setting names.
	Kind SuggestedSearchKind `json:"kind"`
	// The likelihood of the record that holds the value, from 0.01 to 0.99. Enrichment suggests a value only from a record more likely than the setting's min_likelihood, the likeliest first.
	Likelihood float64 `json:"likelihood"`
}

func (s *SuggestedSearch) GetRecordID() string {
	if s == nil {
		return ""
	}
	return s.RecordID
}

func (s *SuggestedSearch) GetField() string {
	if s == nil {
		return ""
	}
	return s.Field
}

func (s *SuggestedSearch) GetKind() SuggestedSearchKind {
	if s == nil {
		return SuggestedSearchKind("")
	}
	return s.Kind
}

func (s *SuggestedSearch) GetLikelihood() float64 {
	if s == nil {
		return 0.0
	}
	return s.Likelihood
}
