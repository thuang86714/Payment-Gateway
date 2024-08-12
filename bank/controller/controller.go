package controller

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/processout-hiring/payment-gateway-thuang86714/bank/service"
	"github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
)

type Controller struct{}

// NewUserController creates a new UserController
func NewController() *Controller {
	return &Controller{}
}

// CreateTransaction
func (ctr *Controller) ProcessTransaction(c *gin.Context) {
	//passed by middlewareBank.SendInitialResponse()
	transactionInterface, exists := c.Get("transaction")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction not found in context"})
		return
	}

	curTransactionWithPSP, ok := transactionInterface.(models.TransactionWithPSP)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid transaction data"})
		return
	}

	if err := service.ProcessTransaction(curTransactionWithPSP); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to process transaction, error = %v", err)})
		return
	}
}