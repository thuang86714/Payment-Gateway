package models

import (
	"bufio"
	"fmt"
	"log"
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
	ExpirationDate string `json:"expirationDate"`
	PricePerItem   int    `json:"pricePerItem"`
	Currency       string `json:"currency"`
	CVV            string `json:"cvv"`
	Item           string `json:"item"`
	Quantity       int    `json:"quantity"`
	Timestamp      string `json:"timestamp"`
	Total          int    `json:"total"`
}

// InvoiceID represents the invoiceID string
type InvoiceID struct {
	ID string `gorm:"primaryKey"`
}

// TakeInputForNewInvoice takes all necessary inputs from the merchant, return a Invoice object
func TakeInputForNewInvoice() Invoice {
	reader := bufio.NewReader(os.Stdin)
	var curInvoice Invoice
	for {
		fmt.Printf("To process payment, please provide: \na. What do you want to check out today? \nb. How many? \nc. Price per item \nd. Currency \ne. Card Number\nf. Expiration Date\ng. CVV\n")

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
		fmt.Printf("If every field of this invoice is correct, enter 1, else enter 2\n")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error reading input:", err)
		}

		input = strings.TrimSpace(input)
		action, err := strconv.Atoi(input)
		if err != nil {
			log.Printf("Incorrect Input: %s. Try again.\n", input)
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
	fmt.Printf("To retrieve a previously-made payment, please provide: InvoiceID\n")
	reader := bufio.NewReader(os.Stdin)
	var invoiceID string
	for {
		fmt.Printf("Please enter the InvoiceID:")
		input, err := reader.ReadString('\n')
		if err != nil {
			logFatalError(err)
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
		fmt.Printf("If you enter 1234567812345678, you will send test request to gateway. But no transaction will be sent to bank\nPlease enter the card number: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			logFatalError(err)
		}

		cardNumber = strings.TrimSpace(input)
		//do input check
		if IsCardNumberInputValid(cardNumber) {
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
		fmt.Printf("Please enter the expeiration date in the MM/YY format, eg. 08/24: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			logFatalError(err)
		}

		expirationDate = strings.TrimSpace(input)
		//do input check
		if IsExpDateInputValid(expirationDate) {
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
	for {
		fmt.Printf("Please enter what are you going to check out today: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			logFatalError(err)
		}

		itemName = strings.TrimSpace(input)
		if isItemNameInputValid(itemName) {
			break
		}
	}

	return itemName
}

// takeQuantityForNewInvoice takes item from the merchant, return quantity
func takeQuantityForNewInvoice(reader *bufio.Reader) int {
	var quantity int
	for {
		fmt.Printf("Please enter how many are you going to check out today: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			logFatalError(err)
		}

		quantityStr := strings.TrimSpace(input)
		quantity, _ = strconv.Atoi(quantityStr)

		//do input check
		if isQuantityInputValid(quantityStr) {
			break
		}
		//print invalid input
		printInvalidInput(quantityStr)
	}

	return quantity
}

// takePricePerItemForNewInvoice takes price per item from the merchant, return price per item
func takePricePerItemForNewInvoice(reader *bufio.Reader) int {
	var pricePerItem int
	for {
		fmt.Printf("Please enter price per item: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			logFatalError(err)
		}

		priceStr := strings.TrimSpace(input)
		//do input check
		if isPricePerItemInputValid(priceStr) {
			pricePerItem, _ = strconv.Atoi(priceStr)
			break
		}
		//print invalid input
		printInvalidInput(priceStr)
	}

	return pricePerItem
}

// takeCurrencyForNewInvoice takes currecny from the merchant, return valid currency
func takeCurrencyForNewInvoice(reader *bufio.Reader) string {
	var currency string

	for {
		fmt.Printf("Please enter the currency, e.g USD: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			logFatalError(err)
		}

		currency = strings.TrimSpace(input)
		currency = strings.ToUpper(currency)
		//do input check
		if isCurrencyInputValid(currency) {
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
		fmt.Printf("Please enter the CVV, e.g 034: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			logFatalError(err)
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
		Currency:       curCurrecny,
		CVV:            curCVV,
		Item:           curItem,
		Quantity:       curQuantity,
		Timestamp:      curTimeStamp,
		Total:          curPricePerItem * curQuantity,
	}
	log.Printf("For this purchase, you invoice is %+v", curInvoice)
	return curInvoice
}

// isItemNameInputValid checks if the input a valid item name, return bool
func isItemNameInputValid(itemName string) bool {
	if len(itemName) == 0 {
		return false
	}
	return true
}

// isCardNumberInputValid checks if the input a valid card number, return bool
func IsCardNumberInputValid(cardNumber string) bool {
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

func IsExpDateInputValid(expirationDate string) bool {
	// Check if the length is exactly 5 characters
	if len(expirationDate) != 5 {
		return false
	}

	if expirationDate[2] != '/' {
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

// isCurrencyInputValid checks if the input is a valid currency, return bool
var getCurrency = money.GetCurrency

func isCurrencyInputValid(code string) bool {
	if len(code) == 0 {
		return false
	}
	currency := getCurrency(code)
	return currency != nil
}

// isQuantityInputValid checks if the input is a valid quantity, return bool
func isQuantityInputValid(quantityStr string) bool {
	if len(quantityStr) == 0 {
		return false
	}
	// Check if all characters are numeric
	for _, char := range quantityStr {
		if !unicode.IsDigit(char) {
			return false
		}
	}

	//quantity can not be 0 or even smaller
	quantity, _ := strconv.Atoi(quantityStr)
	if quantity <= 0 {
		return false
	}

	return true
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

// isPricePerItemInputValid checks the inputs is a valid price, return bool
func isPricePerItemInputValid(pricePerItemStr string) bool {
	if len(pricePerItemStr) == 0 {
		return false
	}
	// Check if all characters are numeric
	for _, char := range pricePerItemStr {
		if !unicode.IsDigit(char) {
			return false
		}
	}

	//price can not be 0 or even smaller
	price, _ := strconv.Atoi(pricePerItemStr)
	if price <= 0 {
		return false
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
	log.Printf("%s is not a valid input, please try again.\n", invalidInput)
}

func logFatalError(err error) {
	log.Fatal("Error reading input:", err)
}
