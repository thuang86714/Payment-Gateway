package models

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/Rhymond/go-money"
)

// Invoice represents the Invoice API object
type Invoice struct {
	CardNumber     string `json:"cardNumber"`
	ExpirationDate string `json: "expirationDate"`
	PricePerItem   int    `json: "pricePerItem"`
	Curreny        string `json: "currency"`
	CVV            string `json: "cvv"`
	Item           string `json: "item"`
	Quantity       int    `json: "quantity"`
	timestamp      string `json: "timestamp"`
	Total          int    `json: "total"`
}

// TakeInputForNewInvoice takes all necessary inputs from the merchant, return a Invoice object
func TakeInputForNewInvoice() Invoice {
	reader := bufio.NewReader(os.Stdin)
	var curInvoice Invoice
	for {
		fmt.Printf("To process payment, please provide: \na. What do you want to check out today? \nb. How many? \nc. Price per item \n d. Currency \n e. Card Number\n f. Expiration Date\n g. CVV\n")

		curItem := takeItemForNewInvoice(reader)
		curQuantity := takeQuantityForNewInvoice(reader)
		curPricePerItem := takePricePerItemForNewInvoice(reader)
		curCardNumber := takeCardNumberForNewInvoice(reader)
		curExpirationDate := takeExpirationDateForNewInvoice(reader)
		curCurrecny := takeCurrencyForNewInvoice(reader)
		curCVV := takeCVVForNewInvoice(reader)
		for {
			if isCardValid(curCardNumber, curCVV, curExpirationDate) {
				break
			}
		}
		curInvoice = createInvoice(curCardNumber, curExpirationDate, curCurrecny, curCVV, curItem, curPricePerItem, curQuantity)
		fmt.Printf("If every field of this invoice is correct, enter 1, else enter 2")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		input = strings.TrimSpace(input)
		action, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Incorrect Input: %s. Try again.\n", input)
			continue
		}

		if action == 1 {
			break
		}
	}
	return curInvoice
}

// TakeInputForOldInvoice takes InvoiceID as input from the merchant, return InvoiceID
func TakeInputForOldInvoice() string {
	fmt.Printf("To retrieve a previously-made payment, please provide: InvoiceID")
	reader := bufio.NewReader(os.Stdin)
	var invoiceID string
	for {
		fmt.Printf("Please enter the InvoiceID:")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		invoiceID = strings.TrimSpace(input)
		if isInvoiceIDInputValid(invoiceID) {
			break
		}
		fmt.Printf("Not a valid input, please try again.\n")
	}
	return invoiceID
}

// takeCardNumberForNewInvoice takes card number from the merchant, return valid card number
func takeCardNumberForNewInvoice(reader *bufio.Reader) string {
	var cardNumber string
	for {
		fmt.Printf("Please enter the card number:")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		cardNumber = strings.TrimSpace(input)
		//do input check
		if isCardNumberInputValid(cardNumber) {
			break
		}
		//print invalid input
		printInvalidInput(cardNumber)
	}
	return cardNumber
}

// takeExpirationDateForNewInvoice takes expiration date from the merchant, return a valid expiration date
func takeExpirationDateForNewInvoice(reader *bufio.Reader) string {
	var expirationDate string
	for {
		fmt.Printf("Please enter the expeiration date in the MM/YY format, eg. 08/24")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		expirationDate = strings.TrimSpace(input)
		//do input check
		if isExpDateInputValid(expirationDate) {
			break
		}
		//print invalid input
		printInvalidInput(expirationDate)
	}
	return expirationDate
}

// takeItemForNewInvoice takes item from the merchant, return item
func takeItemForNewInvoice(reader *bufio.Reader) string {
	var itemName string
	fmt.Printf("Please enter what are you going to check out today:")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	itemName = strings.TrimSpace(input)

	return itemName
}

// takeQuantityForNewInvoice takes item from the merchant, return quantity
func takeQuantityForNewInvoice(reader *bufio.Reader) int {
	var quantity int
	fmt.Printf("Please enter how many are you going to check out today:")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	quantityStr := strings.TrimSpace(input)
	quantity, _ = strconv.Atoi(quantityStr)

	return quantity
}

// takePricePerItemForNewInvoice takes price per item from the merchant, return price per item
func takePricePerItemForNewInvoice(reader *bufio.Reader) int {
	var pricePerItem int
	fmt.Printf("Please enter price per item:")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	priceStr := strings.TrimSpace(input)
	pricePerItem, _ = strconv.Atoi(priceStr)

	return pricePerItem
}

// takeCurrencyForNewInvoice takes currecny from the merchant, return valid currency
func takeCurrencyForNewInvoice(reader *bufio.Reader) string {
	var currency string

	for {
		fmt.Printf("Please enter the currency, e.g USD:")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		currency = strings.TrimSpace(input)
		//do input check
		if isCVVInputValid(currency) {
			break
		}
		//print invalid input
		printInvalidInput(currency)
	}
	return currency
}

// takeCVVForNewInvoice takes CVV from the merchant, return valid CVV
func takeCVVForNewInvoice(reader *bufio.Reader) string {
	var cvv string
	for {
		fmt.Printf("Please enter the CVV, e.g 034:")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		cvv = strings.TrimSpace(input)
		//do input check
		if isCVVInputValid(cvv) {
			break
		}
		//print invalid input
		printInvalidInput(cvv)
	}
	return cvv
}

// createInvoice creates an invoice, which includes a timestamp, return an Invoice object
func createInvoice(curCardNumber, curExpirationDate, curCurrecny, curCVV, curItem string, curPricePerItem, curQuantity int) Invoice {
	curTime := time.Now()
	curTimeStamp := curTime.Format("2006-01-02 15:04:05")
	curInvoice := Invoice{
		CardNumber:     curCardNumber,
		ExpirationDate: curExpirationDate,
		PricePerItem:   curPricePerItem,
		Curreny:        curCurrecny,
		CVV:            curCVV,
		Item:           curItem,
		Quantity:       curQuantity,
		timestamp:      curTimeStamp,
		Total:          curPricePerItem * curQuantity,
	}
	printInvoice(curInvoice)
	return curInvoice
}

// isCardNumberInputValid checks if the input a valid card number, return bool
func isCardNumberInputValid(cardNumber string) bool {
	// Check if the length is exactly 16 characters
	if len(cardNumber) != 16 {
		return false
	}

	// Check if all of the characters are numeric
	for _, char := range cardNumber {
		if !unicode.IsDigit(char) {
			return false
		}
	}

	return true
}

// isExpDateInputValid checks if the input is a valid expiration date, return bool
var TimeNow = time.Now

func isExpDateInputValid(expirationDate string) bool {
	// Check if the length is exactly 5 characters
	if len(expirationDate) != 5 {
		return false
	}

	if expirationDate[2] != '/'{
		return false
	}

	// Get current month and year
	now := TimeNow()
	formatted := now.Format("01/06")
	curMonth, _ := strconv.Atoi(formatted[:2])
	curYear, _ := strconv.Atoi(formatted[3:])

	// Check if the first two characters are a valid month
	month, err := strconv.Atoi(expirationDate[:2])
	if err != nil || month < 1 || month > 12 {
		return false
	}

	// Check if the last two characters are a valid year
	year, err := strconv.Atoi(expirationDate[3:])
	if err != nil {
		return false
	}

	// Check if the expiration year is within the valid range
	if year < curYear || year > curYear+5 {
		return false
	}

	// If the year is the current year, check if the month is valid
	if year == curYear && month < curMonth {
		return false
	}

	// If the year is 5 years from now, check if the month is valid
	if year == curYear+5 && month > curMonth {
		return false
	}

	return true
}

// TODO: isCardValid checks with card issuer to see if it's valid
func isCardValid(cardNumber, cVV, expirationDate string) bool {
	return true
}

var getCurrency = money.GetCurrency

// isCurrencyInputValid checks if the input is a valid currency, return bool
func isCurrencyInputValid(code string) bool {
	currency := getCurrency(code)
	return currency != nil
}

// isCVVInputValid checks if the input is a valid CVV, return bool
func isCVVInputValid(code string) bool {
	// Check if the length is exactly 3 or 4 characters
	if len(code) != 3 && len(code) != 4 {
		return false
	}

	// Check if the all characters are numeric
	for _, char := range code {
		if !unicode.IsDigit(char) {
			return false
		}
	}

	return true
}

// isInvoiceIDInputValid checks the inputs is a valid invoiceID, return bool
func isInvoiceIDInputValid(invoiceID string) bool {
	// Check if the length is exactly 16 characters
	if len(invoiceID) != 16 {
		return false
	}

	// Check if the first 4 characters are "INVD"
	if invoiceID[:4] != "INVD" {
		return false
	}

	// Check if the last 12 characters are numeric
	for _, char := range invoiceID[4:] {
		if !unicode.IsDigit(char) {
			return false
		}
	}

	return true
}

// printInvalidInput prints invalid input
func printInvalidInput(invalidInput string) {
	fmt.Printf("%s is not a valid input, please try again.\n", invalidInput)
}

// printInvalidInput prints content of curInvoice
func printInvoice(curInvoice Invoice) {
	fmt.Printf("For this purchase, you check out: %s\n", curInvoice.Item)
	fmt.Printf("For this purchase, you check out: %d\n item", curInvoice.Quantity)
	fmt.Printf("For this purchase, price per item is: %d\n", curInvoice.PricePerItem)
	fmt.Printf("For this purchase, card number is: %s\n", curInvoice.CardNumber)
	fmt.Printf("For this purchase, expiration date is: %s\n", curInvoice.ExpirationDate)
	fmt.Printf("For this purchase, CVV is: %s\n", curInvoice.CVV)
	fmt.Printf("For this purchase, currecny is: %s\n", curInvoice.Curreny)
	fmt.Printf("For this purchase, total amount is: %d\n", curInvoice.Total)
}
