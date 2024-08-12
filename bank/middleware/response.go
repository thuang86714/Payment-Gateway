package middlewareBank

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/processout-hiring/payment-gateway-thuang86714/bank/service"
	"github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
)

func SendInitialResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		var curTransactionWithPSP models.TransactionWithPSP
		if err := c.ShouldBindJSON(&curTransactionWithPSP); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Check if invoice ID already exists
		if service.DoesInvoiceExists(curTransactionWithPSP.InvoiceID) {
			initialResponse, err := service.CreateResponse(curTransactionWithPSP, "failed: invoiceID already exists")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create response"})
				c.Abort()
				return
			}
	
			c.JSON(http.StatusConflict, initialResponse)
			c.Abort()
			return
		}

		// Store the invoice ID
		if err := service.StoreInvoiceID(curTransactionWithPSP.InvoiceID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store invoice ID"})
			return
		}

		// Create and send initial "processing" response
		initialResponse, err := service.CreateResponse(curTransactionWithPSP, "processing")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create response"})
			c.Abort()
			return
		}

		c.JSON(http.StatusAccepted, initialResponse)

		// Store the transaction in the context for the next handler
		c.Set("transaction", curTransactionWithPSP)

		// Call the next handler
		c.Next()
	}
}