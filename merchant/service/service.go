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
	//load .env parameters
	loadEnv()


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

//load from .env
func loadEnv(){
		// Define possible locations for the .env file
		envLocations := []string{
			".env",
			"../.env",
			"../../.env",
			"/root/.env",
		}
	
		// Try to load the .env file from each location
		var envLoaded bool
		for _, loc := range envLocations {
			if err := godotenv.Load(loc); err == nil {
				log.Printf("Loaded .env file from %s", loc)
				envLoaded = true
				break
			}
		}
	
		if !envLoaded {
			log.Fatal("Error loading .env file")
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