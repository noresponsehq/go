package components

// PersonScheduleRequest - The schedule to set.
type PersonScheduleRequest struct {
	// How often the person is investigated again: daily, weekly or monthly, counted from when the schedule was set; off runs no scheduled investigation.
	Period SchedulePeriod `json:"period"`
}

func (p *PersonScheduleRequest) GetPeriod() SchedulePeriod {
	if p == nil {
		return SchedulePeriod("")
	}
	return p.Period
}
