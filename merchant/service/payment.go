package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
)

// ProcessPayment is a POST function
func processPayment() error {
	//take input, return a invoice object
	curInvoice := models.TakeInputForNewInvoice()

	//do a REST POST to gateway
	curResponse, err := postPayment(curInvoice)
	if err != nil {
		return fmt.Errorf("error posting payment: %v", err)
	}

	//print the response
	fmt.Printf("Response from service: \nInvoiceID: %s\nStatus Code: %s\nRetrieve: %f %s\n", curResponse.InvoiceID, curResponse.StatusCode, curResponse.AmountReceived, curResponse.Currency)
	
	return nil
}

// Retrieve is a GET function
func retrievePayment() error {
	//take input, return a invoice object
	invoiceID := models.TakeInputForOldInvoice()

	//do a rest GET to gateway
	curResponse, err := getPayment(invoiceID)
	if err != nil {
		return fmt.Errorf("error getting payment: %v", err)
	}

	//print the response
	fmt.Printf("Response from service: \nInvoiceID: %s\nStatus Code: %s\nMasked Card Number: %s\nMasked Expiration Date: %s\n", invoiceID, curResponse.StatusCode, curResponse.MaskedCardNumber, curResponse.MaskedExpirationDate)

	return nil
}

// postPayment sends the invoice to the gateway service and handles the response.
func postPayment(curInvoice models.Invoice) (models.PostResponse, error) {
	url := os.Getenv("GATEWAY_URL")
	if url == "" {
		return models.PostResponse{}, fmt.Errorf("GATEWAY_URL not set in environment")
	}
	url += "/processPayment"

	// Marshal the invoice to JSON
	jsonData, err := json.Marshal(curInvoice)
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

// getPayment retrieves the payment record from the gateway service
func getPayment(invoiceID string) (models.GetResponse, error) {
	url := os.Getenv("GATEWAY_URL")
	if url == "" {
		return models.GetResponse{}, fmt.Errorf("GATEWAY_URL not set in environment")
	}
	url += fmt.Sprintf("/retrievePayment?invoiceID=%s", invoiceID)

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.GetResponse{}, fmt.Errorf("error creating HTTP request: %v", err)
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.GetResponse{}, fmt.Errorf("error sending GET request: %v", err)
	}
	defer resp.Body.Close()

	// Handle the response
	if resp.StatusCode != http.StatusOK {
		return models.GetResponse{}, fmt.Errorf("received non-OK response status: %s", resp.Status)
	}

	var getResponse models.GetResponse
	if err := json.NewDecoder(resp.Body).Decode(&getResponse); err != nil {
		return models.GetResponse{}, fmt.Errorf("error decoding response: %v", err)
	}

	// Return the response
	return getResponse, nil
}
