package components

// Stage - ai_typing and analysis appear only on searches recorded before the engine stopped analysing results.
type Stage string

const (
	StageLookup        Stage = "lookup"
	StageLookupTyping  Stage = "lookup_typing"
	StagePlatforms     Stage = "platforms"
	StageWritten       Stage = "written"
	StageLatePlatforms Stage = "late_platforms"
	StageLateLookup    Stage = "late_lookup"
	StageAiTyping      Stage = "ai_typing"
	StageAnalysis      Stage = "analysis"
	StageRating        Stage = "rating"
)

func (e Stage) ToPointer() *Stage {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Stage) IsExact() bool {
	if e != nil {
		switch *e {
		case "lookup", "lookup_typing", "platforms", "written", "late_platforms", "late_lookup", "ai_typing", "analysis", "rating":
			return true
		}
	}
	return false
}
