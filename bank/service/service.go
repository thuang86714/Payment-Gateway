package service

import (
	"log"
	"time"

	"github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
	"github.com/processout-hiring/payment-gateway-thuang86714/bank/repository"
)



type Service struct {
    repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
    return &Service{repo: repo}
}

//validateTransaction simulates the complicated payment processing(KYC, AML, Anti-Terrorism....) in the bank. For simplicity, it would just sleep for 10 seconds
func (s *Service) validateTransaction(){
	//sleep for 10 sec to reprsent time it takes to really process a payment
	time.Sleep(10 * time.Second)
	log.Println("Transaction validation completed!")
}

func (s *Service) CreateResponse(curTransactionWithPSP models.TransactionWithPSP, curStatus string) (models.PostResponse, error){
	var received float64
	var serviceFee float64
	//if it's a valid transaction, do processing
	if curStatus == "done" {
		s.validateTransaction()

		//change account balance
		err := s.AddToBalance(curTransactionWithPSP.Currency, curTransactionWithPSP.AmountReceived)
		if err != nil {
            return models.PostResponse{}, err
        }
		//only if invoiceID is a new one, we add to the balance. Hence return response with AmountReceived and charge service fee
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

func (s *Service) DoesInvoiceExists(invoiceID string) bool {
	exists, err := s.repo.InvoiceExists(invoiceID)
	if err != nil {
		log.Printf("Error checking if invoice exists: %v", err)
		return false // Assume it doesn't exist in case of error
	}
	return exists
}

func (s *Service) StoreInvoiceID(invoiceID string) error {
	return s.repo.StoreInvoice(invoiceID)
}

//AddToBalance adds the specified amount to the current balance for the given currency
func (s *Service) AddToBalance(currency string, amount float64) error {
	return s.repo.UpdateBalance(currency, amount)
}