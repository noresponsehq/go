# Usage

Create an assessment for a new signup. Give every signup its own idempotency key from `noresponse.NewIdempotencyKey()` and store it beside the signup, so a retry after a crash sends the same key. README.md is the full guide.

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	noresponse "github.com/noresponsehq/go/v8"
	"github.com/noresponsehq/go/v8/models/components"
)

func main() {
	client := noresponse.NewClient(os.Getenv("NORESPONSE_API_KEY"))

	// Make one key per signup and save it with the signup record before this
	// call. A retry for the same signup reads the saved key back and sends it.
	key := noresponse.NewIdempotencyKey()

	res, err := client.Assessments.Create(context.Background(), key, components.AssessmentRequest{
		AccountID: "acct_7f3a9c",
		Flow:      components.FlowAccountCreation,
		Identifiers: components.AssessmentRequestIdentifiers{
			Email: noresponse.Pointer("new.user@example.com"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	assessment := res.Result
	if assessment.Account == nil {
		// The account's first assessment is still pending. Apply your own
		// fallback, then wait for the webhook or call noresponse.WaitForAssessment.
		return
	}
	fmt.Println(assessment.Account.Risk.Decision, assessment.Account.Risk.Score)
}
```
