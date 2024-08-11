package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
)

func maskExpirationDate(expirationDate string) (string, error) {
	if !models.IsExpDateInputValid(expirationDate) {
		return "", fmt.Errorf("Expiration Date format incorrect, got = %s", expirationDate)
	}
	return strings.Repeat("*", 2) + expirationDate[2:], nil
}

// maskCardNumber masks the first 12 digits of the card number, return string
func maskCardNumber(cardNumber string) (string, error) {
	if !models.IsCardNumberInputValid(cardNumber) {
		return "", fmt.Errorf("cardNumber format incorrect")
	}

	return strings.Repeat("*", len(cardNumber)-4) + cardNumber[len(cardNumber)-4:], nil
}

// CreateInvoiceID randomly create a 16 char long string for invoiceID, and check for collision, return string
var invoiceIDMap = make(map[string]bool)
var mu sync.Mutex // Mutex to handle concurrent access
func CreateInvoiceID(curInvoice models.Invoice) string {
	counter := 0

	for {
		// Convert the invoice to a string representation (for example, using the JSON representation)
		// Include a counter to ensure different hash values in case of collisions
		invoiceStr := fmt.Sprintf("%v-%d", curInvoice, counter)
		// Initialize a new SHA-256 hash generator
		hash := sha256.New()

		// Write the invoice string to the hash generator
		hash.Write([]byte(invoiceStr))

		// Finalize the hash computation and get the resulting hash as a byte slice
		hashBytes := hash.Sum(nil)

		// Convert the hash to a hexadecimal string
		hashHex := hex.EncodeToString(hashBytes)

		// Convert the hexadecimal string to a big integer
		bigInt := new(big.Int)
		bigInt.SetString(hashHex, 16)

		// Convert the big integer to a base-10 string
		hashDigits := bigInt.String()

		// Ensure the string is exactly 12 characters long, truncated if necessary
		if len(hashDigits) > 12 {
			hashDigits = hashDigits[:12]
		} else {
			hashDigits = fmt.Sprintf("%012s", hashDigits)
		}

		// Attach "INVD" prefix
		invoiceID := "INVD" + hashDigits

		// Debug statements
		fmt.Printf("Counter: %d, InvoiceStr: %s, InvoiceID: %s\n", counter, invoiceStr, invoiceID)

		// Check for collision
		mu.Lock()
		if !invoiceIDMap[invoiceID] {
			// No collision, mark ID as used and return it
			invoiceIDMap[invoiceID] = true
			mu.Unlock()
			return invoiceID
		}
		mu.Unlock()
		counter++
	}

	return ""
}
