package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	rootHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", resp.StatusCode)
	}

	body := w.Body.String()

	if !strings.Contains(body, "Go HTMX Tailwind Starter") {
		t.Errorf("expected response body to contain 'Go HTMX Tailwind Starter', got:\n%s", body)
	}
}
