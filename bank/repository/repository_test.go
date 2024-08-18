package repository

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"

    "github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
)

func setupTestDB(t *testing.T) *gorm.DB {
    db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    require.NoError(t, err)

    err = db.AutoMigrate(&models.Balance{}, &models.InvoiceID{})
    require.NoError(t, err)

    return db
}

func TestGormRepository_UpdateBalance(t *testing.T) {
    db := setupTestDB(t)
    repo := NewGormRepository(db)

    err := repo.UpdateBalance("USD", 100.0)
    assert.NoError(t, err)

    var balance models.Balance
    err = db.Where("currency = ?", "USD").First(&balance).Error
    assert.NoError(t, err)
    assert.Equal(t, 100.0, balance.Amount)

    err = repo.UpdateBalance("USD", 50.0)
    assert.NoError(t, err)

    err = db.Where("currency = ?", "USD").First(&balance).Error
    assert.NoError(t, err)
    assert.Equal(t, 150.0, balance.Amount)
}

func TestGormRepository_InvoiceExists(t *testing.T) {
    db := setupTestDB(t)
    repo := NewGormRepository(db)

    exists, err := repo.InvoiceExists("INV001")
    assert.NoError(t, err)
    assert.False(t, exists)

    err = repo.StoreInvoice("INV001")
    assert.NoError(t, err)

    exists, err = repo.InvoiceExists("INV001")
    assert.NoError(t, err)
    assert.True(t, exists)
}

func TestGormRepository_StoreInvoice(t *testing.T) {
    db := setupTestDB(t)
    repo := NewGormRepository(db)

    err := repo.StoreInvoice("INV001")
    assert.NoError(t, err)

    var invoice models.InvoiceID
    err = db.Where("id = ?", "INV001").First(&invoice).Error
    assert.NoError(t, err)
    assert.Equal(t, "INV001", invoice.ID)
}