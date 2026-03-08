package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newApp(cfg Config) *App {
	return &App{config: cfg}
}

func TestHelloHandler_FlagEnabled(t *testing.T) {
	app := newApp(Config{FfA: true})
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()

	app.HelloHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var resp HelloResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Message == "" {
		t.Error("expected non-empty message")
	}
}

func TestHelloHandler_FlagDisabled(t *testing.T) {
	app := newApp(Config{FfA: false})
	mux := http.NewServeMux()

	if app.config.FfA {
		mux.HandleFunc("GET /hello", app.HelloHandler)
	}

	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404 when FF_A is disabled, got %d", w.Code)
	}
}
