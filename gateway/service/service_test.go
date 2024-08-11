package service

import (
	"strings"
	"testing"

	"github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
	"github.com/stretchr/testify/assert"
)

func TestMaskExpirationDate(t *testing.T) {
	tests := []struct {
		name           string
		expirationDate string
		expected       string
		expectError    bool
	}{
		{"Valid Date", "12/24", "**/24", false},
		{"Invalid Date Format", "1223", "", true},
		{"Another Valid Date", "05/25", "**/25", false},
		{"Another Invalid Date Format", "0525", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := maskExpirationDate(tt.expirationDate)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestMaskCardNumber(t *testing.T) {
	tests := []struct {
		name        string
		cardNumber  string
		expected    string
		expectError bool
	}{
		{"Valid card number", "1234567812345678", "************5678", false},
		{"Invalid card number (short length)", "12345678", "", true},
		{"Invalid card number (non-digit)", "1234abcd5678efgh", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maskedCardNumber, err := maskCardNumber(tt.cardNumber)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, maskedCardNumber)
			}
		})
	}
}

func TestCreateInvoiceID(t *testing.T) {
	tests := []struct {
		name      string
		invoice   models.Invoice
		expectLen int
	}{
		{
			name: "Valid invoice ID generation",
			invoice: models.Invoice{
				CardNumber:     "1234567812345678",
				ExpirationDate: "12/24",
				PricePerItem:   100,
				Currency:        "USD",
				CVV:            "123",
				Item:           "Test Item",
				Quantity:       1,
				Total:          100,
			},
			expectLen: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			invoiceID := CreateInvoiceID(tt.invoice)
			assert.Equal(t, tt.expectLen, len(invoiceID))
			assert.True(t, strings.HasPrefix(invoiceID, "INVD"))
		})
	}
}
