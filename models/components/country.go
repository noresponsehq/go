package components

// Country - IP country, phone country, breach countries and policy_context.region together. When the breach records give a country, the timezone and locale countries count too. The IP country of a proxy or hosting network does not count.
type Country string

const (
	CountryConsistent Country = "consistent"
	CountryMismatch   Country = "mismatch"
	CountryUnknown    Country = "unknown"
)

func (e Country) ToPointer() *Country {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Country) IsExact() bool {
	if e != nil {
		switch *e {
		case "consistent", "mismatch", "unknown":
			return true
		}
	}
	return false
}
