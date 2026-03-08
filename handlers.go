package main

import (
	"encoding/json"
	"net/http"
)

type App struct {
	config Config
}

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (a *App) HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:  "UP",
		Message: "Service is healthy",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type VersionResponse struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuildTime string `json:"build_time"`
}

func (a *App) VersionHandler(w http.ResponseWriter, r *http.Request) {
	response := VersionResponse{
		Version:   a.config.Version,
		Commit:    a.config.Commit,
		BuildTime: a.config.BuildTime,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type FeatureFlagsResponse struct {
	FfA bool `json:"ff_a"`
	FfB bool `json:"ff_b"`
	FfC bool `json:"ff_c"`
}

func (a *App) FeatureFlagsHandler(w http.ResponseWriter, r *http.Request) {
	response := FeatureFlagsResponse{
		FfA: a.config.FfA,
		FfB: a.config.FfB,
		FfC: a.config.FfC,
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
