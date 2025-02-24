<!--
Author: Ashwin Nair
Date: 2025-02-24
Project name: latency.md
Summary: Enter summary here.
-->

# Latency with Golang

-   Latency refers to the time taken for a request to travel from the **client to
    the server and back.**

---

## Necessary Features

-   Login
-   Settings Page
    -   Enter company logo (displayed on navbar on dashboard)
    -   Enter Endpoints
    -   Dark or light theme
    -   Number portal (for notification about high latency > 500)
    -   email portal (for same notification above)
    -   downloadable logs
-   SQLite Database
    -   Holds user info and company logo
    -   Holds endpoints
-   Dashboard
    -   Displays user info
-   Written in Golang, with maybe some js framework for making the site look 
    better.

---

### What is Middleware?
-   Middleware is code that acts as a bridge between different parts of an
    application, handling tasks like **request processing, authentication, and 
    error handling** before reaching the final destination.
-   What constitutes middleware?
    -   A request passes through one or more middleware functions
    -   Each middleware modifies the request, pass it to the next middleware, or
        stopping the request if he needed to.
    -   Finally, the request reaches the **main handler** (e.g., a route handler
        in a web server).
    -   After the response is generated, it also passes through middleware before
        being sent back to the client.
        -   **Logging** - Recording request details (e.g. IP, method, endpoint).
        -   **Authentication and Authorization** - Checking if the user is
            logged in or has permissions.
        -   **Request Validation** - Ensuring request data is valid before
            processing.
        -   **Rate limiting** - Preventing abuse by limiting requests per user.
        -   **Compressions** - Reducing response size using Gzip or Brotli.
        -   **Error Handling** - Catching and formatting errors before they
            reach the client.
---

### How Does Network Latency Work?

When a device (client) sends a request to another device (server), the following
steps occur:
-   Packet transmission: Data is broken into **packets** and sent over the
    network.
-   Routing: The packets travel through multiple **routers** and **switches** to
    reach the destination.
-   Processing: The destination device processes the request and generates a
    response.
-   Return Trip: The response travels back through the network to the client.

The total time taken for this process is the latency.

---

### How is Latency Measured?

-   Latency is usually measured in milliseconds (ms), and can broken down into
    different components:
    -   **DNS Resolution Time** - Time taken to resolve the domain name.
    -   **TCP Connection Time** - Time taken to establish a TCP connection.
    -   **TLS Handshake time (only if using HTTPS)** - Time to negotiate
        encryption.
    -   **Request Time** - Time taken for the server to receive the request.
    -   **Processing Time** - Time spent processing the request (in the backend).
    -   **Response Time** - Time taken to send the response back to the client.
    -   **Total Round-Trip Time** - The complete time taken from request
        initiation to response reception.
-   In GoLang, latency can be measured using `time.Since(startTime)`, where
    `startTime` is the moment where a request is received.

---

### Measuring Latency with GoLang.

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func latencyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Printf("Request to %s took %v", r.URL.Path, duration)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	
	// Wrap with latency middleware
	loggedMux := latencyMiddleware(mux)
	
	server := &http.Server{
		Addr:    ":8080",
		Handler: loggedMux,
	}

	log.Println("Server running on port 8080")
	log.Fatal(server.ListenAndServe())
}

```

---

### Features For a Latency Dashboard in Go

-   A latency dashboard for monitoring web services in GoLang should include:
    -   Real Time Metrics:
        -   Average response time per endpoint
        -   95th and 99th percentile latencies
    -   Graphical Visualization
        -   Line charts for Latency trends
        -   Histograms for request distribution
    -   Alerting System
        -   Threshold-based alerts (e.g. notify if latency exceed 500ms)
    -   Requesting Breakdown
        -   Separating DNS resolution, processing, and response latencies
    -   Endpoint-Specific Monitoring
        -   Latency per API route (e.g. `/api/login`, `/api/users`)
    -   Exporting Metrics
        -   Integration with **Prometheus** and **Grafana** for monitoring

---

### Implementing a Latency Dashboard in GoLang

```go
package main

import (
    "net/http"
    "time"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var requestDuration = prometheus.NewHistogramVec(
    prometheus.HistogramOpts{
        Name:    "http_request_duration_seconds",
        Help:    "Histogram of response time for handler in seconds",
        Buckets: prometheus.DefBuckets,
    },
    []string{"path"},
)

func init() {
    prometheus.MustRegister(requestDuration)
}

func latencyMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        duration := time.Since(start).Seconds()
        requestDuration.WithLabelValues(r.URL.Path).Observe(duration)
    })
}

func handler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", handler)

    http.Handle("/metrics", promhttp.Handler())
    http.Handle("/", latencyMiddleware(mux))

    http.ListenAndServe(":8080", nil)
}
```

---
