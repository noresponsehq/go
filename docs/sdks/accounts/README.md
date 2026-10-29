# Accounts

## Overview

Your accounts, each with one current risk, and their assessment history.

### Available Operations

* [ListChanged](#listchanged) - Poll the accounts whose risk changed.
* [Get](#get) - Read an account's one current risk.
* [Delete](#delete) - Erase an account and its history.
* [ListAssessments](#listassessments) - Page through one account's unexpired assessments.
* [RecordEnforcement](#recordenforcement) - Record what you did to an account in your identity provider.
* [ListEnforcements](#listenforcements) - Page through account enforcements
* [RecordOutcome](#recordoutcome) - Record or correct an account outcome
* [ListOutcomes](#listoutcomes) - List retained account feedback
* [RetrieveDecision](#retrievedecision) - Retrieve an exact historical account decision
* [AssociateCredential](#associatecredential) - Associate a customer-issued API key
* [ListCredentials](#listcredentials) - List issued credential associations
* [RetireCredential](#retirecredential) - Retire an issued credential association
* [Export](#export) - Export an account
* [ListDecisions](#listdecisions) - Page through recorded account decisions
* [Search](#search) - Search distinct accounts by submitted identifier or account ID

## ListChanged

Your accounts whose risk changed, oldest change first, each as accounts.get answers it. Every page carries next_cursor: pass it back as changed_since to read the changes made after that page; without changed_since the list starts from the oldest change. When nothing changed, items is empty and next_cursor is the cursor you sent, so poll by calling again with the latest next_cursor. An account appears again each time its risk changes, with a higher risk_version; keep the copy with the highest. unenforced=true lists only accounts with a decision other than allow for some flow and no applied enforcement at their current risk_version; such a page can hold fewer than limit items, or none, while later changes exist, so keep polling with next_cursor. Requires the history scope. Read-only and safe to retry. A tampered or foreign cursor answers 400 invalid_cursor.

### Example Usage: caughtUp

<!-- UsageSnippet language="go" operationID="listAccounts" method="get" path="/v1/accounts" example="caughtUp" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
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
	noresponse "github.com/noresponsehq/go/v9"
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

| Parameter                                                                                                                                                                               | Type                                                                                                                                                                                    | Required                                                                                                                                                                                | Description                                                                                                                                                                             |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                      | The context to use for the request.                                                                                                                                                     |
| `changedSince`                                                                                                                                                                          | `*string`                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                      | next_cursor from an earlier page, unchanged. Absent, the list starts from the oldest change.                                                                                            |
| `limit`                                                                                                                                                                                 | `*int64`                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                      | Most accounts per page.                                                                                                                                                                 |
| `unenforced`                                                                                                                                                                            | `*bool`                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                      | true lists only accounts with no applied enforcement at their current risk_version that have a decision other than allow for some flow or a hold no later applied enforcement resolved. |
| `opts`                                                                                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                      | The options for this request.                                                                                                                                                           |

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
	noresponse "github.com/noresponsehq/go/v9"
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
	noresponse "github.com/noresponsehq/go/v9"
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
	noresponse "github.com/noresponsehq/go/v9"
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
	noresponse "github.com/noresponsehq/go/v9"
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

The account's history: the same page as POST /v1/assessments/query for this account_id, newest first, each at its latest assessment_version. Requires the history scope. Read-only and safe to retry. An account with no assessments answers 200 with empty items. The cursor is opaque and bound to your workspace and the account; a tampered or expired cursor answers 400 invalid_cursor. response_mode=detailed needs the detailed scope. Each item's identifiers holds the identifiers its assessment was sent, exactly as received. The account_id travels in the URL; use POST /v1/assessments/query to keep it out of access logs.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountAssessments" method="get" path="/v1/accounts/{account_id}/assessments" example="page" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
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
	noresponse "github.com/noresponsehq/go/v9"
	"github.com/noresponsehq/go/v9/models/components"
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
	noresponse "github.com/noresponsehq/go/v9"
	"github.com/noresponsehq/go/v9/models/components"
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

## ListEnforcements

List retained enforcement reports for one account with an authenticated cursor.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountEnforcements" method="get" path="/v1/accounts/{account_id}/enforcements" example="page" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Accounts.ListEnforcements(ctx, "acct_7f3a9c", noresponse.Pointer[int64](20), nil)
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
| `opts`                                                                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                 | The options for this request.                                                                                                                                                                                      |                                                                                                                                                                                                                    |

### Response

**[*operations.ListAccountEnforcementsResponse](../../models/operations/listaccountenforcementsresponse.md), error**

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

## RecordOutcome

Record a factual report, not an enforcement action. A known outcome replaces the effective report. unknown alone records uncertainty without changing it; unknown with supersedes explicitly retracts it. supersedes must name the current effective report or the request answers 409 outcome_conflict. Reports remain append-only. Idempotency-Key is required. observed_at defaults to receipt time and cannot be over five minutes ahead.

### Example Usage

<!-- UsageSnippet language="go" operationID="recordAccountOutcome" method="post" path="/v1/accounts/{account_id}/outcomes" -->
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

    res, err := s.Accounts.RecordOutcome(ctx, "acct_7f3a9c", "account-acct_7f3a9c", components.AccountOutcomeRequest{
        Outcome: components.OutcomeUnknown,
        ObservedAt: noresponse.Pointer("2026-01-01T00:00:00Z"),
        ExternalRef: noresponse.Pointer("case_123"),
        Supersedes: noresponse.Pointer("out_01JA2Z8K6Q4V7W9XBCDEFGHJKM"),
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
| `idempotencyKey`                                                                                                                                                                                                   | `string`                                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                 | Opaque client key, 16 to 128 characters of [A-Za-z0-9._:-] (a UUID or ULID fits). Scoped to the workspace principal and bound to the canonical body for 24 hours.                                                  | account-acct_7f3a9c                                                                                                                                                                                                |
| `body`                                                                                                                                                                                                             | [components.AccountOutcomeRequest](../../models/components/accountoutcomerequest.md)                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                 | N/A                                                                                                                                                                                                                |                                                                                                                                                                                                                    |
| `opts`                                                                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                 | The options for this request.                                                                                                                                                                                      |                                                                                                                                                                                                                    |

### Response

**[*operations.RecordAccountOutcomeResponse](../../models/operations/recordaccountoutcomeresponse.md), error**

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

## ListOutcomes

List retained account reports, newest first. The page-level effective outcome id names the current report. Each receipt preserves the effective report and risk version from its original write. Page with the returned cursor.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountOutcomes" method="get" path="/v1/accounts/{account_id}/outcomes" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Accounts.ListOutcomes(ctx, "acct_7f3a9c", noresponse.Pointer[int64](20), nil)
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
| `opts`                                                                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                 | The options for this request.                                                                                                                                                                                      |                                                                                                                                                                                                                    |

### Response

**[*operations.ListAccountOutcomesResponse](../../models/operations/listaccountoutcomesresponse.md), error**

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

## RetrieveDecision

Retrieve the account decision at one recorded risk version. The answer uses the original assessment and account-specific risk inputs. Earlier versions that were not captured return 404. The assessment version, policy version and account calculation version identify the decision.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccountDecision" method="get" path="/v1/accounts/{account_id}/decisions/{risk_version}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Accounts.RetrieveDecision(ctx, "acct_7f3a9c", 269840)
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
| `riskVersion`                                                                                                                                                                                                      | `int64`                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                 | The canonical decimal risk_version originally served.                                                                                                                                                              |                                                                                                                                                                                                                    |
| `opts`                                                                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                 | The options for this request.                                                                                                                                                                                      |                                                                                                                                                                                                                    |

### Response

**[*operations.GetAccountDecisionResponse](../../models/operations/getaccountdecisionresponse.md), error**

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

## AssociateCredential

Associate the SHA-256 of a high-entropy key issued by your service. Send the hash instead of the key. Repeating a hash returns its original association, including its retirement state. A hash associated with another account returns a conflict.

### Example Usage

<!-- UsageSnippet language="go" operationID="associateCredential" method="post" path="/v1/accounts/{account_id}/credentials" -->
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

    res, err := s.Accounts.AssociateCredential(ctx, "acct_7f3a9c", components.CredentialRequest{
        Sha256: "5da2158606a4cb396280c842b578872bb64b96c4eb80dc02e6fb4e0e864b77c6",
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
| `body`                                                                                                                                                                                                             | [components.CredentialRequest](../../models/components/credentialrequest.md)                                                                                                                                       | :heavy_check_mark:                                                                                                                                                                                                 | N/A                                                                                                                                                                                                                |                                                                                                                                                                                                                    |
| `opts`                                                                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                 | The options for this request.                                                                                                                                                                                      |                                                                                                                                                                                                                    |

### Response

**[*operations.AssociateCredentialResponse](../../models/operations/associatecredentialresponse.md), error**

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

## ListCredentials

Retained credential associations, ordered by id. Page with the authenticated cursor. Associations are removed on account erasure.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountCredentials" method="get" path="/v1/accounts/{account_id}/credentials" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Accounts.ListCredentials(ctx, "acct_7f3a9c", noresponse.Pointer[int64](20), nil)
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
| `opts`                                                                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                 | The options for this request.                                                                                                                                                                                      |                                                                                                                                                                                                                    |

### Response

**[*operations.ListAccountCredentialsResponse](../../models/operations/listaccountcredentialsresponse.md), error**

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

## RetireCredential

Mark this association retired, idempotently. This does not revoke the actual key. Requires delete scope.

### Example Usage

<!-- UsageSnippet language="go" operationID="retireAccountCredential" method="delete" path="/v1/accounts/{account_id}/credentials/{credential_id}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Accounts.RetireCredential(ctx, "acct_7f3a9c", "<id>")
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
| `credentialID`                                                                                                                                                                                                     | `string`                                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                 | N/A                                                                                                                                                                                                                |                                                                                                                                                                                                                    |
| `opts`                                                                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                 | The options for this request.                                                                                                                                                                                      |                                                                                                                                                                                                                    |

### Response

**[*operations.RetireAccountCredentialResponse](../../models/operations/retireaccountcredentialresponse.md), error**

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

## Export

Retrieve one consistent snapshot of retained customer-accessible records, including assessment revisions, account decisions, feedback, key associations, enforcements and request diagnostics. Each assessment revision's detail.identifiers holds the identifiers that assessment was sent, exactly as received. Secrets, restricted source information, other customers' data and legacy free-text notes are excluded. The same standard/detailed permissions apply. An export exceeding 10000 records or 8 MiB returns export_too_large without a partial result. Historical decisions that were not captured are identified as an exclusion.

### Example Usage

<!-- UsageSnippet language="go" operationID="exportAccount" method="get" path="/v1/accounts/{account_id}/export" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Accounts.Export(ctx, "acct_7f3a9c", nil)
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
| `responseMode`                                                                                                                                                                                                     | [*components.ResponseMode](../../models/components/responsemode.md)                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                 | standard returns the consolidated risk response; detailed adds merged provider-neutral facts when the key has detailed scope. There is no raw mode.                                                                |                                                                                                                                                                                                                    |
| `opts`                                                                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                 | The options for this request.                                                                                                                                                                                      |                                                                                                                                                                                                                    |

### Response

**[*operations.ExportAccountResponse](../../models/operations/exportaccountresponse.md), error**

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

## ListDecisions

List retained immutable risk decisions for one account with an authenticated cursor.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountDecisions" method="get" path="/v1/accounts/{account_id}/decisions" example="page" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Accounts.ListDecisions(ctx, "acct_7f3a9c", noresponse.Pointer[int64](20), nil)
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
| `opts`                                                                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                 | The options for this request.                                                                                                                                                                                      |                                                                                                                                                                                                                    |

### Response

**[*operations.ListAccountDecisionsResponse](../../models/operations/listaccountdecisionsresponse.md), error**

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

## Search

Search distinct accounts observed in retained assessments in this environment. identifier matches an exact, canonical or explicitly selected mailbox form of a submitted email, phone or customer-namespace username. It matches historical retained inputs and returns matched_identifier with the original value and matching assessment. Each row has its newest retained assessment, current account decision when available, retained assessment count, and total distinct-account count. identifier and legacy email are mutually exclusive; account_id and needs_action combine with either using AND. Set needs_action to select accounts with a retained open or in-progress review, or a current review/block account decision without an effective retained account outcome. This filter applies before counting and pagination and adds waiting_for reasons. Current decisions include account list overrides and shared-key scoring and preserve admission holds. Typed identifier search uses stable opaque account order; ordinary and legacy email searches retain newest-seen order. Results are live across pages; each page's items and total share a database snapshot, while current account risk is read separately. Cursors are expiring and bound to scope, operation and filters. Totals are exact below 10000 and capped at 10000. Requires history or investigate authority and identifier-read permission for identifier or email filtering. needs_action additionally requires management scope. Source-discovered associations are not searched and equality is not verified ownership.

### Example Usage

<!-- UsageSnippet language="go" operationID="searchAccounts" method="post" path="/v1/accounts/search" -->
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

    res, err := s.Accounts.Search(ctx, components.AccountDirectorySearchRequest{
        Identifier: &components.IdentifierSearch{
            Kind: components.IdentifierSearchKindUsername,
            Value: "janedoe",
            Match: components.MatchCanonical.ToPointer(),
            Namespace: components.NamespaceCustomer.ToPointer(),
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [components.AccountDirectorySearchRequest](../../models/components/accountdirectorysearchrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.SearchAccountsResponse](../../models/operations/searchaccountsresponse.md), error**

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