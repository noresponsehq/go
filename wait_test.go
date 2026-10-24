package noresponse

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/noresponsehq/go/v7/models/components"
)

func TestWaitForAssessmentReturnsSettledResponse(t *testing.T) {
	var calls int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		body, err := os.ReadFile(filepath.Join("tests", "fixtures", "valid", "assessment-view-complete.json"))
		if err != nil {
			t.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(body)
	}))
	defer server.Close()

	client := New(WithServerURL(server.URL), WithSecurity("nrk_test_kkkkkkkkkkkkkkkkkkkk"))
	const id = "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM"
	assessment, err := WaitForAssessment(context.Background(), client, id, nil)
	if err != nil {
		t.Fatal(err)
	}
	if assessment.Status != components.StatusComplete || calls != 1 {
		t.Fatalf("status %s after %d calls", assessment.Status, calls)
	}
}

func TestWaitForAssessmentTimesOutWithLastPendingResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := os.ReadFile(filepath.Join("tests", "fixtures", "valid", "assessment-view-pending.json"))
		if err != nil {
			t.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Retry-After", "60")
		_, _ = w.Write(body)
	}))
	defer server.Close()

	client := New(WithServerURL(server.URL), WithSecurity("nrk_test_kkkkkkkkkkkkkkkkkkkk"))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	const id = "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM"
	_, err := WaitForAssessment(ctx, client, id, nil)
	var timeout *WaitTimeoutError
	if !errors.As(err, &timeout) || timeout.Last.Status != components.StatusPending {
		t.Fatalf("want pending WaitTimeoutError, got %v", err)
	}
}

func serveFixtures(t *testing.T, names ...string) (*httptest.Server, *int) {
	t.Helper()
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := names[len(names)-1]
		if calls < len(names) {
			name = names[calls]
		}
		calls++
		body, err := os.ReadFile(filepath.Join("tests", "fixtures", "valid", name))
		if err != nil {
			t.Error(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(body)
	}))
	t.Cleanup(server.Close)
	return server, &calls
}

func TestWaitForAssessmentKeepsPollingPartialUntilComplete(t *testing.T) {
	server, calls := serveFixtures(t, "assessment-view-partial.json", "assessment-view-partial.json", "assessment-view-complete.json")
	client := New(WithServerURL(server.URL), WithSecurity("nrk_test_kkkkkkkkkkkkkkkkkkkk"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	assessment, err := WaitForAssessment(ctx, client, "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM", nil)
	if err != nil {
		t.Fatal(err)
	}
	if assessment.Status != components.StatusComplete || *calls != 3 {
		t.Fatalf("status %s after %d calls", assessment.Status, *calls)
	}
}

func TestWaitForAssessmentReturnsLatestPartialAtTimeout(t *testing.T) {
	server, calls := serveFixtures(t, "assessment-view-pending.json", "assessment-view-partial.json")
	client := New(WithServerURL(server.URL), WithSecurity("nrk_test_kkkkkkkkkkkkkkkkkkkk"))
	ctx, cancel := context.WithTimeout(context.Background(), 2500*time.Millisecond)
	defer cancel()
	started := time.Now()
	assessment, err := WaitForAssessment(ctx, client, "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM", nil)
	if err != nil {
		t.Fatalf("want the partial answer, got %v", err)
	}
	if assessment.Status != components.StatusPartial || *calls < 2 {
		t.Fatalf("status %s after %d calls", assessment.Status, *calls)
	}
	if elapsed := time.Since(started); elapsed > 3*time.Second {
		t.Fatalf("wait took %s", elapsed)
	}
}
