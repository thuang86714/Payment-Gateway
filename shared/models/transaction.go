package models

type TransactionWithPSP struct {
	InvoiceID       string  `json:"invoiceID"`
	AmountPayable   float64 `json:"amountPayable"`
	AmountReceived  float64 `json:"amountReceived"`
	CardNumber      string  `json:"cardNumber"`
	CVV             string  `json:"cvv"`
	ExpirationDate  string  `json:"expirationDate"`
	ServiceProvider string  `json:"serviceProvider"`
	ServiceFee      float64 `json:"serviceFee"`
	Currency       string `json:"currency"`
}