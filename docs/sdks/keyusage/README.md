# KeyUsage

## Overview

### Available Operations

* [Record](#record) - Report requests made with API keys your platform issued.

## Record

Records requests made with API keys your platform issued to your accounts, one event per request: the key's SHA-256 (the issued_credential.sha256 the account's assessment carried), the client IP, the time and optionally the user agent, model and token counts. Most callers use the SDK's usage buffer, which batches events and sends them in the background. Sampling is allowed: sample_rate weights an event's request and token counts. Requires the assess scope. Idempotency-Key is required and handled as on POST /v1/assessments: a retry with the same key and body answers the same record with Idempotent-Replayed: true and counts nothing again; a different body answers 409. A batch holds 1 to 500 events and at most 256 KiB. No raw event is kept: usage is aggregated per key, and IPs, user agents and models are kept only as keyed digests, with the IP's country, for 24 hours. When one key is used from 20 or more IP addresses, or 3 or more countries, within 24 hours, every account whose assessment carried that key gets the shared_api_key reason for 30 days.

### Example Usage: batch

<!-- UsageSnippet language="go" operationID="recordKeyUsage" method="post" path="/v1/key_usage/batches" example="batch" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v7"
	"github.com/noresponsehq/go/v7/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.KeyUsage.Record(ctx, "account-acct_7f3a9c", components.KeyUsageBatchRequest{
        Events: []components.KeyUsageEvent{
            components.KeyUsageEvent{
                KeySha256: "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
                IP: "203.0.113.24",
                OccurredAt: "2026-10-23T09:14:07Z",
                UserAgent: noresponse.Pointer("python-requests/2.32.3"),
                Model: noresponse.Pointer("gpt-4o-mini"),
                InputTokens: noresponse.Pointer[int64](812),
                OutputTokens: noresponse.Pointer[int64](164),
            },
            components.KeyUsageEvent{
                KeySha256: "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
                IP: "2001:db8:4a::17",
                OccurredAt: "2026-10-23T09:14:09.250Z",
                SampleRate: noresponse.Pointer[float64](0.25),
            },
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
### Example Usage: recorded

<!-- UsageSnippet language="go" operationID="recordKeyUsage" method="post" path="/v1/key_usage/batches" example="recorded" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v7"
	"github.com/noresponsehq/go/v7/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.KeyUsage.Record(ctx, "account-acct_7f3a9c", components.KeyUsageBatchRequest{
        Events: []components.KeyUsageEvent{
            components.KeyUsageEvent{
                KeySha256: "2bdbf5c72f2a36df0ef40fb92203612274d4546f56dba946070a6a79c2401a1b",
                IP: "203.0.113.24",
                OccurredAt: "2026-10-23T09:14:07.000Z",
                UserAgent: noresponse.Pointer("python-requests/2.32.3"),
                Model: noresponse.Pointer("gpt-4o-mini"),
            },
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

### Parameters

| Parameter                                                                                                                                                         | Type                                                                                                                                                              | Required                                                                                                                                                          | Description                                                                                                                                                       | Example                                                                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                                                             | :heavy_check_mark:                                                                                                                                                | The context to use for the request.                                                                                                                               |                                                                                                                                                                   |
| `idempotencyKey`                                                                                                                                                  | `string`                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                | Opaque client key, 16 to 128 characters of [A-Za-z0-9._:-] (a UUID or ULID fits). Scoped to the workspace principal and bound to the canonical body for 24 hours. | account-acct_7f3a9c                                                                                                                                               |
| `body`                                                                                                                                                            | [components.KeyUsageBatchRequest](../../models/components/keyusagebatchrequest.md)                                                                                | :heavy_check_mark:                                                                                                                                                | N/A                                                                                                                                                               |                                                                                                                                                                   |
| `opts`                                                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                                                          | :heavy_minus_sign:                                                                                                                                                | The options for this request.                                                                                                                                     |                                                                                                                                                                   |

### Response

**[*operations.RecordKeyUsageResponse](../../models/operations/recordkeyusageresponse.md), error**

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