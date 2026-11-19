package components

import (
	"github.com/noresponsehq/go/v17/internal/utils"
)

// Record - One answer from one source: where it came from, what found it and every field it holds. A person's lookup records also say how likely each is the person.
type Record struct {
	// Made from the source and the fields' names and original values, so the same answer found twice is one record with the same id. Its fields are typed as its latest finding typed them.
	ID string `json:"id"`
	// Where a record came from.
	Source RecordSource `json:"source"`
	// Each kind of searched value that found the record.
	Matched []RecordMatch `json:"matched"`
	// Every field the source returned, in the order the source's connector lists them; a breach row's columns in name order. Nothing is withheld: a breach row's passwords, hashes and salts are fields like any other.
	Fields []RecordField `json:"fields"`
	// When the source last gave this answer.
	ObservedAt string `json:"observed_at"`
	// How likely the record is the person it was found for, from 0.01 to 0.99, to two places. A record found by the person's own email or phone starts as theirs, and one found by a search of a value another record holds starts at that record's likelihood divided by one plus the records its search found on its site; any other starts lower the more people its search found on its site, or anywhere for a name search. Each value it shares with what is known of the person raises it, a rare value more than a common one, and a different birth date or full name lowers it. What is known is what the person's row gives and the values of their joined records: those found by their email or phone, those a reviewer joined and those whose likelihood reaches 0.9. Each record that joins adds its values and the rest are scored again until the joined records stop changing; a record whose likelihood falls below 0.9 is not joined and adds nothing, so a reviewer's rejection also undoes the joins that came through the rejected record. A record whose source kind starts with domain_ describes an email's domain, not the person, so it is not compared and adds nothing to what is known: it stays at 0.99 when the person's own email found it and at 0.5 otherwise. Present on a person's lookup records, with agreed, clashed and joined.
	Likelihood *float64 `json:"likelihood,omitzero"`
	// The fields of this record whose values agreed with what is known of the person, by their field names, in the record's order. The value that found the record is not among them.
	Agreed []string `json:"agreed,omitzero"`
	// The fields of this record that clashed with what is known of the person: a birth date or a full name that differs from the person's.
	Clashed []string `json:"clashed,omitzero"`
	// Whether the record joins the person: its likelihood reached 0.9 or a reviewer joined it, and no reviewer rejected it. A record that does not join stays a lead.
	Joined *bool `json:"joined,omitzero"`
	// A reviewer's overrule (POST .../people/{person_id}/records), which decides joined whatever the likelihood and sticks for as long as a lookup finds the same record. Absent when nobody overruled it.
	Verdict *Verdict `json:"verdict,omitzero"`
}

func (r Record) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *Record) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (r *Record) GetID() string {
	if r == nil {
		return ""
	}
	return r.ID
}

func (r *Record) GetSource() RecordSource {
	if r == nil {
		return RecordSource{}
	}
	return r.Source
}

func (r *Record) GetMatched() []RecordMatch {
	if r == nil {
		return []RecordMatch{}
	}
	return r.Matched
}

func (r *Record) GetFields() []RecordField {
	if r == nil {
		return []RecordField{}
	}
	return r.Fields
}

func (r *Record) GetObservedAt() string {
	if r == nil {
		return ""
	}
	return r.ObservedAt
}

func (r *Record) GetLikelihood() *float64 {
	if r == nil {
		return nil
	}
	return r.Likelihood
}

func (r *Record) GetAgreed() []string {
	if r == nil {
		return nil
	}
	return r.Agreed
}

func (r *Record) GetClashed() []string {
	if r == nil {
		return nil
	}
	return r.Clashed
}

func (r *Record) GetJoined() *bool {
	if r == nil {
		return nil
	}
	return r.Joined
}

func (r *Record) GetVerdict() *Verdict {
	if r == nil {
		return nil
	}
	return r.Verdict
}
