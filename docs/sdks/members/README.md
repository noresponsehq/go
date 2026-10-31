# Members

## Overview

Scoped workspace membership automation.

### Available Operations

* [List](#list) - List a workspace's members.
* [Update](#update) - Change a member's role.
* [Remove](#remove) - Remove a member.

## List

Read the workspace directory with a live key carrying members or members_manage. Member identities and grants belong to this workspace. Owner-controlled issuance limits team automation to authorised service credentials.

### Example Usage

<!-- UsageSnippet language="go" operationID="listMembers" method="get" path="/v1/workspaces/{workspace_id}/members" example="list" -->
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

    res, err := s.Members.List(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM")
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
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |                                                                                                                                    |

### Response

**[*operations.ListMembersResponse](../../models/operations/listmembersresponse.md), error**

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

## Update

Update a non-owner member role, environment access or sensitive permission exclusions with a live members_manage key. Owner changes, identity binding and support access remain outside this scope. The current delegation ceiling is checked within the change transaction.

### Example Usage: member

<!-- UsageSnippet language="go" operationID="updateMember" method="patch" path="/v1/workspaces/{workspace_id}/members/{member_id}" example="member" -->
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

    res, err := s.Members.Update(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "mem_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.MemberUpdateRequest{
        Role: components.RoleViewer,
        Environments: []components.MemberUpdateRequestEnvironment{
            components.MemberUpdateRequestEnvironmentTest,
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
### Example Usage: role

<!-- UsageSnippet language="go" operationID="updateMember" method="patch" path="/v1/workspaces/{workspace_id}/members/{member_id}" example="role" -->
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

    res, err := s.Members.Update(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "mem_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.MemberUpdateRequest{
        Role: components.RoleAdmin,
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
| `memberID`                                                                                                                         | `string`                                                                                                                           | :heavy_check_mark:                                                                                                                 | The member's mem_ id. A member of another workspace, or a removed one, answers 404.                                                | mem_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                     |
| `body`                                                                                                                             | [components.MemberUpdateRequest](../../models/components/memberupdaterequest.md)                                                   | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                | {<br/>"role": "viewer",<br/>"environments": [<br/>"test"<br/>]<br/>}                                                               |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |                                                                                                                                    |

### Response

**[*operations.UpdateMemberResponse](../../models/operations/updatememberresponse.md), error**

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

## Remove

Remove a non-owner member with a live members_manage key. Subsequent access is refused and personal credentials are revoked. Independently owned service credentials follow their own lifecycle.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeMember" method="delete" path="/v1/workspaces/{workspace_id}/members/{member_id}" -->
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

    res, err := s.Members.Remove(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "mem_01JA2Z8K6Q4V7W9XBCDEFGHJKM")
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
| `memberID`                                                                                                                         | `string`                                                                                                                           | :heavy_check_mark:                                                                                                                 | The member's mem_ id. A member of another workspace, or a removed one, answers 404.                                                | mem_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                     |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |                                                                                                                                    |

### Response

**[*operations.RemoveMemberResponse](../../models/operations/removememberresponse.md), error**

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