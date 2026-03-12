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

func TestHealthHandler(t *testing.T) {
	app := newApp(Config{})
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	app.HealthHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var resp map[string]string
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp["status"] != "ok" {
		t.Errorf("expected status ok, got %q", resp["status"])
	}
}

func TestVersionHandler(t *testing.T) {
	app := newApp(Config{Version: "1.0.0", Commit: "abc1234", Env: "test"})
	req := httptest.NewRequest(http.MethodGet, "/version", nil)
	w := httptest.NewRecorder()

	app.VersionHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var resp VersionResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Version != "1.0.0" {
		t.Errorf("expected version 1.0.0, got %q", resp.Version)
	}
	if resp.Commit != "abc1234" {
		t.Errorf("expected commit abc1234, got %q", resp.Commit)
	}
	if resp.Env != "test" {
		t.Errorf("expected env test, got %q", resp.Env)
	}
}

func TestFeatureFlagsHandler(t *testing.T) {
	app := newApp(Config{FeatureHello: true})
	req := httptest.NewRequest(http.MethodGet, "/feature-flags", nil)
	w := httptest.NewRecorder()

	app.FeatureFlagsHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var resp FeatureFlagsResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if !resp.FeatureHello {
		t.Error("expected feature_hello to be true")
	}
}

func TestHelloHandler_FlagEnabled(t *testing.T) {
	app := newApp(Config{FeatureHello: true})
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
	app := newApp(Config{FeatureHello: false})
	mux := http.NewServeMux()

	if app.config.FeatureHello {
		mux.HandleFunc("GET /hello", app.HelloHandler)
	}

	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404 when FEATURE_HELLO is disabled, got %d", w.Code)
	}
}
