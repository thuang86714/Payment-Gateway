package repository

import (
	"github.com/processout-hiring/payment-gateway-thuang86714/gateway/db"
	"github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
)

// CreateTransaction creates a new transaction record in the database
// It takes a pointer to a PostResponse model and returns an error if the operation fails
func CreateTransaction(response *models.PostResponse) error {
	return db.DB.Create(response).Error
}

// GetTransaction retrieves a transaction record from the database by its invoice ID
// It takes an invoice ID string and returns a pointer to a PostResponse model and an error
// If the transaction is not found, it returns a nil pointer and a "record not found" error
func GetTransaction(invoiceID string) (*models.PostResponse, error) {
	var response models.PostResponse
	err := db.DB.Where("invoice_id = ?", invoiceID).First(&response).Error
	return &response, err
}

// UpdateTransaction updates an existing transaction record in the database
// It takes a pointer to a PostResponse model and returns an error if the operation fails
// This function will update all fields of the record with the values in the provided model
func UpdateTransaction(response *models.PostResponse) error {
	return db.DB.Save(response).Error
}
