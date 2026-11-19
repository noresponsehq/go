package components

// SchedulePeriod - How often the person is investigated again: daily, weekly or monthly, counted from when the schedule was set; off runs no scheduled investigation.
type SchedulePeriod string

const (
	SchedulePeriodDaily   SchedulePeriod = "daily"
	SchedulePeriodWeekly  SchedulePeriod = "weekly"
	SchedulePeriodMonthly SchedulePeriod = "monthly"
	SchedulePeriodOff     SchedulePeriod = "off"
)

func (e SchedulePeriod) ToPointer() *SchedulePeriod {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *SchedulePeriod) IsExact() bool {
	if e != nil {
		switch *e {
		case "daily", "weekly", "monthly", "off":
			return true
		}
	}
	return false
}
