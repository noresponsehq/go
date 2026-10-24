<img src="https://raw.githubusercontent.com/noresponsehq/.github/main/profile/sdk-header.png" alt="NoResponse" width="100%">

# NoResponse Go SDK

The Go client for the [NoResponse API](https://docs.noresponse.org): assess an account creation, login or purchase, wait for a pending result, page through an account's history, record what you later learned and verify webhooks.

Call it from your server only. The API key must never reach a browser or an app.

## Install

```sh
go get github.com/noresponsehq/go/v7
```

Go 1.24 or later.

## Quick start

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	noresponse "github.com/noresponsehq/go/v7"
	"github.com/noresponsehq/go/v7/models/components"
)

func main() {
	client := noresponse.NewClient(os.Getenv("NORESPONSE_API_KEY"))

	res, err := client.Assessments.Create(context.Background(), noresponse.NewIdempotencyKey(), components.AssessmentRequest{
		AccountID: "acct_7f3a9c",
		Flow:      components.FlowAccountCreation,
		Identifiers: components.AssessmentRequestIdentifiers{
			Email: noresponse.Pointer("new.user@example.com"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	assessment := res.Result
	if assessment.Account == nil {
		// The account's first assessment is still pending. Apply your own
		// fallback, then wait for the webhook or call noresponse.WaitForAssessment.
		return
	}
	fmt.Println(assessment.Account.Risk.Decision, assessment.Account.Risk.Score)
}
```

`NewClient` reads nothing from the environment itself; pass the key. `noresponse.New()` with no key option reads `NORESPONSE_API_KEY`.

Every request names its `Flow`: `components.FlowAccountCreation`, `FlowLogin` or `FlowPurchase`. Act on `assessment.Account.Risk.Decision`: `allow`, `step_up`, `review` or `block`, where `block` means your block threshold was reached. The assessment's own facts (score, recommendation, signals, `AssessmentVersion`) are under `assessment.Detail`. `Account` is nil while the account's first assessment is pending, and a 429, a 503 or a timeout also needs your own fallback.

Every method takes a `context.Context` first. Cancelling it cancels the request and any retry wait.

## Methods

| Call | Does |
|---|---|
| `client.Assessments.Create(ctx, idempotencyKey, request)` | Assess identifiers. Returns 200 (complete or partial) or 202 (pending) as `res.Result`. |
| `client.Assessments.Get(ctx, assessmentID, nil)` | Fetch the latest version of one assessment. |
| `client.Assessments.RecordOutcome(ctx, assessmentID, idempotencyKey, request)` | Report what you later learned, for example `components.OutcomeConfirmedAbuse`. |
| `client.Assessments.Delete(ctx, assessmentID)` | Delete one assessment (`history` scope). `res.Result.Account.CurrentAssessmentID` names the assessment the account's risk now comes from, nil when none is left. A repeat answers 404. |
| `client.Assessments.Query(ctx, query)` | One page of an account's history. `res.Next()` fetches the next page. |
| `noresponse.AllAssessments(ctx, client, query)` | Range over every page of an account's history. |
| `client.Accounts.Get(ctx, accountID)` | The account's one current risk: score, a decision for each flow and the assessment it comes from (`history` scope). An account with no decided assessment is a 404. |
| `client.Accounts.ListAssessments(ctx, accountID, limit, cursor, responseMode)` | One page of the account's history with the account id in the URL; pass `nil` for the defaults. `res.Next()` fetches the next page. |
| `client.Accounts.Delete(ctx, accountID)` | Erase the account and its history in your key's environment (`history` scope). Idempotent: a repeat answers `AssessmentsDeleted` 0. An `account.deleted` webhook follows when something was deleted. |
| `noresponse.WaitForAssessment(ctx, client, assessmentID, nil)` | Poll until the assessment is `complete`, honouring `Retry-After`, until `ctx`'s deadline (30 s when it has none). Then a `partial` answer is returned as it stands, and a `pending` one returns `*WaitTimeoutError`. |
| `noresponse.ConstructEvent(payload, headers, secrets)` | Verify a webhook delivery and return the typed event. |
| `client.Keys.Rotate(ctx, request)` | Rotate the API key that signs the request. |
| `client.Integrations.Create(ctx, request)` | Register a hosted WorkOS or Auth0 integration (`integrations` scope). For Auth0, `res.Result.Secret` is shown once. Never retried. |
| `client.Integrations.List(ctx)` | Every integration, oldest first, with no secrets. |
| `client.Integrations.Delete(ctx, integrationID)` | Delete an integration. |
| `client.Integrations.Rotate(ctx, integrationID, body)` | Replace an integration's secret. Never retried. |

Each response carries `Headers` (including `X-Request-ID`) and `Result`.

## Idempotency keys

`Create` and `RecordOutcome` require an idempotency key. The same key with the same body replays the stored answer and starts no new work. The same key with a different body is a 409. Use `noresponse.NewIdempotencyKey()` and store the key beside the signup so a retry after a crash sends the same one.

## Errors

Every failure is an `error` you can inspect with `errors.As`. Problem documents from the API are typed and carry the request id:

```go
_, err := client.Assessments.Get(ctx, id, nil)

var notFound *apierrors.NotFoundError
var limited *apierrors.RateLimitError
var other *apierrors.APIStatusError
switch {
case errors.As(err, &notFound):
	log.Println("no such assessment, request", notFound.RequestID)
case errors.As(err, &limited):
	log.Println("rate limited, request", limited.RequestID)
case errors.As(err, &other):
	log.Println("HTTP", other.StatusCode)
case err != nil:
	log.Fatal(err) // network failure or cancelled context
}
```

The typed errors are `BadRequestError` (400), `AuthenticationError` (401), `PermissionDeniedError` (403), `NotFoundError` (404), `ConflictError` (409) and `RateLimitError` (429), all in `github.com/noresponsehq/go/v7/models/apierrors`. Any other status is an `*apierrors.APIStatusError`.

## Retries and timeouts

The client retries 408, 429, 5xx and connection errors with exponential backoff and jitter: 500 ms first, 8 s at most, 30 s in total. The same idempotency key is sent each time, so a retry never repeats work. Change it for every call or for one:

```go
client := noresponse.NewClient(key,
	noresponse.WithTimeout(10*time.Second),
	noresponse.WithRetryConfig(retry.Config{
		Strategy:              "backoff",
		Backoff:               &retry.BackoffStrategy{InitialInterval: 200, MaxInterval: 2000, Exponent: 2, MaxElapsedTime: 5000}, // milliseconds
		RetryConnectionErrors: true,
	}),
)

res, err := client.Assessments.Get(ctx, id, nil, operations.WithRetries(retry.Config{Strategy: "none"}))
```

## Paging

```go
query := components.AssessmentQuery{AccountID: "acct_7f3a9c"}
for assessment, err := range noresponse.AllAssessments(ctx, client, query) {
	if err != nil {
		return err
	}
	fmt.Println(assessment.AssessmentID, assessment.Status)
}
```

## Issued keys

When your platform issues an API key to the new account, send only its hash, as `issued_credential.sha256`; the key itself never leaves your backend.

```go
sha256, err := noresponse.HashIssuedCredential(issuedKey) // 64 lower-case hex characters
```

## Webhooks

Register your endpoint with a key that has the `webhooks` scope. The signing secret is shown only in this answer, so store it, then send the verification challenge:

```go
created, err := client.WebhookEndpoints.Create(ctx, components.WebhookEndpointRequest{URL: "https://app.example.com/webhooks/noresponse"})
if err != nil {
	return err
}
secret := created.Result.SigningSecret // store secret.KeyID and secret.Secret
check, err := client.WebhookEndpoints.SendTestEvent(ctx, created.Result.Endpoint.EndpointID,
	components.WebhookTestRequest{Event: components.WebhookTestRequestEventEndpointVerification}) // check.Result.Verified
```

`List`, `RotateSigningKey` (the old key verifies for 7 more days) and `Delete` manage it afterwards. An `assessment.updated` test event arrives with `event.Test` true.

Pass the raw request body. Re-encoding parsed JSON breaks the signature.

```go
secrets := map[string]string{"whk_example0001": os.Getenv("NORESPONSE_WEBHOOK_SECRET")}

http.HandleFunc("/webhooks/noresponse", func(w http.ResponseWriter, r *http.Request) {
	payload, err := io.ReadAll(http.MaxBytesReader(w, r.Body, 1<<20))
	if err != nil {
		http.Error(w, "unreadable body", http.StatusBadRequest)
		return
	}
	event, err := noresponse.ConstructEvent(payload, r.Header, secrets)
	if err != nil {
		http.Error(w, "invalid signature", http.StatusBadRequest)
		return
	}
	switch data := event.Data.(type) {
	case *components.Assessment:
		if event.Test {
			break // SendTestEvent's fixed example assessment: apply nothing
		}
		// The full assessment, as Get returns it: act on data.Account.Risk.Decision.
	case *noresponse.AssessmentExpired:
		// Delete what you hold for data.AssessmentID.
	case *noresponse.AccountDeleted:
		// You erased data.AccountID with Accounts.Delete: delete what you hold for it.
	case *noresponse.WorkspaceDeleted:
		// The workspace is deleted: the last event this endpoint receives. Delete what you hold for data.Environment.
	case *noresponse.EndpointDisabled:
		// data.EndpointID failed for 72 hours and is disabled: fix it and verify it again.
	case *noresponse.AccountRiskChanged:
		// A flow's decision changed. data.Account is the account as Accounts.Get returns it,
		// and data.PreviousDecisionsByFlow holds the decisions it replaced.
	case *noresponse.EndpointVerification:
		body, err := noresponse.VerificationResponse(event)
		if err != nil {
			http.Error(w, "invalid verification", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(body)
		return
	}
	w.WriteHeader(http.StatusNoContent)
})
```

Deliveries arrive at least once and unordered. Deduplicate on `event.ID` and keep the highest `Detail.AssessmentVersion` per assessment.

### Late results

A later decision reaches you as `account.risk_changed`. The event only wakes your handler: `HandleRiskChange` verifies it with `ConstructEvent`, reads the account's current risk with `Accounts.Get` (the event's own decision is ignored), asks your `GetAccountState` where the account is in your flow, and acts through your identity provider:

```go
provider, err := noresponse.IdentityProvider("clerk", noresponse.IdentityProviderOptions{Token: os.Getenv("CLERK_SECRET_KEY")}) // or "workos", "auth0" with Domain
result, err := noresponse.HandleRiskChange(r.Context(), client, payload, r.Header, secrets, &noresponse.EnforcementOptions{
	Provider:        provider,
	GetAccountState: accountState, // func(ctx, components.Account) (*noresponse.AccountStateAnswer, error)
})
```

An `onboarding` account fails onboarding and a `provisional` one is refused activation on an `account_creation` decision of `review` or `block`; an `active` account has its sessions revoked on a `login` decision of `review` and is restricted and signed out on `block`. A `restricted` account is unrestricted when its `login` decision is `allow` again only if you set `Flows: map[components.Flow]noresponse.FlowPolicy{components.FlowLogin: {Reinstate: true}}`. Nothing deletes a user: Clerk bans (undone by unban), Auth0 sets `blocked`, and WorkOS, which has no block, writes `noresponse_state` to the user's metadata for your sign-in to check and revokes every session. Every action, or `none`, is recorded with `Accounts.RecordEnforcement` at the `RiskVersion` read; store `result.RiskVersion` and return it as `EnforcedRiskVersion` so a replayed or older event changes nothing. A refusal is recorded as failed and returned as `*ProviderActionError`; answer 5xx so the delivery is retried. `EnforceAccount` enforces an account from `Accounts.ListChanged` when catching up, and `CustomProvider` runs your own functions for your own user store.

## Other options

`noresponse.WithServerURL`, `noresponse.WithClient` (any `Do(*http.Request)` implementation) and `noresponse.WithSecuritySource` (a key fetched per call, for rotation) are available on both `New` and `NewClient`.


## Changes since 2.0.0

No name was renamed. New in the root package: `NewClient`, `WithAPIKey` (the same as `WithSecurity`), `NewIdempotencyKey` and `AllAssessments`. The module now needs Go 1.24, where 2.0.0 declared Go 1.26.8.
