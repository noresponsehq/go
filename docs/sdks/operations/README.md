# Operations

## Overview

Retained outcomes, enforcement reports and review populations.

### Available Operations

* [Summarize](#summarize) - Summarize retained account outcomes, enforcement reports and reviews.

## Summarize

Read one retained snapshot of account outcomes, enforcement reports and review states for the environment and a window of at most 31 days. Each section declares its own population, numerator, denominator and label coverage. Current outcome heads preserve explicit unknown labels. Enforcement rows are reports from your application and do not independently verify execution. Requires history or investigate together with management scope. A result exceeding 100000 population rows or the query budget is refused without truncation.

### Example Usage

<!-- UsageSnippet language="go" operationID="summarizeWorkspaceOperations" method="post" path="/v1/workspaces/{workspace_id}/operations/summary" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"github.com/noresponsehq/go/v9/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Operations.Summarize(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.WorkspaceOperationsSummaryRequest{
        Environment: components.AuditSettingsEnvironmentLive,
        From: "2026-01-01T00:00:00Z",
        To: "2026-01-02T00:00:00Z",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        | Example                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |                                                                                                                                    |
| `workspaceID`                                                                                                                      | `string`                                                                                                                           | :heavy_check_mark:                                                                                                                 | Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404. | tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                     |
| `body`                                                                                                                             | [components.WorkspaceOperationsSummaryRequest](../../models/components/workspaceoperationssummaryrequest.md)                       | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |                                                                                                                                    |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |                                                                                                                                    |

### Response

**[*operations.SummarizeWorkspaceOperationsResponse](../../models/operations/summarizeworkspaceoperationsresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| apierrors.BadRequestError       | 400                             | application/problem+json        |
| apierrors.AuthenticationError   | 401                             | application/problem+json        |
| apierrors.PermissionDeniedError | 403                             | application/problem+json        |
| apierrors.NotFoundError         | 404                             | application/problem+json        |
| apierrors.ConflictError         | 409                             | application/problem+json        |
| apierrors.RateLimitError        | 429                             | application/problem+json        |
| apierrors.APIStatusError        | 4XX, 5XX                        | \*/\*                           |