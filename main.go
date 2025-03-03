package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/anair14/latency-dashboard/pkg/database"
	"github.com/anair14/latency-dashboard/pkg/middleware"
    "github.com/anair14/latency-dashboard/routes"
)

func main() {
	// Initialize the database
	database.InitDB()

	// Create a new router
	r := mux.NewRouter()

	// Apply authentication middleware
	r.Use(middleware.AuthMiddleware)

	// Register routes
	routes.RegisterRoutes(r)

	// Start the server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed:", err)
	}
}

