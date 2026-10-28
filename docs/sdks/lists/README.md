# Lists

## Overview

### Available Operations

* [List](#list) - List an environment's list entries.
* [Create](#create) - Add an entry to an environment's list.
* [Delete](#delete) - Remove an entry from an environment's list.

## List

Lists every entry of the signing key's environment's list, newest first: each an email or an email domain that decides the accounts it matches, block or allow, whatever their score. Requires the management scope; the other environment answers 404 not_found. Read-only and safe to retry.

### Example Usage

<!-- UsageSnippet language="go" operationID="listListEntries" method="get" path="/v1/workspaces/{workspace_id}/lists/{environment}" example="default" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v8"
	"github.com/noresponsehq/go/v8/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Lists.List(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentTest)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                              | Type                                                                                                                                                                   | Required                                                                                                                                                               | Description                                                                                                                                                            | Example                                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                  | :heavy_check_mark:                                                                                                                                                     | The context to use for the request.                                                                                                                                    |                                                                                                                                                                        |
| `workspaceID`                                                                                                                                                          | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.                                     | tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `environment`                                                                                                                                                          | [components.KeyEnvironment](../../models/components/keyenvironment.md)                                                                                                 | :heavy_check_mark:                                                                                                                                                     | The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404. |                                                                                                                                                                        |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.ListListEntriesResponse](../../models/operations/listlistentriesresponse.md), error**

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

## Create

Adds an exact email (kind email) or an email domain (kind email_domain) to the signing key's environment's list with the action block or allow, and answers the entry with its value in canonical form. An account matches when one of its assessments sent the email, or an email of the domain: a block entry decides block for every flow and an allow entry allow, block winning when both match, and the account's risk carries list_match with the entry's id while its score is unchanged. Every matching account is re-decided at once and sends account.risk_changed when its decision changes. A value already on the list answers 409 list_entry_exists, a full list 409 list_limit_reached, and an entry matching more than 5,000 accounts 409 list_entry_too_broad. Requires the management scope; the other environment answers 404 not_found.

### Example Usage

<!-- UsageSnippet language="go" operationID="createListEntry" method="post" path="/v1/workspaces/{workspace_id}/lists/{environment}" example="default" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v8"
	"github.com/noresponsehq/go/v8/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Lists.Create(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentLive, components.ListEntryRequest{
        Kind: components.ListEntryKindEmailDomain,
        Value: "throwaway.example",
        Action: components.ListEntryActionBlock,
        Note: noresponse.Pointer("Signup burst from one domain"),
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

| Parameter                                                                                                                                                              | Type                                                                                                                                                                   | Required                                                                                                                                                               | Description                                                                                                                                                            | Example                                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                  | :heavy_check_mark:                                                                                                                                                     | The context to use for the request.                                                                                                                                    |                                                                                                                                                                        |
| `workspaceID`                                                                                                                                                          | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.                                     | tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `environment`                                                                                                                                                          | [components.KeyEnvironment](../../models/components/keyenvironment.md)                                                                                                 | :heavy_check_mark:                                                                                                                                                     | The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404. |                                                                                                                                                                        |
| `body`                                                                                                                                                                 | [components.ListEntryRequest](../../models/components/listentryrequest.md)                                                                                             | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.CreateListEntryResponse](../../models/operations/createlistentryresponse.md), error**

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

Removes an entry from the signing key's environment's list and re-decides the accounts it matched, the 5,000 most recently assessed at most, each sending account.risk_changed when its decision changes. An unknown or already removed entry answers 404 not_found. Requires the management scope.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteListEntry" method="delete" path="/v1/workspaces/{workspace_id}/lists/{environment}/{list_entry_id}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v8"
	"github.com/noresponsehq/go/v8/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Lists.Delete(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentLive, "lst_01JA2ZE4K5M6N7P8Q9R0S1T2V3")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                              | Type                                                                                                                                                                   | Required                                                                                                                                                               | Description                                                                                                                                                            | Example                                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                  | :heavy_check_mark:                                                                                                                                                     | The context to use for the request.                                                                                                                                    |                                                                                                                                                                        |
| `workspaceID`                                                                                                                                                          | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.                                     | tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `environment`                                                                                                                                                          | [components.KeyEnvironment](../../models/components/keyenvironment.md)                                                                                                 | :heavy_check_mark:                                                                                                                                                     | The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404. |                                                                                                                                                                        |
| `listEntryID`                                                                                                                                                          | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | The entry's id. An entry of another environment or workspace answers 404.                                                                                              | lst_01JA2ZE4K5M6N7P8Q9R0S1T2V3                                                                                                                                         |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.DeleteListEntryResponse](../../models/operations/deletelistentryresponse.md), error**

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