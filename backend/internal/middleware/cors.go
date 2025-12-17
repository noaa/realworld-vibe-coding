package middleware

import (
	"net/http"
)

// CORS middleware handles Cross-Origin Resource Sharing
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Origin header from the request
		origin := r.Header.Get("Origin")

		// List of allowed origins
		allowedOrigins := map[string]bool{
			"http://localhost:5173":                    true, // Local frontend
			"http://localhost:4173":                    true, // Local preview
			"https://noaa.github.io":                   true, // Production frontend
			"https://dohyunjung.github.io":             true, // Previous frontend (just in case)
		}

		// Check if the origin is allowed, otherwise default to "*" (or restrict it)
		// For development simplicity, we can keep "*" if specific origin check fails,
		// or strictly allow only the list. Here we prioritize the specific origin.
		if allowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			// Fallback for other origins (e.g. mobile apps or tools)
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true") // Important for authenticated requests

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
