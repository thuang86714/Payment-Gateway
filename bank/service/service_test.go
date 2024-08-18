// File: bank/service/service_test.go

package service

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"

    "github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
)

type MockRepository struct {
    mock.Mock
}

func (m *MockRepository) UpdateBalance(currency string, amountToAdd float64) error {
    args := m.Called(currency, amountToAdd)
    return args.Error(0)
}

func (m *MockRepository) InvoiceExists(invoiceID string) (bool, error) {
    args := m.Called(invoiceID)
    return args.Bool(0), args.Error(1)
}

func (m *MockRepository) StoreInvoice(invoiceID string) error {
    args := m.Called(invoiceID)
    return args.Error(0)
}

func TestService_CreateResponse(t *testing.T) {
    mockRepo := new(MockRepository)
    svc := NewService(mockRepo)

    transaction := models.TransactionWithPSP{
        CardNumber:     "1234567890123456",
        ExpirationDate: "12/25",
        Currency:       "USD",
        AmountReceived: 100.0,
        ServiceFee:     5.0,
        InvoiceID:      "INV001",
    }

    mockRepo.On("UpdateBalance", "USD", 100.0).Return(nil)

    response, err := svc.CreateResponse(transaction, "done")
    assert.NoError(t, err)
    assert.Equal(t, "done", response.StatusCode)
    assert.Equal(t, 100.0, response.AmountReceived)
    assert.Equal(t, 5.0, response.ServiceFee)

    mockRepo.AssertExpectations(t)
}

func TestService_DoesInvoiceExists(t *testing.T) {
    mockRepo := new(MockRepository)
    svc := NewService(mockRepo)

    mockRepo.On("InvoiceExists", "INV001").Return(true, nil)
    mockRepo.On("InvoiceExists", "INV002").Return(false, nil)

    exists := svc.DoesInvoiceExists("INV001")
    assert.True(t, exists)

    exists = svc.DoesInvoiceExists("INV002")
    assert.False(t, exists)

    mockRepo.AssertExpectations(t)
}

func TestService_StoreInvoiceID(t *testing.T) {
    mockRepo := new(MockRepository)
    svc := NewService(mockRepo)

    mockRepo.On("StoreInvoice", "INV001").Return(nil)

    err := svc.StoreInvoiceID("INV001")
    assert.NoError(t, err)

    mockRepo.AssertExpectations(t)
}

func TestService_AddToBalance(t *testing.T) {
    mockRepo := new(MockRepository)
    svc := NewService(mockRepo)

    mockRepo.On("UpdateBalance", "USD", 100.0).Return(nil)

    err := svc.AddToBalance("USD", 100.0)
    assert.NoError(t, err)

    mockRepo.AssertExpectations(t)
}