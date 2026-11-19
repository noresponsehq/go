# SchedulePeriod

How often the person is investigated again: daily, weekly or monthly, counted from when the schedule was set; off runs no scheduled investigation.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v19/models/components"
)

value := components.SchedulePeriodDaily

// Open enum: custom values can be created with a direct type cast
custom := components.SchedulePeriod("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `SchedulePeriodDaily`   | daily                   |
| `SchedulePeriodWeekly`  | weekly                  |
| `SchedulePeriodMonthly` | monthly                 |
| `SchedulePeriodOff`     | off                     |