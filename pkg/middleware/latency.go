package middleware

import (
    "log"
    "net/http"
    "time"

    "github.com/anair14/latency-dashboard/pkg/metrics"
)

func LatencyMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        duration := time.Since(start)

        // log Latency 
        log.Printf("Request to %s took %v", r.URL.Path, duration)

        // Observe latency in Prometheus
        metrics.RequestDuration.WithLabelValues(r.URL.Path).Observe(duration.Seconds())

    })

}
