# ScreeningLookupCheckStatus

found: a source found the identifier. not_found: every source answered and none found it. partial: the sources that answered found nothing and the others did not answer. not_covered: no source answered. A lookup that failed reads partial or not_covered, not not_found.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v11/models/components"
)

value := components.ScreeningLookupCheckStatusFound

// Open enum: custom values can be created with a direct type cast
custom := components.ScreeningLookupCheckStatus("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `ScreeningLookupCheckStatusFound`      | found                                  |
| `ScreeningLookupCheckStatusNotFound`   | not_found                              |
| `ScreeningLookupCheckStatusPartial`    | partial                                |
| `ScreeningLookupCheckStatusNotCovered` | not_covered                            |