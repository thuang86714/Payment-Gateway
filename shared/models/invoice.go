package shared

import (
	"fmt"
	"time"

)

// Invoice represents the Invoice API object
type Invoice struct {
	cardNumber     string `json:"cardNumber"`
	expirationDate string `json: "expirationDate"`
	amount         int    `json: "amount"`
	curreny        string `json: "currency"`
	cVV            string `json: "CVV"`
	item           string `json: "item"`
	quantity       int    `json: "quantity"`
	timestamp      string `json: "timestamp"`
}

// getInputForInvoice take input from the merchant and return a Invoice object
func GetInputForInvoice() *Invoice {
	fmt.Printf("To process payment, please provide:\n a. card number\n b. expiration date\n c. amount of money\n d. currency\n e. CVV\n")

	return createInvoice()
}

func createInvoice(curCardNumber, curExpirationDate, curCurrecny, curCVV, curItem string, curAmount, curQuantity int) *Invoice {
	curTime := time.Now()
	curTimeStamp := curTime.Format("2006-01-02 15:04:05")
	curInvoice := &Invoice{
		cardNumber:     curCardNumber,
		expirationDate: curExpirationDate,
		amount:         curAmount,
		curreny:        curCurrecny,
		cVV:            curCVV,
		item:           curItem,
		quantity:       curQuantity,
		timestamp:      curTimeStamp,
	}
	return curInvoice
}
