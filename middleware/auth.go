package middleware

import (
	"context"

	"net/http"
	"strings"

	"github.com/d3vtech/graphQL/confiq"
)

// Define a GraphQL middleware to handle authentication and authorization
func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the API key from the request headers

		apiKeyHeader := r.Header.Get("Authorization")
		apiKeyParts := strings.Split(apiKeyHeader, " ")
		if len(apiKeyParts) != 2 || apiKeyParts[0] != "Bearer" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		apiKey := apiKeyParts[1]

		// Check if the API key is valid
		user, ok := confiq.Users[apiKey]
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Set the user information in the request context
		ctx := r.Context()
		ctx = context.WithValue(ctx, "user", user)

		// Determine the user role (client or admin)
		role := confiq.RoleClient
		if user == "admin" {
			role = confiq.RoleAdmin
		}
		ctx = context.WithValue(ctx, "role", role)

		// Create a new request with the updated context
		r = r.WithContext(ctx)

		// Call the original handler with the updated request
		h.ServeHTTP(w, r)
	})
}
