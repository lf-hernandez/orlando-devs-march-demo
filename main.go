package main

import "net/http"

func main() {
	http.HandleFunc("/healthz", HealthHandler)
	http.ListenAndServe(":5000", nil)
}
