package pkg
import(
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/processout-hiring/payment-gateway-thuang86714/merchant/merchant_service"
)
func Exec(){
	//create connection

	//take input from the merchant, by cli
	fmt.Printf("Dear Merchant, Welcome to Tommy's Payment Gateway!\n")
	reader := bufio.NewReader(os.Stdin)
	var action int
	for {
		for {
			fmt.Print("Select 1 to process a payment. Select 2 to retrieve a previously-made payment.\n")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				os.Exit(1)
			}

			input = strings.TrimSpace(input)
			action, err = strconv.Atoi(input)
			if err != nil {
				fmt.Printf("Incorrect Input: %s. Try again.\n", input)
				continue
			}

			if action == 1 || action == 2 {
				break
			} else {
				fmt.Printf("Incorrect Input: %d. Try again.\n", action)
			}
		}

		if action == 1 {
			//process a payment
			processPayment()
		} else {
			//retrieve a payment detail
			retrievePayment()
		}
	}
}