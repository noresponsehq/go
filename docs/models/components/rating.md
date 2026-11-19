# Rating

unique: an email or phone, which belongs to one person by design. distinctive: rare enough that a new search on it likely finds one person. common: shared by many people, so a search on it finds strangers. A new search is suggested for unique and distinctive values only; any value can still be searched.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v16/models/components"
)

value := components.RatingUnique

// Open enum: custom values can be created with a direct type cast
custom := components.Rating("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `RatingUnique`      | unique              |
| `RatingDistinctive` | distinctive         |
| `RatingCommon`      | common              |