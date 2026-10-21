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

	"github.com/noresponsehq/go/v2/models/components"
)

func TestWaitForAssessmentReturnsSettledResponse(t *testing.T) {
	var calls int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		body, err := os.ReadFile(filepath.Join("tests", "fixtures", "valid", "assessment-complete.json"))
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
		body, err := os.ReadFile(filepath.Join("tests", "fixtures", "valid", "assessment-pending.json"))
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
