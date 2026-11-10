# FindingCategory

What a finding is about: a direct threat to the principal or the event, intent to harm without a named target, fixation on the principal, a grievance against the host, research into the venue, route or security, weapons tied to any of these, disruption or protest planning toward this event, violent extremist content or identification, prior harassment or stalking of staff or principals, or posting the principal's addresses or movements.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.FindingCategoryDirectThreat

// Open enum: custom values can be created with a direct type cast
custom := components.FindingCategory("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `FindingCategoryDirectThreat`     | direct_threat                     |
| `FindingCategoryHarmIntent`       | harm_intent                       |
| `FindingCategoryFixation`         | fixation                          |
| `FindingCategoryGrievance`        | grievance                         |
| `FindingCategorySecurityResearch` | security_research                 |
| `FindingCategoryWeapons`          | weapons                           |
| `FindingCategoryDisruption`       | disruption                        |
| `FindingCategoryViolentExtremism` | violent_extremism                 |
| `FindingCategoryHarassment`       | harassment                        |
| `FindingCategoryMovementExposure` | movement_exposure                 |