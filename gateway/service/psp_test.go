package service

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
)

func TestNewTransaction(t *testing.T) {
	invoice := models.Invoice{
		CardNumber:     "1234567812345678",
		ExpirationDate: "12/24",
		CVV:            "123",
		Total:          1000,
	}
	invoiceID := "INVD123456789012"

	transaction := NewTransactionWithPSP(invoice, invoiceID)

	assert.Equal(t, invoiceID, transaction.InvoiceID, "InvoiceID should match")
	assert.Equal(t, float64(invoice.Total), transaction.AmountPayable, "AmountPayable should match invoice total")
	assert.True(t, transaction.ServiceFee > 0, "ServiceFee should be greater than 0")
	assert.Equal(t, transaction.AmountPayable-transaction.ServiceFee, transaction.AmountReceived, "AmountReceived should be correct")
	assert.Equal(t, invoice.CardNumber, transaction.CardNumber, "CardNumber should match")
	assert.Equal(t, invoice.ExpirationDate, transaction.ExpirationDate, "ExpirationDate should match")
	assert.Equal(t, invoice.CVV, transaction.CVV, "CVV should match")
	assert.NotEmpty(t, transaction.ServiceProvider, "ServiceProvider should not be empty")
}
