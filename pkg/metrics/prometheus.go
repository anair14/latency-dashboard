package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
)

var RequestDuration = prometheus.NewHistogramVec(
    prometheus.HistogramOpts{
        Name:   "http_request_duration_seconds",
        Help:   "Histogram of response time for handler (in seconds)",
        Buckets: prometheus.DefBuckets,
    },
    []string{"path"},
)

func init() {
    prometheus.MustRegister(RequestDuration)
}
