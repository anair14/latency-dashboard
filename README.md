# latency-dashboard
A latency dashboard, made for adding endpoints and customizable settings.

Architecture:
```go
/latency-dashboard
├── /cmd
│   └── /server
│       └── main.go                  # Main entry point to start the server
├── /pkg
│   └── /middleware
│       └── latency.go               # Middleware to measure latency
│   └── /handlers
│       └── handler.go               # Request handler logic
│   └── /metrics
│       └── prometheus.go            # Prometheus metrics integration
├── /assets
│   └── /css
│       └── style.css                # CSS files for UI styling (light/dark theme)
│   └── /js
│       └── script.js                # JavaScript for frontend interaction
├── /templates
│   └── index.html                  # HTML template for the dashboard
│   └── settings.html               # Settings page for user customization
├── /config
│   └── config.go                   # Configuration file for settings (e.g., DB, logo, themes)
├── /database
│   └── db.go                       # SQLite database connection and schema setup
├── /logs
│   └── server.log                  # Log files for server activity
├── /prometheus
│   └── prometheus_config.yml       # Prometheus config to scrape metrics
├── /grafana
│   └── grafana_dashboard.json      # Grafana dashboard JSON template
├── /docs
│   └── README.md                   # Project documentation
└── go.mod                           # Go module dependencies
└── go.sum                           # Go module checksum

```
