package service

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
	"github.com/stretchr/testify/assert"
)

func TestPostTransactionToBank(t *testing.T) {
	// Save the original env var and defer its restoration
	originalBankURL := os.Getenv("BANK_URL")
	defer os.Setenv("BANK_URL", originalBankURL)

	tests := []struct {
		name               string
		transaction        models.TransactionWithPSP
		bankResponse       models.PostResponse
		bankResponseStatus int
		expectedError      bool
	}{
		{
			name: "Successful transaction",
			transaction: models.TransactionWithPSP{
				CardNumber: "4111111111111111",
				ExpirationDate: "12/25",
				AmountPayable: 1000,
				Currency: "USD",
			},
			bankResponse: models.PostResponse{
				StatusCode: "SUCCESS",
				AmountReceived: 990,
				ServiceFee: 10,
			},
			bankResponseStatus: http.StatusOK,
			expectedError: false,
		},
		{
			name: "Bank error response",
			transaction: models.TransactionWithPSP{
				CardNumber: "4111111111111111",
				ExpirationDate: "12/25",
				AmountPayable: 1000,
				Currency: "USD",
			},
			bankResponseStatus: http.StatusInternalServerError,
			expectedError: true,
		},
		{
			name: "Test card number",
			transaction: models.TransactionWithPSP{
				CardNumber: "1234567812345678",
				ExpirationDate: "12/25",
				AmountPayable: 3998,
				Currency: "USD",
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test server to mock the bank's response
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Check if it's the test card number
				if tt.transaction.CardNumber == "1234567812345678" {
					// Don't do anything, let the function handle it
					return
				}

				// Set the response status
				w.WriteHeader(tt.bankResponseStatus)

				// If it's a successful response, write the mock response body
				if tt.bankResponseStatus == http.StatusOK {
					json.NewEncoder(w).Encode(tt.bankResponse)
				}
			}))
			defer server.Close()

			// Set the BANK_URL environment variable to our test server
			os.Setenv("BANK_URL", server.URL)

			// Call the function
			response, err := PostTransactionToBank(tt.transaction)

			// Check the results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.transaction.CardNumber == "1234567812345678" {
					assert.Equal(t, "1234567812345678", response.CardNumber)
					assert.Equal(t, "SUCCESS", response.StatusCode)
				} else {
					assert.Equal(t, tt.bankResponse.StatusCode, response.StatusCode)
					assert.Equal(t, tt.bankResponse.AmountReceived, response.AmountReceived)
					assert.Equal(t, tt.bankResponse.ServiceFee, response.ServiceFee)
				}
			}
		})
	}
}
