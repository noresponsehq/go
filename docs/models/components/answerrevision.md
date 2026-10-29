# AnswerRevision

Which answer of each assessment the filters describe: latest, the latest stored revision, or first, the first answer, as served when the assessment was first answered (POST /v1/assessments/aggregate counts first answers). A cell of an aggregate drills down to POST /v1/assessments/search with the same filters, revision first and the cell's bucket as from and to.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v9/models/components"
)

value := components.AnswerRevisionLatest
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `AnswerRevisionLatest` | latest                 |
| `AnswerRevisionFirst`  | first                  |