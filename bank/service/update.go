package service

import (
	"github.com/processout-hiring/payment-gateway-thuang86714/bank/repository"
)

//addToBalance add transactionWithPSP.AmountReceived to merchant account
func addToBalance(currency string, amount float64) error {
	return repository.UpdateBalance(currency, amount)
}