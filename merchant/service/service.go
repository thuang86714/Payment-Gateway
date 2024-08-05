package service

import (
	"fmt"
	"log"
	"strconv"
	"bufio"
	"os"
	"strings"
	"github.com/joho/godotenv"
)

func Exec() {
	// Load environment variables from .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//Greet the merchant
	fmt.Printf("Dear Merchant, Welcome to Tommy's Payment Gateway!\n")

	// Main loop to handle merchant actions
	for {
		action := getMerchantAction()
		if err := handleAction(action); err != nil {
			log.Printf("Error: %v\n", err)
		}
	}
}

// getMerchantAction prompts the merchant to select an action.
func getMerchantAction() int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Select 1 to process a payment. Select 2 to retrieve a previously-made payment.\n")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %v", err)
		}

		input = strings.TrimSpace(input)
		action, err := strconv.Atoi(input)
		if err != nil || (action != 1 && action != 2) {
			fmt.Printf("Incorrect Input: %s. Try again.\n", input)
			continue
		}
		return action
	}
}

// handleAction processes the selected action.
func handleAction(action int) error {
	switch action {
	case 1:
		return processPayment()
	case 2:
		return retrievePayment()
	default:
		return fmt.Errorf("unknown action: %d", action)
	}
}