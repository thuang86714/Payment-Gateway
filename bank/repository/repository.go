package repository

import (
    "gorm.io/gorm"

	"github.com/processout-hiring/payment-gateway-thuang86714/bank/db"
	"github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
)

/*
TODO: current implementation brings several problems:
1. Testing can be more challenging because you can't easily inject a mock database for unit tests.
2. It creates a tight coupling between your repository and the global DB instance.
3. It's less flexible if you ever need to use different DB connections for different operations.
*/

// UpdateBalance adds the specified amount to the current balance for the given currency.
func UpdateBalance(currency string, amountToAdd float64) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
        var balance models.Balance
        
        // Try to find the existing balance or create a new one
        if err := tx.Where(models.Balance{Currency: currency}).FirstOrCreate(&balance).Error; err != nil {
            return err
        }

        // Add the new amount to the balance
        balance.Amount += amountToAdd

		// Save the updated balance
		return tx.Save(&balance).Error
	})
}

func InvoiceExists(invoiceID string) (bool, error) {
	var count int64
	err := db.DB.Model(&models.InvoiceID{}).Where("id = ?", invoiceID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func StoreInvoice(invoiceID string) error {
	return db.DB.Create(&models.InvoiceID{ID: invoiceID}).Error
}