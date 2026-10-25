package components

// SourceID - A NoResponse source: github (GitHub accounts), phone (the phone lookup), breach (breach exposure), rdap (the email domain's registration), mx (the email domain's mail servers) or ipintel (IP intelligence).
type SourceID string

const (
	SourceIDGithub  SourceID = "github"
	SourceIDPhone   SourceID = "phone"
	SourceIDBreach  SourceID = "breach"
	SourceIDRdap    SourceID = "rdap"
	SourceIDMx      SourceID = "mx"
	SourceIDIpintel SourceID = "ipintel"
)

func (e SourceID) ToPointer() *SourceID {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *SourceID) IsExact() bool {
	if e != nil {
		switch *e {
		case "github", "phone", "breach", "rdap", "mx", "ipintel":
			return true
		}
	}
	return false
}
