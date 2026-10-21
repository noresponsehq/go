<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	noresponse "github.com/noresponsehq/go/v2"
	"github.com/noresponsehq/go/v2/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := noresponse.New(
		noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
	)

	res, err := s.Assessments.Create(ctx, "<value>", components.AssessmentRequest{
		AccountID: "acct_7f3a9c",
		Identifiers: components.AssessmentRequestIdentifiers{
			Email: noresponse.Pointer("new.user@example.com"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->