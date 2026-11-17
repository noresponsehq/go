package components

// EnrichmentMode - off: enrichment starts no search and suggests none. suggest: enrichment starts no search by itself, and the person report lists the searches it would start in suggested_searches. automatic: after a workspace person's own investigation or a group's screening, and each time a search on that person finishes, enrichment starts searches by itself. A group takes suggest while the workspace has not set a mode.
type EnrichmentMode string

const (
	EnrichmentModeOff       EnrichmentMode = "off"
	EnrichmentModeSuggest   EnrichmentMode = "suggest"
	EnrichmentModeAutomatic EnrichmentMode = "automatic"
)

func (e EnrichmentMode) ToPointer() *EnrichmentMode {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *EnrichmentMode) IsExact() bool {
	if e != nil {
		switch *e {
		case "off", "suggest", "automatic":
			return true
		}
	}
	return false
}
