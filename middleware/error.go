package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

func ErrorHandler(schema graphql.Schema, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body
		var reqBody struct {
			Query     string                 `json:"query"`
			Variables map[string]interface{} `json:"variables"`
		}
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Execute the GraphQL query and catch any errors
		result := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  reqBody.Query,
			VariableValues: reqBody.Variables,
			Context:        r.Context(),
		})

		// Check if any errors occurred
		if len(result.Errors) > 0 {
			// Handle the errors
			// Format and transform the errors into GraphQL error objects
			gqlErrors := make([]map[string]interface{}, len(result.Errors))
			for i, err := range result.Errors {
				gqlError := map[string]interface{}{
					"message": err.Message,
				}

				// Add any additional information you want to include in the error object
				// For example, you can add the locations of the errors:
				if len(err.Locations) > 0 {
					gqlError["locations"] = err.Locations
				}

				gqlErrors[i] = gqlError

				// Log the errors, send notifications, or perform other error-specific actions
				// ...

				// You can also set appropriate status code and response headers for different types of errors
				// For example, if it's a validation error, you can set 400 Bad Request status code:
				if err.Message == "Validation error" {
					w.WriteHeader(http.StatusBadRequest)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			}

			// Marshal the error objects into JSON
			errResponse := map[string]interface{}{
				"errors": gqlErrors,
			}
			jsonResponse, err := json.Marshal(errResponse)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			// Set the response headers
			w.Header().Set("Content-Type", "application/json")

			// Return the error response
			w.Write(jsonResponse)
			return
		}

		// Marshal the data into JSON
		dataResponse := map[string]interface{}{
			"data": result.Data,
		}
		jsonResponse, err := json.Marshal(dataResponse)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Set the response headers
		w.Header().Set("Content-Type", "application/json")

		// Return the data response
		w.Write(jsonResponse)
	})
}
