package routes

import (
	"database/sql"
	"net/http"

	"github.com/anair14/latency-dashboard/pkg/database"
)

// SettingsHandler displays the settings page
func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/settings.html")
}

// AddEndpointHandler adds a new endpoint
func AddEndpointHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/settings", http.StatusSeeOther)
		return
	}

	url := r.FormValue("url")
	threshold := r.FormValue("threshold")

	_, err := database.DB.Exec("INSERT INTO endpoints (url, threshold) VALUES (?, ?)", url, threshold)
	if err != nil {
		http.Error(w, "Failed to add endpoint", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/settings", http.StatusFound)
}

// DeleteEndpointHandler deletes an endpoint
func DeleteEndpointHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	_, err := database.DB.Exec("DELETE FROM endpoints WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Failed to delete endpoint", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/settings", http.StatusFound)
}

