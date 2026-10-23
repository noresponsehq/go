package noresponse_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	noresponse "github.com/noresponsehq/go/v5"
	"github.com/noresponsehq/go/v5/models/apierrors"
	"github.com/noresponsehq/go/v5/models/components"
)

// These examples are compiled by go test and shown on pkg.go.dev. They have
// no Output comment, so they are never run against the live API.

func Example() {
	ctx := context.Background()
	client := noresponse.NewClient(os.Getenv("NORESPONSE_API_KEY"))

	res, err := client.Assessments.Create(ctx, noresponse.NewIdempotencyKey(), components.AssessmentRequest{
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
		// The account's first assessment is still pending: wait for the
		// webhook or poll, and apply your own fallback in the meantime.
		return
	}
	switch assessment.Account.Risk.Decision {
	case components.AccountDecisionBlock:
		fmt.Println("block")
	case components.AccountDecisionReview, components.AccountDecisionStepUp:
		fmt.Println("challenge", assessment.Account.Risk.Score)
	default:
		fmt.Println("allow")
	}
}

func ExampleAssessments_Get() {
	client := noresponse.NewClient(os.Getenv("NORESPONSE_API_KEY"))

	res, err := client.Assessments.Get(context.Background(), "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Result.Status)
}

func ExampleAssessments_RecordOutcome() {
	client := noresponse.NewClient(os.Getenv("NORESPONSE_API_KEY"))

	_, err := client.Assessments.RecordOutcome(context.Background(),
		"asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM",
		noresponse.NewIdempotencyKey(),
		components.AssessmentOutcomeRequest{Outcome: components.OutcomeConfirmedAbuse},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleAllAssessments() {
	client := noresponse.NewClient(os.Getenv("NORESPONSE_API_KEY"))

	query := components.AssessmentQuery{AccountID: "acct_7f3a9c"}
	for assessment, err := range noresponse.AllAssessments(context.Background(), client, query) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(assessment.AssessmentID, assessment.Status)
	}
}

func ExampleWaitForAssessment() {
	client := noresponse.NewClient(os.Getenv("NORESPONSE_API_KEY"))

	assessment, err := noresponse.WaitForAssessment(context.Background(), client, "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM", nil)
	var timeout *noresponse.WaitTimeoutError
	if errors.As(err, &timeout) {
		// Still pending: the webhook delivers the result later.
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(assessment.Status)
}

func ExampleNewClient_errors() {
	client := noresponse.NewClient(os.Getenv("NORESPONSE_API_KEY"))

	_, err := client.Assessments.Get(context.Background(), "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM", nil)

	var notFound *apierrors.NotFoundError
	var limited *apierrors.RateLimitError
	var other *apierrors.APIStatusError
	switch {
	case errors.As(err, &notFound):
		fmt.Println("no such assessment, request", notFound.RequestID)
	case errors.As(err, &limited):
		fmt.Println("rate limited, request", limited.RequestID)
	case errors.As(err, &other):
		fmt.Println("HTTP", other.StatusCode)
	case err != nil:
		// A network error or a cancelled context.
		log.Fatal(err)
	}
}

func ExampleConstructEvent() {
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
			fmt.Println(data.ID, data.Detail.AssessmentVersion)
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
}
