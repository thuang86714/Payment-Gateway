package service

import (
	"log"
	"time"

	"github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
	"github.com/processout-hiring/payment-gateway-thuang86714/bank/repository"
)

// validateTransaction simulates the complicated payment processing (KYC, AML, Anti-Terrorism....) in the bank.
// For simplicity, it just sleeps for 10 seconds to represent the time it takes to really process a payment.
func validateTransaction(){
	//sleep for 10 sec to reprsent time it takes to really process a payment
	time.Sleep(10 * time.Second)
	log.Println("Transaction validation completed!")
}


// CreateResponse generates a PostResponse based on the provided TransactionWithPSP and status.
// It performs transaction validation and balance update for valid transactions.
func CreateResponse(curTransactionWithPSP models.TransactionWithPSP, curStatus string) (models.PostResponse, error){
	var received float64
	var serviceFee float64
	//if it's a valid transaction, do processing
	if curStatus == "done" {
		validateTransaction()

		//change account balance
		addToBalance(curTransactionWithPSP.Currency, curTransactionWithPSP.AmountReceived)
		// Only if invoiceID is a new one, we add to the balance. 
		// Hence return response with AmountReceived and charge service fee
		received = curTransactionWithPSP.AmountReceived
		serviceFee = curTransactionWithPSP.ServiceFee
	}

	curResponse := models.PostResponse{
		CardNumber: curTransactionWithPSP.CardNumber,
		ExpirationDate: curTransactionWithPSP.ExpirationDate,
		InvoiceID: curTransactionWithPSP.InvoiceID,
		StatusCode: curStatus,
		AmountPayable: int(curTransactionWithPSP.AmountPayable),
		Currency: curTransactionWithPSP.Currency,         
		ServiceFee: serviceFee,
		ServiceProvider: curTransactionWithPSP.ServiceProvider,
		AmountReceived: received,
	}

	return curResponse, nil
}

// DoesInvoiceExists checks if an invoice with the given ID already exists.
// It returns true if the invoice exists, false otherwise.
// In case of an error, it logs the error and returns false.
func DoesInvoiceExists(invoiceID string) bool {
	exists, err := repository.InvoiceExists(invoiceID)
	if err != nil {
		log.Printf("Error checking if invoice exists: %v", err)
		return false // Assume it doesn't exist in case of error
	}
	return exists
}

// StoreInvoiceID stores a new invoice ID in the repository.
// It returns an error if the operation fails.
func StoreInvoiceID(invoiceID string) error {
	return repository.StoreInvoice(invoiceID)
}