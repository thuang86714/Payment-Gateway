package service
import(
	"encoding/json"
	"log"
	
	"github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
)

//ProcessPayment is a POST function
func ProcessPayment(){
	//take input, return a invoice object
	curInvoice := shared.GetInputForInvoice()

	//do a rest POST to gateway
	jsonData, err := json.Marshal(curInvoice)
    if err != nil {
        log.Fatalf("Error marshalling JSON: %v", err)
    }

	//print out response
}

//Retrieve is a GET function
func RetrievePayment(){
	
}