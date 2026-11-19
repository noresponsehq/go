package components

type ErasureReceiptTargetName string

const (
	ErasureReceiptTargetNameServing ErasureReceiptTargetName = "serving"
	ErasureReceiptTargetNameOutbox  ErasureReceiptTargetName = "outbox"
)

func (e ErasureReceiptTargetName) ToPointer() *ErasureReceiptTargetName {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ErasureReceiptTargetName) IsExact() bool {
	if e != nil {
		switch *e {
		case "serving", "outbox":
			return true
		}
	}
	return false
}
