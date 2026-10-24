# Country

IP country, phone country, breach countries and policy_context.region together.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v6/models/components"
)

value := components.CountryConsistent

// Open enum: custom values can be created with a direct type cast
custom := components.Country("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `CountryConsistent` | consistent          |
| `CountryMismatch`   | mismatch            |
| `CountryUnknown`    | unknown             |