package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type App struct {
	config Config
}

func (a *App) RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	flagStatus := func(enabled bool) string {
		if enabled {
			return `<span style="color:green">&#10003; enabled</span>`
		}
		return `<span style="color:gray">&#10007; disabled</span>`
	}

	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head><title>ODevs Demo</title></head>
<body style="font-family:monospace;max-width:600px;margin:40px auto;padding:0 20px">
  <h1>Orlando Devs Demo</h1>
  <h2>Build</h2>
  <table>
    <tr><td>version</td><td>&nbsp;&nbsp;</td><td><strong>%s</strong></td></tr>
    <tr><td>commit</td><td></td><td><strong>%s</strong></td></tr>
    <tr><td>built</td><td></td><td>%s</td></tr>
    <tr><td>env</td><td></td><td>%s</td></tr>
  </table>
  <h2>Feature Flags</h2>
  <table>
    <tr><td>FEATURE_HELLO</td><td>&nbsp;&nbsp;</td><td>%s</td></tr>
  </table>
</body>
</html>`,
		a.config.Version,
		a.config.Commit,
		a.config.BuildTime,
		a.config.Env,
		flagStatus(a.config.FeatureHello),
	)
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
