package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "latency-dashboard/pkg/database"
    "latency-dashboard/pkg/handlers"
    "latency-dashboard/pkg/middleware"
    "latency-dasbhaord/pkg/metrics"

    "github.com/prometheus/client_golang/prometheus/promhttp"

)

func main() {
    // Initialize DB
    database.InitDB()

    // Create a new ServeMux
    mux := http.NewServeMux()

    // Define routes
    mux.HandleFunc("/", handlers.HomeHandler)
    mux.HandleFunc("/", handlers.SettingsHandler)
    mux.Handle("/metrics", promhttp.Handler())
