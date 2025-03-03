package handlers

import (
    "fmt"
    "net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "Welcome to the Latency Dashboard!")
}

func SettingsHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "Settings Page - Work in progress")
}
