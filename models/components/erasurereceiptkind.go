package components

type ErasureReceiptKind string

const (
	ErasureReceiptKindAssessment ErasureReceiptKind = "assessment"
	ErasureReceiptKindAccount    ErasureReceiptKind = "account"
	ErasureReceiptKindWorkspace  ErasureReceiptKind = "workspace"
)

func (e ErasureReceiptKind) ToPointer() *ErasureReceiptKind {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ErasureReceiptKind) IsExact() bool {
	if e != nil {
		switch *e {
		case "assessment", "account", "workspace":
			return true
		}
	}
	return false
}
