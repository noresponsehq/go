package components

// ErasureReceipt - Acknowledged propagation of a retained erasure record in the calling environment. Pending means the target has not acknowledged completion. Completion covers these serving-system targets only; backup and point-in-time recovery copies expire according to their own retention policy. This receipt contains no subject id or digest. Workspace teardown revokes customer credentials, so obtain required receipts before teardown.
type ErasureReceipt struct {
	// Opaque, server-issued erasure receipt id. It contains no subject identifier.
	ID                 string                 `json:"id"`
	Kind               ErasureReceiptKind     `json:"kind"`
	Status             ListChangeState        `json:"status"`
	AssessmentsDeleted int64                  `json:"assessments_deleted"`
	RequestedAt        string                 `json:"requested_at"`
	CompletedAt        *string                `json:"completed_at,omitzero"`
	Targets            []ErasureReceiptTarget `json:"targets"`
	BackupPolicy       string                 `json:"backup_policy"`
}

func (e *ErasureReceipt) GetID() string {
	if e == nil {
		return ""
	}
	return e.ID
}

func (e *ErasureReceipt) GetKind() ErasureReceiptKind {
	if e == nil {
		return ErasureReceiptKind("")
	}
	return e.Kind
}

func (e *ErasureReceipt) GetStatus() ListChangeState {
	if e == nil {
		return ListChangeState("")
	}
	return e.Status
}

func (e *ErasureReceipt) GetAssessmentsDeleted() int64 {
	if e == nil {
		return 0
	}
	return e.AssessmentsDeleted
}

func (e *ErasureReceipt) GetRequestedAt() string {
	if e == nil {
		return ""
	}
	return e.RequestedAt
}

func (e *ErasureReceipt) GetCompletedAt() *string {
	if e == nil {
		return nil
	}
	return e.CompletedAt
}

func (e *ErasureReceipt) GetTargets() []ErasureReceiptTarget {
	if e == nil {
		return []ErasureReceiptTarget{}
	}
	return e.Targets
}

func (e *ErasureReceipt) GetBackupPolicy() string {
	if e == nil {
		return ""
	}
	return e.BackupPolicy
}
