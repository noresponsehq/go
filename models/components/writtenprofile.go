package components

type WrittenProfile struct {
	// queued: the profile is being written. written: summary, statements and leads hold it. failed: it could not be written; error_code says why and no content is returned.
	Status WrittenProfileStatus `json:"status"`
	// The investigation whose evidence the profile is written from.
	InvestigationID string `json:"investigation_id"`
	// Present only when status is failed. no_writer: no profile writer serves this environment. refused: the model refused. not_written: the answer was not a profile, or no statement survived grounding. uncited_statement: a statement or lead cited a record the person does not hold. not_answered: the writer did not answer. worker_attempts_exhausted: the worker stopped retrying.
	ErrorCode  *WrittenProfileErrorCode `json:"error_code,omitzero"`
	QueuedAt   string                   `json:"queued_at"`
	FinishedAt *string                  `json:"finished_at"`
	// Present when written: a few plain sentences with the most useful picture first.
	Summary *string `json:"summary,omitzero"`
	// Each fact or inference with its certainty and the evidence it cites. Statements that cited evidence the writer was not sent are dropped before storage.
	Statements []ProfileStatement `json:"statements"`
	// Values worth searching next, each held by a linked record.
	Leads []ProfileLead `json:"leads"`
}

func (w *WrittenProfile) GetStatus() WrittenProfileStatus {
	if w == nil {
		return WrittenProfileStatus("")
	}
	return w.Status
}

func (w *WrittenProfile) GetInvestigationID() string {
	if w == nil {
		return ""
	}
	return w.InvestigationID
}

func (w *WrittenProfile) GetErrorCode() *WrittenProfileErrorCode {
	if w == nil {
		return nil
	}
	return w.ErrorCode
}

func (w *WrittenProfile) GetQueuedAt() string {
	if w == nil {
		return ""
	}
	return w.QueuedAt
}

func (w *WrittenProfile) GetFinishedAt() *string {
	if w == nil {
		return nil
	}
	return w.FinishedAt
}

func (w *WrittenProfile) GetSummary() *string {
	if w == nil {
		return nil
	}
	return w.Summary
}

func (w *WrittenProfile) GetStatements() []ProfileStatement {
	if w == nil {
		return []ProfileStatement{}
	}
	return w.Statements
}

func (w *WrittenProfile) GetLeads() []ProfileLead {
	if w == nil {
		return []ProfileLead{}
	}
	return w.Leads
}
