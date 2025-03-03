package middleware

import (
	"net/http"
	"github.com/anair14/latency-dashboard/routes"
)

// AuthMiddleware ensures a user is authenticated before accessing certain routes
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := routes.Store.Get(r, "session")
		if err != nil || session.Values["authenticated"] != true {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

