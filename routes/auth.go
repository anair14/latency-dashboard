package routes
import (
	"database/sql"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"github.com/anair14/latency-dashboard/pkg/database"
)

var Store = sessions.NewCookieStore([]byte("super-secret-key"))

// LoginHandler processes user login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.ServeFile(w, r, "templates/login.html")
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	var storedPassword string
	err := database.DB.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedPassword)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	session, _ := Store.Get(r, "session")
	session.Values["authenticated"] = true
	session.Save(r, w)

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

// LogoutHandler logs the user out
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session")
	session.Values["authenticated"] = false
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusFound)
}

