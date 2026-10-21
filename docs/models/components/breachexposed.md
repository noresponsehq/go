# BreachExposed

The kind of data a breach table held beside the identifier. password means a plaintext password was present and password_hash means a hash, salt or token was present. The value itself is not stored or returned; the schema has no member for one.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v2/models/components"
)

value := components.BreachExposedEmail

// Open enum: custom values can be created with a direct type cast
custom := components.BreachExposed("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `BreachExposedEmail`        | email                       |
| `BreachExposedUsername`     | username                    |
| `BreachExposedName`         | name                        |
| `BreachExposedPassword`     | password                    |
| `BreachExposedPasswordHash` | password_hash               |
| `BreachExposedPhone`        | phone                       |
| `BreachExposedAddress`      | address                     |
| `BreachExposedDob`          | dob                         |
| `BreachExposedIP`           | ip                          |
| `BreachExposedPayment`      | payment                     |
| `BreachExposedSecurityQa`   | security_qa                 |