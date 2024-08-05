package models

type PostResponse struct {
	InvoiceID  string `json:"invoiceID"`
	StatusCode string `json:"statusCode"`
	Amount     int    `json:"amount"`
	Curreny    string `json:"currency"`
}

type GetResponse struct {
	MaskedCardNumber     string `json:"maskedCardNumber"`
	MaskedExpirationDate string `json:"maskedExpirationDate"`
	StatusCode           string `json:"statusCode"`
}
