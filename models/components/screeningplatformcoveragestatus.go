package components

// ScreeningPlatformCoverageStatus - covered read at least one profile; none_found looked and found no account; inaccessible found accounts that are private or locked; failed could not finish looking; unavailable has no resolver or collector enabled.
type ScreeningPlatformCoverageStatus string

const (
	ScreeningPlatformCoverageStatusCovered      ScreeningPlatformCoverageStatus = "covered"
	ScreeningPlatformCoverageStatusNoneFound    ScreeningPlatformCoverageStatus = "none_found"
	ScreeningPlatformCoverageStatusInaccessible ScreeningPlatformCoverageStatus = "inaccessible"
	ScreeningPlatformCoverageStatusFailed       ScreeningPlatformCoverageStatus = "failed"
	ScreeningPlatformCoverageStatusUnavailable  ScreeningPlatformCoverageStatus = "unavailable"
)

func (e ScreeningPlatformCoverageStatus) ToPointer() *ScreeningPlatformCoverageStatus {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ScreeningPlatformCoverageStatus) IsExact() bool {
	if e != nil {
		switch *e {
		case "covered", "none_found", "inaccessible", "failed", "unavailable":
			return true
		}
	}
	return false
}
