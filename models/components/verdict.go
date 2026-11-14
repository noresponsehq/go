package components

// Verdict - A reviewer's overrule (POST .../people/{person_id}/records), which decides joined whatever the likelihood and sticks for as long as a lookup finds the same record. Absent when nobody overruled it.
type Verdict string

const (
	VerdictJoined   Verdict = "joined"
	VerdictRejected Verdict = "rejected"
)

func (e Verdict) ToPointer() *Verdict {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Verdict) IsExact() bool {
	if e != nil {
		switch *e {
		case "joined", "rejected":
			return true
		}
	}
	return false
}
