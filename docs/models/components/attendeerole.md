# AttendeeRole

The attendee's role at the event: guest (a ticket holder), crew, vendor or press. Leave it out when you do not know it: an attendee without one has no role, not guest. A later row naming the same attendee replaces it, and a change of role alone keeps the attendee's screening results.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.AttendeeRoleGuest

// Open enum: custom values can be created with a direct type cast
custom := components.AttendeeRole("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `AttendeeRoleGuest`  | guest                |
| `AttendeeRoleCrew`   | crew                 |
| `AttendeeRoleVendor` | vendor               |
| `AttendeeRolePress`  | press                |