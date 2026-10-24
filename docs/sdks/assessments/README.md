# Assessments

## Overview

Assessments of account creation, login and purchase.

### Available Operations

* [Create](#create) - Assess identifiers for a signup or gateway decision.
* [Get](#get) - Poll the authoritative assessment of one assessment.
* [Delete](#delete) - Delete one assessment.
* [RecordOutcome](#recordoutcome) - Report what was later learned about an assessed account.
* [Query](#query) - Page through one account's unexpired assessments.

## Create

Assesses one account creation, login or purchase (flow) and answers with the account's one current risk: account.risk.decision is the decision for this flow (allow, step_up, review or block, where block means your block threshold was reached), and detail holds this assessment's own facts. depth sets how long the answer waits for sources: quick waits only for the fast ones (500 ms), full for the whole budget (3000 ms at most); it defaults to full for account_creation and purchase and quick for login. 200 carries complete or partial, 202 carries pending with Location and Retry-After. While the account's first assessment is pending, account is absent: apply your own fallback for that, 429, 503 or a timeout. NoResponse blocks nothing itself; you decide. Requires the assess scope. Idempotency-Key is required and is bound to the body for 24 hours: the same key with the same body replays the latest assessment with Idempotent-Replayed: true and starts no new work; the same key with a different body answers 409 idempotency_key_reused. response_mode defaults to standard; detailed needs the detailed scope.

### Example Usage: complete

<!-- UsageSnippet language="go" operationID="createAssessment" method="post" path="/v1/assessments" example="complete" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v6"
	"github.com/noresponsehq/go/v6/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Assessments.Create(ctx, "account-acct_7f3a9c", components.AssessmentRequest{
        AccountID: "acct_7f3a9c",
        Flow: components.FlowPurchase,
        Identifiers: components.AssessmentRequestIdentifiers{
            Email: noresponse.Pointer("new.user@example.com"),
            Username: noresponse.Pointer("new.user"),
            Phone: noresponse.Pointer("+61412345678"),
            Name: noresponse.Pointer("Alex Example"),
            IP: noresponse.Pointer("203.0.113.45"),
        },
        PolicyContext: &components.AssessmentRequestPolicyContext{
            Region: noresponse.Pointer("AU"),
        },
        RequestID: noresponse.Pointer("req-2026-10-14-0001"),
        ClientContext: &components.AssessmentRequestClientContext{
            Timezone: noresponse.Pointer("Australia/Sydney"),
            Locale: noresponse.Pointer("en-AU"),
            AppVersion: noresponse.Pointer("4.12.0"),
            UserAgent: noresponse.Pointer("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0 Safari/537.36"),
            DeviceID: noresponse.Pointer("dev_5c1e8a"),
            SessionID: noresponse.Pointer("sess_9b2d41"),
            ClientObservedAt: noresponse.Pointer("2026-10-14T08:59:58.000Z"),
        },
        AuthContext: &components.AuthContext{
            Provider: "clerk",
            Issuer: "https://example-app.clerk.accounts.dev",
            Audience: noresponse.Pointer("https://app.example.com"),
            Audiences: []string{
                "https://app.example.com",
            },
            Subject: "user_2EXAMPLE",
            SessionID: noresponse.Pointer("sess_2EXAMPLE"),
            OrganizationID: noresponse.Pointer("org_2EXAMPLE"),
            TenantID: noresponse.Pointer("tenant_2EXAMPLE"),
            TokenID: noresponse.Pointer("2EXAMPLEJTI"),
            IssuedAt: noresponse.Pointer("2026-10-14T08:59:40.000Z"),
            ExpiresAt: noresponse.Pointer("2026-10-14T09:00:40.000Z"),
            SessionCreatedAt: noresponse.Pointer("2026-10-14T08:57:00.000Z"),
            AuthTime: noresponse.Pointer("2026-10-14T08:55:00.000Z"),
            AccountCreatedAt: noresponse.Pointer("2026-10-14T08:56:00.000Z"),
            Verified: true,
            Mfa: &components.AuthContextMfa{
                Methods: []string{
                    "totp",
                },
            },
            Roles: []string{
                "admin",
            },
            Scopes: []string{
                "org:sys_profile:manage",
            },
        },
        IssuedCredential: &components.AssessmentRequestIssuedCredential{
            Sha256: "2bdbf5c72f2a36df0ef40fb92203612274d4546f56dba946070a6a79c2401a1b",
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
### Example Usage: detailed

<!-- UsageSnippet language="go" operationID="createAssessment" method="post" path="/v1/assessments" example="detailed" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v6"
	"github.com/noresponsehq/go/v6/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Assessments.Create(ctx, "account-acct_7f3a9c", components.AssessmentRequest{
        AccountID: "acct_7f3a9c",
        Flow: components.FlowPurchase,
        Identifiers: components.AssessmentRequestIdentifiers{
            Email: noresponse.Pointer("new.user@example.com"),
            Username: noresponse.Pointer("new.user"),
            Phone: noresponse.Pointer("+61412345678"),
            Name: noresponse.Pointer("Alex Example"),
            IP: noresponse.Pointer("203.0.113.45"),
        },
        PolicyContext: &components.AssessmentRequestPolicyContext{
            Region: noresponse.Pointer("AU"),
        },
        RequestID: noresponse.Pointer("req-2026-10-14-0001"),
        ClientContext: &components.AssessmentRequestClientContext{
            Timezone: noresponse.Pointer("Australia/Sydney"),
            Locale: noresponse.Pointer("en-AU"),
            AppVersion: noresponse.Pointer("4.12.0"),
            UserAgent: noresponse.Pointer("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0 Safari/537.36"),
            DeviceID: noresponse.Pointer("dev_5c1e8a"),
            SessionID: noresponse.Pointer("sess_9b2d41"),
            ClientObservedAt: noresponse.Pointer("2026-10-14T08:59:58.000Z"),
        },
        AuthContext: &components.AuthContext{
            Provider: "clerk",
            Issuer: "https://example-app.clerk.accounts.dev",
            Audience: noresponse.Pointer("https://app.example.com"),
            Audiences: []string{
                "https://app.example.com",
            },
            Subject: "user_2EXAMPLE",
            SessionID: noresponse.Pointer("sess_2EXAMPLE"),
            OrganizationID: noresponse.Pointer("org_2EXAMPLE"),
            TenantID: noresponse.Pointer("tenant_2EXAMPLE"),
            TokenID: noresponse.Pointer("2EXAMPLEJTI"),
            IssuedAt: noresponse.Pointer("2026-10-14T08:59:40.000Z"),
            ExpiresAt: noresponse.Pointer("2026-10-14T09:00:40.000Z"),
            SessionCreatedAt: noresponse.Pointer("2026-10-14T08:57:00.000Z"),
            AuthTime: noresponse.Pointer("2026-10-14T08:55:00.000Z"),
            AccountCreatedAt: noresponse.Pointer("2026-10-14T08:56:00.000Z"),
            Verified: true,
            Mfa: &components.AuthContextMfa{
                Methods: []string{
                    "totp",
                },
            },
            Roles: []string{
                "admin",
            },
            Scopes: []string{
                "org:sys_profile:manage",
            },
        },
        IssuedCredential: &components.AssessmentRequestIssuedCredential{
            Sha256: "2bdbf5c72f2a36df0ef40fb92203612274d4546f56dba946070a6a79c2401a1b",
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
### Example Usage: full

<!-- UsageSnippet language="go" operationID="createAssessment" method="post" path="/v1/assessments" example="full" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v6"
	"github.com/noresponsehq/go/v6/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Assessments.Create(ctx, "account-acct_7f3a9c", components.AssessmentRequest{
        AccountID: "acct_7f3a9c",
        Flow: components.FlowAccountCreation,
        Identifiers: components.AssessmentRequestIdentifiers{
            Email: noresponse.Pointer("new.user@example.com"),
            Username: noresponse.Pointer("new_user_42"),
            Phone: noresponse.Pointer("+447700900123"),
        },
        PolicyContext: &components.AssessmentRequestPolicyContext{
            Region: noresponse.Pointer("GB"),
        },
        RequestID: noresponse.Pointer("req-2026-10-14-0001"),
        DeadlineMs: noresponse.Pointer[int64](300),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: login

<!-- UsageSnippet language="go" operationID="createAssessment" method="post" path="/v1/assessments" example="login" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v6"
	"github.com/noresponsehq/go/v6/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Assessments.Create(ctx, "account-acct_7f3a9c", components.AssessmentRequest{
        AccountID: "acct_7f3a9c",
        Flow: components.FlowPurchase,
        Identifiers: components.AssessmentRequestIdentifiers{
            Email: noresponse.Pointer("new.user@example.com"),
            Username: noresponse.Pointer("new.user"),
            Phone: noresponse.Pointer("+61412345678"),
            Name: noresponse.Pointer("Alex Example"),
            IP: noresponse.Pointer("203.0.113.45"),
        },
        PolicyContext: &components.AssessmentRequestPolicyContext{
            Region: noresponse.Pointer("AU"),
        },
        RequestID: noresponse.Pointer("req-2026-10-14-0001"),
        ClientContext: &components.AssessmentRequestClientContext{
            Timezone: noresponse.Pointer("Australia/Sydney"),
            Locale: noresponse.Pointer("en-AU"),
            AppVersion: noresponse.Pointer("4.12.0"),
            UserAgent: noresponse.Pointer("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0 Safari/537.36"),
            DeviceID: noresponse.Pointer("dev_5c1e8a"),
            SessionID: noresponse.Pointer("sess_9b2d41"),
            ClientObservedAt: noresponse.Pointer("2026-10-14T08:59:58.000Z"),
        },
        AuthContext: &components.AuthContext{
            Provider: "clerk",
            Issuer: "https://example-app.clerk.accounts.dev",
            Audience: noresponse.Pointer("https://app.example.com"),
            Audiences: []string{
                "https://app.example.com",
            },
            Subject: "user_2EXAMPLE",
            SessionID: noresponse.Pointer("sess_2EXAMPLE"),
            OrganizationID: noresponse.Pointer("org_2EXAMPLE"),
            TenantID: noresponse.Pointer("tenant_2EXAMPLE"),
            TokenID: noresponse.Pointer("2EXAMPLEJTI"),
            IssuedAt: noresponse.Pointer("2026-10-14T08:59:40.000Z"),
            ExpiresAt: noresponse.Pointer("2026-10-14T09:00:40.000Z"),
            SessionCreatedAt: noresponse.Pointer("2026-10-14T08:57:00.000Z"),
            AuthTime: noresponse.Pointer("2026-10-14T08:55:00.000Z"),
            AccountCreatedAt: noresponse.Pointer("2026-10-14T08:56:00.000Z"),
            Verified: true,
            Mfa: &components.AuthContextMfa{
                Methods: []string{
                    "totp",
                },
            },
            Roles: []string{
                "admin",
            },
            Scopes: []string{
                "org:sys_profile:manage",
            },
        },
        IssuedCredential: &components.AssessmentRequestIssuedCredential{
            Sha256: "2bdbf5c72f2a36df0ef40fb92203612274d4546f56dba946070a6a79c2401a1b",
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
### Example Usage: minimal

<!-- UsageSnippet language="go" operationID="createAssessment" method="post" path="/v1/assessments" example="minimal" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v6"
	"github.com/noresponsehq/go/v6/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Assessments.Create(ctx, "account-acct_7f3a9c", components.AssessmentRequest{
        AccountID: "acct_7f3a9c",
        Flow: components.FlowAccountCreation,
        Identifiers: components.AssessmentRequestIdentifiers{
            Email: noresponse.Pointer("new.user@example.com"),
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
### Example Usage: name_and_ip

<!-- UsageSnippet language="go" operationID="createAssessment" method="post" path="/v1/assessments" example="name_and_ip" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v6"
	"github.com/noresponsehq/go/v6/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Assessments.Create(ctx, "account-acct_7f3a9c", components.AssessmentRequest{
        AccountID: "acct_8412",
        Flow: components.FlowAccountCreation,
        Identifiers: components.AssessmentRequestIdentifiers{
            Email: noresponse.Pointer("jane.doe@example.com"),
            Username: noresponse.Pointer("janedoe"),
            Phone: noresponse.Pointer("+14155550123"),
            Name: noresponse.Pointer("Jane A. Doe"),
            IP: noresponse.Pointer("73.162.18.45"),
        },
        PolicyContext: &components.AssessmentRequestPolicyContext{
            Region: noresponse.Pointer("US"),
        },
        RequestID: noresponse.Pointer("req_01"),
        DeadlineMs: noresponse.Pointer[int64](500),
        ResponseMode: components.ResponseModeStandard.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: partial

<!-- UsageSnippet language="go" operationID="createAssessment" method="post" path="/v1/assessments" example="partial" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v6"
	"github.com/noresponsehq/go/v6/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Assessments.Create(ctx, "account-acct_7f3a9c", components.AssessmentRequest{
        AccountID: "acct_7f3a9c",
        Flow: components.FlowPurchase,
        Identifiers: components.AssessmentRequestIdentifiers{
            Email: noresponse.Pointer("new.user@example.com"),
            Username: noresponse.Pointer("new.user"),
            Phone: noresponse.Pointer("+61412345678"),
            Name: noresponse.Pointer("Alex Example"),
            IP: noresponse.Pointer("203.0.113.45"),
        },
        PolicyContext: &components.AssessmentRequestPolicyContext{
            Region: noresponse.Pointer("AU"),
        },
        RequestID: noresponse.Pointer("req-2026-10-14-0001"),
        ClientContext: &components.AssessmentRequestClientContext{
            Timezone: noresponse.Pointer("Australia/Sydney"),
            Locale: noresponse.Pointer("en-AU"),
            AppVersion: noresponse.Pointer("4.12.0"),
            UserAgent: noresponse.Pointer("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0 Safari/537.36"),
            DeviceID: noresponse.Pointer("dev_5c1e8a"),
            SessionID: noresponse.Pointer("sess_9b2d41"),
            ClientObservedAt: noresponse.Pointer("2026-10-14T08:59:58.000Z"),
        },
        AuthContext: &components.AuthContext{
            Provider: "clerk",
            Issuer: "https://example-app.clerk.accounts.dev",
            Audience: noresponse.Pointer("https://app.example.com"),
            Audiences: []string{
                "https://app.example.com",
            },
            Subject: "user_2EXAMPLE",
            SessionID: noresponse.Pointer("sess_2EXAMPLE"),
            OrganizationID: noresponse.Pointer("org_2EXAMPLE"),
            TenantID: noresponse.Pointer("tenant_2EXAMPLE"),
            TokenID: noresponse.Pointer("2EXAMPLEJTI"),
            IssuedAt: noresponse.Pointer("2026-10-14T08:59:40.000Z"),
            ExpiresAt: noresponse.Pointer("2026-10-14T09:00:40.000Z"),
            SessionCreatedAt: noresponse.Pointer("2026-10-14T08:57:00.000Z"),
            AuthTime: noresponse.Pointer("2026-10-14T08:55:00.000Z"),
            AccountCreatedAt: noresponse.Pointer("2026-10-14T08:56:00.000Z"),
            Verified: true,
            Mfa: &components.AuthContextMfa{
                Methods: []string{
                    "totp",
                },
            },
            Roles: []string{
                "admin",
            },
            Scopes: []string{
                "org:sys_profile:manage",
            },
        },
        IssuedCredential: &components.AssessmentRequestIssuedCredential{
            Sha256: "2bdbf5c72f2a36df0ef40fb92203612274d4546f56dba946070a6a79c2401a1b",
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
| `body`                                                                                                                                                            | [components.AssessmentRequest](../../models/components/assessmentrequest.md)                                                                                      | :heavy_check_mark:                                                                                                                                                | N/A                                                                                                                                                               |                                                                                                                                                                   |
| `opts`                                                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                                                          | :heavy_minus_sign:                                                                                                                                                | The options for this request.                                                                                                                                     |                                                                                                                                                                   |

### Response

**[*operations.CreateAssessmentResponse](../../models/operations/createassessmentresponse.md), error**

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

Returns the highest assessment_version, with the account's current risk. Requires the assess or history scope. While status is pending, or partial with a source still running, the response carries Retry-After; poll no faster than it says. Another workspace's assessment, an unknown or expired assessment and a malformed assessment_id all answer 404 not_found. response_mode=detailed needs the detailed scope.

### Example Usage: complete

<!-- UsageSnippet language="go" operationID="getAssessment" method="get" path="/v1/assessments/{assessment_id}" example="complete" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v6"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Assessments.Get(ctx, "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: detailed

<!-- UsageSnippet language="go" operationID="getAssessment" method="get" path="/v1/assessments/{assessment_id}" example="detailed" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v6"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Assessments.Get(ctx, "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: login

<!-- UsageSnippet language="go" operationID="getAssessment" method="get" path="/v1/assessments/{assessment_id}" example="login" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v6"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Assessments.Get(ctx, "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: partial

<!-- UsageSnippet language="go" operationID="getAssessment" method="get" path="/v1/assessments/{assessment_id}" example="partial" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v6"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Assessments.Get(ctx, "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: pending

<!-- UsageSnippet language="go" operationID="getAssessment" method="get" path="/v1/assessments/{assessment_id}" example="pending" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v6"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Assessments.Get(ctx, "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: pending_known_account

<!-- UsageSnippet language="go" operationID="getAssessment" method="get" path="/v1/assessments/{assessment_id}" example="pending_known_account" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v6"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Assessments.Get(ctx, "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                         | Type                                                                                                              | Required                                                                                                          | Description                                                                                                       | Example                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                             | :heavy_check_mark:                                                                                                | The context to use for the request.                                                                               |                                                                                                                   |
| `assessmentID`                                                                                                    | `string`                                                                                                          | :heavy_check_mark:                                                                                                | Path segment, matched case-sensitively. A value that does not match the assessment_id shape answers 404, not 400. | asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                    |
| `responseMode`                                                                                                    | [*components.ResponseMode](../../models/components/responsemode.md)                                               | :heavy_minus_sign:                                                                                                | Optional standard or separately authorised detailed customer view; raw is rejected.                               |                                                                                                                   |
| `opts`                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                          | :heavy_minus_sign:                                                                                                | The options for this request.                                                                                     |                                                                                                                   |

### Response

**[*operations.GetAssessmentResponse](../../models/operations/getassessmentresponse.md), error**

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

Deletes one assessment in the environment of your key, with its versions, evidence and outcomes and the identity-graph marks its outcomes set, and scores its account again from what remains. Afterwards the assessment answers 404 and leaves the account's history. The answer names the account and the assessment its current risk now comes from; current_assessment_id is absent when the account has no assessment with a decision left. Requires the history scope. An unknown, expired or already deleted assessment answers 404, so a repeat answers 404; to erase a whole account idempotently, use accounts.delete.

### Example Usage: last

<!-- UsageSnippet language="go" operationID="deleteAssessment" method="delete" path="/v1/assessments/{assessment_id}" example="last" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v6"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Assessments.Delete(ctx, "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: next

<!-- UsageSnippet language="go" operationID="deleteAssessment" method="delete" path="/v1/assessments/{assessment_id}" example="next" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v6"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Assessments.Delete(ctx, "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                         | Type                                                                                                              | Required                                                                                                          | Description                                                                                                       | Example                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                             | :heavy_check_mark:                                                                                                | The context to use for the request.                                                                               |                                                                                                                   |
| `assessmentID`                                                                                                    | `string`                                                                                                          | :heavy_check_mark:                                                                                                | Path segment, matched case-sensitively. A value that does not match the assessment_id shape answers 404, not 400. | asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                    |
| `opts`                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                          | :heavy_minus_sign:                                                                                                | The options for this request.                                                                                     |                                                                                                                   |

### Response

**[*operations.DeleteAssessmentResponse](../../models/operations/deleteassessmentresponse.md), error**

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

Records what you later learned about an assessed account: confirmed_abuse, account_takeover and chargeback report abuse; false_positive and legitimate withdraw your earlier report; unknown is recorded only. Outcomes are append-only and the latest one wins. Requires the assess scope. Idempotency-Key is required and handled as on POST /v1/assessments: a retry with the same key and body answers the same record with Idempotent-Replayed: true; a different body answers 409. An unknown, expired or foreign assessment answers 404 not_found.

### Example Usage: abuse

<!-- UsageSnippet language="go" operationID="recordAssessmentOutcome" method="post" path="/v1/assessments/{assessment_id}/outcomes" example="abuse" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v6"
	"github.com/noresponsehq/go/v6/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Assessments.RecordOutcome(ctx, "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "account-acct_7f3a9c", components.AssessmentOutcomeRequest{
        Outcome: components.OutcomeConfirmedAbuse,
        ObservedAt: noresponse.Pointer("2026-10-20T05:00:00Z"),
        ExternalRef: noresponse.Pointer("case-123"),
        Notes: noresponse.Pointer("Card testing from this account; closed by trust and safety."),
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

<!-- UsageSnippet language="go" operationID="recordAssessmentOutcome" method="post" path="/v1/assessments/{assessment_id}/outcomes" example="recorded" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v6"
	"github.com/noresponsehq/go/v6/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Assessments.RecordOutcome(ctx, "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "account-acct_7f3a9c", components.AssessmentOutcomeRequest{
        Outcome: components.OutcomeLegitimate,
        ObservedAt: noresponse.Pointer("2026-10-14T09:04:30.000Z"),
        ExternalRef: noresponse.Pointer("case-48213"),
        Notes: noresponse.Pointer("Chargeback filed two days after account creation."),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: withdrawn

<!-- UsageSnippet language="go" operationID="recordAssessmentOutcome" method="post" path="/v1/assessments/{assessment_id}/outcomes" example="withdrawn" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v6"
	"github.com/noresponsehq/go/v6/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Assessments.RecordOutcome(ctx, "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "account-acct_7f3a9c", components.AssessmentOutcomeRequest{
        Outcome: components.OutcomeFalsePositive,
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
| `assessmentID`                                                                                                                                                    | `string`                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                | Path segment, matched case-sensitively. A value that does not match the assessment_id shape answers 404, not 400.                                                 | asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                    |
| `idempotencyKey`                                                                                                                                                  | `string`                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                | Opaque client key, 16 to 128 characters of [A-Za-z0-9._:-] (a UUID or ULID fits). Scoped to the workspace principal and bound to the canonical body for 24 hours. | account-acct_7f3a9c                                                                                                                                               |
| `body`                                                                                                                                                            | [components.AssessmentOutcomeRequest](../../models/components/assessmentoutcomerequest.md)                                                                        | :heavy_check_mark:                                                                                                                                                | N/A                                                                                                                                                               |                                                                                                                                                                   |
| `opts`                                                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                                                          | :heavy_minus_sign:                                                                                                                                                | The options for this request.                                                                                                                                     |                                                                                                                                                                   |

### Response

**[*operations.RecordAssessmentOutcomeResponse](../../models/operations/recordassessmentoutcomeresponse.md), error**

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

## Query

History is a POST with account_id in the body so an account identifier stays out of URLs and access logs. Requires the history scope. Read-only and safe to retry. Newest first, each at its latest assessment_version. An account with no assessments answers 200 with empty items. The cursor is opaque and bound to your workspace and the account; a tampered or expired cursor answers 400 invalid_cursor. response_mode=detailed needs the detailed scope.

### Example Usage: firstPage

<!-- UsageSnippet language="go" operationID="queryAccountAssessments" method="post" path="/v1/assessments/query" example="firstPage" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v6"
	"github.com/noresponsehq/go/v6/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Assessments.Query(ctx, components.AssessmentQuery{
        AccountID: "acct_7f3a9c",
    })
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
### Example Usage: page

<!-- UsageSnippet language="go" operationID="queryAccountAssessments" method="post" path="/v1/assessments/query" example="page" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v6"
	"github.com/noresponsehq/go/v6/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Assessments.Query(ctx, components.AssessmentQuery{
        AccountID: "acct_7f3a9c",
        Cursor: noresponse.Pointer("eyJvIjoyMH0"),
    })
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [components.AssessmentQuery](../../models/components/assessmentquery.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.QueryAccountAssessmentsResponse](../../models/operations/queryaccountassessmentsresponse.md), error**

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