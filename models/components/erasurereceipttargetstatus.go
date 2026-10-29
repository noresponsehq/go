package components

type ErasureReceiptTargetStatus string

const (
	ErasureReceiptTargetStatusPending   ErasureReceiptTargetStatus = "pending"
	ErasureReceiptTargetStatusCompleted ErasureReceiptTargetStatus = "completed"
)

func (e ErasureReceiptTargetStatus) ToPointer() *ErasureReceiptTargetStatus {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ErasureReceiptTargetStatus) IsExact() bool {
	if e != nil {
		switch *e {
		case "pending", "completed":
			return true
		}
	}
	return false
}
