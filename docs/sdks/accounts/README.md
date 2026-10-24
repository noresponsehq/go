# Accounts

## Overview

Your accounts, each with one current risk, and their assessment history.

### Available Operations

* [ListChanged](#listchanged) - Poll the accounts whose risk changed.
* [Get](#get) - Read an account's one current risk.
* [Delete](#delete) - Erase an account and its history.
* [ListAssessments](#listassessments) - Page through one account's unexpired assessments.
* [RecordEnforcement](#recordenforcement) - Record what you did to an account in your identity provider.

## ListChanged

Your accounts whose risk changed, oldest change first, each as accounts.get answers it. Every page carries next_cursor: pass it back as changed_since to read the changes made after that page; without changed_since the list starts from the oldest change. When nothing changed, items is empty and next_cursor is the cursor you sent, so poll by calling again with the latest next_cursor. An account appears again each time its risk changes, with a higher risk_version; keep the copy with the highest. unenforced=true lists only accounts with a decision other than allow for some flow and no applied enforcement at their current risk_version; such a page can hold fewer than limit items, or none, while later changes exist, so keep polling with next_cursor. Requires the history scope. Read-only and safe to retry. A tampered or foreign cursor answers 400 invalid_cursor.

### Example Usage: caughtUp

<!-- UsageSnippet language="go" operationID="listAccounts" method="get" path="/v1/accounts" example="caughtUp" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v7"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Accounts.ListChanged(ctx, nil, noresponse.Pointer[int64](50), noresponse.Pointer(false))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: page

<!-- UsageSnippet language="go" operationID="listAccounts" method="get" path="/v1/accounts" example="page" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v7"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Accounts.ListChanged(ctx, nil, noresponse.Pointer[int64](50), noresponse.Pointer(false))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                         | Type                                                                                                                              | Required                                                                                                                          | Description                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                             | :heavy_check_mark:                                                                                                                | The context to use for the request.                                                                                               |
| `changedSince`                                                                                                                    | `*string`                                                                                                                         | :heavy_minus_sign:                                                                                                                | next_cursor from an earlier page, unchanged. Absent, the list starts from the oldest change.                                      |
| `limit`                                                                                                                           | `*int64`                                                                                                                          | :heavy_minus_sign:                                                                                                                | Most accounts per page.                                                                                                           |
| `unenforced`                                                                                                                      | `*bool`                                                                                                                           | :heavy_minus_sign:                                                                                                                | true lists only accounts with a decision other than allow for some flow and no applied enforcement at their current risk_version. |
| `opts`                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                          | :heavy_minus_sign:                                                                                                                | The options for this request.                                                                                                     |

### Response

**[*operations.ListAccountsResponse](../../models/operations/listaccountsresponse.md), error**

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

An account is the account_id you send with your assessments. Its current risk is the risk of its latest assessment with a decision (status complete or partial), scored as GET /v1/assessments/{assessment_id} scores that assessment, with a decision for each flow; risk.decision is the decision for the flow of that assessment. Every assessment answers with this same account. Every assessment updates its account: once an assessment completes, the account carries its risk. Requires the history scope. Read-only and safe to retry. An account with no unexpired assessment that has a decision answers 404 not_found, and a malformed account_id answers 400 invalid_request.

### Example Usage: blocked

<!-- UsageSnippet language="go" operationID="getAccount" method="get" path="/v1/accounts/{account_id}" example="blocked" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v7"
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
	noresponse "github.com/noresponsehq/go/v7"
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

| Parameter                                                                                                                                                                                                          | Type                                                                                                                                                                                                               | Required                                                                                                                                                                                                           | Description                                                                                                                                                                                                        | Example                                                                                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                 | The context to use for the request.                                                                                                                                                                                |                                                                                                                                                                                                                    |
| `accountID`                                                                                                                                                                                                        | `string`                                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                 | The account_id the workspace sends with its assessments, as one path segment (percent-encode '\|' and ':' if the client requires it). A value that does not match the account_id shape answers 400 invalid_request. | acct_7f3a9c                                                                                                                                                                                                        |
| `opts`                                                                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                 | The options for this request.                                                                                                                                                                                      |                                                                                                                                                                                                                    |

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

## Delete

Erases the account in the environment of your key, for an erasure request from your user: every assessment of it, with its revisions, evidence and outcomes, and the identity-graph marks its outcomes set. An assessment of the account begun before the call is refused or stopped. Afterwards the account and each erased assessment answer 404 and the history is empty. When the call deleted at least one assessment, an account.deleted event goes to your verified webhook endpoints. Requires the history scope. Idempotent and safe to retry: a repeat, or an account with nothing to erase, answers 200 with assessments_deleted 0 and sends no event. The account_id can be assessed again later, as a new account.

### Example Usage: erased

<!-- UsageSnippet language="go" operationID="deleteAccount" method="delete" path="/v1/accounts/{account_id}" example="erased" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v7"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Accounts.Delete(ctx, "acct_7f3a9c")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: repeat

<!-- UsageSnippet language="go" operationID="deleteAccount" method="delete" path="/v1/accounts/{account_id}" example="repeat" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v7"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Accounts.Delete(ctx, "acct_7f3a9c")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                          | Type                                                                                                                                                                                                               | Required                                                                                                                                                                                                           | Description                                                                                                                                                                                                        | Example                                                                                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                 | The context to use for the request.                                                                                                                                                                                |                                                                                                                                                                                                                    |
| `accountID`                                                                                                                                                                                                        | `string`                                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                 | The account_id the workspace sends with its assessments, as one path segment (percent-encode '\|' and ':' if the client requires it). A value that does not match the account_id shape answers 400 invalid_request. | acct_7f3a9c                                                                                                                                                                                                        |
| `opts`                                                                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                 | The options for this request.                                                                                                                                                                                      |                                                                                                                                                                                                                    |

### Response

**[*operations.DeleteAccountResponse](../../models/operations/deleteaccountresponse.md), error**

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

The account's history: the same page as POST /v1/assessments/query for this account_id, newest first, each at its latest assessment_version. Requires the history scope. Read-only and safe to retry. An account with no assessments answers 200 with empty items. The cursor is opaque and bound to your workspace and the account; a tampered or expired cursor answers 400 invalid_cursor. response_mode=detailed needs the detailed scope. The account_id travels in the URL; use POST /v1/assessments/query to keep it out of access logs.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountAssessments" method="get" path="/v1/accounts/{account_id}/assessments" example="page" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v7"
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

| Parameter                                                                                                                                                                                                          | Type                                                                                                                                                                                                               | Required                                                                                                                                                                                                           | Description                                                                                                                                                                                                        | Example                                                                                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                 | The context to use for the request.                                                                                                                                                                                |                                                                                                                                                                                                                    |
| `accountID`                                                                                                                                                                                                        | `string`                                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                 | The account_id the workspace sends with its assessments, as one path segment (percent-encode '\|' and ':' if the client requires it). A value that does not match the account_id shape answers 400 invalid_request. | acct_7f3a9c                                                                                                                                                                                                        |
| `limit`                                                                                                                                                                                                            | `*int64`                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                 | Items per page.                                                                                                                                                                                                    |                                                                                                                                                                                                                    |
| `cursor`                                                                                                                                                                                                           | `*string`                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                 | next_cursor from the previous page, unchanged.                                                                                                                                                                     |                                                                                                                                                                                                                    |
| `responseMode`                                                                                                                                                                                                     | [*components.ResponseMode](../../models/components/responsemode.md)                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                 | Optional standard or separately authorised detailed customer view; raw is rejected.                                                                                                                                |                                                                                                                                                                                                                    |
| `opts`                                                                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                 | The options for this request.                                                                                                                                                                                      |                                                                                                                                                                                                                    |

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

## RecordEnforcement

Records what you did to the account in your identity provider for one flow's decision at one risk_version: the action (or none), the account's state before it, the provider and whether it was applied or failed. It is your action, not an outcome, and it leaves the account's risk unchanged. A record is kept once per account, risk_version, flow and action: the first report answers 201, and a repeat answers 200 with the stored record unchanged, so retries are safe without an Idempotency-Key. An applied report after a failed one replaces the failed record, keeping its id and created_at, and answers 201. A risk_version older than the account's current one is accepted and stored, but only an applied record at the current risk_version takes the account out of listChanged with unenforced=true; a newer one answers 409 risk_version_ahead. An account with no current risk answers 404 not_found. error_code is allowed only with status failed. Requires the history scope.

### Example Usage: applied

<!-- UsageSnippet language="go" operationID="recordEnforcement" method="post" path="/v1/accounts/{account_id}/enforcements" example="applied" -->
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

    res, err := s.Accounts.RecordEnforcement(ctx, "acct_7f3a9c", components.EnforcementRequest{
        RiskVersion: 3,
        Flow: components.FlowAccountCreation,
        Action: components.EnforcementActionRestrict,
        StateBefore: components.EnforcementStateProvisional,
        Status: components.EnforcementStatusApplied,
        Provider: components.EnforcementProviderClerk,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: failed

<!-- UsageSnippet language="go" operationID="recordEnforcement" method="post" path="/v1/accounts/{account_id}/enforcements" example="failed" -->
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

    res, err := s.Accounts.RecordEnforcement(ctx, "acct_7f3a9c", components.EnforcementRequest{
        RiskVersion: 3,
        Flow: components.FlowAccountCreation,
        Action: components.EnforcementActionRestrict,
        StateBefore: components.EnforcementStateProvisional,
        Status: components.EnforcementStatusFailed,
        Provider: components.EnforcementProviderClerk,
        ErrorCode: noresponse.Pointer("clerk.rate_limited"),
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

| Parameter                                                                                                                                                                                                          | Type                                                                                                                                                                                                               | Required                                                                                                                                                                                                           | Description                                                                                                                                                                                                        | Example                                                                                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                 | The context to use for the request.                                                                                                                                                                                |                                                                                                                                                                                                                    |
| `accountID`                                                                                                                                                                                                        | `string`                                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                 | The account_id the workspace sends with its assessments, as one path segment (percent-encode '\|' and ':' if the client requires it). A value that does not match the account_id shape answers 400 invalid_request. | acct_7f3a9c                                                                                                                                                                                                        |
| `body`                                                                                                                                                                                                             | [components.EnforcementRequest](../../models/components/enforcementrequest.md)                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                 | N/A                                                                                                                                                                                                                |                                                                                                                                                                                                                    |
| `opts`                                                                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                 | The options for this request.                                                                                                                                                                                      |                                                                                                                                                                                                                    |

### Response

**[*operations.RecordEnforcementResponse](../../models/operations/recordenforcementresponse.md), error**

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