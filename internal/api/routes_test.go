package api

import (
	"bytes"
	"mime/multipart"
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

func TestHandleAddRouteGet(t *testing.T) {
	urlMap := map[string]string{"abcd1234": "https://google.com/"}
	codeLength := 8

	handler := func(w http.ResponseWriter, r *http.Request) {
		HandleAddRoute(w, r, urlMap, codeLength)
	}

	req := httptest.NewRequest("GET", "/add", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf(`Expected Method Not Allowed, Got %d`, resp.StatusCode)
	}
}

func TestHandleAddRoutePostWrongContent(t *testing.T) {
	urlMap := map[string]string{"abcd1234": "https://google.com/"}
	codeLength := 8

	handler := func(w http.ResponseWriter, r *http.Request) {
		HandleAddRoute(w, r, urlMap, codeLength)
	}

	req := httptest.NewRequest("POST", "/add", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf(`Expected Bad Request, Got %v`, resp.Status)
	}
}

func TestHandleAddRoutePostNoUrl(t *testing.T) {
	urlMap := map[string]string{"abcd1234": "https://google.com/"}
	codeLength := 8

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	_ = writer.WriteField("noturl", "http")

	writer.Close()

	handler := func(w http.ResponseWriter, r *http.Request) {
		HandleAddRoute(w, r, urlMap, codeLength)
	}

	req := httptest.NewRequest("POST", "/add", &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf(`Expected Bad Request, Got %v`, resp.Status)
	}
}

func TestHandleAddRoutePostUrl(t *testing.T) {
	urlMap := map[string]string{"abcd1234": "https://google.com/"}
	codeLength := 8

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	_ = writer.WriteField("url", "https://google.com/")

	writer.Close()

	handler := func(w http.ResponseWriter, r *http.Request) {
		HandleAddRoute(w, r, urlMap, codeLength)
	}

	req := httptest.NewRequest("POST", "/add", &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusFound {
		t.Fatalf(`Expected Redirect, Got %v`, resp.Status)
	}
}
