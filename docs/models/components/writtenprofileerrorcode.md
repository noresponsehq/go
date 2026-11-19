# WrittenProfileErrorCode

Present only when status is failed. no_writer: no profile writer serves this environment. refused: the model refused. not_written: the answer was not a profile, or no statement survived grounding. uncited_statement: a statement or lead cited a record the person does not hold. not_answered: the writer did not answer. worker_attempts_exhausted: the worker stopped retrying.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v17/models/components"
)

value := components.WrittenProfileErrorCodeNoWriter

// Open enum: custom values can be created with a direct type cast
custom := components.WrittenProfileErrorCode("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `WrittenProfileErrorCodeNoWriter`                | no_writer                                        |
| `WrittenProfileErrorCodeRefused`                 | refused                                          |
| `WrittenProfileErrorCodeNotWritten`              | not_written                                      |
| `WrittenProfileErrorCodeUncitedStatement`        | uncited_statement                                |
| `WrittenProfileErrorCodeNotAnswered`             | not_answered                                     |
| `WrittenProfileErrorCodeWorkerAttemptsExhausted` | worker_attempts_exhausted                        |