# Flow

The provider flow the assessment gates. Use the top-level flow instead; this member is ignored when the top-level flow is present. Absent means signup.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v2/models/components"
)

value := components.FlowSignup
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `FlowSignup`        | signup              |
| `FlowLogin`         | login               |
| `FlowRecovery`      | recovery            |
| `FlowProfileChange` | profile_change      |