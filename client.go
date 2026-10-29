package noresponse

import (
	"context"
	"crypto/rand"
	"fmt"
	"iter"

	"github.com/noresponsehq/go/v9/models/components"
)

// NewClient returns a client that authenticates with apiKey. It is New with
// WithAPIKey applied first, so later options can still override it.
//
//	client := noresponse.NewClient(os.Getenv("NORESPONSE_API_KEY"))
func NewClient(apiKey string, opts ...SDKOption) *NoResponse {
	return New(append([]SDKOption{WithAPIKey(apiKey)}, opts...)...)
}

// WithAPIKey sets the API key sent as the bearer token. Keep the key on your
// server. Without this option the client reads NORESPONSE_API_KEY.
func WithAPIKey(apiKey string) SDKOption {
	return WithSecurity(apiKey)
}

// NewIdempotencyKey returns a random key for Assessments.Create and
// Assessments.RecordOutcome: a version 4 UUID, which is 36 characters from
// the allowed set. Generate one per logical request and store it beside the
// account creation so that a retry after a crash sends the same key.
func NewIdempotencyKey() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		panic(fmt.Sprintf("noresponse: no randomness for an idempotency key: %v", err))
	}
	b[6] = b[6]&0x0f | 0x40
	b[8] = b[8]&0x3f | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

// AllAssessments iterates over every unexpired assessment of one account,
// newest first, fetching pages as it goes. It stops after the first error,
// which it yields once.
//
//	for assessment, err := range noresponse.AllAssessments(ctx, client, components.AssessmentQuery{AccountID: "acct_7f3a9c"}) {
//		if err != nil {
//			return err
//		}
//		fmt.Println(assessment.AssessmentID)
//	}
func AllAssessments(ctx context.Context, client *NoResponse, query components.AssessmentQuery) iter.Seq2[components.AssessmentDetail, error] {
	return func(yield func(components.AssessmentDetail, error) bool) {
		if client == nil || client.Assessments == nil {
			yield(components.AssessmentDetail{}, fmt.Errorf("noresponse: client is nil"))
			return
		}
		page, err := client.Assessments.Query(ctx, query)
		for {
			if err != nil {
				yield(components.AssessmentDetail{}, err)
				return
			}
			for _, item := range page.Result.Items {
				if !yield(item, nil) {
					return
				}
			}
			if page.Next == nil {
				return
			}
			next, nextErr := page.Next()
			if next == nil && nextErr == nil {
				return
			}
			page, err = next, nextErr
		}
	}
}
