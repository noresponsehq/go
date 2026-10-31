# SignIn

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.SignInPassword

// Open enum: custom values can be created with a direct type cast
custom := components.SignIn("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `SignInPassword`  | password          |
| `SignInGoogle`    | google            |
| `SignInFacebook`  | facebook          |
| `SignInApple`     | apple             |
| `SignInMicrosoft` | microsoft         |
| `SignInSso`       | sso               |
| `SignInOther`     | other             |