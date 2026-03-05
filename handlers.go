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

	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
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

	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
