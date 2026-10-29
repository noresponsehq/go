package components

type CustomerExportStatus string

const (
	CustomerExportStatusQueued    CustomerExportStatus = "queued"
	CustomerExportStatusRunning   CustomerExportStatus = "running"
	CustomerExportStatusCompleted CustomerExportStatus = "completed"
	CustomerExportStatusFailed    CustomerExportStatus = "failed"
	CustomerExportStatusCancelled CustomerExportStatus = "cancelled"
	CustomerExportStatusExpired   CustomerExportStatus = "expired"
)

func (e CustomerExportStatus) ToPointer() *CustomerExportStatus {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *CustomerExportStatus) IsExact() bool {
	if e != nil {
		switch *e {
		case "queued", "running", "completed", "failed", "cancelled", "expired":
			return true
		}
	}
	return false
}
