# Accounts

## Overview

Your accounts, each with one current risk, and their assessment history.

### Available Operations

* [Get](#get) - Read an account's one current risk.
* [ListAssessments](#listassessments) - Page through one account's unexpired assessments.

## Get

An account is the account_id you send with your assessments. Its current risk is the risk of its latest assessment with a decision (status complete or partial), scored as GET /v1/assessments/{assessment_id} scores that assessment, with a decision for each flow. Every assessment updates its account: once an assessment completes, the account carries its risk. Requires the history scope. Read-only and safe to retry. An account with no unexpired assessment that has a decision answers 404 not_found, and a malformed account_id answers 400 invalid_request.

### Example Usage: blocked

<!-- UsageSnippet language="go" operationID="getAccount" method="get" path="/v1/accounts/{account_id}" example="blocked" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Accounts.Get(ctx, "acct_7f3a9c")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: current

<!-- UsageSnippet language="go" operationID="getAccount" method="get" path="/v1/accounts/{account_id}" example="current" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Accounts.Get(ctx, "acct_7f3a9c")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                       | Type                                                                                                                                                                                                            | Required                                                                                                                                                                                                        | Description                                                                                                                                                                                                     | Example                                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                              | The context to use for the request.                                                                                                                                                                             |                                                                                                                                                                                                                 |
| `accountID`                                                                                                                                                                                                     | `string`                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                              | The account_id the tenant sends with its assessments, as one path segment (percent-encode '\|' and ':' if the client requires it). A value that does not match the account_id shape answers 400 invalid_request. | acct_7f3a9c                                                                                                                                                                                                     |
| `opts`                                                                                                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                              | The options for this request.                                                                                                                                                                                   |                                                                                                                                                                                                                 |

### Response

**[*operations.GetAccountResponse](../../models/operations/getaccountresponse.md), error**

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

## ListAssessments

The account's history: the same page as POST /v1/assessments/query for this account_id, newest first, each at its latest assessment_version. Requires the history scope. Read-only and safe to retry. An account with no assessments answers 200 with empty items. The cursor is opaque and bound to your tenant and the account; a tampered or expired cursor answers 400 invalid_cursor. response_mode=detailed needs the detailed scope. The account_id travels in the URL; use POST /v1/assessments/query to keep it out of access logs.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountAssessments" method="get" path="/v1/accounts/{account_id}/assessments" example="page" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Accounts.ListAssessments(ctx, "acct_7f3a9c", noresponse.Pointer[int64](20), nil, nil)
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

| Parameter                                                                                                                                                                                                       | Type                                                                                                                                                                                                            | Required                                                                                                                                                                                                        | Description                                                                                                                                                                                                     | Example                                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                              | The context to use for the request.                                                                                                                                                                             |                                                                                                                                                                                                                 |
| `accountID`                                                                                                                                                                                                     | `string`                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                              | The account_id the tenant sends with its assessments, as one path segment (percent-encode '\|' and ':' if the client requires it). A value that does not match the account_id shape answers 400 invalid_request. | acct_7f3a9c                                                                                                                                                                                                     |
| `limit`                                                                                                                                                                                                         | `*int64`                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                              | Items per page.                                                                                                                                                                                                 |                                                                                                                                                                                                                 |
| `cursor`                                                                                                                                                                                                        | `*string`                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                              | next_cursor from the previous page, unchanged.                                                                                                                                                                  |                                                                                                                                                                                                                 |
| `responseMode`                                                                                                                                                                                                  | [*components.ResponseMode](../../models/components/responsemode.md)                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                              | Optional standard or separately authorised detailed customer view; raw is rejected.                                                                                                                             |                                                                                                                                                                                                                 |
| `opts`                                                                                                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                              | The options for this request.                                                                                                                                                                                   |                                                                                                                                                                                                                 |

### Response

**[*operations.ListAccountAssessmentsResponse](../../models/operations/listaccountassessmentsresponse.md), error**

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