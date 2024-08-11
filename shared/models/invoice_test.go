package models

import (
	"testing"
	"time"

	"github.com/Rhymond/go-money"
)

func TestIsCardNumberInputValid(t *testing.T) {
	var testcase = []struct {
		name     string
		input    string
		expected bool
	}{
		{"Less than 16 digits", "1234", false},
		{"More than 16 digits", "12121212121212121212", false},
		{"Contains non-digit (alphabet)", "121212D123345456", false},
		{"Contains non-digit (question mark)", "?121212121212121", false},
		{"Valid card number", "1212121212121212", true},
		{"Empty string", "", false},
		{"Exactly 16 non-digit characters", "abcdabcdabcdabcd", false},
	}

	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			got := IsCardNumberInputValid(tc.input)

			if tc.expected != got {
				t.Errorf("expected %t, but got %t", tc.expected, got)
			}
		})
	}
}

func TestIsExpDateInputValid(t *testing.T) {
	// Mock the current time to ensure consistent test results
	originalTimeNow := TimeNow
	defer func() { TimeNow = originalTimeNow }() // Restore the original function after the test

	mockNow, _ := time.Parse("01/06", "08/24")
	TimeNow = func() time.Time {
		return mockNow
	}

	testCases := []struct {
		name           string
		expirationDate string
		expected       bool
	}{
		{"Valid current month", "08/24", true},
		{"Valid future month same year", "12/24", true},
		{"Valid next year", "01/25", true},
		{"Valid 5 years from now, same month", "08/29", true},
		{"Valid 5 years from now, earlier month", "07/29", true},
		{"Invalid past month same year", "07/24", false},
		{"Invalid past year", "12/23", false},
		{"Invalid more than 5 years", "09/29", false},
		{"Invalid exactly 5 years plus one month", "09/29", false},
		{"Invalid format no slash", "0824", false},
		{"Invalid format wrong position slash", "082/4", false},
		{"Invalid month 00", "00/25", false},
		{"Invalid month 13", "13/25", false},
		{"Invalid format non-digit", "0a/2025", false},
		{"Invalid format too short", "8/2024", false},
		{"Invalid format too long", "08/245", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
      got := IsExpDateInputValid(tc.expirationDate)
			if tc.expected != got {
				t.Errorf("isExpDateInputValid(%q) = %t; expected %t", tc.expirationDate, got, tc.expected)
			}
		})
	}
}

func TestIsCardValid(t *testing.T) {
	return
}

// Mock function for testing
func mockGetCurrency(code string) *money.Currency {
	validCurrencies := map[string]*money.Currency{
		"USD": money.GetCurrency("USD"),
		"TWD": money.GetCurrency("TWD"),
		"EUR": money.GetCurrency("EUR"),
		"GBP": money.GetCurrency("GBP"),
	}

	return validCurrencies[code]
}

func TestIsCurrencyInputValid(t *testing.T) {
	// Override the getCurrency function with the mock function
	originalGetCurrency := getCurrency
	getCurrency = mockGetCurrency
	defer func() { getCurrency = originalGetCurrency }() // Restore the original function after the test

	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid currency - USD", "USD", true},
		{"Valid currency - TWD", "TWD", true},
		{"Valid currency - EUR", "EUR", true},
		{"Valid currency - GBP", "GBP", true},
		{"Invalid currency - ABC", "ABC", false},
		{"Invalid currency - 123", "123", false},
		{"Empty string", "", false},
		{"Lowercase valid currency - usd", "usd", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isCurrencyInputValid(tc.input)
			if tc.expected != got {
				t.Errorf("isCurrencyInputValid(%q) = %t; expected %t", tc.input, got, tc.expected)
			}
		})
	}
}

func TestIsCVVInputValid(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid CVV - 3 digits", "123", true},
		{"Valid CVV - 4 digits", "1234", true},
		{"Invalid CVV - 2 digits", "12", false},
		{"Invalid CVV - 5 digits", "12345", false},
		{"Invalid CVV - contains letters", "12a", false},
		{"Invalid CVV - contains special character", "12!", false},
		{"Empty CVV", "", false},
		{"Valid CVV - leading zero", "012", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isCVVInputValid(tc.input)
			if tc.expected != got {
				t.Errorf("isCVVInputValid(%q) = %t; expected %t", tc.input, got, tc.expected)
			}
		})
	}
}

func TestIsInvoiceIDInputValid(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid invoice ID", "INVD123456789012", true},
		{"Less than 16 characters", "INVD12345678901", false},
		{"More than 16 characters", "INVD1234567890123", false},
		{"Invalid prefix", "ABCD123456789012", false},
		{"Non-numeric characters in suffix", "INVD12345678ABCD", false},
		{"Empty string", "", false},
		{"Valid invoice ID with leading zeros", "INVD000000000001", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isInvoiceIDInputValid(tc.input)
			if tc.expected != got {
				t.Errorf("isInvoiceIDInputValid(%q) = %t; expected %t", tc.input, got, tc.expected)
			}
		})
	}
}
