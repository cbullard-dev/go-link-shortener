package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleRootOrDefault(t *testing.T) {
	urlMap := map[string]string{"abcd1234": "http://google.com/"}

	handler := func(w http.ResponseWriter, r *http.Request) {
		HandleRootOrDefault(w, r, urlMap)
	}

	req := httptest.NewRequest("GET", "/abcd1234", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusFound {
		t.Errorf("Expected redirect, got %d", resp.StatusCode)
	}
}

