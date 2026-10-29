package components

// AggregateGroupListMatch - The list action captured with the first answer. none means no matching entry was found. For retained_first_answers an absent field means the older assessment did not capture list state; this remains a separate group.
type AggregateGroupListMatch string

const (
	AggregateGroupListMatchBlock AggregateGroupListMatch = "block"
	AggregateGroupListMatchAllow AggregateGroupListMatch = "allow"
	AggregateGroupListMatchNone  AggregateGroupListMatch = "none"
)

func (e AggregateGroupListMatch) ToPointer() *AggregateGroupListMatch {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *AggregateGroupListMatch) IsExact() bool {
	if e != nil {
		switch *e {
		case "block", "allow", "none":
			return true
		}
	}
	return false
}
