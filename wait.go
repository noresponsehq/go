package noresponse

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/noresponsehq/go/v5/models/components"
	"github.com/noresponsehq/go/v5/models/operations"
)

const DefaultWaitTimeout = 30 * time.Second

// WaitTimeoutError reports the last pending assessment when the wait deadline
// expires. The webhook may deliver the eventual result later.
type WaitTimeoutError struct {
	AssessmentID string
	Last         components.Assessment
}

func (e *WaitTimeoutError) Error() string {
	return fmt.Sprintf("assessment %s is still pending", e.AssessmentID)
}

// WaitForAssessment polls GET until the assessment's status is complete. It
// honours Retry-After, bounds each request by the remaining wait time, and
// stops promptly when ctx is cancelled.
//
// The wait lasts until ctx's deadline, or DefaultWaitTimeout when ctx has
// none. When it runs out, a partial answer is returned as it stands, with
// Status still components.StatusPartial and a nil error, and a pending one
// returns a *WaitTimeoutError, since a pending assessment has no
// recommendation yet.
func WaitForAssessment(ctx context.Context, client *NoResponse, assessmentID string, responseMode *components.ResponseMode) (*components.Assessment, error) {
	if client == nil || client.Assessments == nil {
		return nil, fmt.Errorf("noresponse: client is nil")
	}
	waitCtx := ctx
	cancel := func() {}
	if _, hasDeadline := ctx.Deadline(); !hasDeadline {
		waitCtx, cancel = context.WithTimeout(ctx, DefaultWaitTimeout)
	}
	defer cancel()

	var last *components.Assessment
	for {
		remaining := time.Until(deadlineOf(waitCtx))
		if remaining <= 0 {
			// The deadline can pass before the context's timer fires, while
			// waitCtx.Err() is still nil, so never return that nil.
			if last != nil && ctx.Err() != context.Canceled {
				return timedOut(assessmentID, last)
			}
			if err := ctx.Err(); err != nil {
				return nil, err
			}
			return nil, context.DeadlineExceeded
		}
		response, err := client.Assessments.Get(waitCtx, assessmentID, responseMode, operations.WithOperationTimeout(remaining))
		if err != nil {
			// The last poll ran into the wait's deadline: a partial answer
			// still stands. A cancelled ctx is the caller's error.
			if last != nil && last.Status == components.StatusPartial && deadlinePassed(waitCtx) {
				return last, nil
			}
			return nil, err
		}
		if response.Result.Status == components.StatusComplete {
			return &response.Result, nil
		}
		last = &response.Result
		delay := retryAfter(response.Headers)
		if time.Until(deadlineOf(waitCtx)) <= delay {
			return timedOut(assessmentID, last)
		}
		timer := time.NewTimer(delay)
		select {
		case <-waitCtx.Done():
			timer.Stop()
			if waitCtx.Err() == context.DeadlineExceeded {
				return timedOut(assessmentID, last)
			}
			return nil, waitCtx.Err()
		case <-timer.C:
		}
	}
}

// timedOut ends a wait whose time ran out: a partial answer stands, and a
// pending one is a *WaitTimeoutError.
func timedOut(assessmentID string, last *components.Assessment) (*components.Assessment, error) {
	if last.Status == components.StatusPartial {
		return last, nil
	}
	return nil, &WaitTimeoutError{AssessmentID: assessmentID, Last: *last}
}

// deadlinePassed reports whether the wait ran out of time rather than being
// cancelled.
func deadlinePassed(ctx context.Context) bool {
	switch ctx.Err() {
	case context.DeadlineExceeded:
		return true
	case nil:
		return time.Until(deadlineOf(ctx)) <= 0
	}
	return false
}

func deadlineOf(ctx context.Context) time.Time {
	if deadline, ok := ctx.Deadline(); ok {
		return deadline
	}
	return time.Now().Add(DefaultWaitTimeout)
}

func retryAfter(headers map[string][]string) time.Duration {
	for name, values := range headers {
		if !strings.EqualFold(name, "retry-after") || len(values) == 0 {
			continue
		}
		if seconds, err := strconv.ParseInt(strings.TrimSpace(values[0]), 10, 64); err == nil {
			if seconds < 1 {
				seconds = 1
			}
			return time.Duration(seconds) * time.Second
		}
		if date, err := http.ParseTime(values[0]); err == nil {
			delay := time.Until(date)
			if delay < time.Second {
				return time.Second
			}
			return delay
		}
	}
	return time.Second
}
