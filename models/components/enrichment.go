package components

type EnrichmentKind string

const (
	EnrichmentKindUsername   EnrichmentKind = "username"
	EnrichmentKindEmail      EnrichmentKind = "email"
	EnrichmentKindPhone      EnrichmentKind = "phone"
	EnrichmentKindPersonName EnrichmentKind = "person_name"
)

func (e EnrichmentKind) ToPointer() *EnrichmentKind {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *EnrichmentKind) IsExact() bool {
	if e != nil {
		switch *e {
		case "username", "email", "phone", "person_name":
			return true
		}
	}
	return false
}

// Enrichment - How enrichment searches the values that a workspace person's own investigation (POST /v1/investigations) and a group's screening find. Defaults: mode automatic, min_likelihood 0.5, kinds username, max_spend null and deadline null. Group screening follows the same setting, but takes mode suggest while the workspace has not set a mode. An investigation request may override each field for that investigation. A change applies to the next search enrichment starts, investigations already run included. Enrichment also keeps to limits of its own: each value once per person (once more if that search failed), at most 5 searches from one search's records, 20 per person and 6 searches deep.
type Enrichment struct {
	// off: enrichment starts no search and suggests none. suggest: enrichment starts no search by itself, and the person report lists the searches it would start in suggested_searches. automatic: after a workspace person's own investigation or a group's screening, and each time a search on that person finishes, enrichment starts searches by itself. A group takes suggest while the workspace has not set a mode.
	Mode EnrichmentMode `json:"mode"`
	// Enrichment searches values only from records more likely than this to be the person's, 0 to 0.99.
	MinLikelihood float64 `json:"min_likelihood"`
	// The kinds of value enrichment searches: distinctive usernames, email addresses and phone numbers, which belong to one person each, and distinctive person names, searched in the breach data.
	Kinds []EnrichmentKind `json:"kinds"`
	// The most enrichment may spend on a person from one investigation, in minor units of USD (cents): each search it started after that investigation started counts its estimate or what it spent, whichever is more. At the bound enrichment starts nothing more for that investigation. null: only the environment's screening_daily_cap bounds it.
	MaxSpend *int64 `json:"max_spend"`
	// How long enrichment may start searches after an investigation starts, 60 to 2592000 seconds. null: no deadline.
	Deadline *int64 `json:"deadline"`
}

func (e *Enrichment) GetMode() EnrichmentMode {
	if e == nil {
		return EnrichmentMode("")
	}
	return e.Mode
}

func (e *Enrichment) GetMinLikelihood() float64 {
	if e == nil {
		return 0.0
	}
	return e.MinLikelihood
}

func (e *Enrichment) GetKinds() []EnrichmentKind {
	if e == nil {
		return []EnrichmentKind{}
	}
	return e.Kinds
}

func (e *Enrichment) GetMaxSpend() *int64 {
	if e == nil {
		return nil
	}
	return e.MaxSpend
}

func (e *Enrichment) GetDeadline() *int64 {
	if e == nil {
		return nil
	}
	return e.Deadline
}
