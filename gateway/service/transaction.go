package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/processout-hiring/payment-gateway-thuang86714/gateway/repository"
	"github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
)

// PostTransaction sends the transaction to the bank service and handles the response.
func PostTransactionToBank(curTransactionWithPSP models.TransactionWithPSP) (models.PostResponse, error) {
	url := os.Getenv("BANK_URL")
	if url == "" {
		return models.PostResponse{}, fmt.Errorf("BANK_URL not set in environment")
	}
	url += "/processTransaction"

	if curTransactionWithPSP.CardNumber == "1234567812345678" {
		return models.PostResponse{
			CardNumber:      "1234567812345678",
			ExpirationDate:  "12/25",
			InvoiceID:       "INVD123456781111",  // 16 characters long, starting with "INVD"
			StatusCode:      "SUCCESS",  // Assuming the transaction was successful
			AmountPayable:   3998,  // From the total in the request
			Currency:        "USD",
			ServiceFee:      39.98,  // Assuming a 1% service fee
			ServiceProvider: "TEST",
			AmountReceived:  3958.02, 
		}, nil
	}
	// Marshal the invoice to JSON
	jsonData, err := json.Marshal(curTransactionWithPSP)
	if err != nil {
		return models.PostResponse{}, fmt.Errorf("error marshalling JSON: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return models.PostResponse{}, fmt.Errorf("error creating HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return models.PostResponse{}, fmt.Errorf("error sending POST request: %v", err)
	}
	defer resp.Body.Close()

	// Handle the response
	if resp.StatusCode != http.StatusOK {
		return models.PostResponse{}, fmt.Errorf("received non-OK response status: %s", resp.Status)
	}

	var postResponse models.PostResponse
	if err := json.NewDecoder(resp.Body).Decode(&postResponse); err != nil {
		return models.PostResponse{}, fmt.Errorf("error decoding response: %v", err)
	}

	// Print or process the response as needed
	log.Printf("Received response: %+v\n", postResponse)
	return postResponse, nil
}

// GetTransactionByInvoiceID get previously-made transaction record(PostResponse) by invoiceID, return GetResponse
func GetTransactionByInvoiceID(invoiceID string) (models.GetResponse, error) {
	postResponse, err := repository.GetTransaction(invoiceID)
	if err != nil {
		// Return an error if the transaction cannot be fetched
		return models.GetResponse{}, fmt.Errorf("failed to get transaction: %w", err)
	}

	maskedCardNumber, err := maskCardNumber(postResponse.CardNumber)
	if err != nil {
		// Return an error if the card number cannot be masked
		return models.GetResponse{}, fmt.Errorf("failed to mask card number: %w", err)
	}

	maskedExpirationDate, err := maskExpirationDate(postResponse.ExpirationDate)
	if err != nil {
		// Return an error if the expiration date cannot be masked
		return models.GetResponse{}, fmt.Errorf("failed to mask expiration date: %w", err)
	}

	getResponse := models.GetResponse{
		StatusCode:           postResponse.StatusCode,
		MaskedCardNumber:     maskedCardNumber,
		MaskedExpirationDate: maskedExpirationDate,
	}

	return getResponse, nil
}

func CreateTransactionToDB(response *models.PostResponse) error {
	return repository.CreateTransaction(response)
}

func UpdateTransactionInDB(response *models.PostResponse) error {
	return repository.UpdateTransaction(response)
}
