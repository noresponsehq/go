# ScreeningMatchEvidenceKind

What matched. email, phone, supplied_handle and source_fact are hard identifiers; company, city and school are soft. supplied_handle is a handle on the person row; source_fact is an account a profile found by the person's lookups lists.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v11/models/components"
)

value := components.ScreeningMatchEvidenceKindEmail

// Open enum: custom values can be created with a direct type cast
custom := components.ScreeningMatchEvidenceKind("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `ScreeningMatchEvidenceKindEmail`          | email                                      |
| `ScreeningMatchEvidenceKindPhone`          | phone                                      |
| `ScreeningMatchEvidenceKindSuppliedHandle` | supplied_handle                            |
| `ScreeningMatchEvidenceKindSourceFact`     | source_fact                                |
| `ScreeningMatchEvidenceKindName`           | name                                       |
| `ScreeningMatchEvidenceKindCompany`        | company                                    |
| `ScreeningMatchEvidenceKindCity`           | city                                       |
| `ScreeningMatchEvidenceKindSchool`         | school                                     |