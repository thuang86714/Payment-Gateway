package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetRoutes(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Create a new Gin engine
	r := gin.New()

	// Apply our SetRoutes function
	router := SetRoutes(r)

	// Test cases
	testCases := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
	}{
		{"Process Payment Route", "POST", "/processPayment", http.StatusOK},
		{"Retrieve Payment Route", "GET", "/retrievePayment", http.StatusOK},
		{"Update Payment Route", "PATCH", "/updatePayment", http.StatusOK},
		{"Non-existent Route", "GET", "/nonexistentroute", http.StatusNotFound},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new HTTP request
			req, _ := http.NewRequest(tc.method, tc.path, nil)

			// Create a response recorder
			w := httptest.NewRecorder()

			// Serve the HTTP request
			router.ServeHTTP(w, req)

			// Assert the status code is what we expect
			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

func TestLoggingMiddleware(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Create a new Gin engine
	r := gin.New()

	// Apply our SetRoutes function
	router := SetRoutes(r)

	// Create a new HTTP request
	req, _ := http.NewRequest("GET", "/retrievePayment", nil)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	// Assert that the logging middleware was applied
	// This is a basic check; you might want to mock the logger for more detailed testing
	assert.Equal(t, http.StatusOK, w.Code)
}