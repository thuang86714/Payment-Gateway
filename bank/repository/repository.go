package repository

import (
	"gorm.io/gorm"

	"github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
)

type Repository interface {
	UpdateBalance(currency string, amountToAdd float64) error
	InvoiceExists(invoiceID string) (bool, error)
	StoreInvoice(invoiceID string) error
}

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

// UpdateBalance adds the specified amount to the current balance for the given currency.
func (r *GormRepository) UpdateBalance(currency string, amountToAdd float64) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
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

func (r *GormRepository) InvoiceExists(invoiceID string) (bool, error) {
	var count int64
	err := r.db.Model(&models.InvoiceID{}).Where("id = ?", invoiceID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *GormRepository) StoreInvoice(invoiceID string) error {
	return r.db.Create(&models.InvoiceID{ID: invoiceID}).Error
}
