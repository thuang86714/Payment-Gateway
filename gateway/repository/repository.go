package repository

import (
	"github.com/processout-hiring/payment-gateway-thuang86714/gateway/db"
	"github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
)

func CreateTransaction(response *models.PostResponse) error {
	return db.DB.Create(response).Error
}

func GetTransaction(invoiceID string) (*models.PostResponse, error) {
	var response models.PostResponse
	err := db.DB.Where("invoice_id = ?", invoiceID).First(&response).Error
	return &response, err
}

func UpdateTransaction(response *models.PostResponse) error {
	return db.DB.Save(response).Error
}
