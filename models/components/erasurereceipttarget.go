package components

type ErasureReceiptTarget struct {
	Name        ErasureReceiptTargetName   `json:"name"`
	Status      ErasureReceiptTargetStatus `json:"status"`
	CompletedAt *string                    `json:"completed_at,omitzero"`
}

func (e *ErasureReceiptTarget) GetName() ErasureReceiptTargetName {
	if e == nil {
		return ErasureReceiptTargetName("")
	}
	return e.Name
}

func (e *ErasureReceiptTarget) GetStatus() ErasureReceiptTargetStatus {
	if e == nil {
		return ErasureReceiptTargetStatus("")
	}
	return e.Status
}

func (e *ErasureReceiptTarget) GetCompletedAt() *string {
	if e == nil {
		return nil
	}
	return e.CompletedAt
}
