package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

func main() {
	// Define GraphQL types
	userType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"id":    &graphql.Field{Type: graphql.ID},
				"name":  &graphql.Field{Type: graphql.String},
				"email": &graphql.Field{Type: graphql.String},
			},
		},
	)

	// Define root query
	rootQuery := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"user": &graphql.Field{
					Type: userType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{Type: graphql.ID},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						// Resolve user data based on the provided ID
						id, ok := p.Args["id"].(string)
						if !ok {
							return nil, nil
						}

						// Retrieve user data from the database or other sources
						// Here, we're returning dummy data for demonstration purposes
						if id == "1" {
							return map[string]interface{}{
								"id":    "1",
								"name":  "John Doe",
								"email": "john@example.com",
							}, nil
						}

						return nil, nil
					},
				},
			},
		},
	)

	// Define the GraphQL schema
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: rootQuery,
		},
	)
	if err != nil {
		log.Fatalf("Failed to create GraphQL schema: %v", err)
	}

	// Handle GraphQL requests
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body
		var reqBody struct {
			Query string `json:"query"`
		}
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Execute the GraphQL query
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: reqBody.Query,
		})

		// Convert the result to JSON
		jsonResult, err := json.Marshal(result)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Set the response headers
		w.Header().Set("Content-Type", "application/json")

		// Write the response
		w.Write(jsonResult)
	})

	// Start the server
	log.Println("Server is running on http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
