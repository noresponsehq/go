package components

// IPSummaryType - The address's network type.
type IPSummaryType string

const (
	IPSummaryTypeResidential IPSummaryType = "residential"
	IPSummaryTypeHosting     IPSummaryType = "hosting"
	IPSummaryTypeVpn         IPSummaryType = "vpn"
	IPSummaryTypeProxy       IPSummaryType = "proxy"
	IPSummaryTypeTor         IPSummaryType = "tor"
	IPSummaryTypeMobile      IPSummaryType = "mobile"
	IPSummaryTypeUnknown     IPSummaryType = "unknown"
)

func (e IPSummaryType) ToPointer() *IPSummaryType {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *IPSummaryType) IsExact() bool {
	if e != nil {
		switch *e {
		case "residential", "hosting", "vpn", "proxy", "tor", "mobile", "unknown":
			return true
		}
	}
	return false
}
