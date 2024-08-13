package service

import (
	"math/rand"

	"github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
)

// PSPFactory creates a new PSP transaction
type PSPFactory struct {
	ProviderName string
}

// newTransaction creates a new TransactionWithPSP instance
func (p PSPFactory) newTransaction(invoice models.Invoice, invoiceID string, feePercentage float64) models.TransactionWithPSP {
	return models.TransactionWithPSP{
		InvoiceID:       invoiceID,
		AmountPayable:   float64(invoice.Total),
		ServiceFee:      float64(invoice.Total) * feePercentage,
		AmountReceived:  float64(invoice.Total) - float64(invoice.Total)*feePercentage,
		CardNumber:      invoice.CardNumber,
		CVV:             invoice.CVV,
		ExpirationDate:  invoice.ExpirationDate,
		ServiceProvider: p.ProviderName,
		Currency:        invoice.Currency,
	}
}

// funcMap is a map of PSP factories, keyed by an integer identifier
var funcMap = map[int]PSPFactory{
	0: {ProviderName: "Checkout.com"},
	1: {ProviderName: "Stripe"},
	2: {ProviderName: "Klarna"},
	3: {ProviderName: "Square"},
	4: {ProviderName: "WorldPay"},
	5: {ProviderName: "Paypal"},
}

// NewTransactionWithPSP creates a new transaction with a selected PSP
func NewTransactionWithPSP(curInvoice models.Invoice, invoiceID string) models.TransactionWithPSP {
	//decided which PSP benefit the merchant the most
	curPSP, feePercentage := decidePSP()
	var createTransactionWithPSP = funcMap[curPSP]

	return createTransactionWithPSP.newTransaction(curInvoice, invoiceID, feePercentage)
}

// TODO: decidePSP is the core value of PSP platform, requiring complicated processes
// Here we use rand to make a simple representation
// decidePSP() decideds best PSP for this transaction, return int
func decidePSP() (int, float64) {
	percentage := decideServiceFeePercentage()
	return rand.Intn(6), percentage
}

// TODO: decideServiceFee is the core value of PSP platform, requiring complicated processes
// Here we use rand to make a simple representation
// decidePSP() decideds best PSP for this transaction, return float64
func decideServiceFeePercentage() float64 {
	feePercentage := float64(rand.Intn(100)+1) / float64(10000)
	return feePercentage
}
