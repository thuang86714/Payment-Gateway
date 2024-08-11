package models

import (
	"gorm.io/gorm"
)

type PostResponse struct {
	gorm.Model
    CardNumber      string  `json:"cardNumber"`
    ExpirationDate  string  `json:"expirationDate"`
    InvoiceID       string  `json:"invoiceID"`
    StatusCode      string  `json:"statusCode"`
    AmountPayable   int     `json:"amountPayable"`
    Currency        string  `json:"currency"`          
    ServiceFee      float64 `json:"serviceFee"`
    ServiceProvider string  `json:"serviceProvider"`
    AmountReceived  float64 `json:"amountReceived"`
}

type GetResponse struct {
	MaskedCardNumber     string `json:"maskedCardNumber"`
	MaskedExpirationDate string `json:"maskedExpirationDate"`
	StatusCode           string `json:"statusCode"`
}
