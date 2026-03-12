package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type App struct {
	config Config
}

func (a *App) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}

type VersionResponse struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuildTime string `json:"build_time"`
	Env       string `json:"env"`
}

func (a *App) VersionHandler(w http.ResponseWriter, r *http.Request) {
	response := VersionResponse{
		Version:   a.config.Version,
		Commit:    a.config.Commit,
		BuildTime: a.config.BuildTime,
		Env:       a.config.Env,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type FeatureFlagsResponse struct {
	FeatureHello bool `json:"feature_hello"`
}

func (a *App) FeatureFlagsHandler(w http.ResponseWriter, r *http.Request) {
	response := FeatureFlagsResponse{
		FeatureHello: a.config.FeatureHello,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type HelloResponse struct {
	Message string `json:"message"`
}

func (a *App) HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := HelloResponse{
		Message: "Hello, ODevs!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (a *App) SlowHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(30 * time.Second)

	http.Error(w, "server is hanging", http.StatusServiceUnavailable)
}
