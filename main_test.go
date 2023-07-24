package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/graphql-go/graphql"

	"github.com/d3vtech/graphQL/graph"
	"github.com/d3vtech/graphQL/middleware"
)

func TestResolvePost(t *testing.T) {
	// Create a new GraphQL schema
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graph.RootQuery,
		},
	)
	if err != nil {
		t.Fatalf("Failed to create GraphQL schema: %v", err)
	}

	// Create a new request with the GraphQL query
	requestBody := `{"query": "query { post(id: \"1\") { id title} }"}`
	req, err := http.NewRequest("POST", "/graphql", bytes.NewBufferString(requestBody))
	if err != nil {
		t.Fatalf("Failed to create HTTP request: %v", err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create a new HTTP handler with the error handling middleware and the GraphQL schema
	handler := middleware.ErrorHandler(schema, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Execute the GraphQL query
		http.DefaultServeMux.ServeHTTP(w, r)
	}))

	// Serve the HTTP request
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %v, but got %v", http.StatusOK, rr.Code)
	}

	// Parse the response body
	var response map[string]interface{}
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	// Check the response data
	expectedData := map[string]interface{}{
		"data": map[string]interface{}{
			"post": map[string]interface{}{
				"id":    "1",
				"title": "Post 1",
			},
		},
	}
	if !reflect.DeepEqual(response, expectedData) {
		t.Errorf("Expected response %+v, but got %+v", expectedData, response)
	}
}

// Add more test functions for other resolver functions if needed
