# Flow

The provider flow the assessment gates. Absent means signup.

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