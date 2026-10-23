# Lane

An evidence lane: accounts (account presence for email and username), phone, breach and ip.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v3/models/components"
)

value := components.LaneAccounts

// Open enum: custom values can be created with a direct type cast
custom := components.Lane("custom_value")
```


## Values

| Name           | Value          |
| -------------- | -------------- |
| `LaneAccounts` | accounts       |
| `LanePhone`    | phone          |
| `LaneBreach`   | breach         |
| `LaneIP`       | ip             |