# Usage

## Overview

Recorded usage for your environment.

### Available Operations

* [Get](#get) - Read a workspace's usage.

## Get

Read daily usage units for the environments you can access. These counts are unpriced. With live access it also returns the daily console lookup cap: used, limit and resets_at.

### Example Usage

<!-- UsageSnippet language="go" operationID="getUsage" method="get" path="/v1/workspaces/{workspace_id}/usage" example="usage" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v10"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Usage.Get(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", nil, nil)
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
| `from`                                                                                                                             | `*string`                                                                                                                          | :heavy_minus_sign:                                                                                                                 | The first UTC day, as YYYY-MM-DD.                                                                                                  |                                                                                                                                    |
| `to`                                                                                                                               | `*string`                                                                                                                          | :heavy_minus_sign:                                                                                                                 | The last UTC day, included, as YYYY-MM-DD; no later than today.                                                                    |                                                                                                                                    |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |                                                                                                                                    |

### Response

**[*operations.GetUsageResponse](../../models/operations/getusageresponse.md), error**

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