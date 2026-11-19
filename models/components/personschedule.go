package components

// PersonSchedule - The person's rerun schedule. Each due run starts a full investigation of the person, as POST /v1/investigations does, and its changes are stored as changes. The schedule lives in the database and the worker runs it, so it survives restarts and releases; each due run starts exactly one investigation.
type PersonSchedule struct {
	// How often the person is investigated again: daily, weekly or monthly, counted from when the schedule was set; off runs no scheduled investigation.
	Period SchedulePeriod `json:"period"`
	// When the next scheduled investigation starts, within a minute. null when the schedule is off.
	NextRunAt *string `json:"next_run_at"`
	// The investigation the schedule last started. null before its first run.
	LastInvestigationID *string `json:"last_investigation_id"`
	// Why the last due run started no investigation, as the problem code its start would have answered. null when the last due run started one, or none has come due. A run that started none is not retried: the schedule moves to its next period.
	LastError *LastError `json:"last_error"`
	UpdatedAt *string    `json:"updated_at"`
	// Who last set the schedule.
	UpdatedBy *string `json:"updated_by"`
}

func (p *PersonSchedule) GetPeriod() SchedulePeriod {
	if p == nil {
		return SchedulePeriod("")
	}
	return p.Period
}

func (p *PersonSchedule) GetNextRunAt() *string {
	if p == nil {
		return nil
	}
	return p.NextRunAt
}

func (p *PersonSchedule) GetLastInvestigationID() *string {
	if p == nil {
		return nil
	}
	return p.LastInvestigationID
}

func (p *PersonSchedule) GetLastError() *LastError {
	if p == nil {
		return nil
	}
	return p.LastError
}

func (p *PersonSchedule) GetUpdatedAt() *string {
	if p == nil {
		return nil
	}
	return p.UpdatedAt
}

func (p *PersonSchedule) GetUpdatedBy() *string {
	if p == nil {
		return nil
	}
	return p.UpdatedBy
}
