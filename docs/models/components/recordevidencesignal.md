# RecordEvidenceSignal

found_by: what found the record, which sets where its likelihood starts. agrees: a value the record shares with what the investigator gave or with a joined record. clashes: a birth date, full name, gender or country that differs from theirs.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v18/models/components"
)

value := components.RecordEvidenceSignalFoundBy

// Open enum: custom values can be created with a direct type cast
custom := components.RecordEvidenceSignal("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `RecordEvidenceSignalFoundBy` | found_by                      |
| `RecordEvidenceSignalAgrees`  | agrees                        |
| `RecordEvidenceSignalClashes` | clashes                       |