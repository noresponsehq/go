# Exports

## Overview

Bounded customer export jobs and encrypted artifacts.

### Available Operations

* [Create](#create) - Request an asynchronous customer export.
* [List](#list) - List exports for one collection.
* [Get](#get) - Read export state and progress.
* [Cancel](#cancel) - Cancel and remove an export artifact.
* [Retry](#retry) - Retry a failed export.
* [DownloadChunk](#downloadchunk) - Download one authorized export chunk.

## Create

Request an encrypted JSON Lines snapshot of retained account or audit records. Requires exports plus history or audit scope. An account export without legacy history needs investigate with evidence or detailed, because it holds submitted identifiers; detailed account data always needs evidence or detailed. Criteria are frozen and hashed. Each environment permits four active jobs, 64 MiB per artifact, and 256 KiB per chunk. Artifacts expire after 24 hours. Send an Idempotency-Key to safely repeat a request. Every operation checks current permissions and source retention.

### Example Usage

<!-- UsageSnippet language="go" operationID="createCustomerExport" method="post" path="/v1/workspaces/{workspace_id}/exports" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v10"
	"github.com/noresponsehq/go/v10/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Exports.Create(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.CustomerExportRequest{
        Kind: components.CustomerExportRequestKindAccount,
        Environment: components.EnvironmentTest,
        AccountID: noresponse.Pointer("acct_example"),
        ResponseMode: components.ResponseModeStandard.ToPointer(),
    }, nil)
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
| `body`                                                                                                                             | [components.CustomerExportRequest](../../models/components/customerexportrequest.md)                                               | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                | {<br/>"kind": "account",<br/>"environment": "live",<br/>"account_id": "acct_7f3a9c",<br/>"response_mode": "standard"<br/>}         |
| `idempotencyKey`                                                                                                                   | `*string`                                                                                                                          | :heavy_minus_sign:                                                                                                                 | N/A                                                                                                                                |                                                                                                                                    |

### Response

**[*operations.CreateCustomerExportResponse](../../models/operations/createcustomerexportresponse.md), error**

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

## List

List authorised jobs for one kind and environment. Detailed jobs require current detailed-data permission and are filtered before pagination. Keep the kind and environment unchanged when following next_cursor.

### Example Usage

<!-- UsageSnippet language="go" operationID="listCustomerExports" method="get" path="/v1/workspaces/{workspace_id}/exports" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v10"
	"github.com/noresponsehq/go/v10/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Exports.List(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.CustomerExportRequestKindAudit, components.EnvironmentLive, noresponse.Pointer[int64](25), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        | Example                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |                                                                                                                                    |
| `workspaceID`                                                                                                                      | `string`                                                                                                                           | :heavy_check_mark:                                                                                                                 | Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404. | tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                     |
| `kind`                                                                                                                             | [components.CustomerExportRequestKind](../../models/components/customerexportrequestkind.md)                                       | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |                                                                                                                                    |
| `environment`                                                                                                                      | [components.Environment](../../models/components/environment.md)                                                                   | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |                                                                                                                                    |
| `limit`                                                                                                                            | `*int64`                                                                                                                           | :heavy_minus_sign:                                                                                                                 | N/A                                                                                                                                |                                                                                                                                    |
| `cursor`                                                                                                                           | `*string`                                                                                                                          | :heavy_minus_sign:                                                                                                                 | N/A                                                                                                                                |                                                                                                                                    |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |                                                                                                                                    |

### Response

**[*operations.ListCustomerExportsResponse](../../models/operations/listcustomerexportsresponse.md), error**

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

## Get

Read the status of an authorised export. Records and bytes describe committed progress in the current generation. Only completed, unexpired artifacts can be downloaded. Source erasure, retention changes and workspace removal can invalidate a job.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCustomerExport" method="get" path="/v1/workspaces/{workspace_id}/exports/{export_id}" -->
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

    res, err := s.Exports.Get(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "<id>")
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
| `exportID`                                                                                                                         | `string`                                                                                                                           | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |                                                                                                                                    |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |                                                                                                                                    |

### Response

**[*operations.GetCustomerExportResponse](../../models/operations/getcustomerexportresponse.md), error**

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

## Cancel

Cancel an export and remove its stored chunks. Repeating cancellation is safe. An already admitted download may finish; bytes already returned cannot be recalled. Send an empty JSON object.

### Example Usage

<!-- UsageSnippet language="go" operationID="cancelCustomerExport" method="post" path="/v1/workspaces/{workspace_id}/exports/{export_id}/cancel" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v10"
	"github.com/noresponsehq/go/v10/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Exports.Cancel(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "<id>", components.CustomerExportActionRequest{})
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
| `exportID`                                                                                                                         | `string`                                                                                                                           | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |                                                                                                                                    |
| `body`                                                                                                                             | [components.CustomerExportActionRequest](../../models/components/customerexportactionrequest.md)                                   | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |                                                                                                                                    |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |                                                                                                                                    |

### Response

**[*operations.CancelCustomerExportResponse](../../models/operations/cancelcustomerexportresponse.md), error**

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

## Retry

Start a new generation of a failed export before its original expiry, up to ten generations. The request criteria stay fixed and the new generation takes a fresh source snapshot. Cancelled, erased and expired jobs cannot be retried. Send an empty JSON object.

### Example Usage

<!-- UsageSnippet language="go" operationID="retryCustomerExport" method="post" path="/v1/workspaces/{workspace_id}/exports/{export_id}/retry" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v10"
	"github.com/noresponsehq/go/v10/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Exports.Retry(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "<id>", components.CustomerExportActionRequest{})
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
| `exportID`                                                                                                                         | `string`                                                                                                                           | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |                                                                                                                                    |
| `body`                                                                                                                             | [components.CustomerExportActionRequest](../../models/components/customerexportactionrequest.md)                                   | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |                                                                                                                                    |

### Response

**[*operations.RetryCustomerExportResponse](../../models/operations/retrycustomerexportresponse.md), error**

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

## DownloadChunk

Download one chunk of a completed, unexpired export after a fresh permission and privacy check. The response contains base64-encoded JSON Lines, a SHA-256, generation and next index. Decode and concatenate chunks in order from one generation. Audit storage failure withholds the bytes.

### Example Usage

<!-- UsageSnippet language="go" operationID="downloadCustomerExportChunk" method="get" path="/v1/workspaces/{workspace_id}/exports/{export_id}/chunks/{chunk_index}" -->
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

    res, err := s.Exports.DownloadChunk(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "<id>", 768311)
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
| `exportID`                                                                                                                         | `string`                                                                                                                           | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |                                                                                                                                    |
| `chunkIndex`                                                                                                                       | `int64`                                                                                                                            | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |                                                                                                                                    |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |                                                                                                                                    |

### Response

**[*operations.DownloadCustomerExportChunkResponse](../../models/operations/downloadcustomerexportchunkresponse.md), error**

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