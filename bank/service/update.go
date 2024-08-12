package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
	"github.com/processout-hiring/payment-gateway-thuang86714/bank/repository"
)

func addToBalance(currency string, amount float64) error {
	return repository.UpdateBalance(currency, amount)
}

//updateTransactionToGateway 
func updateTransactionToGateway(curResponse models.PostResponse) error {
	url := os.Getenv("GATEWAY_URL")
	if url == "" {
		return fmt.Errorf("GATEWAY_URL not set in environment")
	}
	url += "/updatePayment"

	// Marshal the invoice to JSON
	jsonData, err := json.Marshal(curResponse)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending POST request: %v", err)
	}
	defer resp.Body.Close()

	return nil
}