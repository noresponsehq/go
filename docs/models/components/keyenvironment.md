# KeyEnvironment

The environment the key belongs to. A live key assesses with live sources; a test key answers from the deterministic simulator.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v7/models/components"
)

value := components.KeyEnvironmentLive

// Open enum: custom values can be created with a direct type cast
custom := components.KeyEnvironment("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `KeyEnvironmentLive` | live                 |
| `KeyEnvironmentTest` | test                 |