package components

type LineType string

const (
	LineTypeMobile   LineType = "mobile"
	LineTypeVoip     LineType = "voip"
	LineTypeLandline LineType = "landline"
	LineTypeTollFree LineType = "toll_free"
	LineTypeOther    LineType = "other"
)

func (e LineType) ToPointer() *LineType {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *LineType) IsExact() bool {
	if e != nil {
		switch *e {
		case "mobile", "voip", "landline", "toll_free", "other":
			return true
		}
	}
	return false
}
