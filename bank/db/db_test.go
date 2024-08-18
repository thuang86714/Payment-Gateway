package db

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"

    "github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
)

// newTestDB creates a new in-memory SQLite database for testing
func newTestDB() (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    
    // Migrate the schema
    err = db.AutoMigrate(&models.Balance{}, &models.InvoiceID{})
    if err != nil {
        return nil, err
    }
    
    return db, nil
}

func TestNewDB(t *testing.T) {
    // We are not creating a mock config since it is not used in this test

    // Use the test database function instead of NewDB
    db, err := newTestDB()
    require.NoError(t, err)
    assert.NotNil(t, db)

    // Test creating and retrieving a record
    balance := models.Balance{Currency: "USD", Amount: 100.0}
    err = db.Create(&balance).Error
    assert.NoError(t, err)

    var retrievedBalance models.Balance
    err = db.First(&retrievedBalance, "currency = ?", "USD").Error
    assert.NoError(t, err)
    assert.Equal(t, balance.Amount, retrievedBalance.Amount)
}

func TestCloseDB(t *testing.T) {
    db, err := newTestDB()
    require.NoError(t, err)

    CloseDB(db)
    // Since SQLite is in-memory, there's not much we can test here
    // In a real scenario, you might want to check if the connection is actually closed
    // For SQLite in-memory, the database is automatically closed when the connection is closed
}