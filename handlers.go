package main

import (
	"encoding/json"
	"log/slog"
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
	w.Header().Set("Content-Type", "application/json")

	response := HealthResponse{
		Status:  "UP",
		Message: "Service is healthy",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("Failed to encode response", "error", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

type VersionResponse struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuildTime string `json:"buildtime"`
}

func (a *App) VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := VersionResponse{
		Version:   a.config.Version,
		Commit:    a.config.Commit,
		BuildTime: a.config.BuildTime,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("Failed to encode response", "error", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}
