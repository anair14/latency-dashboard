package handlers

import (
    "database/sql"
    "log"
    "net/http"
    "github.com/gorillas/sessions"
    "github.com/anair14/latency-dashboard/templates"
    "github.com/anair14/latency-dashboard/config"

    _ "github.com/mattn/go-sqlite3"
)

// SettingsPage handles displaying the settings page
func SettingsPage(w http.ResponseWriter, r *http.Request) {
    userID, authenticated := session.GetUserID(r)
    if !authenticated {
        http.Redirect(w, r, "/login", http.StatusFound)
        return
    }

    db, err := sql.Open("sqlite3", config.DBPath)
    if err != nil {
        log.Println("Error opening database:", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    endpoints, err := getUserEndpoints(db, userID)
    if err != nil {
        log.Println("Error fetching endpoints:", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    templates.RenderTemplate(w, "settings.html", map[string]interface{}{
        "Endpoints": endpoints,
    })
}

// AddEndpoint handles adding a new endpoint
func AddEndpoint(w http.ResponseWriter, r *http.Request) {
    userID, authenticated := session.GetUserID(r)
    if !authenticated {
        http.Redirect(w, r, "/login", http.StatusFound)
        return
    }

    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    endpoint := r.FormValue("endpoint")
    if endpoint == "" {
        http.Error(w, "Endpoint URL cannot be empty", http.StatusBadRequest)
        return
    }

    db, err := sql.Open("sqlite3", config.DBPath)
    if err != nil {
        log.Println("Error opening database:", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    _, err = db.Exec("INSERT INTO endpoints (user_id, url, threshold) VALUES (?, ?, ?)", userID, endpoint, config.DefaultLatencyThreshold)
    if err != nil {
        log.Println("Error adding endpoint:", err)
        http.Error(w, "Failed to add endpoint", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/settings", http.StatusSeeOther)
}

// RemoveEndpoint handles removing an endpoint
func RemoveEndpoint(w http.ResponseWriter, r *http.Request) {
    userID, authenticated := session.GetUserID(r)
    if !authenticated {
        http.Redirect(w, r, "/login", http.StatusFound)
        return
    }

    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    endpointID := r.FormValue("endpoint_id")
    if endpointID == "" {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    db, err := sql.Open("sqlite3", config.DBPath)
    if err != nil {
        log.Println("Error opening database:", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    _, err = db.Exec("DELETE FROM endpoints WHERE id = ? AND user_id = ?", endpointID, userID)
    if err != nil {
        log.Println("Error removing endpoint:", err)
        http.Error(w, "Failed to remove endpoint", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/settings", http.StatusSeeOther)
}

// Helper function to get endpoints for a specific user
func getUserEndpoints(db *sql.DB, userID int) ([]map[string]string, error) {
    rows, err := db.Query("SELECT id, url, threshold FROM endpoints WHERE user_id = ?", userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var endpoints []map[string]string
    for rows.Next() {
        var id, threshold int
        var url string
        if err := rows.Scan(&id, &url, &threshold); err != nil {
            return nil, err
        }
        endpoints = append(endpoints, map[string]string{
            "ID":        string(id),
            "URL":       url,
            "Threshold": string(threshold),
        })
    }
    return endpoints, nil
}
