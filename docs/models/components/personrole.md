# PersonRole

The person's role in the group: guest (a ticket holder), crew, vendor or press. Leave it out when you do not know it: a person without one has no role, not guest. A later row naming the same person replaces it, and a change of role alone keeps the person's screening results.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v11/models/components"
)

value := components.PersonRoleGuest

// Open enum: custom values can be created with a direct type cast
custom := components.PersonRole("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `PersonRoleGuest`  | guest              |
| `PersonRoleCrew`   | crew               |
| `PersonRoleVendor` | vendor             |
| `PersonRolePress`  | press              |