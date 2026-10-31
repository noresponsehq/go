# Attribute

name_gender is the records' gender against the submitted first name's usual one.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.AttributeBirthDate

// Open enum: custom values can be created with a direct type cast
custom := components.Attribute("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `AttributeBirthDate`   | birth_date             |
| `AttributeGender`      | gender                 |
| `AttributeNationality` | nationality            |
| `AttributeNameGender`  | name_gender            |