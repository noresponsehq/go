# ScreeningPlatformCoverageStatus

covered read at least one profile; none_found looked and found no account; inaccessible found accounts that are private or locked; failed could not finish looking; unavailable has no resolver or collector enabled.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.ScreeningPlatformCoverageStatusCovered

// Open enum: custom values can be created with a direct type cast
custom := components.ScreeningPlatformCoverageStatus("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `ScreeningPlatformCoverageStatusCovered`      | covered                                       |
| `ScreeningPlatformCoverageStatusNoneFound`    | none_found                                    |
| `ScreeningPlatformCoverageStatusInaccessible` | inaccessible                                  |
| `ScreeningPlatformCoverageStatusFailed`       | failed                                        |
| `ScreeningPlatformCoverageStatusUnavailable`  | unavailable                                   |