# FootprintFamily

The source-neutral kind of evidence: account_presence (an account found on a site), auth_provider (the account at your auth provider), breach (breach and credential tables), breach_abuse (abuse and fraud communities' tables) and code_hosting (a code hosting profile).

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v7/models/components"
)

value := components.FootprintFamilyAccountPresence

// Open enum: custom values can be created with a direct type cast
custom := components.FootprintFamily("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `FootprintFamilyAccountPresence` | account_presence                 |
| `FootprintFamilyAuthProvider`    | auth_provider                    |
| `FootprintFamilyBreach`          | breach                           |
| `FootprintFamilyBreachAbuse`     | breach_abuse                     |
| `FootprintFamilyCodeHosting`     | code_hosting                     |